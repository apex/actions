#!/bin/sh

set -e

# Respect AWS_DEFAULT_REGION if specified
[ -n "$AWS_DEFAULT_REGION" ] || export AWS_DEFAULT_REGION=us-east-1

# Respect AWS_DEFAULT_OUTPUT if specified
[ -n "$AWS_DEFAULT_OUTPUT" ] || export AWS_DEFAULT_OUTPUT=json

# Capture output
start=$SECONDS
output=$( sh -c "sam $*" )

# Preserve output for consumption by downstream actions
echo "$output" > "${HOME}/${GITHUB_ACTION}.${AWS_DEFAULT_OUTPUT}"

# Write output to STDOUT
echo "$output"

# Slack message
duration=$(($SECONDS - $start))
icon="https://newmathdata.com/wp-content/uploads/2018/03/aws_sam_local.png"
title=$GITHUB_REPOSITORY
text="$1 completed"
footer="Completed in ${duration}s"

cat <<EOF > slack.json
{
  "icon_url": $icon,
  "attachments": [
    {
      "title": "$title",
      "text": "$text",
      "footer": "$footer"
    }
  ]
}
EOF
