# FROM fuzzers/go-fuzz:1.2.0
FROM golang:1.18 as builder

COPY . /go-unixfs
WORKDIR /go-unixfs/mayhem

RUN go-fuzz-build -libfuzzer -o fuzz.a && \
    clang -fsanitize=fuzzer fuzz.a -o fuzz.libfuzzer

FROM debian:bookworm-slim
COPY --from=builder /go-unixfs/mayhem/fuzz.libfuzzer /