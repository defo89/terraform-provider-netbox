package netbox

import (
	"regexp"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	pkgerrors "github.com/pkg/errors"
	netboxclient "github.com/smutel/go-netbox/netbox/client"
	"github.com/smutel/go-netbox/netbox/client/tenancy"
	"github.com/smutel/go-netbox/netbox/models"
)

func resourceNetboxTenancyTenantGroup() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxTenancyTenantGroupCreate,
		Read:   resourceNetboxTenancyTenantGroupRead,
		Update: resourceNetboxTenancyTenantGroupUpdate,
		Delete: resourceNetboxTenancyTenantGroupDelete,
		Exists: resourceNetboxTenancyTenantGroupExists,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 50),
			},
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

func resourceNetboxTenancyTenantGroupCreate(d *schema.ResourceData,
	m interface{}) error {
	client := m.(*netboxclient.NetBox)

	tenantGroupName := d.Get("name").(string)
	tenantGroupSlug := d.Get("slug").(string)

	newTenantGroup := &models.TenantGroup{
		Name: &tenantGroupName,
		Slug: &tenantGroupSlug,
	}

	p := tenancy.NewTenancyTenantGroupsCreateParams().WithData(newTenantGroup)

	tenantGroupCreated, err := client.Tenancy.TenancyTenantGroupsCreate(p, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(tenantGroupCreated.Payload.ID, 10))
	return resourceNetboxTenancyTenantGroupRead(d, m)
}

func resourceNetboxTenancyTenantGroupRead(d *schema.ResourceData,
	m interface{}) error {
	client := m.(*netboxclient.NetBox)

	tenantGroups, err := client.Tenancy.TenancyTenantGroupsList(nil, nil)
	if err != nil {
		return err
	}

	for _, tenantGroup := range tenantGroups.Payload.Results {
		if strconv.FormatInt(tenantGroup.ID, 10) == d.Id() {
			d.Set("name", tenantGroup.Name)
			d.Set("slug", tenantGroup.Slug)
			return nil
		}
	}

	d.SetId("")
	return nil
}

func resourceNetboxTenancyTenantGroupUpdate(d *schema.ResourceData,
	m interface{}) error {
	client := m.(*netboxclient.NetBox)
	updatedParams := &models.TenantGroup{}

	name := d.Get("name").(string)
	updatedParams.Name = &name

	slug := d.Get("slug").(string)
	updatedParams.Slug = &slug

	p := tenancy.NewTenancyTenantGroupsPartialUpdateParams().WithData(
		updatedParams)

	tenantID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return pkgerrors.New("Unable to convert tenant ID into int64")
	}

	p.SetID(tenantID)

	_, err = client.Tenancy.TenancyTenantGroupsPartialUpdate(p, nil)
	if err != nil {
		return err
	}

	return resourceNetboxTenancyTenantGroupRead(d, m)
}

func resourceNetboxTenancyTenantGroupDelete(d *schema.ResourceData,
	m interface{}) error {
	client := m.(*netboxclient.NetBox)

	resourceExists, err := resourceNetboxTenancyTenantGroupExists(d, m)
	if err != nil {
		return err
	}

	if resourceExists == false {
		return nil
	}

	tenantGroupID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return pkgerrors.New("Unable to convert tenant ID into int64")
	}

	p := tenancy.NewTenancyTenantGroupsDeleteParams().WithID(tenantGroupID)
	if _, err := client.Tenancy.TenancyTenantGroupsDelete(p, nil); err != nil {
		return err
	}

	return nil
}

func resourceNetboxTenancyTenantGroupExists(d *schema.ResourceData,
	m interface{}) (b bool,
	e error) {
	client := m.(*netboxclient.NetBox)
	tenantGroupExist := false

	tenantGroups, err := client.Tenancy.TenancyTenantGroupsList(nil, nil)
	if err != nil {
		return tenantGroupExist, err
	}

	for _, tenantGroup := range tenantGroups.Payload.Results {
		if strconv.FormatInt(tenantGroup.ID, 10) == d.Id() {
			tenantGroupExist = true
		}
	}

	return tenantGroupExist, nil
}
