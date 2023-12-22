# This file is responsible for configuring your application
# and its dependencies with the aid of the Config module.
#
# This configuration file is loaded before any dependency and
# is restricted to this project.

# General application configuration
import Config

config :dotsync,
  ecto_repos: [Dotsync.Repo]

# Configures the endpoint
config :dotsync, DotsyncWeb.Endpoint,
  url: [host: "0.0.0.0"],
  render_errors: [
    formats: [json: DotsyncWeb.ErrorJSON],
    layout: false
  ],
  pubsub_server: Dotsync.PubSub,
  live_view: [signing_salt: "1O+mfL/D"]

# Configures the mailer
config :swoosh, :api_client, false

config :dotsync, Dotsync.Mailer,
  adapter: Swoosh.Adapters.SMTP,
  relay: System.get_env("SMTP_SERVER"),
  port: System.get_env("SMTP_PORT") || "0" |> String.to_integer(),
  username: System.get_env("SMTP_USERNAME"),
  password: System.get_env("SMTP_PASSWORD"),
  auth: :always,
  ssl: false,
  tls: :always,
  retries: 0,
  no_mx_lookups: false,
  tls_options: [
    versions: [:"tlsv1.1", :"tlsv1.2", :"tlsv1.3"],
    verify: :verify_peer,
    cacerts: :public_key.cacerts_get(),
    server_name_indication: ~c"#{System.get_env("SMTP_SERVER")}",
    depth: 99
  ]

# Configures Elixir's Logger
config :logger, :console,
  format: "$time $metadata[$level] $message\n",
  metadata: [:request_id]

# Use Jason for JSON parsing in Phoenix
config :phoenix, :json_library, Jason

# Joken, jwt signer
config :joken, default_signer: System.get_env("JWT_SECRET")
# config :dotsync, Dotsync.Jwt, default_signer: System.get_env("JWT_SECRET")

# Import environment specific config. This must remain at the bottom
# of this file so it overrides the configuration defined above.
import_config "#{config_env()}.exs"
