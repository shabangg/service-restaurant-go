FROM golang:alpine AS builder
LABEL maintainer="rohanluthra <rohanluthra.1123@gmail.com>"

# define system module and update them
RUN apk update \
  && apk upgrade \
  && apk add --no-cache \
  git \
  dep \
  ca-certificates \
  && update-ca-certificates

# Create user for accessing job service.
RUN adduser -D -g '' restaurantuser
WORKDIR /go/src/github.com/rohan-luthra/microservice-grpc-go/service-restaurants-go

# install dependencies
COPY Gopkg.* ./
RUN dep ensure -vendor-only -v

# copy source and generate a build
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main main.go

# use scratch (base for a docker image)
FROM scratch
COPY --from=builder /etc/passwd /etc/passwd

# set working directory
WORKDIR /root/

# copy necessary config files
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /go/src/github.com/rohan-luthra/microservice-grpc-go/service-restaurants-go/config.yml .
COPY --from=builder /go/src/github.com/rohan-luthra/microservice-grpc-go/service-restaurants-go/main .

# Use an unprivileged user.
USER restaurantuser

# expose port and start the app
EXPOSE 3000
CMD ["./main", "serve"]