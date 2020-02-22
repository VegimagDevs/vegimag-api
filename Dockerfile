FROM golang:alpine as build-stage

ENV GO111MODULE=on \
		CGO_ENABLED=0 \
		GOOS=linux \
		GOARCH=amd64

WORKDIR /go/src/app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .
RUN go build -o /go/bin/app

FROM gcr.io/distroless/static as production-stage

ENV GIN_MODE=release \
		PORT=80

COPY --from=build-stage /go/bin/app /

EXPOSE 80
CMD ["/app"]
