FROM golang:1.22.2 AS build-stage

WORKDIR /app

# Fetch
COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o /goapp

FROM gcr.io/distroless/base-debian11 AS build-release-stage 

WORKDIR /

# Executable
COPY --from=build-stage /goapp /goapp
COPY .env .env

EXPOSE 8080

USER nonroot:nonroot

CMD [ "./goapp" ]
