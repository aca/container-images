FROM public.ecr.aws/docker/library/alpine:edge
RUN apk add --no-cache \
		ca-certificates \
		tzdata
RUN apk add --no-cache --repository=http://dl-cdn.alpinelinux.org/alpine/edge/testing/ \
        watchexec
