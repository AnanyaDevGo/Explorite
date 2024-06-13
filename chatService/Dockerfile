FROM golang:1.22.1-alpine3.18 AS build-stage
WORKDIR /chat_svc
COPY ./ /chat_svc
RUN mkdir -p /chat_svc/build
RUN go mod download
RUN go build -v -o /chat_svc/build/api ./cmd


FROM scratch
COPY --from=build-stage /chat_svc/build/api /
COPY --from=build-stage /chat_svc/.env /
COPY --from=build-stage /chat_svc/Logging /Logging/
EXPOSE 50054
CMD ["/api"]