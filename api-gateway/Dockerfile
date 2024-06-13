FROM golang:1-alpine3.19 AS build-stage
WORKDIR /explorite
COPY ./ /explorite
RUN mkdir -p /explorite/build
RUN go mod download
RUN go build -v -o /explorite/build/api ./cmd


FROM scratch
COPY --from=build-stage /explorite/build/api /
COPY --from=build-stage /explorite/template/ /template/
COPY --from=build-stage /explorite/static /static/
COPY --from=build-stage /explorite/.env /
EXPOSE 8080
CMD ["/api"]

