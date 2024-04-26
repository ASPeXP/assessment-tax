FROM golang:1.22.1-alpine as build-base

WORKDIR /app 

COPY .env .
COPY go.mod . 
COPY go.sum .
RUN go mod download

COPY . . 


RUN CGO_ENABLED=0 go test -v 

RUN go build -o ./out/ktax .

# ==============================

FROM alpine:3.19
COPY --from=build-base /app/out/ktax /app/ktax 

EXPOSE 8080 5432

CMD ["/app/ktax"]