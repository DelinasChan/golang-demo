FROM golang:1.13.4-alpine

RUN apk update && apk upgrade && \
    apk add git

# codegangsta/gin 自動編譯套件
RUN go get github.com/codegangsta/gin

# Web App Framework
RUN go get github.com/gin-gonic/gin

#建立依賴項目
# RUN go mod init app/demo

#Docker 啟動時 預設目錄
WORKDIR /var/www/html
# CMD ["gin","-p","7070","-a","7000","run","main.go"]
