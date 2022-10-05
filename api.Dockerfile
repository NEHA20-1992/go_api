FROM golang:1.19-alpine

WORKDIR /app
RUN apk update && apk add libc-dev && apk add gcc apk add make

COPY ./go.mod go.sum./
RUN go mod download && go mod verify
 RUN  go get github.com/githubnemo/CompileDemon
COPY . .
COPY ./entrypoint.sh /entrypoint.sh

COPY http://raw.githubusercontent/eficode/wait-for/v2.1.0/wait-for usr/local/bin/wait-for
RUN chmod +rx usr/local/bin/wait-for /entrypoint.sh

ENTRYPOINT [ "sh" "/entrypoint.sh" ]
