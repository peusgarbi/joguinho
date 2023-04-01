FROM alpine:latest as builder
RUN apk update && apk add --no-cache cargo
COPY . /usr/src/app
WORKDIR /usr/src/app
RUN cargo build --release

FROM alpine:latest as runtime
RUN apk update && apk add libgcc
COPY --from=builder /usr/src/app/target/release/joguinho /usr/bin/joguinho
CMD ["joguinho"]
