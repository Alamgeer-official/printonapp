FROM golang:1.21.5-alpine3.19
WORKDIR /printonapp/backend
COPY . /printonapp/backend
RUN go build /printonapp/backend
EXPOSE 4000
ENTRYPOINT [ "./printonapp" ]