FROM golang:1.18-alpine as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download && go mod verify
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -installsuffix cgo -o /cmd/financing

FROM ubuntu:latest
RUN apt-get update && apt-get install -y curl
COPY --from=build /app/.env /
COPY --from=build /cmd/financing /
EXPOSE 8082
CMD ["/financing"]