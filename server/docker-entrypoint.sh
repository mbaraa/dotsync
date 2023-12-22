#!/bin/sh

set -e

if [[ "$1" = "run" ]]
then
    mix ecto.create
    mix ecto.migrate
    mix phx.server
else
    exec "$@"
fi
