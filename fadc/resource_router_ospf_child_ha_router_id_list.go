// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router ospf child ha router id list.

package fortiadc

import (
	"fmt"
	"log"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceRouterOspfChildHaRouterIdList() *schema.Resource {
	return &schema.Resource{
		Read:   resourceRouterOspfChildHaRouterIdListRead,
		Update: resourceRouterOspfChildHaRouterIdListUpdate,
		Create: resourceRouterOspfChildHaRouterIdListCreate,
		Delete: resourceRouterOspfChildHaRouterIdListDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"node": &schema.Schema{
				Type:     schema.TypeString,
				Computed: true,
				Optional: true,
			},
			"router_id": &schema.Schema{
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
func resourceRouterOspfChildHaRouterIdListCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing RouterOspfChildHaRouterIdList: type error")
	}

	obj, err := getObjectRouterOspfChildHaRouterIdList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspfChildHaRouterIdList resource while getting object: %v", err)
	}

	_, err = c.CreateRouterOspfChildHaRouterIdList(obj, vdom)
	if err != nil {
		return fmt.Errorf("Error creating RouterOspfChildHaRouterIdList resource: %v", err)
	}

	d.SetId(mkey)

	return resourceRouterOspfChildHaRouterIdListRead(d, m)
}
func resourceRouterOspfChildHaRouterIdListUpdate(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()
	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	obj, err := getObjectRouterOspfChildHaRouterIdList(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspfChildHaRouterIdList resource while getting object: %v", err)
	}

	_, err = c.UpdateRouterOspfChildHaRouterIdList(obj, mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error updating RouterOspfChildHaRouterIdList resource: %v", err)
	}

	return resourceRouterOspfChildHaRouterIdListRead(d, m)
}
func resourceRouterOspfChildHaRouterIdListDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""
	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	err := c.DeleteRouterOspfChildHaRouterIdList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error deleting RouterOspfChildHaRouterIdList resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceRouterOspfChildHaRouterIdListRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdom := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdom = s
		}
	}

	o, err := c.ReadRouterOspfChildHaRouterIdList(mkey, vdom)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspfChildHaRouterIdList resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectRouterOspfChildHaRouterIdList(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading RouterOspfChildHaRouterIdList resource from API: %v", err)
	}
	return nil
}
func flattenRouterOspfChildHaRouterIdListNode(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfChildHaRouterIdListRouterId(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func flattenRouterOspfChildHaRouterIdListMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectRouterOspfChildHaRouterIdList(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	if err = d.Set("node", flattenRouterOspfChildHaRouterIdListNode(o["node"], d, "node", sv)); err != nil {
		if !fortiAPIPatch(o["node"]) {
			return fmt.Errorf("Error reading node: %v", err)
		}
	}

	if err = d.Set("router_id", flattenRouterOspfChildHaRouterIdListRouterId(o["router_id"], d, "router_id", sv)); err != nil {
		if !fortiAPIPatch(o["router_id"]) {
			return fmt.Errorf("Error reading router_id: %v", err)
		}
	}

	if err = d.Set("mkey", flattenRouterOspfChildHaRouterIdListMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandRouterOspfChildHaRouterIdListNode(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfChildHaRouterIdListRouterId(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandRouterOspfChildHaRouterIdListMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectRouterOspfChildHaRouterIdList(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	obj := make(map[string]interface{})

	if v, ok := d.GetOk("node"); ok {
		t, err := expandRouterOspfChildHaRouterIdListNode(d, v, "node", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["node"] = t
		}
	}

	if v, ok := d.GetOk("router_id"); ok {
		t, err := expandRouterOspfChildHaRouterIdListRouterId(d, v, "router_id", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["router_id"] = t
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandRouterOspfChildHaRouterIdListMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			obj["mkey"] = t
		}
	}

	return &obj, nil
}
