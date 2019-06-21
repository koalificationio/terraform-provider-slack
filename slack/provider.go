package slack

import (
	"github.com/hashicorp/terraform/helper/schema"
	"github.com/hashicorp/terraform/terraform"
)

func Provider() terraform.ResourceProvider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"api_token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("SLACK_TOKEN", nil),
				Description: "The Slack API token.",
			},
		},

		DataSourcesMap: map[string]*schema.Resource{
			"slack_user":    dataSourceSlackUser(),
			"slack_channel": dataSourceSlackChannel(),
		},

		ConfigureFunc: providerConfigure,
	}
}

func providerConfigure(d *schema.ResourceData) (interface{}, error) {
	config := &Config{
		APIToken: d.Get("api_token").(string),
	}

	return config.Client()
}
