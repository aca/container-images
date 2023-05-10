FROM public.ecr.aws/docker/library/golang:1.19-alpine

COPY go.mod .
COPY go.sum .

RUN go mod download

