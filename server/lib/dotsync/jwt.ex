defmodule Dotsync.Jwt do
  use Joken.Config

  def token_config() do
    default_claims()
    |> add_claim("iss", fn -> "Dotsync" end)
    |> add_claim("aud", fn -> "Dotsync" end)
    |> add_claim("exp", fn -> :os.system_time(:seconds) + 30 * 60 end)
  end

  @spec check_callback(token :: String.t(), callback :: function()) ::
          {:ok, any()} | {:error, atom()}
  def(check_callback(token, callback)) do
    case check(token) do
      {:ok, claims} -> claims |> callback.()
      {:error, reason} -> {:error, reason}
    end
  end

  @spec check(token :: String.t()) :: {:ok, any()} | {:error, atom()}
  def check(token) do
    case token
         |> String.trim()
         |> verify() do
      {:ok, claims} ->
        cond do
          Map.get(claims, "exp") < :os.system_time(:seconds) -> {:error, :expired_token}
          true -> {:ok, claims}
        end

      {:error, _reason} ->
        {:error, :invalid_token}
    end
  end
end
