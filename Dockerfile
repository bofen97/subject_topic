# syntax=docker/dockerfile:1


FROM golang:1.21 AS BUILD

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY *.go ./

RUN CGO_ENABLED=0 GOOS=linux go build -o /subject_server



FROM scratch
COPY --from=BUILD /subject_server /subject_server

EXPOSE 8081
CMD [ "/subject_server" ]