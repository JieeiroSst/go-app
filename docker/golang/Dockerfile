FROM golang
WORKDIR ./app
ENV GO111MODULE=on
COPY ../../go.mod go.sum ./
RUN go mod download
COPY ../.. .
RUN go get -u gorm.io/driver/mysql
RUN go get -u gorm.io/gorm
RUN go get -u github.com/kelseyhightower/envconfig
RUN export WEB_MYSQL_DSN="root:1234@tcp(localhost:3306)/db?parseTime=True"
RUN export WEB_PORT=:4040
EXPOSE 8080
RUN go build -o main .
CMD ["./main"]