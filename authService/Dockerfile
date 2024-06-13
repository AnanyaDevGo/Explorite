FROM golang:1.22.1-alpine3.18 AS build-stage
WORKDIR /auth_svc
COPY ./ /auth_svc
RUN mkdir -p /auth_svc/build
RUN go mod download
RUN go build -v -o /auth_svc/build/api ./cmd


FROM  scratch
COPY --from=build-stage /auth_svc/build/api /
COPY --from=build-stage /auth_svc/.env /
COPY --from=build-stage /auth_svc/Logging /Logging/
EXPOSE 50052
CMD ["/api"]