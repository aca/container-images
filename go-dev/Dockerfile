FROM public.ecr.aws/docker/library/golang:1.19-alpine

WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

ENTRYPOINT ["go", "run", "."]
