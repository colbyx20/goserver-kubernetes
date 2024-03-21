FROM golang:1.22.1-alpine3.19 as base 

WORKDIR /app 

COPY go.mod /app/

COPY . .
RUN go build -o /bin/server main.go

# copy binary from /bin/server into /bin/server on scratch image from base
FROM scratch AS server
COPY --from=base /bin/server /bin/server 

ENTRYPOINT [ "/bin/server" ]