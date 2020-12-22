#1st build
FROM golang:1.15
WORKDIR /go/src/github.com/emailtovamos/GoAPI
COPY cli ./cli
COPY vendor ./vendor
COPY accounts ./accounts
COPY authentication ./authentication
COPY utils ./utils
COPY handlers ./handlers
COPY .env ./.env
RUN CGO_ENABLED=0 GOOS=linux go install ./cli/server

#2nd Stage
FROM alpine:latest
WORKDIR /app/
COPY --from=0 /go/bin/server ./binary
CMD ./binary