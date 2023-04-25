FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN ["apt-get", "update"]
RUN ["apt-get", "-y", "install", "vim"]
RUN go get github.com/go-sql-driver/mysql
RUN go get gopkg.in/gomail.v2
RUN go build -o main .
CMD ["/app/main"]
