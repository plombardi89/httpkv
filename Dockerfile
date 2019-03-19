FROM golang:1.12 as foundation

WORKDIR /build
COPY go.mod .
COPY go.sum .
RUN go mod download

FROM foundation as builder

COPY . .
RUN make

FROM gcr.io/distroless/base as runtime

COPY --from=builder /build/bin/httpkv /bin/httpkv

ENTRYPOINT ["/bin/httpkv"]
