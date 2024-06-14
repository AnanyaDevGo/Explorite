FROM golang:1-alpine3.19 AS build-stage
WORKDIR /explorite
COPY ./ /explorite
RUN mkdir -p /explorite/build
RUN go mod download
RUN go build -v -o /explorite/build/api ./cmd
RUN apk add -U --no-cache ca-certificates


FROM scratch
COPY --from=build-stage /explorite/build/api /
COPY --from=build-stage /explorite/template/ /template/
COPY --from=build-stage /explorite/static /static/
COPY --from=build-stage /explorite/Logging /Logging/
COPY --from=build-stage /explorite/.env /
COPY --from=build-stage /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
EXPOSE 8080
CMD ["/api"]

