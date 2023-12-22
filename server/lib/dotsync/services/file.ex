defmodule Dotsync.Services.File do
  alias Dotsync.Repo
  alias Dotsync.Schemas
  alias Dotsync.Jwt
  import Ecto.Query, only: [from: 2]

  def add_file(session_token, %{"path" => path, "content" => content}) do
    cond do
      path |> String.length() == 0 ->
        {:error, :empty_file_path}

      content |> String.length() == 0 ->
        {:error, :empty_file_content}

      not is_allowed_content_size?(content) ->
        {:error, :file_content_is_too_large}

      true ->
        session_token
        |> get_user_callback(fn user ->
          case user
               |> Ecto.build_assoc(:files)
               |> Schemas.File.changeset(%{
                 path: path,
                 content: encrypt_file_content(content),
                 user_id: user.id
               })
               |> Repo.insert() do
            {:ok, _} -> {:ok, "uploaded #{path}"}
            {:error, _reason} -> {:error, :file_exits_for_user}
          end
        end)
    end
  end

  def add_directory(session_token, files) do
    large_files =
      Enum.map(files, fn %{"path" => path, "content" => content} ->
        if not is_allowed_content_size?(content) do
          path
        end
      end)
      |> Enum.filter(&(not is_nil(&1)))

    cond do
      is_nil(files) || length(files) == 0 ->
        {:error, :empty_files_list}

      length(large_files) != 0 ->
        {:error, "the files [#{large_files |> Enum.join(", ")}] are larger than 256KiB"}

      true ->
        session_token
        |> get_user_callback(fn user ->
          for %{"path" => path, "content" => content} <- files |> Enum.filter(&(not is_nil(&1))) do
            case user
                 |> Ecto.build_assoc(:files)
                 |> Schemas.File.changeset(%{
                   path: path,
                   content: encrypt_file_content(content),
                   user_id: user.id
                 })
                 |> Repo.insert() do
              # HACK: ignoring erorrs cuz there's no problem with duplicates here!
              _ -> {:ok, "uploaded #{path}"}
            end
          end

          {:ok, "done"}
        end)
    end
  end

  def get_synced_files(session_token) do
    session_token
    |> get_user_callback(fn user ->
      q =
        from f in Schemas.File,
          where: f.user_id == ^user.id,
          select: f.path

      case Repo.all(q) do
        nil ->
          {:error, :no_files_were_found}

        [] ->
          {:error, :no_files_were_found}

        files ->
          {:ok, files}
      end
    end)
  end

  def delete_file(session_token, file_path) do
    cond do
      file_path |> String.length() == 0 ->
        {:error, :empty_file_path}

      true ->
        session_token
        |> get_user_callback(fn user ->
          file = Repo.get_by(Schemas.File, path: file_path, user_id: user.id)

          cond do
            is_nil(file) ->
              case delete_directory(user, file_path) do
                :ok -> {:ok, "deleted #{file_path}"}
                :error -> {:error, :file_doesnt_exits_for_user}
              end

            true ->
              case file
                   |> Repo.delete() do
                {:ok, _} ->
                  {:ok, "deleted #{file_path}"}

                {:error, _reason} ->
                  {:error, :file_doesnt_exits_for_user}
              end
          end
        end)
    end
  end

  def delete_directory(user, dir_path) do
    files_to_delete =
      from(f in Schemas.File,
        where: f.user_id == ^user.id
      )
      |> Repo.all()
      |> Enum.filter(fn file ->
        file.path |> String.starts_with?(dir_path)
      end)
      |> Enum.filter(&(not is_nil(&1)))

    cond do
      length(files_to_delete) == 0 ->
        {:error, :directory_is_empty_or_unsynced}

      true ->
        files_to_delete |> Enum.each(&Repo.delete(&1))
    end
  end

  def download_synced_files(session_token) do
    session_token
    |> get_user_callback(fn user ->
      q =
        from f in Schemas.File,
          where: f.user_id == ^user.id,
          select: %{"path" => f.path, "content" => f.content}

      case Repo.all(q) do
        nil ->
          {:error, :no_files_were_found}

        [] ->
          {:error, :no_files_were_found}

        files ->
          {:ok,
           files
           |> Enum.map(fn file ->
             %{file | "content" => Map.get(file, "content") |> decrypt_file_content()}
           end)}
      end
    end)
  end

  def upload_synced_files(session_token, files) do
    large_files =
      Enum.map(files, fn %{"path" => path, "content" => content} ->
        if not is_allowed_content_size?(content) do
          path
        end
      end)
      |> Enum.filter(&(not is_nil(&1)))

    cond do
      is_nil(files) || length(files) == 0 ->
        {:error, :empty_files_list}

      length(large_files) != 0 ->
        {:error, "the files [#{large_files |> Enum.join(", ")}] are larger than 256KiB"}

      true ->
        session_token
        |> get_user_callback(fn user ->
          changed =
            files
            |> Enum.map(fn %{"path" => path, "content" => content} ->
              case Repo.get_by(Schemas.File, path: path, user_id: user.id) do
                nil -> nil
                file -> {file, content}
              end
            end)
            |> Enum.filter(&(not is_nil(&1)))
            |> Enum.map(fn {file, content} ->
              if file.content != content do
                file
                |> Ecto.Changeset.change(content: encrypt_file_content(content))
                |> Repo.update()
              end
            end)
            |> Enum.filter(&(not is_nil(&1)))
            |> Enum.map(fn {_, changed} -> changed.path end)

          {:ok, changed}
        end)
    end
  end

  defp is_allowed_content_size?(content) do
    case Base.decode64(content) do
      {:ok, _} -> content
      :error -> content
    end
    |> String.length() < 256 * 1024
  end

  defp decrypt_file_content(content) do
    case Phoenix.Token.decrypt(
           DotsyncWeb.Endpoint,
           System.get_env("SECRET_KEY_BASE"),
           content
         ) do
      {:ok, content} -> content
      {:error, _} -> ""
    end
  end

  defp encrypt_file_content(content) do
    Phoenix.Token.encrypt(
      DotsyncWeb.Endpoint,
      System.get_env("SECRET_KEY_BASE"),
      content |> encode_file_content_b64()
    )
  end

  defp encode_file_content_b64(content) do
    case Base.decode64(content) do
      {:ok, _} -> content
      :error -> Base.encode64(content)
    end
  end

  defp get_user_callback(token, callback) do
    token
    |> Jwt.check_callback(fn claims ->
      user = Repo.get_by(Schemas.User, email: Map.get(claims, "email"))

      cond do
        is_nil(user) -> {:error, :user_not_found}
        true -> user |> callback.()
      end
    end)
  end
end
