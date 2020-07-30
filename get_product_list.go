// Copyright 2020 go-myjvn authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package myjvn

import (
	"context"
	"encoding/xml"
	"net/url"
)

// ParamsGetProductList specifies the parameters of a HTTP request for GetProductList.
type ParamsGetProductList struct {
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

// ProductList stores the data from API response.
type ProductList struct {
	XMLName    xml.Name   `xml:"Result"`
	VendorInfo VendorInfo `xml:"VendorInfo"`
	Status     Status     `xml:"Status"`
}

// NewParamsGetProductList creates an instance of ParamsGetProductList.
func NewParamsGetProductList(params *Parameter) *ParamsGetProductList {
	if params == nil {
		params = &Parameter{}
	}

	p := &ParamsGetProductList{
		Method: "getProductList",
		Feed:   "hnd",
	}

	p.StartItem = params.StartItem
	p.MaxCountItem = params.MaxCountItem
	p.CPEName = params.CPEName
	p.VendorID = params.VendorID
	p.ProductID = params.ProductID
	p.Keyword = url.QueryEscape(params.Keyword)
	p.Language = params.Language

	return p
}

// GetProductList downloads a product list.
// See: https://jvndb.jvn.jp/apis/getProductList_api_hnd.html
func (c *Client) GetProductList(ctx context.Context, params *ParamsGetProductList) (*ProductList, error) {
	if params == nil {
		params = NewParamsGetProductList(nil)
	}

	u, err := addOptions(defaultAPIPath, params)
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
