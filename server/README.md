# Dotsync's server

[![server-rex-deploy](https://github.com/mbaraa/dotsync/actions/workflows/server-rex-deploy.yml/badge.svg)](https://github.com/mbaraa/dotsync/actions/workflows/server-rex-deploy.yml)

This is the server of [Dotsync](https://github.com/mbaraa/dotsync), where it stores all the synced dotfiles securely.

### Self-hosting:

#### You'll need:

1. Docker and Docker Compose
1. An internet connection
1. A valid SMTP server
1. ~314 seconds, Elixir likes to take a lot of time to build...

#### You'll do:

1. Clone the repo
1. Modify database's password in [docker-compose.yml](./docker-compose.yml) (just do a find and replace)
1. Copy `.env.example` into `.env.docker`
1. Set the environmental values, port, secrets, credentials...
1. Run `docker compose up -d`

- (Optional) use [Rex](https://github.com/mbaraa/rex) for a minimalistic deployment automation.
