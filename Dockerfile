FROM golang:alpine

WORKDIR /home/app

# Install required packages
RUN apk add --no-cache \
    git \
    make \
    vim

# Keep container running using JSON format
CMD ["tail", "-f", "/dev/null"]
