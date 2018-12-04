# Up

GitHub Action for deploying and operating serverless applications with [Up](https://github.com/apex/up).

## Secrets

- `AWS_ACCESS_KEY_ID` - *Required* The AWS Access Key ID.
- `AWS_SECRET_ACCESS_KEY` - *Required* The AWS Secret Key.

See the [AWS Security Credentials](https://docs.aws.amazon.com/general/latest/gr/aws-security-credentials.html) page for more information.

## Example

```hcl
workflow "Deploy to Up" {
  on = "push"
  resolves = ["Deploy"]
}

action "Deploy" {
  uses = "apex/actions/up@master"
  args = "deploy production"
  secrets = ["AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY"]
}
```
