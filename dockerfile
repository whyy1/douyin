# 最低版本go的依赖
FROM golang:1.17.7 

# 配置模块代理
ENV GOPROXY=https://goproxy.cn,direct

# 复制文件目录下的所有代码 
ADD . /whygo

#默认进入文件目录
WORKDIR /whygo

# 复制代码
ADD /whygo/go.mod .
ADD /whygo/go.sum .

# 下载依赖
RUN go mod download

# 运行命令，安装依赖
# 例如 RUN npm install && cd /app && mkdir logs
ADD ./whygo .
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

# RUN 命令可以有多个，但是可以用 && 连接多个命令来减少层级。
# 构建exe文件，名字为demo
RUN go build  -o  main 

# CMD 指令只能一个，是容器启动后执行的命令，算是程序的入口。
# 默认运行 demo
ENTRYPOINT [ "./main" ]