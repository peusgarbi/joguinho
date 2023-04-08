FROM alpine:latest as builder
RUN apk update && apk add --no-cache cargo sqlite-dev
WORKDIR /app
COPY . .
RUN cargo build --release

FROM alpine:latest
RUN apk update && apk add --no-cache libgcc sqlite-dev
ENV IMAGES_DIR=/app/src/images
ENV MIGRATIONS_DIR=/app/migrations
ENV DATABASE_URL=/app/src/database/joguinho.db
COPY ./migrations /app/migrations
COPY ./src/images /app/src/images
RUN mkdir /app/src/database
COPY --from=builder /app/target/release/joguinho /usr/bin/joguinho
WORKDIR /app
CMD ["joguinho"]
