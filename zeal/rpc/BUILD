go_library(
    name = "rpc",
    srcs = ["rpc.go"],
    deps = [
        "//grpcutil",
        "//rexclient",
        "//third_party/go:cli-init",
        "//third_party/go:grpc",
        "//third_party/go/prometheus",
        "//third_party/go:go-multierror",
        "//third_party/go:protobuf",
        "//third_party/go:remote-apis",
        "//third_party/go:remote-apis-sdks",
        "//third_party/go:retryablehttp",
        "//third_party/go:rpcstatus",
        "//third_party/go:sri",
    ],
    visibility = ["//zeal/..."],
)
