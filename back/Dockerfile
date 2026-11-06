FROM golang:1.27.0
RUN go install github.com/air-verse/air@latest
COPY . /app
WORKDIR /app
ENTRYPOINT ["air"]