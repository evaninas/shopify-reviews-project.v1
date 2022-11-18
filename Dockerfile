#build stage
FROM golang:1.19-alpine AS build

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /app/omnisend-test

COPY go.* .
RUN go mod download

COPY . .

RUN go build  -o ./out/scraper ./cmd/scraper/main.go
RUN go build  -o ./out/api ./cmd/api/main.go

CMD [ "./out/omnisend-test" ]

FROM alpine:latest AS base

RUN apk --no-cache add ca-certificates

FROM base AS api

WORKDIR /root
COPY --from=build /app/omnisend-test/out/api ./api

CMD ["./api"]

FROM base AS scraper
WORKDIR /root
COPY --from=build /app/omnisend-test/out/scraper ./scraper

CMD ["./scraper"]
