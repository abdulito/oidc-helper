http_archive(
  name = "io_bazel_rules_go",
  url = "https://github.com/bazelbuild/rules_go/releases/download/0.10.1/rules_go-0.10.1.tar.gz",
  sha256 = "4b14d8dd31c6dbaf3ff871adcd03f28c3274e42abc855cb8fb4d01233c0154dc",
)

http_archive(
  name = "bazel_gazelle",
  url = "https://github.com/bazelbuild/bazel-gazelle/releases/download/0.10.1/bazel-gazelle-0.10.1.tar.gz",
  sha256 = "d03625db67e9fb0905bbd206fa97e32ae9da894fe234a493e7517fd25faec914",
)


load("@io_bazel_rules_go//go:def.bzl", "go_rules_dependencies", "go_register_toolchains", "go_repository")
go_rules_dependencies()
go_register_toolchains()
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")
gazelle_dependencies()

go_repository(
  name = "org_golang_x_oauth2",
  importpath = "golang.org/x/oauth2",
  commit = "c57b0facaced709681d9f90397429b9430a74754",
)

go_repository(
  name = "com_github_pquerna_cachecontrol",
  importpath = "github.com/pquerna/cachecontrol",
  commit = "1555304b9b35fdd2b425bccf1a5613677705e7d0",
)

go_repository(
  name = "com_github_pquerna_cachecontrol",
  importpath = "github.com/pquerna/cachecontrol",
  commit = "1555304b9b35fdd2b425bccf1a5613677705e7d0",
)

go_repository(
    name = "in_gopkg_square_go_jose_v2",
    commit = "b25e6cab129e4a54675b42ea49d38e9c33ade9e6",  # Jun 7, 2017 (v2.1.2)
    importpath = "gopkg.in/square/go-jose.v2",
)

go_repository(
    name = "com_github_coreos_gooidc",
    importpath = "github.com/coreos/go-oidc",
    tag = "v2.0.0",
)