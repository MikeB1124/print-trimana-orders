# Use an official Alpine Linux distribution as a parent image
FROM alpine:latest

# Install golang
COPY --from=golang:1.13-alpine /usr/local/go/ /usr/local/go/

# Install all needed services for the container
RUN apk update && apk add --no-cache vim \
    openrc

# Set ENV variables for container
ENV PATH="/usr/local/go/bin:${PATH}"
ENV EDITOR=vim
ENV VISUAL=vim
ENV CONTAINER=true
ENV LOG_PATH=/home/app.log

# Copy the current directory contents into the container at /home/print-trimana-orders
COPY . /home/print-trimana-orders

# Set the working directory to root of go project
WORKDIR /home/print-trimana-orders

# Build go binary and create it in the /home directory
RUN go build -o /home/print-orders

# Set working directory to /home
WORKDIR /home

# Copy the config.yaml to the /home directory
RUN cp /home/print-trimana-orders/config.yaml /home/config.yaml

# Remove the print-trimana-orders directory
RUN rm -rf /home/print-trimana-orders

# Add cron job to the crontab to execute the go binary
RUN crontab -l | { cat; echo "* * * * * /home/print-orders >> /home/cron.log 2>&1"; } | crontab -

# Run crond -f -d 8 command when contianer starts up
CMD ["crond", "-f", "-d", "8"]
