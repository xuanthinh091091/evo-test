FROM golang

WORKDIR /api

COPY src/go.mod .

COPY src/go.sum .

RUN go mod vendor

COPY src/ .

ENV PORT 3000

RUN go build

run ls -l

CMD ["./src"]