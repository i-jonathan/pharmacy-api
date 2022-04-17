FROM golang:1.16.8-alpine as builder
RUN apk --no-cache -U add libc-dev build-base
WORKDIR /pharmacy
COPY go.mod .
COPY go.sum .
RUN go mod download
COPY . .
RUN go build -ldflags "-linkmode external -extldflags -static" -o main .

FROM scratch
COPY --from=builder /pharmacy/docs /docs
COPY --from=builder /pharmacy/main ./main
CMD ["./main"]
