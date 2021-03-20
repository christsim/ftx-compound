FROM golang:1.16-buster as build

WORKDIR /go/src/app
ADD . /go/src/app

RUN go get -d -v ./...

RUN go build -o /go/bin/app

# Now copy it into our base image.
FROM gcr.io/distroless/base-debian10
COPY --from=build /go/bin/app /

#ENV API_KEY
#ENV API_SECRET
#ENV API_SUBACCOUNT
#ENV COIN
#ENV YEARLY_RATE

CMD ["/app"]
