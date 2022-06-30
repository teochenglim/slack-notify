FROM golang:1.18-alpine AS build
WORKDIR /app
COPY . .
RUN go mod tidy
RUN CGO_ENABLED=0 GOOS=linux go build

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=build /app/slack-notify /bin/

CMD /bin/slack-notify
