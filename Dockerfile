FROM alpine:latest as builder
RUN apk update && apk add --no-cache cargo
WORKDIR /app
COPY . .
RUN cargo build --release

FROM alpine:latest
RUN apk update && apk add --no-cache libgcc
COPY --from=builder /app/target/release/joguinho /usr/bin/joguinho
CMD ["joguinho"]