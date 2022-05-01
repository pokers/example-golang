FROM golang:alpine
RUN mkdir /webApp
COPY . /webApp
WORKDIR /webApp
RUN go build -o launcher .
CMD ["/webApp/launcher"]