package netbox

import (
	"regexp"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	netboxclient "github.com/smutel/go-netbox/netbox/client"
	"github.com/smutel/go-netbox/netbox/client/tenancy"
)

func dataNetboxTenancyTenantGroup() *schema.Resource {
	return &schema.Resource{
		Read: dataNetboxTenancyTenantGroupRead,

		Schema: map[string]*schema.Schema{
			"slug": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile("^[-a-zA-Z0-9_]{1,50}$"),
					"Must be like ^[-a-zA-Z0-9_]{1,50}$"),
			},
		},
	}
}

func dataNetboxTenancyTenantGroupRead(d *schema.ResourceData, m interface{}) error {
	client := m.(*netboxclient.NetBox)

	tenantGroupSlug := d.Get("slug").(string)

	p := tenancy.NewTenancyTenantGroupsListParams().WithSlug(&tenantGroupSlug)

	tenantGroupsList, err := client.Tenancy.TenancyTenantGroupsList(p, nil)
	if err != nil {
		return err
	}

	if *tenantGroupsList.Payload.Count == 1 {
		d.SetId(strconv.FormatInt(tenantGroupsList.Payload.Results[0].ID, 10))
	}

	return nil
}
