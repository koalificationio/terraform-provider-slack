package slack

import (
	"fmt"
	"log"
	"time"

	"github.com/hashicorp/terraform/helper/schema"
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

	var allChannels []slack.Channel

	for {
		var cs []slack.Channel

		cs, cursor, err := api.GetConversations(&params)
		if err != nil {
			// Slack rate limit tire 2 https://api.slack.com/docs/rate-limits#tier_t2
			if r, ok := err.(*slack.RateLimitedError); ok {
				log.Printf("[INFO] got rate limited by Slack API, sleep for %v", r.RetryAfter)
				time.Sleep(r.RetryAfter)
				// Add 3s because its never enough
				time.Sleep(3 * time.Second)
				continue
			} else {
				return fmt.Errorf("Could not fetch channels: %v", err)
			}
		}

		allChannels = append(allChannels, cs...)

		if cursor == "" {
			break
		}

		params = slack.GetConversationsParameters{
			Limit:  100,
			Cursor: cursor,
		}

		time.Sleep(3 * time.Second)
	}

	var channel slack.Channel
	for _, c := range allChannels {
		if c.Name == name {
			channel = c
			break
		}
	}

	if channel.Name == "" {
		return fmt.Errorf("Channel '%s' not found", name)
	}

	d.SetId(channel.ID)
	d.Set("is_general", channel.IsGeneral)
	d.Set("is_archived", channel.IsArchived)
	d.Set("is_shared", channel.IsShared)
	d.Set("is_private", channel.IsPrivate)

	return nil
}
