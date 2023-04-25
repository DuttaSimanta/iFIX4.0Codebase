FROM golang:alpine as builder
RUN mkdir /masterapp
ADD . /masterapp
WORKDIR /masterapp
RUN go get github.com/go-sql-driver/mysql
RUN go get golang.org/x/crypto/bcrypt
RUN go get github.com/Azure/azure-storage-blob-go/azblob
RUN go get github.com/gofrs/uuid
RUN go get github.com/dgrijalva/jwt-go
RUN go get github.com/tealeg/xlsx;
RUN go get github.com/disintegration/imaging
RUN go get github.com/pquerna/otp/totp
RUN go get github.com/natefinch/lumberjack
RUN go get gopkg.in/gomail.v2
RUN go build -o main .
CMD ["/app/main"]

FROM alpine:latest
WORKDIR /app
COPY --from=builder /masterapp/ /app
CMD ["/app/main"]

