FROM golang:1.15

WORKDIR /app
COPY . /app
RUN cd /app && go build -o main
ENTRYPOINT [ "./main" ]