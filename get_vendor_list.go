// Copyright 2020 go-myjvn authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package myjvn

import (
	"context"
	"encoding/xml"
)

// VLVendor stores the data from API response.
type VLVendor struct {
	Text  string `xml:",chardata"`
	VName string `xml:"vname,attr"`
	CPE   string `xml:"cpe,attr"`
	VID   string `xml:"vid,attr"`
}

// VLVendorInfo stores the data from API response.
type VLVendorInfo struct {
	Text    string      `xml:",chardata"`
	Lang    string      `xml:"lang,attr"`
	Vendors []*VLVendor `xml:"Vendor"`
}

// VendorList stores the data from API response.
type VendorList struct {
	XMLName        xml.Name     `xml:"Result"`
	Text           string       `xml:",chardata"`
	Version        string       `xml:"version,attr"`
	XSI            string       `xml:"xsi,attr"`
	XMLNS          string       `xml:"xmlns,attr"`
	MJRes          string       `xml:"mjres,attr"`
	AttrStatus     string       `xml:"status,attr"`
	SchemaLocation string       `xml:"schemaLocation,attr"`
	VendorInfo     VLVendorInfo `xml:"VendorInfo"`
	Status         Status       `xml:"Status"`
}

// GetVendorList downloads a vendor list.
// See: https://jvndb.jvn.jp/apis/getVendorList_api_hnd.html
func (c *Client) GetVendorList(ctx context.Context, opts ...Option) (*VendorList, error) {
	p := &parameter{

		Method: "getVendorList",
		Feed:   "hnd",
	}

	for _, opt := range opts {
		opt(p)
	}

	u, err := addOptions(defaultAPIPath, p)
	if err != nil {
		return nil, err
	}
	req, err := c.newRequest("GET", u)
	if err != nil {
		return nil, err
	}

	vendorList := new(VendorList)
	err = c.do(ctx, req, nil, vendorList)
	if err != nil {
		return nil, err
	}

	return vendorList, nil
}
