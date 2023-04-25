FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN ["apt-get", "update"]
RUN ["apt-get", "install", "-y", "vim"]
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/jasonlvhit/gocron
RUN go get github.com/natefinch/lumberjack
RUN go build -o main .
CMD ["/app/main"]
