FROM golang:1.22

WORKDIR /src/backend

COPY go.mod go.sum ./
RUN go mod download && go mod verify
RUN go get github.com/go-chi/chi/v5
RUN go get github.com/go-chi/cors
RUN go get github.com/go-chi/render

COPY . .

RUN go build -o main .

CMD ["./main"]