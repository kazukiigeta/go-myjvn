// Copyright 2020 go-myjvn authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package myjvn

import (
	"context"
	"encoding/xml"
)

// SResData stores the data from API response.
type SResData struct {
	Text   string `xml:",chardata"`
	Date   string `xml:"date,attr"`
	CntAll string `xml:"cntAll,attr"`
	CntC   string `xml:"cntC,attr"`
	CntH   string `xml:"cntH,attr"`
	CntM   string `xml:"cntM,attr"`
	CntL   string `xml:"cntL,attr"`
	CntN   string `xml:"cntN,attr"`
}

// SResDataTotal stores the data from API response.
type SResDataTotal struct {
	Text    string `xml:",chardata"`
	VulInfo string `xml:"vulinfo,attr"`
	Vendor  string `xml:"vendor,attr"`
	Product string `xml:"product,attr"`
}

// STitle stores the data from API response.
type STitle struct {
	Text string `xml:",chardata"`
	Lang string `xml:"lang,attr"`
}

// SSumJVNDB stores the data from API response.
type SSumJVNDB struct {
	Text         string        `xml:",chardata"`
	Titles       []*STitle     `xml:"title"`
	ResDataTotal SResDataTotal `xml:"resDataTotal"`
	ResData      []*SResData   `xml:"resData"`
}

// SSumCVSS stores the data from API response.
type SSumCVSS struct {
	Text         string        `xml:",chardata"`
	Titles       []*STitle     `xml:"title"`
	ResDataTotal SResDataTotal `xml:"resDataTotal"`
	ResData      []*SResData   `xml:"resData"`
}

// SSumCWE stores the data from API response.
type SSumCWE struct {
	Text         string        `xml:",chardata"`
	CWEID        string        `xml:"cweId,attr"`
	Titles       []*STitle     `xml:"title"`
	ResDataTotal SResDataTotal `xml:"resDataTotal"`
	ResData      []*SResData   `xml:"resData"`
}

// Statistics stores the data from API response.
type Statistics struct {
	XMLName        xml.Name  `xml:"Result"`
	Text           string    `xml:",chardata"`
	Version        string    `xml:"version,attr"`
	XSI            string    `xml:"xsi,attr"`
	XMLNS          string    `xml:"xmlns,attr"`
	MJRes          string    `xml:"mjres,attr"`
	MJStat         string    `xml:"mjstat,attr"`
	AttrStatus     string    `xml:"status,attr"`
	SchemaLocation string    `xml:"schemaLocation,attr"`
	SumJVNDB       SSumJVNDB `xml:"sumJvnDb"`
	SumCVSS        SSumCVSS  `xml:"sumCvss"`
	SumCWE         SSumCWE   `xml:"sumCwe"`
	Status         Status    `xml:"Status"`
}

// ParamsGetStatistics specifies the parameters of a HTTP request for GetStatistics.
type ParamsGetStatistics struct {
	Method           string `url:"method"`
	Feed             string `url:"feed"`
	Theme            string `url:"theme"`
	DataType         string `url:"type,omitempty"`
	CWEID            string `url:"cweId,omitempy"`
	PID              uint   `url:"pid,omitempty"`
	CPEName          string `url:"cpeName,omitempty"`
	DatePublicStartY uint16 `url:"datePublicStartY,omitempty"`
	DatePublicStartM uint8  `url:"datePublicStartM,omitempty"`
	DatePublicEndY   uint16 `url:"datePublicEndY,omitempty"`
	DatePublicEndM   uint8  `url:"datePublicEndM,omitempty"`
}

// NewParamsGetStatisticsHND creates an instance of ParamsGetStatistics.
func NewParamsGetStatisticsHND(params *Parameter) *ParamsGetStatistics {
	if params == nil {
		params = &Parameter{}
	}

	p := &ParamsGetStatistics{
		Method: "getStatistics",
		Feed:   "hnd",
	}

	p.Theme = params.Theme
	p.DataType = params.DataType
	p.CWEID = params.CWEID
	p.PID = params.PID
	p.CPEName = params.CPEName
	p.DatePublicStartY = params.DatePublicStartY
	p.DatePublicStartM = params.DatePublicStartM
	p.DatePublicEndY = params.DatePublicEndY
	p.DatePublicEndM = params.DatePublicEndM

	return p
}

// NewParamsGetStatisticsITM creates an instance of ParamsGetStatistics.
func NewParamsGetStatisticsITM(params *Parameter) *ParamsGetStatistics {
	if params == nil {
		params = &Parameter{}
	}

	p := &ParamsGetStatistics{
		Method: "getStatistics",
		Feed:   "itm",
	}

	p.Theme = params.Theme
	p.DataType = params.DataType
	p.CWEID = params.CWEID
	p.PID = params.PID
	p.CPEName = params.CPEName
	p.DatePublicStartY = params.DatePublicStartY
	p.DatePublicStartM = params.DatePublicStartM
	p.DatePublicEndY = params.DatePublicEndY
	p.DatePublicEndM = params.DatePublicEndM

	return p
}

// GetStatistics downloads a product list.
// See: https://jvndb.jvn.jp/apis/getStatistics_api_hnd.html
// See: https://jvndb.jvn.jp/apis/getStatistics_api_itm.html
func (c *Client) GetStatistics(ctx context.Context, params *ParamsGetStatistics) (*Statistics, error) {
	if params == nil {
		params = NewParamsGetStatisticsHND(nil)
	}

	u, err := addOptions(defaultAPIPath, params)
	if err != nil {
		return nil, err
	}
	req, err := c.newRequest("GET", u)
	if err != nil {
		return nil, err
	}

	statisticsHND := new(Statistics)
	err = c.do(ctx, req, nil, statisticsHND)
	if err != nil {
		return nil, err
	}

	return statisticsHND, nil
}
