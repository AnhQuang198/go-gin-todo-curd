FROM social-todo-service-cached as builder

ADD . /app/
WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux go buid -a -installsuffix cgo -o demoApp .

FROM alpine
WORKDIR /app/
COPY --from=builder /app/demoApp .
ENTRYPOINT ["/app/demoApp"]