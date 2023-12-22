defmodule Dotsync.Repo do
  use Ecto.Repo,
    otp_app: :dotsync,
    adapter: Ecto.Adapters.MyXQL
end
