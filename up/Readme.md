# Up

GitHub Action for deploying and operating serverless applications with [Up](https://github.com/apex/up).

## Secrets

- `AWS_ACCESS_KEY_ID` - *Required* The AWS Access Key ID.
- `AWS_SECRET_ACCESS_KEY` - *Required* The AWS Secret Key.
- `UP_CONFIG` — Up authentication 

See the [AWS Security Credentials](https://docs.aws.amazon.com/general/latest/gr/aws-security-credentials.html) page for more information regarding AWS credentials, and see [Creating encrypted secrets](https://docs.github.com/en/free-pro-team@latest/actions/reference/encrypted-secrets#creating-encrypted-secrets-for-a-repository)to learn how to create secrets for your GitHub repository.

If you're using Up Pro you can run `up team ci` and paste that as the `UP_CONFIG` secret's value, which will authenticate and install Up Pro for you.

## Example

Deploy an application in the root directory to `production`:

```yaml
name: Deploy

on:
  push:
    branches: [ master ]

jobs:
  build:
    name: Build
    runs-on: ubuntu-latest
    steps:
      - name: Set up Go 1.x
        uses: actions/setup-go@v2
        with:
          go-version: ^1.15
        id: go

      - name: Check out code into the Go module directory
        uses: actions/checkout@v2

      - name: Build
        run: go build -v .

      - name: Test
        run: go test -v .
      
      - name: Deploy
        uses: apex/actions/up@v0.5.1
        env:
          AWS_ACCESS_KEY_ID: ${{ secrets.AWS_ACCESS_KEY_ID }}
          AWS_SECRET_ACCESS_KEY: ${{ secrets.AWS_SECRET_ACCESS_KEY }}
          UP_CONFIG: ${{ secrets.UP_CONFIG }}
        with:
          stage: production

```
