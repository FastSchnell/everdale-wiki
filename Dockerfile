FROM golang:1.15-alpine as build

ENV GO111MODULE=on

WORKDIR /go/release

ADD . .

RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -ldflags="-s -w" -installsuffix cgo -o nt cmd/api.go

FROM scratch as prod

COPY --from=build /go/release/ew /

COPY --from=build /go/release/conf ./conf
COPY --from=build /go/release/static ./static


CMD ["/ew"]