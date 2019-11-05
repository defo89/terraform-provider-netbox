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

func resourceNetboxTenancyTenant() *schema.Resource {
	return &schema.Resource{
		Create: resourceNetboxTenancyTenantCreate,
		Read:   resourceNetboxTenancyTenantRead,
		Update: resourceNetboxTenancyTenantUpdate,
		Delete: resourceNetboxTenancyTenantDelete,
		Exists: resourceNetboxTenancyTenantExists,

		Schema: map[string]*schema.Schema{
			"name": {
				Type:         schema.TypeString,
				Required:     true,
				ValidateFunc: validation.StringLenBetween(1, 30),
			},
			"slug": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringMatch(
					regexp.MustCompile("^[-a-zA-Z0-9_]{1,50}$"),
					"Must be like ^[-a-zA-Z0-9_]{1,50}$"),
			},
			"tenant_group_id": {
				Type:     schema.TypeInt,
				Optional: true,
			},
			"description": {
				Type:         schema.TypeString,
				Optional:     true,
				ValidateFunc: validation.StringLenBetween(1, 100),
			},
			"comments": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"tags": {
				Type: schema.TypeSet,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
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
		},
	}
}

func resourceNetboxTenancyTenantCreate(d *schema.ResourceData,
	m interface{}) error {
	client := m.(*netboxclient.NetBox)

	tenantName := d.Get("name").(string)
	tenantSlug := d.Get("slug").(string)
	tenantDescription := d.Get("description").(string)
	tenantComments := d.Get("comments").(string)
	tenantGroupID := int64(d.Get("tenant_group_id").(int))
	tenantTags := d.Get("tags").(*schema.Set).List()
	tenantCustomFields := d.Get("custom_field").(*schema.Set).List()

	cf, err := convertCFToAPI(tenantCustomFields)
	if err != nil {
		return err
	}

	newTenant := &models.WritableTenant{
		Name:         &tenantName,
		Slug:         &tenantSlug,
		Description:  tenantDescription,
		Comments:     tenantComments,
		Tags:         expandToStringSlice(tenantTags),
		CustomFields: cf,
	}

	if tenantGroupID != 0 {
		newTenant.Group = &tenantGroupID
	}

	p := tenancy.NewTenancyTenantsCreateParams().WithData(newTenant)

	tenantCreated, err := client.Tenancy.TenancyTenantsCreate(p, nil)
	if err != nil {
		return err
	}

	d.SetId(strconv.FormatInt(tenantCreated.Payload.ID, 10))

	return resourceNetboxTenancyTenantRead(d, m)
}

func resourceNetboxTenancyTenantRead(d *schema.ResourceData,
	m interface{}) error {
	client := m.(*netboxclient.NetBox)

	tenants, err := client.Tenancy.TenancyTenantsList(nil, nil)
	if err != nil {
		return err
	}

	for _, tenant := range tenants.Payload.Results {
		if strconv.FormatInt(tenant.ID, 10) == d.Id() {
			d.Set("name", tenant.Name)
			d.Set("slug", tenant.Slug)
			d.Set("description", tenant.Description)
			d.Set("comments", tenant.Comments)
			d.Set("tags", tenant.Tags)

			if tenant.Group == nil {
				d.Set("tenant_group_id", nil)
			} else {
				d.Set("tenant_group_id", tenant.Group.ID)
			}

			d.Set("custom_fields", convertAPIToCF(tenant.CustomFields))
			return nil
		}
	}

	d.SetId("")
	return nil
}

func resourceNetboxTenancyTenantUpdate(d *schema.ResourceData,
	m interface{}) error {
	client := m.(*netboxclient.NetBox)
	updatedParams := &models.WritableTenant{}

	name := d.Get("name").(string)
	updatedParams.Name = &name

	slug := d.Get("slug").(string)
	updatedParams.Slug = &slug

	if d.HasChange("description") {
		description := d.Get("description").(string)
		updatedParams.Description = description
	}

	if d.HasChange("comments") {
		comments := d.Get("comments").(string)
		updatedParams.Comments = comments
	}

	if d.HasChange("tenant_group_id") {
		tenantGroupID := int64(d.Get("tenant_group_id").(int))
		updatedParams.Group = &tenantGroupID
	}

	tenantTags := d.Get("tags").(*schema.Set).List()
	updatedParams.Tags = expandToStringSlice(tenantTags)

	p := tenancy.NewTenancyTenantsPartialUpdateParams().WithData(updatedParams)

	tenantID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return pkgerrors.New("Unable to convert tenant ID into int64")
	}

	p.SetID(tenantID)

	_, err = client.Tenancy.TenancyTenantsPartialUpdate(p, nil)
	if err != nil {
		return err
	}

	return resourceNetboxTenancyTenantRead(d, m)
}

func resourceNetboxTenancyTenantDelete(d *schema.ResourceData,
	m interface{}) error {
	client := m.(*netboxclient.NetBox)

	resourceExists, err := resourceNetboxTenancyTenantExists(d, m)
	if err != nil {
		return err
	}

	if resourceExists == false {
		return nil
	}

	tenantID, err := strconv.ParseInt(d.Id(), 10, 64)
	if err != nil {
		return pkgerrors.New("Unable to convert tenant ID into int64")
	}

	p := tenancy.NewTenancyTenantsDeleteParams().WithID(tenantID)
	if _, err := client.Tenancy.TenancyTenantsDelete(p, nil); err != nil {
		return err
	}

	return nil
}

func resourceNetboxTenancyTenantExists(d *schema.ResourceData,
	m interface{}) (b bool,
	e error) {
	client := m.(*netboxclient.NetBox)
	tenantExist := false

	tenants, err := client.Tenancy.TenancyTenantsList(nil, nil)
	if err != nil {
		return tenantExist, err
	}

	for _, tenant := range tenants.Payload.Results {
		if strconv.FormatInt(tenant.ID, 10) == d.Id() {
			tenantExist = true
		}
	}

	return tenantExist, nil
}
