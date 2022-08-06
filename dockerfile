FROM golang:1.18.5 AS builder
WORKDIR /app
ENV GOPROXY=https://goproxy.cn,direct
ENV PATH="/opt/gtk/bin:$env/development.env"
COPY . .

# RUN go mod tidy
# RUN go build -o main .

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

CMD [ "main" ]
# ENTRYPOINT [ "/app/start.sh", "ls"]

