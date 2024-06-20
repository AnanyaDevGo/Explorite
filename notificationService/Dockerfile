FROM golang:1.22.1-alpine3.18 AS build-stage
WORKDIR /notification_svc
COPY ./ /notification_svc
RUN mkdir -p /notification_svc/build
RUN go mod download
RUN go build -v -o /notification_svc/build/api ./cmd
RUN apk add -U --no-cache ca-certificates



FROM scratch
COPY --from=build-stage /notification_svc/build/api /
COPY --from=build-stage /notification_svc/.env /
COPY --from=build-stage /notification_svc/Logging /Logging/
COPY --from=build-stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 50055
CMD ["/api"]