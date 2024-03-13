FROM gcr.io/distroless/static-debian12

COPY ./dist/server ./app/server
ENTRYPOINT ["/app/server"]
LABEL description="Docker image for school purpose only."