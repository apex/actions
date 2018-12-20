# SAM

GitHub Action for packaging and deploying [AWS SAM](https://github.com/awslabs/serverless-application-model) applications.

## Secrets

- `AWS_ACCESS_KEY_ID` - *Required* The AWS Access Key ID.
- `AWS_SECRET_ACCESS_KEY` - *Required* The AWS Secret Key.

See the [AWS Security Credentials](https://docs.aws.amazon.com/general/latest/gr/aws-security-credentials.html) page for more information.

## Environment Variables

- `AWS_DEFAULT_REGION`- **Optional** The AWS region name, defaults to `us-east-1` ([more info](https://docs.aws.amazon.com/general/latest/gr/rande.html))
- `AWS_DEFAULT_OUTPUT`- **Optional** The CLI's output output format, defaults to `json` ([more info](https://docs.aws.amazon.com/cli/latest/userguide/cli-environment.html))

## Example

Package and deploy AWS SAM application with Slack notifications:

```hcl
workflow "Deployment" {
  on = "push"
  resolves = [
    "Build Notification",
    "Deploy Notification",
  ]
}

action "Build" {
  uses = "apex/actions/aws/sam@master"
  secrets = ["AWS_SECRET_ACCESS_KEY", "AWS_ACCESS_KEY_ID"]
  args = "package --template-file template.yml --output-template-file out.yml --s3-bucket my-bucket-name"
}

action "Build Notification" {
  needs = "Build"
  uses = "apex/actions/slack@master"
  secrets = ["SLACK_WEBHOOK_URL"]
}

action "Deploy" {
  uses = "apex/actions/aws/sam@master"
  needs = ["Build"]
  secrets = ["AWS_ACCESS_KEY_ID", "AWS_SECRET_ACCESS_KEY"]
  args = "deploy --stack-name myapp --capabilities CAPABILITY_IAM --template-file out.yml"
  env = {
    AWS_DEFAULT_REGION = "us-west-2"
  }
}

action "Deploy Notification" {
  uses = "apex/actions/slack@master"
  needs = ["Deploy"]
  secrets = ["SLACK_WEBHOOK_URL"]
}
```

## Notes

This action generates a Slack message upon deployment.
