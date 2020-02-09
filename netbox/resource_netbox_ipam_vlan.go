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

var vlanStatusToInt = map[string]int{
	"Active":     1,
	"Reserved":   2,
	"Deprecated": 3,
}

var vlanIntToStatus = map[int]string{
	1: "Active",
	2: "Reserved",
	3: "Deprecated",
}

func resourceNetboxIpamVlan() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxIpamVlanCreate,
		Read:   resourceNetboxIpamVlanRead,
		Update: resourceNetboxIpamVlanUpdate,
		Delete: resourceNetboxIpamVlanDelete,
		Exists: resourceNetboxIpamVlanExists,

		Schema: map[string]*schema.Schema{
			"vlan_id": {
				Type:     schema.TypeInt,
				Required: true,
			},
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 50),
			},
			"status": {
				Type:     schema.TypeString,
				Optional: true,
				Default:  "Active",
				ValidateFunc: validation.StringInSlice([]string{"Active", "Reserved",
					"Deprecated"}, false),
			},
			"site_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"vlan_group_id": {
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

func resourceNetboxIpamVlanCreate(d *schema.ResourceData,
	m interface{}) error {
	client := m.(*netboxclient.NetBox)

	vlanVid := int64(d.Get("vlan_id").(int))
	vlanName := d.Get("name").(string)
	vlanStatus := int64(vlanStatusToInt[d.Get("status").(string)])
	vlanSiteID := int64(d.Get("site_id").(int))
	vlanGroupID := int64(d.Get("vlan_group_id").(int))
	vlanRoleID := int64(d.Get("role_id").(int))
	vlanDescription := d.Get("description").(string)
	vlanTenantID := int64(d.Get("tenant_id").(int))
	vlanTags := d.Get("tags").(*schema.Set).List()

	newVlan := &models.WritableVLAN{
		Vid:         &vlanVid,
		Name:        &vlanName,
		Status:      vlanStatus,
		Tags:        expandToStringSlice(vlanTags),
		Description: vlanDescription,
	}

	if vlanSiteID != 0 {
		newVlan.Site = &vlanSiteID
	}

	if vlanGroupID != 0 {
		newVlan.Group = &vlanGroupID
	}

	if vlanRoleID != 0 {
		newVlan.Role = &vlanRoleID
	}

	if vlanTenantID != 0 {
		newVlan.Tenant = &vlanTenantID
	}

	p := ipam.NewIpamVlansCreateParams().WithData(newVlan)

	vlanCreated, err := client.Ipam.IpamVlansCreate(p, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(vlanCreated.Payload.ID, 10))
	return resourceNetboxIpamVlanRead(d, m)
}

func resourceNetboxIpamVlanRead(d *schema.ResourceData,
	m interface{}) error {
	client := m.(*netboxclient.NetBox)

	vlans, err := client.Ipam.IpamVlansList(nil, nil)
	if err != nil {
		return err
	}

	for _, vlan := range vlans.Payload.Results {
		if strconv.FormatInt(vlan.ID, 10) == d.Id() {
			d.Set("vlan_id", vlan.Vid)
			d.Set("status", vlanIntToStatus[int(*vlan.Status.Value)])
			d.Set("name", vlan.Name)
			d.Set("description", vlan.Description)
			d.Set("tags", vlan.Tags)

			if vlan.Site == nil {
				d.Set("site_id", nil)
			} else {
				d.Set("site_id", vlan.Site.ID)
			}

			if vlan.Group == nil {
				d.Set("vlan_group_id", nil)
			} else {
				d.Set("vlan_group_id", vlan.Group.ID)
			}

			if vlan.Tenant == nil {
				d.Set("tenant_id", nil)
			} else {
				d.Set("tenant_id", vlan.Tenant.ID)
			}

			if vlan.Role == nil {
				d.Set("role_id", nil)
			} else {
				d.Set("role_id", vlan.Role.ID)
			}

			return nil
		}
	}

	d.SetId("")
	return nil
}

func resourceNetboxIpamVlanUpdate(d *schema.ResourceData,
	m interface{}) error {
	client := m.(*netboxclient.NetBox)
	updatedParams := &models.WritableVLAN{}

	vlanVid := int64(d.Get("vlan_id").(int))
	updatedParams.Vid = &vlanVid

	name := d.Get("name").(string)
	updatedParams.Name = &name

	vlanTags := d.Get("tags").(*schema.Set).List()
	updatedParams.Tags = expandToStringSlice(vlanTags)

	if d.HasChange("status") {
		status := int64(vlanStatusToInt[d.Get("status").(string)])
		updatedParams.Status = status
	}

	if d.HasChange("description") {
		description := d.Get("description").(string)
		updatedParams.Description = description
	}

	if d.HasChange("site_id") {
		siteID := int64(d.Get("site_id").(int))
		if siteID != 0 {
			updatedParams.Site = &siteID
		}
	}

	if d.HasChange("vlan_group_id") {
		vlanGroupID := int64(d.Get("vlan_group_id").(int))
		if vlanGroupID != 0 {
			updatedParams.Group = &vlanGroupID
		}
	}

	if d.HasChange("tenant_id") {
		vlanTenantID := int64(d.Get("tenant_id").(int))
		if vlanTenantID != 0 {
			updatedParams.Tenant = &vlanTenantID
		}
	}

	if d.HasChange("role_id") {
		vlanRoleID := int64(d.Get("role_id").(int))
		if vlanRoleID != 0 {
			updatedParams.Role = &vlanRoleID
		}
	}

	p := ipam.NewIpamVlansPartialUpdateParams().WithData(
		updatedParams)

	tenantID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return pkgerrors.New("Unable to convert tenant ID into int64")
	}

	p.SetID(tenantID)

	_, err = client.Ipam.IpamVlansPartialUpdate(p, nil)
	if err != nil {
		return err
	}

	return resourceNetboxIpamVlanRead(d, m)
}

func resourceNetboxIpamVlanDelete(d *schema.ResourceData, m interface{}) error {
	client := m.(*netboxclient.NetBox)

	resourceExists, err := resourceNetboxIpamVlanExists(d, m)
	if err != nil {
		return err
	}

	if resourceExists == false {
		return nil
	}

	vlanID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return pkgerrors.New("Unable to convert vlan group ID into int64")
	}

	p := ipam.NewIpamVlansDeleteParams().WithID(vlanID)
	if _, err := client.Ipam.IpamVlansDelete(p, nil); err != nil {
		return err
	}

	return nil
}

func resourceNetboxIpamVlanExists(d *schema.ResourceData, m interface{}) (b bool,
	e error) {
	client := m.(*netboxclient.NetBox)
	vlanExist := false

	vlans, err := client.Ipam.IpamVlansList(nil, nil)
	if err != nil {
		return vlanExist, err
	}

	for _, vlan := range vlans.Payload.Results {
		if strconv.FormatInt(vlan.ID, 10) == d.Id() {
			vlanExist = true
		}
	}

	return vlanExist, nil
}
