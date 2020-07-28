// Copyright 2020 go-myjvn authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package myjvn

import (
	"context"
	"encoding/xml"
)

// Link stores the data from API response.
type Link struct {
	Href string `xml:"href,attr"`
}

// Author stores the data from API response.
type Author struct {
	Name string `xml:"name"`
	URI  string `xml:"uri"`
}

// MarkingStructure stores the data from API response.
type MarkingStructure struct {
	XSIType          string `xml:"type,attr"`
	MarkingModelName string `xml:"marking_model_name,attr"`
	MarkingModelRef  string `xml:"marking_model_ref,attr"`
	Color            string `xml:"color,attr"`
}

// Marking stores the data from API response.
type Marking struct {
	MarkingStructure MarkingStructure `xml:"Marking_Structure"`
}

// SecHandling stores the data from API response.
type SecHandling struct {
	Marking Marking `xml:"Marking"`
}

// Category stores the data from API response.
type Category struct {
	Label string `xml:"label,attr"`
	Term  string `xml:"term,attr"`
}

// Entry stores the data from API response.
type Entry struct {
	Title     string   `xml:"title"`
	ID        string   `xml:"id"`
	Published string   `xml:"published"`
	Updated   string   `xml:"updated"`
	Category  Category `xml:"category"`
}

// Status stores the data from API response.
type Status struct {
	Version     string `xml:"version,attr"`
	Method      string `xml:"method,attr"`
	RetCd       uint   `xml:"retCd,attr"`
	RetMax      string `xml:"retMax,attr"`
	ErrCd       string `xml:"errCd,attr"`
	ErrMsg      string `xml:"errMsg,attr"`
	TotalRes    string `xml:"totalRes,attr"`
	TotalResRet string `xml:"totalResRet,attr"`
	FirstRes    string `xml:"firstRes,attr"`
	Feed        string `xml:"feed,attr"`
}

// AlertList stores the data from API response.
type AlertList struct {
	XMLName     xml.Name    `xml:"feed"`
	Title       string      `xml:"title"`
	Updated     string      `xml:"updated"`
	ID          string      `xml:"id"`
	Link        Link        `xml:"link"`
	Author      Author      `xml:"author"`
	SecHandling SecHandling `xml:"handling"`
	Entries     []*Entry    `xml:"entry"`
	Status      Status      `xml:"Status"`
}

// Parameter represents all the parameters of API except method and feed.
type Parameter struct {
	StartItem               uint
	MaxCountItem            uint8
	DatePublished           uint16
	DateFirstPublished      uint16
	CpeName                 string
	Format                  string
	Keyword                 string
	Language                string
	VendorID                string
	ProductID               string
	Severity                string
	Vector                  string
	RangeDatePublic         string
	RangeDatePublished      string
	RangeDateFirstPublished string
	DatePublicStartY        uint16
	DatePublicStartM        uint8
	DatePublicStartD        uint8
	DatePublicEndY          uint16
	DatePublicEndM          uint8
	DatePublicEndD          uint8
	DateFirstPublicStartY   uint16
	DateFirstPublicStartM   uint8
	DateFirstPublicStartD   uint8
	DateFirstPublicEndY     uint16
	DateFirstPublicEndM     uint8
	DateFirstPublicEndD     uint8
}

// ParamsGetAlertList specifies the parameters of a HTTP request for GetAlertList.
type ParamsGetAlertList struct {
	Method             string `url:"method"`
	Feed               string `url:"feed"`
	StartItem          uint   `url:"startItem,omitempty"`
	MaxCountItem       uint8  `url:"maxCountItem,omitempty"`
	DatePublished      uint16 `url:"datePublished,omitempty"`
	DateFirstPublished uint16 `url:"dateFirstPublished,omitempty"`
	CpeName            string `url:"cpeName,omitempty"`
	Format             string `url:"ft,omitempty"`
}

// NewParamsGetAlertList creates an instance of ParamsGetAlertList.
func NewParamsGetAlertList(params *Parameter) *ParamsGetAlertList {
	if params == nil {
		params = &Parameter{}
	}

	p := &ParamsGetAlertList{
		Method: "getAlertList",
		Feed:   "hnd",
	}

	p.StartItem = params.StartItem
	p.MaxCountItem = params.MaxCountItem
	p.DatePublished = params.DatePublished
	p.DateFirstPublished = params.DateFirstPublished
	p.CpeName = params.CpeName
	p.Format = params.Format

	return p
}

// GetAlertList downloads an alert list.
// See: https://jvndb.jvn.jp/apis/getAlertList_api_hnd.html
func (c *Client) GetAlertList(
	ctx context.Context, params *ParamsGetAlertList) (*AlertList, error) {
	if params == nil {
		params = NewParamsGetAlertList(nil)
	}

	u, err := addOptions(defaultAPIPath, params)
	if err != nil {
		return nil, err
	}
	req, err := c.newRequest("GET", u)
	if err != nil {
		return nil, err
	}

	alertList := new(AlertList)
	err = c.do(ctx, req, nil, alertList)
	if err != nil {
		return nil, err
	}

	return alertList, nil
}
