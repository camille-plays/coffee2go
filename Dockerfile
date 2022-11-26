FROM golang:alpine
LABEL maintainer="Camille Plays"

RUN apk update && apk add --no-cache git && apk add --no-cach bash && apk add build-base
RUN mkdir /app
WORKDIR /app
COPY . .
RUN go get -d -v ./... && go install -v ./... && go build -o /coffee2go

FROM scratch
COPY --from=builder /go/bin/coffee2go /go/bin/coffee2go
ENTRYPOINT ["/go/bin/coffee2go"]