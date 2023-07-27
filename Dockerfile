FROM gcr.io/distroless/static:nonroot

# `nonroot` coming from distroless
USER 65532:65532

# Copy the binary that goreleaser built
COPY  resource-provider-api /resource-provider-api

# Run the web service on container startup.
ENTRYPOINT ["/resource-provider-api"]
CMD ["serve"]