FROM golang:1.24.1

# Выведет текущую версию гошки в логи
RUN go version

# Устанавлиает переменную окружения GOPATH = /
ENV GOPATH=/

COPY ./ ./

# Подгружает все зависимости проекта из go.mod и go.sum
RUN go mod download 

RUN go build -o todo-app ./cmd/main.go

CMD ["./todo-app"]