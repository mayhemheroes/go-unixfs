FROM golang:1.18 as builder

COPY . /go-unixfs
WORKDIR /go-unixfs/mayhem

RUN go install github.com/dvyukov/go-fuzz/go-fuzz@latest github.com/dvyukov/go-fuzz/go-fuzz-build@latest
RUN go get github.com/dvyukov/go-fuzz/go-fuzz-dep
RUN apt update && apt install -y clang

RUN go-fuzz-build -libfuzzer -o fuzzUnwrapData.a && \
    clang -fsanitize=fuzzer fuzzUnwrapData.a -o fuzzUnwrapData.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /go-unixfs/mayhem/fuzzUnwrapData.libfuzzer /