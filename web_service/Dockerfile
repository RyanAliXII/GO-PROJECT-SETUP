FROM node:18-alpine3.16 as node-builder



FROM golang:1.19rc2-alpine3.16

WORKDIR /go/app

COPY . .
#copy node-js binaries
COPY --from=node-builder /usr/lib /usr/lib
COPY --from=node-builder /usr/local/share /usr/local/share
COPY --from=node-builder /usr/local/lib /usr/local/lib
COPY --from=node-builder /usr/local/include /usr/local/include
COPY --from=node-builder /usr/local/bin /usr/local/bin

RUN npm install

RUN go mod download && \
    go install github.com/cosmtrek/air@latest &&\
    go build -o ./dist/main

EXPOSE 8080

CMD [ "air" ]
