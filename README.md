# golang学习

1. 运行某一个go文件的命令

   ```shell
   go run mux.go
   ```

2. 运行某一个测试方法

   ```shell
   go test -v -run 'TestSimpleGet'
   ```

3. 编写Dockerfile文件

   1. 创建名字为`Dockerfile`的文件在项目中

   2. 编写`Dockerfile`文件

      ```dockerfile
      FROM golang:1.12.5-alpine3.9
      
      RUN mkdir -p /go/src/app
      WORKDIR /go/src/app
      COPY . .
      
      RUN apk update && apk upgrade && apk add --no-cache bash git openssh
      RUN go get -v -u github.com/gorilla/mux
      RUN go build .
      
      EXPOSE 8000
      ENTRYPOINT ["./app"]
      ```

   3. 生成镜像

      ```shell
      docker build -t mux .
      ```

   4. 运行

      ```shell
      docker run -d -p 8000:8000 --name mux --rm mux
      ```

4. 开发规范

   * 传递结构体时，都传递指针过去

5. docker下载mysql镜像并启动

   ```shell
   docker run --name mysql -p 3306:3306  -v /Users/...:/var/lib/mysql -e MYSQL_ROOT_PASSWORD="root" -d mysql
   ```

   进入容器创建一个新的数据库

   ```shell
   docker exec -it mysql sh # 进入mysql容器
   mysql -u root -p'root' # 进入mysql shell
   
   CREATE DATABASE sqlx; # 创建名为sqlx的数据库
   quit # 退出mysql命令行
   ```

6. 下载sql开发相关的包

   ```shell
   go get github.com/jmoiron/sqlx
   get github.com/go-sql-driver/mysql
   ```

   



#### 参考资料

* [sqlx参考资料](https://github.com/jmoiron/sqlx)