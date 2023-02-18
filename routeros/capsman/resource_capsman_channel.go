package routeros

import (
	"regexp"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/validation"
)

/*
	{
	  ".id": "*1",
	  "band": "2ghz-b",
	  "control-channel-width": "5mhz",
	  "extension-channel": "disabled",
	  "frequency": "2112",
	  "name": "channel1",
	  "reselect-interval": "10m",
	  "save-selected": "true",
	  "secondary-frequency": "disabled",
	  "skip-dfs-channels": "true",
	  "tx-power": "-20"
	}
*/

// https://help.mikrotik.com/docs/display/ROS/CAPsMAN
func ResourceCapsManChannel() *schema.Resource {

	resSchema := map[string]*schema.Schema{
		MetaResourcePath: PropResourcePath("/caps-man/channel"),
		MetaId:           PropId(Name),

		"band": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Define operational radio frequency band and mode taken from hardware capability of wireless card.",
			ValidateFunc: validation.StringMatch(regexp.MustCompile(`^(2|5)ghz-`),
				"The value must start with the frequency of the band: 2ghz- or 5ghz-"),
		},
		KeyComment: PropCommentRw,
		"control_channel_width": {
			Type:         schema.TypeString,
			Optional:     true,
			Description:  "Control channel width.",
			ValidateFunc: validation.StringInSlice([]string{"5mhz", "10mhz", "20mhz", "40mhz-turbo"}, false),
		},
		"extension_channel": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Extension channel configuration. (E.g. Ce = extension channel is above Control channel, " +
				"eC = extension channel is below Control channel)",
		},
		"frequency": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "Channel frequency value in MHz on which AP will operate. If left blank, CAPsMAN will " +
				"automatically determine the best frequency that is least occupied.",
		},
		KeyName: PropNameRw,
		"reselect_interval": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "The interval after which the least occupied frequency is chosen, can be defined as a random " +
				"interval, ex. as '30m..60m'. Works only if channel.frequency is left blank.",
			// We may need to write a custom DiffSuppressFunc.
			// DiffSuppressFunc: TimeEquall, not for time ranges
		},
		"save_selected": {
			Type:     schema.TypeBool,
			Optional: true,
			Description: "If channel frequency is chosen automatically and channel.reselect-interval is used, then " +
				"saves the last picked frequency.",
		},
		"secondary_frequency": {
			Type:     schema.TypeString,
			Optional: true,
			Description: "Specifies the second frequency that will be used for 80+80MHz configuration. " +
				"Set it to Disabled in order to disable 80+80MHz capability.",
		},
		"skip_dfs_channels": {
			Type:        schema.TypeBool,
			Optional:    true,
			Description: "If channel.frequency is left blank, the selection will skip DFS channels.",
		},
		"tx_power": {
			Type:     schema.TypeInt,
			Optional: true,
			Description: "[-30..40] TX Power for CAP interface (for the whole interface not for individual chains) in dBm. " +
				"It is not possible to set higher than allowed by country regulations or interface. " +
				"By default max allowed by country or interface is used.",
		},
		"width": {
			Type:        schema.TypeString,
			Optional:    true,
			Description: "Channel Width in MHz.",
		},
	}

	return &schema.Resource{
		CreateContext: DefaultCreate(resSchema),
		ReadContext:   DefaultRead(resSchema),
		UpdateContext: DefaultUpdate(resSchema),
		DeleteContext: DefaultDelete(resSchema),

		Importer: &schema.ResourceImporter{
			StateContext: schema.ImportStatePassthroughContext,
		},

		Schema: resSchema,
	}
}
