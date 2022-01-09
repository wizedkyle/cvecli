FROM scratch
USER nonroot
ENTRYPOINT ["/cvecli"]
COPY cvecli /
