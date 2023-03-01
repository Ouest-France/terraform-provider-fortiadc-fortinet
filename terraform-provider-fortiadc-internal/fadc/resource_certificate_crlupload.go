// Copyright 2022 Fortinet, Inc. All rights reserved.
// Author: Shih-Hsin Huang
// Description: Configure  router static.

package fortiadc

import (
	"bytes"
	"fmt"
	"log"
	"mime/multipart"
	"strconv"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

func resourceCertificateCrlUpload() *schema.Resource {
	return &schema.Resource{
		Read:   resourceCertificateCrlUploadRead,
		Update: resourceCertificateCrlUploadUpdate,
		Create: resourceCertificateCrlUploadCreate,
		Delete: resourceCertificateCrlUploadDelete,

		Importer: &schema.ResourceImporter{
			State: schema.ImportStatePassthrough,
		},

		Schema: map[string]*schema.Schema{
			"cert": &schema.Schema{
				Type:     schema.TypeString,
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
func resourceCertificateCrlUploadCreate(d *schema.ResourceData, m interface{}) error {
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
		return fmt.Errorf("Error describing CertificateCrlUpload: type error")
	}

	obj, err := getObjectCertificateCrlUpload(d, c.Fv)
	if err != nil {
		return fmt.Errorf("Error creating CertificateCrlUpload resource while getting object: %v", err)
	}

	_, err = c.CreateCertificateCrlUpload(obj, vdom)

	if err != nil {
		return fmt.Errorf("Error creating CertificateCrlUpload resource: %v", err)
	}

	d.SetId(mkey)

	//return resourceCertificateCrlUploadRead(d, m)
	return nil
}
func resourceCertificateCrlUploadUpdate(d *schema.ResourceData, m interface{}) error {
	return resourceCertificateCrlUploadRead(d, m)
}
func resourceCertificateCrlUploadDelete(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""
	if v, ok := d.GetOk("vdomparam"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	err := c.DeleteCertificateCrlUpload(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error deleting CertificateCrlUpload resource: %v", err)
	}

	d.SetId("")

	return nil
}
func resourceCertificateCrlUploadRead(d *schema.ResourceData, m interface{}) error {
	mkey := d.Id()

	c := m.(*FortiClient).Client
	c.Retries = 1

	vdomparam := ""

	if v, ok := d.GetOk("vdom"); ok {
		if s, ok := v.(string); ok {
			vdomparam = s
		}
	}

	o, err := c.ReadCertificateCrlUpload(mkey, vdomparam)
	if err != nil {
		return fmt.Errorf("Error reading CertificateCrlUpload resource: %v", err)
	}

	if o == nil {
		log.Printf("[WARN] resource (%s) not found, removing from state", d.Id())
		d.SetId("")
		return nil
	}

	err = refreshObjectCertificateCrlUpload(d, o, c.Fv)
	if err != nil {
		return fmt.Errorf("Error reading CertificateCrlUpload resource from API: %v", err)
	}
	return nil
}

func flattenCertificateCrlUploadMkey(v interface{}, d *schema.ResourceData, pre string, sv string) interface{} {
	return v
}

func refreshObjectCertificateCrlUpload(d *schema.ResourceData, o map[string]interface{}, sv string) error {
	var err error

	/*if err = d.Set("subject", flattenCertificateCrlUploadSubject(o["subject"], d, "subject", sv)); err != nil {
		if !fortiAPIPatch(o["subject"]) {
			return fmt.Errorf("Error reading subject: %v", err)
		}
	}

	if err = d.Set("validfrom", flattenCertificateCrlUploadValidfrom(o["validfrom"], d, "validfrom", sv)); err != nil {
		if !fortiAPIPatch(o["validfrom"]) {
			return fmt.Errorf("Error reading validfrom: %v", err)
		}
	}

	if err = d.Set("validto", flattenCertificateCrlUploadValidto(o["validto"], d, "validto", sv)); err != nil {
		if !fortiAPIPatch(o["validto"]) {
			return fmt.Errorf("Error reading validto: %v", err)
		}
	}

	if err = d.Set("type", flattenCertificateCrlUploadType(o["type"], d, "type", sv)); err != nil {
		if !fortiAPIPatch(o["type"]) {
			return fmt.Errorf("Error reading type: %v", err)
		}
	}

	if err = d.Set("ca_type", flattenCertificateCrlUploadiCAtype(o["ca_type"], d, "ca_type", sv)); err != nil {
		if !fortiAPIPatch(o["ca_type"]) {
			return fmt.Errorf("Error reading ca_type: %v", err)
		}
	}

	if err = d.Set("issuer", flattenCertificateCrlUploadIssuer(o["issuer"], d, "issuer", sv)); err != nil {
		if !fortiAPIPatch(o["issuer"]) {
			return fmt.Errorf("Error reading issuer: %v", err)
		}
	}

	if err = d.Set("fingerprint", flattenCertificateCrlUploadFingerprint(o["fingerprint"], d, "fingerprint", sv)); err != nil {
		if !fortiAPIPatch(o["fingerprint"]) {
			return fmt.Errorf("Error reading fingerprint: %v", err)
		}
	}*/

	if err = d.Set("mkey", flattenCertificateCrlUploadMkey(o["mkey"], d, "mkey", sv)); err != nil {
		if !fortiAPIPatch(o["mkey"]) {
			return fmt.Errorf("Error reading mkey: %v", err)
		}
	}

	return nil
}

func expandCertificateCrlUploadCert(d *schema.ResourceData, v interface{}, pre string, sv string, writer *multipart.Writer) (interface{}, error) {
	return v, nil
}

func expandCertificateCrlUploadMkey(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func expandCertificateCrlUploadVdom(d *schema.ResourceData, v interface{}, pre string, sv string) (interface{}, error) {
	return v, nil
}

func getObjectCertificateCrlUpload(d *schema.ResourceData, sv string) (*map[string]interface{}, error) {
	body := &bytes.Buffer{}
	obj := make(map[string]interface{})
	writer := multipart.NewWriter(body)

	if v, ok := d.GetOk("cert"); ok {
		t, err := expandCertificateCrlUploadCert(d, v, "cert", sv, writer)
		if err != nil {
			return &obj, err
		} else if t != nil {
			err = set_multipart_parameter(writer, "cert", t.(string), true)
			if err != nil {
				writer.Close()
				return &obj, err
			}
		}
	}

	if v, ok := d.GetOk("mkey"); ok {
		t, err := expandCertificateCrlUploadMkey(d, v, "mkey", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			err = set_multipart_parameter(writer, "mkey", t.(string), false)
			if err != nil {
				writer.Close()
				return &obj, err
			}
		}
	}

	if v, ok := d.GetOk("vdom"); ok {
		t, err := expandCertificateCrlUploadVdom(d, v, "vdom", sv)
		if err != nil {
			return &obj, err
		} else if t != nil {
			err = set_multipart_parameter(writer, "vdom", t.(string), false)
			if err != nil {
				writer.Close()
				return &obj, err
			}
		}
	}
	writer.Close()
	obj["content_type"] = writer.FormDataContentType()
	obj["data"] = body.Bytes()

	return &obj, nil
}
