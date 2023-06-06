FROM golang:1.20
WORKDIR /app
COPY . .

RUN go mod download
RUN cd cmd/url-shortener/;go build -o app
EXPOSE 80
CMD ["/app/cmd/url-shortener/./app"]