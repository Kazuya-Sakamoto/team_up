FROM golang:1.13.5

ENV GOPATH $GOPATH:/go
ENV PATH $PATH:$GOPATH/bin
WORKDIR /go/src/app

RUN apt-get update
RUN go get "github.com/go-sql-driver/mysql"
RUN go get "github.com/beego/bee"
RUN go get "github.com/astaxie/beego"
RUN go get "github.com/jinzhu/gorm"
RUN go get "golang.org/x/crypto/bcrypt"
RUN go get "gopkg.in/gormigrate.v1"
RUN go get "github.com/natefinch/lumberjack"
RUN go get "github.com/stretchr/testify/assert"
RUN go get "github.com/tidwall/gjson"