FROM golang:alpine as build
WORKDIR /go/src
COPY . .
RUN go build -o /bin/app .

FROM alpine
COPY --from=build /bin/app /bin/app
ENTRYPOINT ["/bin/app"]