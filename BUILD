load("@io_bazel_rules_go//go:def.bzl", "go_library", "go_test")
load("@bazel_gazelle//:def.bzl", "gazelle")

# gazelle:prefix github.com/supercairos/perfectcoverage
# gazelle:build_file_name BUILD,BUILD.bazel
gazelle(name = "gazelle")

# adding rule to update deps
gazelle(
    name = "gazelle-update-repos",
    args = [
        "-from_file=go.mod",
        "-to_macro=deps.bzl%go_dependencies",
        "-prune",
    ],
    command = "update-repos",
)

go_library(
    name = "perfectcoverage_lib",
    srcs = ["utils.go"],
    importpath = "github.com/supercairos/perfectcoverage",
    visibility = ["//visibility:private"],
)

go_test(
    name = "perfectcoverage_test",
    srcs = ["utils_test.go"],
    embed = [":perfectcoverage_lib"],
)
