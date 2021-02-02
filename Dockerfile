############################
# STEP 1 build executable binary
############################
FROM golang:1.15-buster as builder

# Create appuser.
RUN useradd -ms /bin/bash appuser
WORKDIR /src

# Build binary
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags="-w -s" -o ./bin/web ./cmd/web

############################
# STEP 2 build a small image
############################
FROM scratch

# Import from builder.
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=builder /etc/passwd /etc/passwd

# Copy our static files and executable.
COPY --from=builder /src/bin/web /web
COPY --from=builder /src/templates /templates
COPY --from=builder /src/public /public

# Use an unprivileged user.
USER appuser

# Run app.
ENTRYPOINT ["/web"]
