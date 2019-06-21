---
layout: "slack"
page_title: "Provider: Slack"
sidebar_current: "docs-slack-index"
description: |-
  Terraform Slack provider.
---

# Slack Provider

The provider needs to be configured
with the proper credentials before it can be used.

Use the navigation to the left to read about the available resources.

## Example Usage

```hcl
# Configure the slack Provider
provider "slack" {
  version   = "~> 0.1"
  api_token = "123abc"
}
```

## Argument Reference

In addition to [generic `provider` arguments](https://www.terraform.io/docs/configuration/providers.html)
(e.g. `alias` and `version`), the following arguments are supported in the slack
 `provider` block:

* `api_token` - (Optional) This is the slack api token. It must be provided, but
  it can also be sourced from the `SLACK_TOKEN` environment variable.
