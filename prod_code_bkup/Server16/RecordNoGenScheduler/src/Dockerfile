FROM golang:latest
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go get github.com/go-sql-driver/mysql
RUN go get github.com/tealeg/xlsx
RUN go get github.com/natefinch/lumberjack
RUN go build -o main .
CMD ["/app/main"]