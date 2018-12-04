# Up

GitHub Action for deploying and operating serverless applications with [Up](https://github.com/apex/up).

## Secrets

- `AWS_ACCESS_KEY_ID` - *Required* The AWS Access Key ID.
- `AWS_SECRET_ACCESS_KEY` - *Required* The AWS Secret Key.

See the [AWS Security Credentials](https://docs.aws.amazon.com/general/latest/gr/aws-security-credentials.html) page for more information.

## Example

Deploy an application in the root directory to `production`:

```hcl
workflow "Deploy Application" {
  on = "push"
  resolves = ["Deploy"]
}

action "Deploy" {
  uses = "apex/actions/up@master"
  args = "deploy production"
  secrets = ["AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY"]
}
```

Deploy an application within a sub-directory to `staging`:

```hcl
workflow "Deploy Application" {
  on = "push"
  resolves = ["Deploy Staging"]
}

action "Deploy Staging" {
  uses = "apex/actions/up@master"
  args = "-C cmd/team-api deploy staging"
  secrets = ["AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY"]
}
```

Deploy an application to `production` after installing NPM dependencies:

```
workflow "Deploy Application" {
  on = "push"
  resolves = ["Deploy"]
}

action "Build" {
  uses = "actions/npm@master"
  args = "install"
}

action "Deploy" {
  needs = "Build"
  uses = "apex/actions/up@master"
  secrets = ["AWS_SECRET_ACCESS_KEY", "AWS_ACCESS_KEY_ID"]
  args = "deploy production"
}
```
