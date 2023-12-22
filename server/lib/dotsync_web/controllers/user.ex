defmodule DotsyncWeb.User do
  use DotsyncWeb, :controller
  alias Dotsync.Services.User

  def login(conn, %{"email" => email}) do
    case User.login_user(email) do
      {:ok, token} ->
        json(conn, %{"token" => token})

      {:error, reason} ->
        {status, msg} =
          case reason do
            :invalid_email -> {400, reason}
            :internal_error -> {500, reason}
          end

        conn
        |> put_resp_content_type("application/json")
        |> put_status(status)
        |> json(%{"error" => msg})
    end
  end

  def verify_login(conn, %{"token" => token}) do
    case User.check_token(token) do
      {:ok, token} ->
        json(conn, %{"msg" => "ok 👍", "token" => token})

      {:error, _reason} ->
        conn
        |> put_resp_content_type("application/json")
        |> put_status(401)
        |> json(%{"error" => "Invalid token!"})
    end
  end

  def delete_user(conn, _params) do
    token =
      case get_req_header(conn, "authorization") do
        [] -> ""
        [token] -> token
      end

    case token |> User.delete_user() do
      {:ok, _} ->
        json(conn, %{"msg" => "ok 👍"})

      {:error, _reason} ->
        conn
        |> put_resp_content_type("application/json")
        |> put_status(400)
        |> json(%{"error" => :user_not_found})
    end
  end
end
