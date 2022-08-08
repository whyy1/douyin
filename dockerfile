FROM golang:1.18 AS builder
WORKDIR /app
ENV GOPROXY=https://goproxy.cn,direct
COPY . .

RUN go version
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64
RUN go build -o main main.go

FROM alpine:3.13
USER root
WORKDIR /app
COPY --from=builder /app/main .
COPY --from=builder /app/app.env .
COPY --from=builder /app/start.sh ./start.sh
COPY --from=builder /app/wait-for.sh ./wait-for.sh

RUN chmod 777 ./start.sh
RUN chmod 777 ./wait-for.sh
RUN chmod 777 ./main

EXPOSE 8080

CMD [ "./main" ]
ENTRYPOINT [ "/app/start.sh"]

