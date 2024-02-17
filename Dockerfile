FROM golang:1.20-buster AS builder

RUN go version

WORKDIR /src

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ../ ./
RUN go build -o /src/app cmd/app/main.go
RUN chmod +x /src/app

FROM gcr.io/distroless/base-debian11

WORKDIR /src

COPY --from=builder /src/app /src/app
COPY --from=builder /src/resources/banner.txt /src/banner.txt

EXPOSE 8080
USER nonroot:nonroot

ENV BANNER_PATH=/src/banner.txt
ENTRYPOINT ["/src/app"]