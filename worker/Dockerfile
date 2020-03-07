FROM golang:alpine as builder
ENV CGO_ENABLED=0

WORKDIR '/app'

COPY . .
RUN go mod tidy
RUN mkdir bin
RUN go build -o ./bin/prod_app
COPY ./view ./bin/view


FROM alpine
WORKDIR '/app'
COPY --from=builder /app/bin .
EXPOSE 80
CMD [ "./prod_app", "80" ]