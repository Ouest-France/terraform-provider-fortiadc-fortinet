// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router md5 ospf child md5 member.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterMd5OspfChildMd5Member() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRouterMd5OspfChildMd5MemberRead,
		Update: resourceRouterMd5OspfChildMd5MemberUpdate,
		Create: resourceRouterMd5OspfChildMd5MemberCreate,
		Delete: resourceRouterMd5OspfChildMd5MemberDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"mkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
			"key": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"vdom": &schema.Schema{
				Type:     schema.TypeString,
				Optional: true,
			},
			"pkey": &schema.Schema{
				Type:     schema.TypeString,
				Required: true,
			},
		},
	}
}
func resourceRouterMd5OspfChildMd5MemberCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing RouterMd5OspfChildMd5Member: type error")
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing RouterMd5OspfChildMd5Member: type error")
	}

	obj, err := getObjectRouterMd5OspfChildMd5Member(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterMd5OspfChildMd5Member resource while getting object: %v", err)
	}

	new_id := fmt.Sprintf("%s_%s", pkey, mkey)
	_, err = c.CreateRouterMd5OspfChildMd5Member(pkey, obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating RouterMd5OspfChildMd5Member resource: %v", err)
	}

	d.SetId(new_id)

	return resourceRouterMd5OspfChildMd5MemberRead(d, m)
}
func resourceRouterMd5OspfChildMd5MemberUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing RouterMd5OspfChildMd5Member: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	obj, err := getObjectRouterMd5OspfChildMd5Member(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterMd5OspfChildMd5Member resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterMd5OspfChildMd5Member(pkey, obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating RouterMd5OspfChildMd5Member resource: %v", err)
	}

	return resourceRouterMd5OspfChildMd5MemberRead(d, m)
}
func resourceRouterMd5OspfChildMd5MemberDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing RouterMd5OspfChildMd5Member: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	err := c.DeleteRouterMd5OspfChildMd5Member(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting RouterMd5OspfChildMd5Member resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceRouterMd5OspfChildMd5MemberRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	pkey := ""

	tp := d.Get("pkey")
	if v, ok := tp.(string); ok {
		pkey = v
	} else if v, ok := tp.(int); ok {
		pkey = strconv.Itoa(v)
	} else {
		return fmt.Errorf("Error describing RouterMd5OspfChildMd5Member: type error")
	}

	mkey = strings.TrimPrefix(mkey, pkey)
	mkey = strings.TrimPrefix(mkey, "_")
	o, err := c.ReadRouterMd5OspfChildMd5Member(pkey, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading RouterMd5OspfChildMd5Member resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterMd5OspfChildMd5Member(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterMd5OspfChildMd5Member resource from API: %v", err)
	}
	return nil
}
func flattenRouterMd5OspfChildMd5MemberMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterMd5OspfChildMd5MemberKey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterMd5OspfChildMd5Member(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("mkey", flattenRouterMd5OspfChildMd5MemberMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	if err = d.Set("key", flattenRouterMd5OspfChildMd5MemberKey(o["key"], d, "key", sv)); err != nil {
		if !fortiAPIPatch(o["key"]) {
			return fmt.Errorf("Error reading key: %v", err)
		}
	}

	return nil
}

func expandRouterMd5OspfChildMd5MemberMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterMd5OspfChildMd5MemberKey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterMd5OspfChildMd5Member(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandRouterMd5OspfChildMd5MemberMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	if v, ok := d.GetOk("key"); ok {
		t, err := expandRouterMd5OspfChildMd5MemberKey(d, v, "key", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["key"] = t
		}
	}

	return &obj, nil
}
