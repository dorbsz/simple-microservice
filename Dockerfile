FROM golang:alpine as simple-microservice-builder
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 
RUN go build -o main . 

#creates a new image with just the executable
FROM alpine
WORKDIR /app
COPY --from=simple-microservice-builder /app/main /app
CMD ["/app/main"]
