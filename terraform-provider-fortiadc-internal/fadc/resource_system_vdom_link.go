// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  system vdom link.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceSystemVdomLink() *schema.Resource {
	return &schema.Resource{
		Read:   resourceSystemVdomLinkRead,
		Update: resourceSystemVdomLinkUpdate,
		Create: resourceSystemVdomLinkCreate,
		Delete: resourceSystemVdomLinkDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"type": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
		},
	}
}
func resourceSystemVdomLinkCreate(d *schema.ResourceData, m interface{}) error {
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	mkey := ""

	t := d.Get("mkey")
	if v, ok := t.(string); ok {
		mkey = v
	} else if v, ok := t.(int); ok {
		mkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing SystemVdomLink: type error")
	}

	obj, err := getObjectSystemVdomLink(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating SystemVdomLink resource while getting object: %v", err)
	}

	_, err = c.CreateSystemVdomLink(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating SystemVdomLink resource: %v", err)
	}

	d.SetId(mkey)

	return resourceSystemVdomLinkRead(d, m)
}
func resourceSystemVdomLinkUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectSystemVdomLink(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomLink resource while getting object: %v", err)
	}

	_, err = c.UpdateSystemVdomLink(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating SystemVdomLink resource: %v", err)
	}

	return resourceSystemVdomLinkRead(d, m)
}
func resourceSystemVdomLinkDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteSystemVdomLink(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting SystemVdomLink resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceSystemVdomLinkRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadSystemVdomLink(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomLink resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectSystemVdomLink(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading SystemVdomLink resource from API: %v", err)
	}
	return nil
}
func flattenSystemVdomLinkType(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenSystemVdomLinkMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectSystemVdomLink(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("type", flattenSystemVdomLinkType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("mkey", flattenSystemVdomLinkMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandSystemVdomLinkType(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandSystemVdomLinkMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectSystemVdomLink(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("type"); ok {
		t, err := expandSystemVdomLinkType(d, v, "type", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["type"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandSystemVdomLinkMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
