FROM golang:1.14.10-alpine3.12 as build

WORKDIR /app
COPY . .
RUN go build -o shhtpm main.go 


FROM alpine:3.12.0 as final

COPY --from=build /app/shhtpm .
EXPOSE 8080
CMD [ "./shhtpm" ]
