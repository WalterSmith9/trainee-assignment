FROM golang:1.21.0-alpine AS builder

RUN apk update &&  \
    apk add --no-cache git ca-certificates tzdata bash && \
    update-ca-certificates

ENV USER=app
ENV UID=10001
ENV TZ=Europe/Moscow

RUN adduser --disabled-password --gecos "" --home "/nonexistent" --shell "/sbin/nologin" --no-create-home \
    --uid "${UID}" "${USER}"

WORKDIR $GOPATH/src/trainee-assignment
COPY . .

RUN go get -d -v ./...
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o /go/bin/trainee-assignment

FROM scratch

COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd
COPY --from=builder /etc/group /etc/group

WORKDIR /go/bin
USER app:app

COPY --from=builder /go/bin/trainee-assignment /go/bin/trainee-assignment