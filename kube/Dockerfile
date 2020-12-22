#1st build
FROM golang:1.15
WORKDIR /go/src/github.com/emailtovamos/GoAPI/kube

COPY cli ./cli
COPY vendor ./vendor

RUN CGO_ENABLED=0 GOOS=linux go install ./cli/server
#RUN CGO_ENABLED=0 GOOS=linux go build -i -o binary github.com/emailtovamos/GoAPI/kube

#2nd Stage

FROM alpine:latest
WORKDIR /app/
COPY --from=0 /go/bin/server ./binary
CMD ./binary

#COPY --from=0 /go/bin/main ./binary
#COPY --from=0 binary ./binary

#CMD ls
#CMD ./binary