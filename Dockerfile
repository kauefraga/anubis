FROM golang:1.22.5-alpine AS build

RUN apk update && apk add --no-cache git

WORKDIR /anubis
COPY . .

RUN go mod download

ENV CGO_ENABLED=0

RUN go build -ldflags="-s -w" -o /bin/anubis ./cmd/main.go

FROM scratch AS production

COPY --from=build /bin/anubis /bin/anubis

ENTRYPOINT [ "/bin/anubis" ]
