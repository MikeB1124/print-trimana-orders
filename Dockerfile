# Use an official Alpine Linux distribution as a parent image
FROM alpine:latest

COPY --from=golang:1.13-alpine /usr/local/go/ /usr/local/go/

RUN apk update && apk add --no-cache vim

ENV PATH="/usr/local/go/bin:${PATH}"
ENV EDITOR=vim
ENV VISUAL=vim
ENV CONTAINER=true

# Copy the current directory contents into the container at /app
COPY . /home/print-trimana-orders

# Set the working directory
WORKDIR /home/print-trimana-orders

RUN GOOS=linux GOARCH=amd64 go build -o /home/print-orders

WORKDIR /home

RUN cp /home/print-trimana-orders/config.yaml /home/config.yaml

RUN rm -rf /home/print-trimana-orders

# Define default command to run when the container starts
CMD ["sh"]
