package(default_visibility = ["PUBLIC"])

go_get(
    name = "logging",
    get = "gopkg.in/op/go-logging.v1",
    revision = "b2cb9fa56473e98db8caba80237377e83fe44db5",
)

go_get(
    name = "terminal",
    get = "golang.org/x/crypto/ssh/terminal",
    revision = "7b85b097bf7527677d54d3220065e966a0e3b613",
)

go_get(
    name = "go-bindata",
    binary = True,
    get = "github.com/kevinburke/go-bindata/...",
    revision = "46eb4c183bfc1ebb527d9d19bcded39476302eb8",
    strip = ["testdata"],
)

go_get(
    name = "go-flags",
    get = "github.com/thought-machine/go-flags",
    revision = "v1.5.0",
    deps = [
        ":unix",
    ],
)

go_get(
    name = "net",
    get = "golang.org/x/net/...",
    revision = "136a25c244d3019482a795d728110278d6ba09a4",
    deps = [
        ":terminal",
        ":text",
    ],
)

go_get(
    name = "tools",
    get = "golang.org/x/tools",
    install = [
        "cover",
    ],
    revision = "2ae76fd1560b622911f444c1e66b70a857e1de67",
)

go_get(
    name = "text",
    get = "golang.org/x/text/...",
    revision = "4e4a3210bb54bb31f6ab2cdca2edcc0b50c420c1",
    deps = [
        ":tools",
    ],
)

go_get(
    name = "grpc",
    exported_deps = [":net"],
    get = "google.golang.org/grpc",
    install = [
        "",
        "attributes",
        "backoff",
        "balancer",
        "balancer/base",
        "balancer/grpclb/...",
        "balancer/roundrobin",
        "binarylog/...",
        "codes",
        "connectivity",
        "credentials/...",
        "encoding",
        "encoding/gzip",
        "encoding/proto",
        "grpclog",
        "health",
        "health/grpc_health_v1",
        "internal/...",
        "keepalive",
        "metadata",
        "peer",
        "reflection/...",
        "resolver",
        "resolver/dns",
        "resolver/passthrough",
        "serviceconfig",
        "stats",
        "status",
        "tap",
    ],
    repo = "github.com/grpc/grpc-go",
    revision = "v1.31.1",
    deps = [
        ":cmp",
        ":oauth2",
        ":protobuf",
        ":rpccode",
        ":rpcstatus",
        ":unix",
    ],
)

#TODO: build from the actual proto.
go_get(
    name = "rpcstatus",
    get = "google.golang.org/genproto/googleapis/rpc/status",
    revision = "a023cd5227bd25fd1f5c5633743ff3eacc93d169",
    deps = [":protobuf"],
)

go_get(
    name = "rpccode",
    get = "google.golang.org/genproto/googleapis/rpc/code",
    revision = "2b5a72b8730b0b16380010cfe5286c42108d88e7",
    deps = [":protobuf"],
)

go_get(
    name = "protoc-gen-go",
    binary = True,
    get = [],
    install = ["github.com/golang/protobuf/protoc-gen-go"],
    deps = [
        ":protobuf",
    ],
)

go_get(
    name = "protobuf",
    exported_deps = [
        ":golang-protobuf",
    ],
    get = "github.com/golang/protobuf/...",
    revision = "v1.4.2",
    strip = [
        "proto/proto3_proto",
        "conformance",
    ],
    deps = [
        ":golang-protobuf",
    ],
)

go_get(
    name = "spew",
    get = "github.com/davecgh/go-spew/spew",
    revision = "d8f796af33cc11cb798c1aaeb27a4ebc5099927d",
)

go_get(
    name = "testify",
    get = "github.com/stretchr/testify",
    install = [
        "assert",
        "require",
        "vendor/github.com/davecgh/go-spew/spew",
        "vendor/github.com/pmezard/go-difflib/difflib",
    ],
    revision = "f390dcf405f7b83c997eac1b06768bb9f44dec18",
    deps = [":spew"],
)

go_get(
    name = "grpc-middleware",
    get = "github.com/grpc-ecosystem/go-grpc-middleware",
    install = [
        "",
        "recovery",
        "retry",
        "util/metautils",
        "util/backoffutils",
    ],
    revision = "fa8fef87dcecac0bda02d36abb3c790ab9e0030b",
    deps = [
        ":grpc",
        ":net",
        ":protobuf",
    ],
)

go_get(
    name = "errgroup",
    get = "golang.org/x/sync/errgroup",
    revision = "09787c993a3ab68e3d1f5c9b2394ab9433f391be",
    deps = [":net"],
)

go_get(
    name = "semaphore",
    get = "golang.org/x/sync/semaphore",
    revision = "09787c993a3ab68e3d1f5c9b2394ab9433f391be",
    deps = [":net"],
)

go_get(
    name = "unix",
    get = "golang.org/x/sys",
    install = [
        "internal/unsafeheader",
        "unix",
    ],
    revision = "ddb9806d33aed8dbaac1cd6f1cba58952e87f933",
)

go_get(
    name = "errors",
    get = "github.com/pkg/errors",
    revision = "30136e27e2ac8d167177e8a583aa4c3fea5be833",
)

go_get(
    name = "skylark",
    get = "github.com/google/skylark/syntax",
    revision = "a5f7082aabed29c0e429c722292c66ec8ecf9591",
)

go_get(
    name = "buildtools",
    get = "github.com/bazelbuild/buildtools/...",
    revision = "0.19.2",
    deps = [
        ":protobuf",
        ":skylark",
    ],
)

go_get(
    name = "buildifier",
    binary = True,
    get = [],
    install = ["github.com/bazelbuild/buildtools/buildifier"],
    deps = [
        ":buildtools",
    ],
)

go_get(
    name = "cli-init",
    get = "github.com/peterebden/go-cli-init/v3",
    module_major_version = "v3",
    revision = "v3.1.0",
    deps = [
        ":go-flags",
        ":humanize",
        ":logging",
        ":terminal",
    ],
)

go_get(
    name = "remote-apis",
    get = "github.com/bazelbuild/remote-apis/build/...",
    repo = "github.com/peterebden/remote-apis",
    revision = "5306a2d66a1b879350f1a474cc639d569953fb48",
    deps = [
        ":annotations",
        ":grpc",
        ":longrunning",
        ":protobuf",
        ":rpcstatus",
    ],
)

go_get(
    name = "remote-apis-sdks",
    get = "github.com/bazelbuild/remote-apis-sdks/go/...",
    repo = "github.com/peterebden/remote-apis-sdks",
    revision = "5e7b88ce7342f4b3a53c29b6e6962f29e0522d12",
    deps = [
        ":annotations",
        ":bytestream",
        ":cmp",
        ":errgroup",
        ":errors",
        ":glog",
        ":golang-protobuf",
        ":grpc",
        ":longrunning",
        ":oauth2",
        ":pborman_uuid",
        ":protobuf",
        ":remote-apis",
        ":rpcstatus",
        ":semaphore",
        ":zstd",
    ],
)

go_get(
    name = "golang-protobuf",
    get = "google.golang.org/protobuf/...",
    revision = "v1.25.0",
    deps = [
        ":cmp",
    ],
)

go_get(
    name = "cmp",
    get = "github.com/google/go-cmp/cmp/...",
    revision = "b1c9c4891a6525d98001fea424c8926c6d77bb56",
)

go_get(
    name = "glog",
    get = "github.com/golang/glog",
    patch = "glog_disable.patch",
    revision = "23def4e6c14b4da8ac2ed8007337bc5eb5007998",
)

go_get(
    name = "metadata",
    get = "cloud.google.com/go/compute/metadata",
    revision = "v0.46.3",
)

go_get(
    name = "oauth2",
    get = "golang.org/x/oauth2",
    install = [
        "",
        "google",
        "internal",
        "jws",
        "jwt",
    ],
    revision = "bf48bf16ab8d622ce64ec6ce98d2c98f916b6303",
    visibility = ["//third_party/go/..."],
    deps = [
        ":metadata",
        "//third_party/go:net",
    ],
)

# TODO(peterebden): Splitting up this & the other bits of genproto is a bit of
#                   a hack. Replace once we have a better way of describing Go modules.
go_get(
    name = "annotations",
    get = "google.golang.org/genproto/googleapis/api/annotations",
    revision = "a023cd5227bd25fd1f5c5633743ff3eacc93d169",
    deps = [
        ":grpc",
        ":protobuf",
    ],
)

go_get(
    name = "longrunning",
    get = "google.golang.org/genproto/googleapis/longrunning",
    revision = "a023cd5227bd25fd1f5c5633743ff3eacc93d169",
    deps = [
        ":annotations",
        ":grpc",
        ":protobuf",
    ],
)

go_get(
    name = "bytestream",
    get = "google.golang.org/genproto/googleapis/bytestream",
    revision = "a023cd5227bd25fd1f5c5633743ff3eacc93d169",
    deps = [
        ":annotations",
        ":grpc",
        ":protobuf",
    ],
)

go_get(
    name = "uuid",
    get = "github.com/google/uuid",
    revision = "v1.1.1",
)

go_get(
    name = "pborman_uuid",
    get = "github.com/pborman/uuid",
    revision = "v1.2.0",
    deps = [":uuid"],
)

go_get(
    name = "ristretto",
    get = "github.com/dgraph-io/ristretto/...",
    revision = "v0.0.1",
    deps = [":xxhash"],
)

go_get(
    name = "xxhash",
    get = "github.com/cespare/xxhash",
    revision = "v2.1.1",
)

go_get(
    name = "humanize",
    get = "github.com/dustin/go-humanize",
    revision = "v1.0.0",
)

go_get(
    name = "retryablehttp",
    get = "github.com/hashicorp/go-retryablehttp",
    revision = "v0.6.1",
    deps = [
        ":cleanhttp",
        ":hclog",
    ],
)

go_get(
    name = "sri",
    get = "github.com/peterebden/go-sri",
    revision = "v1.1.1",
)

go_get(
    name = "go-multierror",
    get = "github.com/hashicorp/go-multierror",
    revision = "v1.1.0",
    deps = [":errwrap"],
)

go_get(
    name = "errwrap",
    get = "github.com/hashicorp/errwrap",
    revision = "v1.0.0",
)

go_get(
    name = "cleanhttp",
    get = "github.com/hashicorp/go-cleanhttp",
    revision = "v0.5.1",
)

go_get(
    name = "hclog",
    get = "github.com/hashicorp/go-hclog",
    revision = "v0.12.0",
    deps = [
        ":color",
        ":isatty",
    ],
)

go_get(
    name = "isatty",
    get = "github.com/mattn/go-isatty",
    revision = "v0.0.12",
    deps = [":unix"],
)

go_get(
    name = "color",
    get = "github.com/fatih/color",
    install = [
        "",
        "vendor/github.com/mattn/go-isatty",
        "vendor/github.com/mattn/go-colorable",
        "vendor/golang.org/x/sys/unix",
    ],
    revision = "v1.9.0",
)

go_get(
    name = "copyfile",
    get = "github.com/peterebden/go-copyfile",
    revision = "bc0baf74909cb55565cf9536cb003f04da55372d",
)

go_get(
    name = "godirwalk",
    get = "github.com/karrick/godirwalk",
    revision = "v1.7.0",
    strip = [
        "examples",
        "testdata",
        "vendor",
    ],
    deps = [":errors"],
)

go_get(
    name = "http-admin",
    get = "github.com/thought-machine/http-admin",
    revision = "v1.1.0",
    deps = [
        ":logging",
        ":mux",
        ":net",
        "//third_party/go/prometheus",
    ],
)

go_get(
    name = "mux",
    get = "github.com/gorilla/mux",
    revision = "v1.7.4",
)

go_get(
    name = "xattr",
    get = "github.com/pkg/xattr",
    revision = "v0.4.1",
    deps = [":unix"],
)

go_get(
    name = "psutil",
    get = "github.com/shirou/gopsutil",
    install = [
        "internal/common",
        "mem",
    ],
    revision = "v2.17.09",
    deps = [":unix"],
)

go_get(
    name = "zstd",
    get = "github.com/klauspost/compress",
    install = [
        "fse",
        "huff0",
        "snappy",
        "zstd/...",
    ],
    revision = "v1.11.0",
)

go_get(
    name = "go-grpc-compression",
    get = "github.com/mostynb/go-grpc-compression/zstd",
    revision = "v1.1.2",
    deps = [
        ":grpc",
        ":zstd",
    ],
)
