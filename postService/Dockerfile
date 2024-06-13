FROM golang:1-alpine3.19 AS build-stage
WORKDIR /post_svc
COPY ./ /post_svc
RUN mkdir -p /post_svc/build
RUN go mod download
RUN go build -v -o /post_svc/build/api ./cmd


FROM scratch
COPY --from=build-stage /post_svc/build/api /
COPY --from=build-stage /post_svc/.env /
COPY --from=build-stage /post_svc/Logging /Logging/
EXPOSE 50053
CMD ["/api"]