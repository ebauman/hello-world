FROM golang:1.20 AS build

COPY . /app

WORKDIR /app

RUN go build -o hello-world

FROM golang:1.20 AS run

COPY --from=build /app/hello-world /hello-world
COPY --from=build /app/html.tmpl /html.tmpl

WORKDIR /

EXPOSE 80

CMD /hello-world