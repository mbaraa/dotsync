defmodule Dotsync.Application do
  # See https://hexdocs.pm/elixir/Application.html
  # for more information on OTP Applications
  @moduledoc false

  use Application

  @impl true
  def start(_type, _args) do
    children = [
      # Start the Telemetry supervisor
      DotsyncWeb.Telemetry,
      # Start the Ecto repository
      Dotsync.Repo,
      # Start the PubSub system
      {Phoenix.PubSub, name: Dotsync.PubSub},
      # Start Finch
      {Finch, name: Swoosh.Finch},
      # Start the Endpoint (http/https)
      DotsyncWeb.Endpoint
      # Start a worker by calling: Dotsync.Worker.start_link(arg)
      # {Dotsync.Worker, arg}
    ]

    # See https://hexdocs.pm/elixir/Supervisor.html
    # for other strategies and supported options
    opts = [strategy: :one_for_one, name: Dotsync.Supervisor]
    Supervisor.start_link(children, opts)
  end

  # Tell Phoenix to update the endpoint configuration
  # whenever the application is updated.
  @impl true
  def config_change(changed, _new, removed) do
    DotsyncWeb.Endpoint.config_change(changed, removed)
    :ok
  end
end
