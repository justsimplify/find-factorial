FROM golang:1.15.7-alpine AS build
WORKDIR /src
COPY . .
RUN export CGO_ENABLED=0 && go build -o /out/main .

FROM alpine AS bin
COPY --from=build /out/main .
RUN touch /hello_world.txt
CMD ["./main"]