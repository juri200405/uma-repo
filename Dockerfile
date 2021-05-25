FROM golang:latest AS Builder
COPY ./app /app
RUN cd /app && go build -o manager cmd/server/main.go

FROM debian:stable-slim
COPY --from=Builder /app/manager /app/manager
COPY ./app/web /app/web
COPY ./app/public /app/public
COPY ./app/json /app/json
WORKDIR /app
