FROM elixir:1.14-alpine

ARG PORT
ARG SECRET_KEY_BASE
ARG DB_NAME
ARG DB_SERVER
ARG DB_USERNAME
ARG DB_PASSWORD
ARG DB_URL
ARG SMTP_SERVER
ARG SMTP_PORT
ARG SMTP_USERNAME
ARG SMTP_PASSWORD

ENV PORT=$PORT
ENV SECRET_KEY_BASE=$SECRET_KEY_BASE
ENV DB_NAME=$DB_NAME
ENV DB_SERVER=$DB_SERVER
ENV DB_USERNAME=$DB_USERNAME
ENV DB_PASSWORD=$DB_PASSWORD
ENV DB_URL=$DB_URL
ENV SMTP_SERVER=$SMTP_SERVER
ENV SMTP_PORT=$SMTP_PORT
ENV SMTP_USERNAME=$SMTP_USERNAME
ENV SMTP_PASSWORD=$SMTP_PASSWORD
ENV MIX_ENV=prod

RUN mkdir /app
WORKDIR /app

RUN apk add git ca-certificates wget gnupg gcc

COPY priv priv

COPY mix.exs mix.lock ./
COPY config config

RUN mix local.hex --force && \
    mix local.rebar --force && \
    mix archive.install hex phx_new && \
    mix deps.get --only prod

COPY lib lib
RUN mix compile

COPY ./docker-entrypoint.sh /entrypoint.sh
RUN chmod a+x /entrypoint.sh

ENTRYPOINT [ "/entrypoint.sh" ]
CMD ["run"]
