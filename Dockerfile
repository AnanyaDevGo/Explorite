FROM golang:1.21.2-alpine3.18 AS build-stage
WORKDIR /explorite
COPY ./ /explorite
RUN mkdir -p /explorite/build
RUN go mod download
RUN go build -v -o /explorite/build/api ./cmd


FROM gcr.io/distroless/static-debian11
COPY --from=build-stage /explorite/build/api /
COPY --from=build-stage /explorite/template/ /template/
COPY --from=build-stage /explorite/static /static/
COPY --from=build-stage /explorite/.env /
EXPOSE 8080
CMD ["/api"]

