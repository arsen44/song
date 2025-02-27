ARG BUILDPATH='/build/song/cmd'

FROM golang:1.23-alpine as builder

ARG BUILDPATH
WORKDIR $BUILDPATH

# Install GCC
RUN apk add --no-cache gcc musl-dev
# Copy code in and fetch dependencies

COPY . ./
RUN go mod download

# Build 
RUN cd ./cmd/ &&\
    CGO_ENABLED=1 GOOS=linux GOARCH=amd64 go build -o backend -tags musl

# Create final image - Export Go Binary onto regular alpine
FROM alpine:latest
ARG BUILDPATH

RUN apk update && \
    apk --no-cache add ca-certificates

COPY --from=builder $BUILDPATH .

EXPOSE 2152

CMD ["./cmd/backend"]