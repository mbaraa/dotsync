defmodule Dotsync.Services.User do
  alias Dotsync.Jwt
  alias Dotsync.Repo
  alias Dotsync.Schemas
  alias Dotsync.Mailer

  @spec login_user(user_email :: String.t()) :: {:ok, any()} | {:error, atom()}
  def login_user(user_email) do
    case check_email(user_email) do
      true ->
        case create_login_token(user_email) do
          {:ok, token} ->
            [header, payload, signature] = String.split(token, ".")

            case Mailer.send_login_token("#{header}🔒#{signature}", user_email) do
              {:ok, _} ->
                {:ok, payload}

              {:error, _reason} ->
                {:error, :internal_error}
            end

          {:error, _reason} ->
            {:error, :internal_error}
        end

      false ->
        {:error, :invalid_email}
    end
  end

  @spec check_token(token :: String.t()) :: {:ok, email :: String.t()} | {:error, atom()}
  def check_token(token) do
    Jwt.check_callback(token, fn claims ->
      Map.get(claims, "email") |> check_or_create_user()
    end)
  end

  def delete_user(token) do
    token
    |> Jwt.check_callback(fn claims ->
      user = Repo.get_by(Schemas.User, email: Map.get(claims, "email"))

      cond do
        is_nil(user) -> {:error, :user_not_found}
        true -> Repo.delete(user)
      end
    end)
  end

  defp check_or_create_user(email) do
    {_, email} =
      case Schemas.User.changeset(%Schemas.User{}, %{email: email})
           |> Repo.insert() do
        {:ok, _} -> {:ok, email}
        # FIXME: handle other occuring errors...
        {:error, _reason} -> {:ok, email}
      end

    create_session_token(email)
  end

  defp create_session_token(email) do
    case Jwt.generate_and_sign(%{
           "email" => email,
           "exp" => :os.system_time(:seconds) + 30 * 24 * 60 * 60
         }) do
      {:ok, token, _claims} -> {:ok, token}
      {:error, _reason} -> {:error, :internal_error}
    end
  end

  defp create_login_token(email) do
    case Jwt.generate_and_sign(%{"email" => email}) do
      {:ok, token, _claims} -> {:ok, token}
      {:error, _reason} -> {:error, :internal_error}
    end
  end

  @type email :: charlist()
  defp check_email(email) do
    email =~
      ~r/(?:[a-zA-Z\d!#$%&'*+\/=?^_{|}~-]+(?:\.[a-zA-Z\d!#$%&'*+\/=?^_{|}~-]+)*|"(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21\x23-\x5b\x5d-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])*")@(?:(?:[a-zA-Z\d](?:[a-zA-Z\d-]*[a-zA-Z\d])?\.)+[a-zA-Z\d](?:[a-zA-Z\d-]*[a-zA-Z\d])?|\[(?:(?:25[0-5]|2[0-4]\d|[01]?\d\d?)\.){3}(?:25[0-5]|2[0-4]\d|[01]?\d\d?|[a-zA-Z\d-]*[a-zA-Z\d]:(?:[\x01-\x08\x0b\x0c\x0e-\x1f\x21-\x5a\x53-\x7f]|\\[\x01-\x09\x0b\x0c\x0e-\x7f])+)])/
  end
end
