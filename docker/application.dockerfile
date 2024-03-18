### BASE GO IMAGE
FROM golang:1.22-alpine as builder

# creates app folder
RUN mkdir /app

# copies everything from current folder into app folder
COPY ../ /app

# sets app folder as working directory for nest instructions
WORKDIR /app

# builds the application service executable into the previously set working directory (app folder)
RUN CGO_ENABLED=0 go build -o application .

# adds the executable rights to application executable
RUN chmod +x /app/application



### BUILD A TINY DOCKER IMAGE
FROM alpine:latest

# creates app folder
RUN mkdir /app

# copies the previously created builder image from /app/application of the previous image into /app of the current image 
COPY --from=builder /app/application /app

# runs the executable
CMD ["/app/application"]