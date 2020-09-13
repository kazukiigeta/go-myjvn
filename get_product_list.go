// Copyright 2020 go-myjvn authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package myjvn

import (
	"context"
	"encoding/xml"
)

// PLProduct stores the data from API response.
type PLProduct struct {
	Text  string `xml:",chardata"`
	PName string `xml:"pname,attr"`
	CPE   string `xml:"cpe,attr"`
	PID   string `xml:"pid,attr"`
}

// PLVendor stores the data from API response.
type PLVendor struct {
	Text     string       `xml:",chardata"`
	VName    string       `xml:"vname,attr"`
	CPE      string       `xml:"cpe,attr"`
	VID      string       `xml:"vid,attr"`
	Products []*PLProduct `xml:"Product"`
}

// PLVendorInfo stores the data from API response.
type PLVendorInfo struct {
	Text    string      `xml:",chardata"`
	Lang    string      `xml:"lang,attr"`
	Vendors []*PLVendor `xml:"Vendor"`
}

// ProductList stores the data from API response.
type ProductList struct {
	XMLName        xml.Name     `xml:"Result"`
	Text           string       `xml:",chardata"`
	Version        string       `xml:"version,attr"`
	XSI            string       `xml:"xsi,attr"`
	XMLNS          string       `xml:"xmlns,attr"`
	MJRes          string       `xml:"mjres,attr"`
	AttrStatus     string       `xml:"status,attr"`
	SchemaLocation string       `xml:"schemaLocation,attr"`
	VendorInfo     PLVendorInfo `xml:"VendorInfo"`
	Status         Status       `xml:"Status"`
}

// GetProductList downloads a product list.
// See: https://jvndb.jvn.jp/apis/getProductList_api_hnd.html
func (c *Client) GetProductList(ctx context.Context, opts ...Option) (*ProductList, error) {
	p := &parameter{

		Method: "getProductList",
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

	productList := new(ProductList)
	err = c.do(ctx, req, nil, productList)
	if err != nil {
		return nil, err
	}

	return productList, nil
}
