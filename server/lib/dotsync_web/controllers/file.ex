defmodule DotsyncWeb.File do
  use DotsyncWeb, :controller
  alias Dotsync.Services

  def add_file(conn, %{"path" => path, "content" => content}) do
    do_request_callback(conn, fn token ->
      Services.File.add_file(token, %{"path" => path, "content" => content})
    end)
  end

  def add_directory(conn, params) do
    do_request_callback(conn, fn token ->
      Services.File.add_directory(token, Map.get(params, "files"))
    end)
  end

  def get_synced_files(conn, _params) do
    do_request_callback(conn, fn token ->
      Services.File.get_synced_files(token)
    end)
  end

  def delete_file(conn, %{"path" => file_path}) do
    do_request_callback(conn, fn token ->
      Services.File.delete_file(token, file_path)
    end)
  end

  def download_files(conn, _params) do
    do_request_callback(conn, fn token ->
      Services.File.download_synced_files(token)
    end)
  end

  def upload_files(conn, params) do
    do_request_callback(conn, fn token ->
      Services.File.upload_synced_files(token, Map.get(params, "files"))
    end)
  end

  defp do_request_callback(conn, callback) do
    case get_auth_token(conn)
         |> callback.() do
      {:ok, resp} ->
        r = if is_map(resp) || is_list(resp), do: resp, else: %{"msg" => resp}
        json(conn, r)

      {:error, reason} ->
        {status, msg} =
          case reason do
            :internal_error -> {500, reason}
            :invalid_token -> {401, reason}
            :expired_token -> {401, reason}
            _ -> {400, reason}
          end

        conn
        |> put_status(status)
        |> json(%{"error" => msg})
    end
  end

  defp get_auth_token(conn) do
    case get_req_header(conn, "authorization") do
      [] -> ""
      [token] -> token
    end
  end
end
