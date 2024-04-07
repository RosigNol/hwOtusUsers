FROM golang:1.22.2-alpine AS build

WORKDIR /app

COPY . .

RUN go mod download
RUN go build main.go
RUN go build migration.go

FROM alpine:latest as release
EXPOSE 8000
COPY --from=build /app/main /server
COPY --from=build /app/migration /migration
COPY --from=build /app/db/migrations /db/migrations
ENTRYPOINT ["/server"]
CMD ["-port", "8000"]
