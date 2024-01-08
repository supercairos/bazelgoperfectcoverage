workspace(name = "com_shadow_analytics_proxy")

# Load the http ruleset and expose the http_archive rule
load("@bazel_tools//tools/build_defs/repo:http.bzl", "http_archive")

######################################
# Download rulesets.
######################################
http_archive(
    name = "io_bazel_rules_go",
    sha256 = "7c76d6236b28ff695aa28cf35f95de317a9472fd1fb14ac797c9bf684f09b37c",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/rules_go/releases/download/v0.44.2/rules_go-v0.44.2.zip",
        "https://github.com/bazelbuild/rules_go/releases/download/v0.44.2/rules_go-v0.44.2.zip",
    ],
)

http_archive(
    name = "bazel_gazelle",
    integrity = "sha256-MpOL2hbmcABjA1R5Bj2dJMYO2o15/Uc5Vj9Q0zHLMgk=",
    urls = [
        "https://mirror.bazel.build/github.com/bazelbuild/bazel-gazelle/releases/download/v0.35.0/bazel-gazelle-v0.35.0.tar.gz",
        "https://github.com/bazelbuild/bazel-gazelle/releases/download/v0.35.0/bazel-gazelle-v0.35.0.tar.gz",
    ],
)

######################################
# Setup rules_go ruleset.
######################################

# Load rules_go ruleset and expose the toolchain and dep rules.
load("@io_bazel_rules_go//go:deps.bzl", "go_register_toolchains", "go_rules_dependencies")

# The following line defines the symbol go_dependencies from the deps.bzl file.
# Having the deps in that file, helps the WORKSPACE file stay less
# cluttered.  The library symbol go_dependencies is then added to
# the envionment. The line below calls that function.
load("//:deps.bzl", "go_dependencies")

# The next comment line includes a macro that gazelle reads.
# This macro tells Gazelle to look for repository rules in a macro in a .bzl file,
# and allows Gazelle to find the correct file to maintain the Go dependencies.
# Then the line after the comment calls go_dependencies(), and that funcation
# contains calls to various go_repository rules.

# gazelle:repository_macro deps.bzl%go_dependencies
go_dependencies()

# go_rules_dependencies is a function that registers external dependencies
# needed by the Go rules.
# https://github.com/bazelbuild/rules_go/blob/master/go/dependencies.rst#go_rules_dependencies
go_rules_dependencies()

# The next rule installs the Go toolchains. The Go version is specified
# using the version parameter. This rule will download the Go SDK.
# https://github.com/bazelbuild/rules_go/blob/master/go/toolchains.rst#go_register_toolchains
go_register_toolchains(version = "1.21.5")

######################################
# Setup bazel_gazelle ruleset.
######################################

# the line below instructs gazelle to save the go dependency definitions
# in the deps.bzl file. Located under '//'.
load("@bazel_gazelle//:deps.bzl", "gazelle_dependencies")

# The following call configured the gazelle dependencies, Go environment and Go SDK.
gazelle_dependencies()
