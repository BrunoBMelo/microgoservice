FROM golang:1.18-alpine as build
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod tidy && go mod verify
COPY . .
RUN CGO_ENABLED=0 go build -ldflags "-s -w" -installsuffix cgo -o /cmd/financing

FROM gcr.io/distroless/static
COPY --from=build /cmd/financing /
EXPOSE 8081
CMD ["/consortium"]