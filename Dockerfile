FROM golang:1.18-alpine as buildStage

WORKDIR /build
COPY ./ ./
RUN go build -ldflags "-s -w" -o /output ./cmd/cronjob

FROM alpine:latest
WORKDIR /
COPY --from=buildStage /output /output
CMD [ "/output" ]