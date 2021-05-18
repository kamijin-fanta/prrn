FROM golang:1.16 as builder
WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . ./
RUN CGO_ENABLED=0 go build -o prrn .


FROM scratch

COPY --from=builder /app/prrn /prrn
CMD ["/prrn"]
