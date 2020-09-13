// Copyright 2020 go-myjvn authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package myjvn

import (
	"context"
	"encoding/xml"
	"net/url"
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

// ParamsGetVendorList specifies the parameters of a HTTP request for GetVendorList.
type ParamsGetVendorList struct {
	Method       string `url:"method"`
	Feed         string `url:"feed"`
	StartItem    uint   `url:"startItem,omitempty"`
	MaxCountItem uint8  `url:"maxCountItem,omitempty"`
	CPEName      string `url:"cpeName,omitempty"`
	VendorID     string `url:"vendorId,omitempty"`
	ProductID    string `url:"productId,omitempty"`
	Keyword      string `url:"keyword,omitempty"`
	Language     string `url:"lang,omitempty"`
}

// NewParamsGetVendorList creates an instance of ParamsGetVendorList.
func NewParamsGetVendorList(params *Parameter) *ParamsGetVendorList {
	if params == nil {
		params = &Parameter{}
	}

	p := &ParamsGetVendorList{
		Method: "getVendorList",
		Feed:   "hnd",
	}

	p.StartItem = params.StartItem
	p.MaxCountItem = params.MaxCountItem
	p.CPEName = params.CPEName
	p.Keyword = url.QueryEscape(params.Keyword)
	p.Language = params.Language

	return p
}

// GetVendorList downloads a vendor list.
// See: https://jvndb.jvn.jp/apis/getVendorList_api_hnd.html
func (c *Client) GetVendorList(ctx context.Context, params *ParamsGetVendorList) (*VendorList, error) {
	if params == nil {
		params = NewParamsGetVendorList(nil)
	}

	u, err := addOptions(defaultAPIPath, params)
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
