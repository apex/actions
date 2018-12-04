# Go

GitHub Action for using the `go` binary. `GO111MODULE` is enabled and encouraged for dependency management.

## Example

Deploy an application to `production` after building the Go binary. By default `./server` is built from `*.go` files, so if you Go source is in root this should work great!

```
workflow "Deployment" {
  on = "push"
  resolves = ["Deploy"]
}

action "Build" {
  uses = "apex/actions/go@master"
}

action "Deploy" {
  needs = "Build"
  uses = "apex/actions/up@master"
  secrets = ["AWS_SECRET_ACCESS_KEY", "AWS_ACCESS_KEY_ID"]
  args = "deploy production"
}
```

Alternatively you can specify the build command:

```
workflow "Deployment" {
  on = "push"
  resolves = ["Deploy"]
}

action "Build" {
  uses = "apex/actions/go@master"
  args = "build -o server cmd/api/main.go"
}

action "Deploy" {
  needs = "Build"
  uses = "apex/actions/up@master"
  secrets = ["AWS_SECRET_ACCESS_KEY", "AWS_ACCESS_KEY_ID"]
  args = "deploy production"
}
```
