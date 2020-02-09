package netbox

import (
	//"log"
	//"regexp"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	pkgerrors "github.com/pkg/errors"
	netboxclient "github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/ipam"
	"github.com/netbox-community/go-netbox/netbox/models"
)

var prefixStatusToInt = map[string]int{
	"Container":  0,
	"Active":     1,
	"Reserved":   2,
	"Deprecated": 3,
}

var prefixIntToStatus = map[int]string{
	0: "Container",
	1: "Active",
	2: "Reserved",
	3: "Deprecated",
}

func resourceNetboxIpamPrefix() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxIpamPrefixCreate,
		Read:   resourceNetboxIpamPrefixRead,
		Update: resourceNetboxIpamPrefixUpdate,
		Delete: resourceNetboxIpamPrefixDelete,
		Exists: resourceNetboxIpamPrefixExists,

		Schema: map[string]*schema.Schema{
			"prefix": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.CIDRNetwork(0, 256),
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Active",
				ValidateFunc: validation.StringInSlice([]string{"Container", "Active",
					"Reserved", "Deprecated"}, false),
			},
			"vrf_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"role_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(1, 100),
			},
			"is_pool": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"site_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vlan_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"tenant_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"custom_field": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"name": {
							Type:     schema.TypeString,
							Required: true,
						},
						"kind": {
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
							ValidateFunc: validation.StringInSlice([]string{"string",
								"bool", "int"}, false),
						},
						"value": {
							Type:     schema.TypeString,
							Required: true,
						},
					},
				},
			},
			"tags": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Optional: true,
			},
		},
	}
}

func resourceNetboxIpamPrefixCreate(d *schema.ResourceData,
	m interface{}) error {
	client := m.(*netboxclient.NetBox)

	prefixPrefix := d.Get("prefix").(string)
	prefixStatus := int64(prefixStatusToInt[d.Get("status").(string)])
	prefixVrfID := int64(d.Get("vrf_id").(int))
	prefixRoleID := int64(d.Get("role_id").(int))
	prefixDescription := d.Get("description").(string)
	prefixIsPool := d.Get("is_pool").(bool)
	prefixSiteID := int64(d.Get("site_id").(int))
	prefixVlanID := int64(d.Get("vlan_id").(int))
	prefixTenantID := int64(d.Get("tenant_id").(int))
	prefixTags := d.Get("tags").(*schema.Set).List()

	newPrefix := &models.WritablePrefix{
		Prefix:      &prefixPrefix,
		Status:      prefixStatus,
		Tags:        expandToStringSlice(prefixTags),
		Description: prefixDescription,
		IsPool:      prefixIsPool,
	}

	if prefixVrfID != 0 {
		newPrefix.Vrf = &prefixVrfID
	}

	if prefixRoleID != 0 {
		newPrefix.Role = &prefixRoleID
	}

	if prefixSiteID != 0 {
		newPrefix.Site = &prefixSiteID
	}

	if prefixVlanID != 0 {
		newPrefix.Vlan = &prefixVlanID
	}

	if prefixTenantID != 0 {
		newPrefix.Tenant = &prefixTenantID
	}

	p := ipam.NewIpamPrefixesCreateParams().WithData(newPrefix)

	prefixCreated, err := client.Ipam.IpamPrefixesCreate(p, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(prefixCreated.Payload.ID, 10))

	return resourceNetboxIpamPrefixRead(d, m)
}

func resourceNetboxIpamPrefixRead(d *schema.ResourceData,
	m interface{}) error {
	client := m.(*netboxclient.NetBox)

	prefixs, err := client.Ipam.IpamPrefixesList(nil, nil)
	if err != nil {
		return err
	}

	for _, prefix := range prefixs.Payload.Results {
		if strconv.FormatInt(prefix.ID, 10) == d.Id() {
			d.Set("prefix", prefix.Prefix)
			d.Set("status", prefixIntToStatus[int(*prefix.Status.Value)])
			d.Set("description", prefix.Description)
			d.Set("tags", prefix.Tags)
			d.Set("is_pool", prefix.IsPool)

			if prefix.Vrf == nil {
				d.Set("vrf_id", nil)
			} else {
				d.Set("vrf_id", prefix.Vrf.ID)
			}

			if prefix.Role == nil {
				d.Set("role_id", nil)
			} else {
				d.Set("role_id", prefix.Role.ID)
			}

			if prefix.Site == nil {
				d.Set("site_id", nil)
			} else {
				d.Set("site_id", prefix.Site.ID)
			}

			if prefix.Vlan == nil {
				d.Set("vlan_id", nil)
			} else {
				d.Set("vlan_id", prefix.Vlan.ID)
			}

			if prefix.Tenant == nil {
				d.Set("tenant_id", nil)
			} else {
				d.Set("tenant_id", prefix.Tenant.ID)
			}

			return nil
		}
	}

	d.SetId("")
	return nil
}

func resourceNetboxIpamPrefixUpdate(d *schema.ResourceData,
	m interface{}) error {
	client := m.(*netboxclient.NetBox)
	updatedParams := &models.WritablePrefix{}

	prefixPrefix := d.Get("prefix").(string)
	updatedParams.Prefix = &prefixPrefix

	prefixTags := d.Get("tags").(*schema.Set).List()
	updatedParams.Tags = expandToStringSlice(prefixTags)

	updatedParams.IsPool = d.Get("is_pool").(bool)

	if d.HasChange("status") {
		status := int64(prefixStatusToInt[d.Get("status").(string)])
		updatedParams.Status = status
	}

	if d.HasChange("description") {
		description := d.Get("description").(string)
		updatedParams.Description = description
	}

	if d.HasChange("vrf_id") {
		vrfID := int64(d.Get("vrf_id").(int))
		if vrfID != 0 {
			updatedParams.Vrf = &vrfID
		}
	}

	if d.HasChange("role_id") {
		prefixRoleID := int64(d.Get("role_id").(int))
		if prefixRoleID != 0 {
			updatedParams.Role = &prefixRoleID
		}
	}

	if d.HasChange("site_id") {
		siteID := int64(d.Get("site_id").(int))
		if siteID != 0 {
			updatedParams.Site = &siteID
		}
	}

	if d.HasChange("vlan_id") {
		vlanID := int64(d.Get("vlan_id").(int))
		if vlanID != 0 {
			updatedParams.Vlan = &vlanID
		}
	}

	if d.HasChange("tenant_id") {
		prefixTenantID := int64(d.Get("tenant_id").(int))
		if prefixTenantID != 0 {
			updatedParams.Tenant = &prefixTenantID
		}
	}

	p := ipam.NewIpamPrefixesPartialUpdateParams().WithData(updatedParams)

	tenantID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return pkgerrors.New("Unable to convert tenant ID into int64")
	}

	p.SetID(tenantID)

	_, err = client.Ipam.IpamPrefixesPartialUpdate(p, nil)
	if err != nil {
		return err
	}

	return resourceNetboxIpamPrefixRead(d, m)
}

func resourceNetboxIpamPrefixDelete(d *schema.ResourceData,
	m interface{}) error {
	client := m.(*netboxclient.NetBox)

	resourceExists, err := resourceNetboxIpamPrefixExists(d, m)
	if err != nil {
		return err
	}

	if resourceExists == false {
		return nil
	}

	prefixID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return pkgerrors.New("Unable to convert prefix group ID into int64")
	}

	p := ipam.NewIpamPrefixesDeleteParams().WithID(prefixID)
	if _, err := client.Ipam.IpamPrefixesDelete(p, nil); err != nil {
		return err
	}

	return nil
}

func resourceNetboxIpamPrefixExists(d *schema.ResourceData,
	m interface{}) (b bool, e error) {
	client := m.(*netboxclient.NetBox)
	prefixExist := false

	prefixs, err := client.Ipam.IpamPrefixesList(nil, nil)
	if err != nil {
		return prefixExist, err
	}

	for _, prefix := range prefixs.Payload.Results {
		if strconv.FormatInt(prefix.ID, 10) == d.Id() {
			prefixExist = true
		}
	}

	return prefixExist, nil
}
