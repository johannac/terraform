package test

import (
	"bytes"
	"fmt"

	"github.com/hashicorp/terraform/helper/hashcode"
	"github.com/hashicorp/terraform/helper/schema"
)

func testResource() *schema.Resource {
	return &schema.Resource{
		Create: testResourceCreate,
		Read:   testResourceRead,
		Update: testResourceUpdate,
		Delete: testResourceDelete,
		Schema: map[string]*schema.Schema{
			"required": {
				Type:     schema.TypeString,
				Required: true,
			},
			"optional": {
				Type:     schema.TypeString,
				Optional: true,
			},
			"optional_bool": {
				Type:     schema.TypeBool,
				Optional: true,
			},
			"optional_force_new": {
				Type:     schema.TypeString,
				Optional: true,
				ForceNew: true,
			},
			"optional_computed_map": {
				Type:     schema.TypeMap,
				Optional: true,
				Computed: true,
			},
			"computed_read_only": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			"computed_read_only_force_new": {
				Type:     schema.TypeString,
				Computed: true,
				ForceNew: true,
			},
			"computed_list": {
				Type:     schema.TypeList,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"set": {
				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"computed_set": {
				Type:     schema.TypeSet,
				Computed: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
				Set: schema.HashString,
			},
			"map": {
				Type:     schema.TypeMap,
				Optional: true,
			},
			"optional_map": {
				Type:     schema.TypeMap,
				Optional: true,
			},
			"required_map": {
				Type:     schema.TypeMap,
				Required: true,
			},
			"map_that_look_like_set": {
				Type:     schema.TypeMap,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeString,
				},
			},
			"computed_map": {
				Type:     schema.TypeMap,
				Computed: true,
			},
			"list_of_map": {
				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Schema{
					Type: schema.TypeMap,
					Elem: &schema.Schema{
						Type: schema.TypeString,
					},
				},
			},
		},
	}
}

func testResourceCreate(d *schema.ResourceData, meta interface{}) error {
	d.SetId("testId")

	// Required must make it through to Create
	if _, ok := d.GetOk("required"); !ok {
		return fmt.Errorf("Missing attribute 'required', but it's required!")
	}
	if _, ok := d.GetOk("required_map"); !ok {
		return fmt.Errorf("Missing attribute 'required_map', but it's required!")
	}

	return testResourceRead(d, meta)
}

func testResourceRead(d *schema.ResourceData, meta interface{}) error {
	d.Set("computed_read_only", "value_from_api")
	d.Set("computed_read_only_force_new", "value_from_api")
	if _, ok := d.GetOk("optional_computed_map"); !ok {
		d.Set("optional_computed_map", map[string]string{})
	}
	d.Set("computed_map", map[string]string{"key1": "value1"})
	d.Set("computed_list", []string{"listval1", "listval2"})
	d.Set("computed_set", []string{"setval1", "setval2"})
	return nil
}

func testResourceUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func testResourceDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}

func computedTestResource() *schema.Resource {
	return &schema.Resource{
		Create: computedTestResourceCreate,
		Read:   computedTestResourceRead,
		Update: computedTestResourceUpdate,
		Delete: computedTestResourceDelete,
		Schema: map[string]*schema.Schema{
			"ebs_block_device": &schema.Schema{
				Type:     schema.TypeSet,
				Optional: true,
				Computed: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"delete_on_termination": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Default:  true,
							ForceNew: true,
						},

						"device_name": &schema.Schema{
							Type:     schema.TypeString,
							Required: true,
							ForceNew: true,
						},

						"encrypted": &schema.Schema{
							Type:     schema.TypeBool,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						"iops": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						"snapshot_id": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						"volume_size": &schema.Schema{
							Type:     schema.TypeInt,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},

						"volume_type": &schema.Schema{
							Type:     schema.TypeString,
							Optional: true,
							Computed: true,
							ForceNew: true,
						},
					},
				},
				Set: func(v interface{}) int {
					var buf bytes.Buffer
					m := v.(map[string]interface{})

					buf.WriteString(fmt.Sprintf("%t-", m["delete_on_termination"].(bool)))
					buf.WriteString(fmt.Sprintf("%s-", m["device_name"].(string)))
					buf.WriteString(fmt.Sprintf("%t-", m["encrypted"].(bool)))
					buf.WriteString(fmt.Sprintf("%d-", m["iops"].(int)))
					buf.WriteString(fmt.Sprintf("%s-", m["snapshot_id"].(string)))
					buf.WriteString(fmt.Sprintf("%d-", m["volume_size"].(int)))
					buf.WriteString(fmt.Sprintf("%s-", m["volume_type"].(string)))

					return hashcode.String(buf.String())
				},
			},
		},
	}
}

func computedTestResourceCreate(d *schema.ResourceData, meta interface{}) error {
	d.SetId("computedTestId")

	return testResourceRead(d, meta)
}

func computedTestResourceRead(d *schema.ResourceData, meta interface{}) error {
	m := make(map[string]interface{})
	m["device_name"] = "/dev/sdc"
	m["delete_on_termination"] = true
	m["volume_size"] = 10
	m["volume_type"] = "gp2"
	m["iops"] = 100
	m["encrypted"] = false

	if err := d.Set("ebs_block_device", []interface{}{m}); err != nil {
		return err
	}
	return nil
}

func computedTestResourceUpdate(d *schema.ResourceData, meta interface{}) error {
	return nil
}

func computedTestResourceDelete(d *schema.ResourceData, meta interface{}) error {
	d.SetId("")
	return nil
}
