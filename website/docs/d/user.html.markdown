---
layout: "slack"
page_title: "Slack: slack_user"
sidebar_current: "docs-aws-datasource-user"
description: |-
  Get information on Slack user.
---

# Data Source: slack_user

Use this data source to get information on an existing slack user.

## Example Usage

```hcl
data "slack_user" "bob" {
  email = "bob@example.com"
}
```

## Argument Reference

* `email` - (Required) Indicate an email of a user to find.


## Attributes Reference

`id` is set to the ID of the found user. In addition, the following attributes are exported:

* `team_id` - ID of a team user belongs to.
* `real_name` - Users full name.
* `is_owner` - a boolean value indicating whether this user is the owner of this enterprise.
* `is_admin` - a boolean value indicating whether this user administers this enterprise.
* `is_bot` - a boolean value indicating whether this user is a bot.


