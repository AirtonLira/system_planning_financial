FROM golang:latest

WORKDIR /app

COPY go.mod ./

RUN go mod download 

COPY . .

RUN CGO_ENABLED=0 go build -v -a -installsuffix cgo -o system_planning_financial ./cmd/server

EXPOSE 8090

CMD [ "./system_planning_financial" ]