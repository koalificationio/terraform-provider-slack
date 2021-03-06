package slack

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/nlopes/slack"
)

func dataSourceSlackChannel() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSlackChannelRead,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"is_general": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_archived": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_shared": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_private": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceSlackChannelRead(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*slack.Client)

	name := d.Get("name").(string)

	params := slack.GetConversationsParameters{
		Limit: 100,
	}

	var channel slack.Channel

	for {
		var cs []slack.Channel

		cs, cursor, err := api.GetConversations(&params)
		if err != nil {
			// Slack rate limit tire 2 https://api.slack.com/docs/rate-limits#tier_t2
			if r, ok := err.(*slack.RateLimitedError); ok {
				log.Printf("[INFO] got rate limited by Slack API, sleep for %v", r.RetryAfter)
				// Add 1s because its never enough
				time.Sleep(r.RetryAfter + 1*time.Second)
				continue
			} else {
				return fmt.Errorf("Could not fetch channels: %v", err)
			}
		}

		for _, c := range cs {
			if c.Name == name {
				channel = c
				break
			}
		}

		if cursor == "" {
			break
		}

		params = slack.GetConversationsParameters{
			Limit:  100,
			Cursor: cursor,
		}

		time.Sleep(3 * time.Second)
	}

	if channel.Name == "" {
		return fmt.Errorf("Channel '%s' not found", name)
	}

	d.SetId(channel.ID)

	if err := d.Set("is_general", channel.IsGeneral); err != nil {
		return fmt.Errorf("Error setting is_general: %v", err)
	}
	if err := d.Set("is_archived", channel.IsArchived); err != nil {
		return fmt.Errorf("Error setting is_archived: %v", err)
	}
	if err := d.Set("is_shared", channel.IsShared); err != nil {
		return fmt.Errorf("Error setting is_shared: %v", err)
	}
	if err := d.Set("is_private", channel.IsPrivate); err != nil {
		return fmt.Errorf("Error setting is_private: %v", err)
	}

	return nil
}
