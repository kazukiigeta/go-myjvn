// Copyright 2020 go-myjvn authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package myjvn

import (
	"context"
	"encoding/xml"
)

// Title stores the data from API response.
type Title struct {
	Text     string `xml:",chardata"`
	Language string `xml:"lang,attr"`
}

// ResDataTotal stores the data from API response.
type ResDataTotal struct {
	VulInfo string `xml:"vulinfo,attr"`
	Vendor  string `xml:"vendor,attr"`
	Product string `xml:"product,attr"`
}

// ResData stores the data from API response.
type ResData struct {
	Date   string `xml:"date,attr"`
	CntAll string `xml:"cntAll,attr"`
	CntC   string `xml:"cntC,attr"`
	CntH   string `xml:"cntH,attr"`
	CntM   string `xml:"cntM,attr"`
	CntL   string `xml:"cntL,attr"`
	CntN   string `xml:"cntN,attr"`
}

// SumJVNDB stores the data from API response.
type SumJVNDB struct {
	Title        []*Title     `xml:"title"`
	ResDataTotal ResDataTotal `xml:"resDataTotal"`
	ResData      []*ResData   `xml:"resData"`
}

// SumCVSS stores the data from API response.
type SumCVSS struct {
	Title        []*Title     `xml:"title"`
	ResDataTotal ResDataTotal `xml:"resDataTotal"`
	ResData      []*ResData   `xml:"resData"`
}

// SumCWE stores the data from API response.
type SumCWE struct {
	Title        []*Title     `xml:"title"`
	ResDataTotal ResDataTotal `xml:"resDataTotal"`
	ResData      []*ResData   `xml:"resData"`
}

// StatisticsHND stores the data from API response.
type StatisticsHND struct {
	XMLName  xml.Name `xml:"Result"`
	SumJVNDB SumJVNDB `xml:"sumJvnDb"`
	SumCVSS  SumCVSS  `xml:"sumCvss"`
	SumCWE   SumCWE   `xml:"sumCwe"`
	Status   Status   `xml:"Status"`
}

// ParamsGetStatisticsHND specifies the parameters of a HTTP request for GetStatisticsHND.
type ParamsGetStatisticsHND struct {
	Method           string `url:"method"`
	Feed             string `url:"feed"`
	Theme            string `url:"theme"`
	DataType         string `url:"type,omitempty"`
	CWEID            string `url:"cweId"`
	PID              uint   `url:"pid,omitempty"`
	CPEName          string `url:"cpeName,omitempty"`
	DatePublicStartY uint16 `url:"datePublicStartY,omitempty"`
	DatePublicStartM uint8  `url:"datePublicStartM,omitempty"`
	DatePublicEndY   uint16 `url:"datePublicEndY,omitempty"`
	DatePublicEndM   uint8  `url:"datePublicEndM,omitempty"`
}

// NewParamsGetStatisticsHND creates an instance of ParamsGetStatisticsHND.
func NewParamsGetStatisticsHND(params *Parameter) *ParamsGetStatisticsHND {
	if params == nil {
		params = &Parameter{}
	}

	p := &ParamsGetStatisticsHND{
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

// GetStatisticsHND downloads a product list.
// See: https://jvndb.jvn.jp/apis/getStatistics_api_hnd.html
func (c *Client) GetStatisticsHND(ctx context.Context, params *ParamsGetStatisticsHND) (*StatisticsHND, error) {
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

	statisticsHND := new(StatisticsHND)
	err = c.do(ctx, req, nil, statisticsHND)
	if err != nil {
		return nil, err
	}

	return statisticsHND, nil
}
