package netbox

import (
	"fmt"

	runtimeclient "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/netbox-community/go-netbox/netbox/client"
)

const authHeaderName = "Authorization"
const authHeaderFormat = "Token %v"

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

	t := runtimeclient.New(url, client.DefaultBasePath, client.DefaultSchemes)
	t.DefaultAuthentication = runtimeclient.APIKeyAuth(authHeaderName, "header", fmt.Sprintf(authHeaderFormat, token))
	return client.New(t, strfmt.Default), nil
}
