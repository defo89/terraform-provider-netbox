package netbox

import (
	"regexp"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	netboxclient "github.com/smutel/go-netbox/netbox/client"
	"github.com/smutel/go-netbox/netbox/client/ipam"
)

func dataNetboxIpamVlanGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataNetboxIpamVlanGroupRead,

		Schema: map[string]*schema.Schema{
			"slug": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile("^[-a-zA-Z0-9_]{1,50}$"),
					"Must be like ^[-a-zA-Z0-9_]{1,50}$"),
			},
			"site_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
		},
	}
}

func dataNetboxIpamVlanGroupRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*netboxclient.NetBox)

	vlanGroupSlug := d.Get("slug").(string)
	vlanSiteID := d.Get("site_id").(int)
	vlanSiteIDStr := strconv.FormatInt(int64(vlanSiteID), 10)

	p := ipam.NewIpamVlanGroupsListParams().WithSlug(&vlanGroupSlug)
	if vlanSiteID != 0 {
		p.SetSiteID(&vlanSiteIDStr)
	}

	vlanGroupsList, err := client.Ipam.IpamVlanGroupsList(p, nil)
	if err != nil {
		return err
	}

	if *vlanGroupsList.Payload.Count == 1 {
		d.SetId(strconv.FormatInt(vlanGroupsList.Payload.Results[0].ID, 10))
	}

	return nil
}
