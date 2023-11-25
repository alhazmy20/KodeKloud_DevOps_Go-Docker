FROM golang:1.21-alpine3.18

WORKDIR /usr/src/app

ARG ENV_DB_CONTAINER_NAME
ARG ENV_NAME=develpoment

ENV DB_CONNECTION=mysql
ENV DB_DATABASE=${ENV_DATABASE}
ENV DB_USERNAME=root
ENV DB_PASSWORD=''
ENV DB_CONTAINER_NAME=${ENV_DB_CONTAINER_NAME}
ENV APP_ENV=${ENV_NAME}

COPY go.mod go.sum ./
RUN go mod download 
RUN go mod verify

COPY . .
RUN go build -v -o /usr/local/bin/app .

CMD [ "app" ]
