# Accept the Go version for the image to be set as a build argument.
# Default to Go 1.11.5
ARG GO_VERSION=1.11.5

# First stage: build the executable.
FROM golang:${GO_VERSION}-alpine AS builder

# Create the user and group files that will be used in the running container to
# run the process as an unprivileged user.
RUN mkdir /user && \
    echo 'nobody:x:65534:65534:nobody:/:' > /user/passwd && \
    echo 'nobody:x:65534:' > /user/group

# Set the working directory outside $GOPATH to enable the support for modules.
RUN mkdir /app 
ADD . /app/ 
WORKDIR /app 

# Build the executable to `/app`. Mark the build as statically linked.
RUN CGO_ENABLED=0 go build \
    -installsuffix 'static' \
    -o webserver .

# Final stage: the running container.
FROM scratch AS final

# Import the user and group files from the first stage.
COPY --from=builder /user/group /user/passwd /etc/

# Import the compiled executable from the first stage.
COPY --from=builder /app/webserver /app/webserver

# Declare the port on which the webserver will be exposed.
# As we're going to run the executable as an unprivileged user, we can't bind
# to ports below 1024.
EXPOSE 8080

# Perform any further action as an unprivileged user.
USER nobody:nobody

# Run the compiled binary.
ENTRYPOINT ["/app/webserver"]