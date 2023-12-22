# build stage
FROM node:18-alpine as build

WORKDIR /app

COPY . .
RUN npm i
RUN npm run build

# run stage
FROM node:18-alpine as run

WORKDIR /app

COPY --from=build /app/package*.json ./
COPY --from=build /app/build ./
COPY --from=build /app/.env ./

ARG PORT
ENV PORT=$PORT
EXPOSE $PORT

CMD ["node", "./index.js"]
