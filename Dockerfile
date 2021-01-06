############################
# STEP 1 build executable binary
############################
FROM golang:1.15-buster as builder

# Install git.
# Git is required for fetching the dependencies.
#RUN apk update && apk add --no-cache git ca-certificates build-base && update-ca-certificates

# Create appuser.
RUN useradd -ms /bin/bash appuser
WORKDIR /home/appuser

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
COPY --from=builder /home/appuser/bin/web /web
COPY --from=builder /home/appuser/templates /templates
COPY --from=builder /home/appuser/public /public

# Use an unprivileged user.
USER appuser

# Heroku port assignment
CMD gunicorn --bind 0.0.0.0:$PORT wsgi

# Run app.
ENTRYPOINT ["/web"]
