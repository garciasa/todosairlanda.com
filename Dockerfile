FROM golang:1.11-alpine3.8

WORKDIR /todosairlandaweb

#   COPY dist/ dist/
COPY app.go .

RUN ["go", "build", "app.go"]
CMD ["./app"]
