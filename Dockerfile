#FROM golang:1.16.8-alpine
#
#WORKDIR /pharmacy
#
#COPY . .
#
#RUN apk --no-cache -U add libc-dev build-base
#RUN go mod download && go mod tidy
#

FROM pharmacybase as builder
WORKDIR /pharmacy 
COPY . .
RUN go build -ldflags "-linkmode external -extldflags -static" -o main .

FROM scratch
COPY --from=builder /pharmacy/main ./main
CMD ["./main"]

