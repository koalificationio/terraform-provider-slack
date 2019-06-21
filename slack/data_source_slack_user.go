package slack

import (
	"fmt"

	"github.com/hashicorp/terraform/helper/schema"
	"github.com/nlopes/slack"
)

func dataSourceSlackUser() *schema.Resource {
	return &schema.Resource{
		Read: dataSourceSlackUserRead,

		Schema: map[string]*schema.Schema{
			"email": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},
			"team_id": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"real_name": {
				Type:     schema.TypeString,
				Computed: true,
			},
			"is_owner": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_admin": {
				Type:     schema.TypeBool,
				Computed: true,
			},
			"is_bot": {
				Type:     schema.TypeBool,
				Computed: true,
			},
		},
	}
}

func dataSourceSlackUserRead(d *schema.ResourceData, meta interface{}) error {
	api := meta.(*slack.Client)

	email := d.Get("email").(string)

	user, err := api.GetUserByEmail(email)
	if err != nil {
		return fmt.Errorf("Could not fetch user: %v", err)
	}

	d.SetId(user.ID)
	d.Set("team_id", user.TeamID)
	d.Set("real_name", user.Profile.RealName)
	d.Set("is_owner", user.IsOwner)
	d.Set("is_admin", user.IsAdmin)
	d.Set("is_bot", user.IsBot)

	return nil
}
