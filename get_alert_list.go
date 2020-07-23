// Copyright 2020 go-myjvn authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package myjvn

import (
	"context"
	"net/http"
)

type params struct {
	Method             string `url:"method"`
	Feed               string `url:"feed"`
	StartItem          uint   `url:"startItem,omitempty"`
	MaxCountItem       uint8  `url:"maxCountItem,omitempty"`
	DatePublished      uint16 `url:"datePublished,omitempty"`
	DateFirstPublished uint16 `url:"dateFirstPublished,omitempty"`
	CpeName            string `url:"cpeName,omitempty"`
	Format             string `url:"ft,omitempty"`
}

// ParamsGetAlertList specifies the parameters of a HTTP request for GetAlertList.
type ParamsGetAlertList params

// NewParamsGetAlertList creates an instance of ParamsGetAlertList.
func NewParamsGetAlertList(
	startItem *uint, maxCountItem *uint8, datePublished, dateFirstPublished *uint16, cpeName *string) *ParamsGetAlertList {
	p := &ParamsGetAlertList{
		Method: "getAlertList",
		Feed:   "hnd",
		Format: "json",
	}

	if startItem != nil {
		p.StartItem = *startItem
	}

	if maxCountItem != nil {
		p.MaxCountItem = *maxCountItem
	}

	if datePublished != nil {
		p.DatePublished = *datePublished
	}

	if dateFirstPublished != nil {
		p.DateFirstPublished = *dateFirstPublished
	}

	if cpeName != nil {
		p.CpeName = *cpeName
	}

	return p
}

// AlertList represents a JSON object of the alert list.
type AlertList struct {
	Feed Feed
}

// Feed represents a JSON value of thefeed in the AlertList.
type Feed struct {
	Title       string      `json:"title"`
	ID          string      `json:"id"`
	Author      Author      `json:"author"`
	Updated     string      `json:"updated"`
	Link        string      `json:"link"`
	SecHandling SecHandling `json:"sec:handling"`
	Entry       []*Entry    `json:"entry"`
	Status      Status      `json:"status:Status"`
}

// Author represents a JSON value of theauthor in the Feed.
type Author struct {
	Name string `json:"name"`
	URI  string `json:"uri"`
}

// SecHandling represents a JSON value of thesec handling in the Feed.
type SecHandling struct {
	Marking Marking `json:"marking:Marking"`
}

// Marking represents a JSON value of themarking in the SecHandling.
type Marking struct {
	MarkingStructure MarkingStructure `json:"marking:Marking_Structure"`
}

// MarkingStructure represents a JSON value of themarking structure in the Marking.
type MarkingStructure struct {
	XSIType   string `json:"xsi:type"`
	ModelName string `json:"marking_model_name"`
	ModelRef  string `json:"marking_model_ref"`
	Color     string `json:"color"`
}

// Entry represents a JSON value of theentry in the Feed.
type Entry struct {
	Title     string      `json:"title"`
	ID        string      `json:"id"`
	Link      string      `json:"link"`
	Summary   string      `json:"summary"`
	Category  Category    `json:"category"`
	Update    string      `json:"update"`
	Published string      `json:"published"`
	SecItems  []*SecItems `json:"sec:items"`
}

// Category represents a JSON value of the category in the Entry.
type Category struct {
	Term  string `json:"term"`
	Label string `json:"lable"`
}

// SecItems represents a JSON value of the sec items in the Entry.
type SecItems struct {
	SecItem SecItem `json:"sec:item"`
}

// SecItem represents a JSON value of the sec item in the SecItems.
type SecItem struct {
	SecTitle      string    `json:"sec:title"`
	SecIdentifier string    `json:"sec:identifier"`
	SecLink       string    `json:"sec:link"`
	SecPublished  string    `json:"sec:published"`
	SecUpdated    string    `json:"sec:updated"`
	SecAuthor     SecAuthor `json:"sec:author"`
	SecCpe        []*SecCpe `json:"sec:cpe"`
}

// SecCpe represents a JSON value of the sec cpe in the SecItem.
type SecCpe struct {
	Value string `json:"value"`
}

// SecAuthor represents a JSON value of the sec author in the SecItem.
type SecAuthor struct {
	Value string `json:"value"`
}

// Status represents a JSON value of the status in the Feed.
type Status struct {
	Version     string `json:"version"`
	Method      string `json:"method"`
	RetCd       int    `json:"retCd"`
	RetMax      string `json:"retMax"`
	ErrCd       string `json:"errCd"`
	ErrMsg      string `json:"errMsg"`
	TotalRes    string `json:"totalRes"`
	TotalResRet string `json:"totalResRet"`
	FirstRes    string `json:"firstRes"`
	Format      string `json:"ft"`
	Feed        string `json:"feed"`
}

// GetAlertList downloads an alert list.
// See: https://jvndb.jvn.jp/apis/getAlertList_api_hnd.html
func (c *Client) GetAlertList(
	ctx context.Context, params *ParamsGetAlertList) (*AlertList, *http.Response, error) {
	if params == nil {
		params = NewParamsGetAlertList(nil, nil, nil, nil, nil)
	}

	u, err := addOptions(defaultAPIPath, params)
	if err != nil {
		return nil, nil, err
	}
	req, err := c.newRequest("GET", u)
	if err != nil {
		return nil, nil, err
	}

	alertList := new(AlertList)
	resp, err := c.do(ctx, req, alertList)
	if err != nil {
		return nil, nil, err
	}

	return alertList, resp, nil
}
