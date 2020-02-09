package netbox

import (
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	netboxapi "github.com/netbox-community/go-netbox/netbox"
)

// Provider exports the actual provider.
func Provider() *schema.Provider {
	return &schema.Provider{
		Schema: map[string]*schema.Schema{
			"url": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_URL", "127.0.0.1:8000"),
				Description: "URL to reach netbox application.",
			},
			"token": {
				Type:        schema.TypeString,
				Required:    true,
				DefaultFunc: schema.EnvDefaultFunc("NETBOX_TOKEN", ""),
				Description: "The token used for API operations.",
			},
		},
		DataSourcesMap: map[string]*schema.Resource{
			"netbox_dcim_site":            dataNetboxDcimSite(),
			"netbox_ipam_role":            dataNetboxIpamRole(),
			"netbox_ipam_vlan_group":      dataNetboxIpamVlanGroup(),
			"netbox_ipam_vlan":            dataNetboxIpamVlan(),
			"netbox_tenancy_tenant_group": dataNetboxTenancyTenantGroup(),
			"netbox_tenancy_tenant":       dataNetboxTenancyTenant(),
		},
		ResourcesMap: map[string]*schema.Resource{
			"netbox_ipam_prefix":          resourceNetboxIpamPrefix(),
			"netbox_ipam_vlan":            resourceNetboxIpamVlan(),
			"netbox_ipam_vlan_group":      resourceNetboxIpamVlanGroup(),
			"netbox_tenancy_tenant":       resourceNetboxTenancyTenant(),
			"netbox_tenancy_tenant_group": resourceNetboxTenancyTenantGroup(),
		},
		ConfigureFunc: configureProvider,
	}
}

func configureProvider(d *schema.ResourceData) (interface{}, error) {
	url := d.Get("url").(string)
	token := d.Get("token").(string)

	return netboxapi.NewNetboxWithAPIKey(url, token), nil
}
