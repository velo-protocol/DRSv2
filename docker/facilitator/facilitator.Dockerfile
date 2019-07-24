#Docker multi-stage builds

# ------------------------------------------------------------------------------
# Development image
# ------------------------------------------------------------------------------

#Builder stage
FROM golang:1.11-alpine3.7 as builder
# Update OS package and install Git
RUN apk update && apk add git && apk add build-base
# Set working directory
WORKDIR /go/src/gitlab.com/velo-lab/core
# Get Golang dependency management tool
RUN go get github.com/golang/dep/cmd/dep
# Install Fresh for local development
RUN go get github.com/pilu/fresh
# Install go tool for convert go test output to junit xml
RUN go get -u github.com/jstemmer/go-junit-report &&\
    go get github.com/axw/gocov/... &&\
    go get github.com/AlekSi/gocov-xml
# Install wait-for
RUN wget https://raw.githubusercontent.com/eficode/wait-for/master/wait-for -O /usr/local/bin/wait-for &&\
    chmod +x /usr/local/bin/wait-for
# Copy Go dependency file
COPY Gopkg.toml Gopkg.toml
COPY Gopkg.lock Gopkg.lock
# Run dep ensure
RUN dep ensure -vendor-only
# Copy Go source code
COPY . .
# Set Docker's entry point commands
RUN cd facilitator/ && go build

# ------------------------------------------------------------------------------
# Deployment image
# ------------------------------------------------------------------------------

#App stage
#FROM golang:1.11-alpine3.7
## Set working directory
#WORKDIR /go/src/gitlab.com/velo-lab/core/facilitator
##Get artifact from buiber stage
#RUN mkdir -p app/migrations app/recovery_migrations
#COPY --from=builder /go/src/gitlab.com/velo-lab/core/facilitator/migrations/ facilitator/migrations/
#COPY --from=builder /go/src/gitlab.com/velo-lab/core/facilitator/recovery_migrations/ app/recovery_migrations/
#COPY --from=builder /usr/local/bin/wait-for /usr/local/bin/wait-for
#COPY --from=builder /go/src/gitlab.com/velo-lab/core/facilitator facilitator/
#COPY --from=builder /go/src/gitlab.com/velo-lab/core/facilitator/vendor vendor
## Set Docker's entry point commands
#CMD ./facilitator;