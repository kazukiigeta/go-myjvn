// Copyright 2020 go-myjvn authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package myjvn

import (
	"context"
	"encoding/xml"
	"net/url"
)

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

// Product stores the data from API response.
type Product struct {
	PName string `xml:"pname,attr"`
	CPE   string `xml:"cpe,attr"`
	PID   string `xml:"pid,attr"`
}

// Vendor stores the data from API response.
type Vendor struct {
	VName    string     `xml:"vname,attr"`
	CPE      string     `xml:"cpe,attr"`
	VID      string     `xml:"vid,attr"`
	Products []*Product `xml:"Product"`
}

// VendorInfo stores the data from API response.
type VendorInfo struct {
	Vendors []*Vendor `xml:"Vendor"`
}

// VendorList stores the data from API response.
type VendorList struct {
	XMLName    xml.Name   `xml:"Result"`
	VendorInfo VendorInfo `xml:"VendorInfo"`
	Status     Status     `xml:"Status"`
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
