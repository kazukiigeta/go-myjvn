// Copyright 2020 go-myjvn authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package myjvn_test

import (
	"context"
	"fmt"

	"github.com/kazukiigeta/go-myjvn"
)

func ExampleClient_GetAlertList() {
	c := myjvn.NewClient(nil)
	params := &myjvn.Parameter{}
	p := myjvn.NewParamsGetAlertList(params)
	alertList, err := c.GetAlertList(context.Background(), p)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(alertList.Title)
}

func ExampleClient_GetVendorList() {
	c := myjvn.NewClient(nil)
	params := &myjvn.Parameter{}
	p := myjvn.NewParamsGetVendorList(params)
	vendorList, err := c.GetVendorList(context.Background(), p)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range vendorList.VendorInfo.Vendors {
		fmt.Println(*v)
	}
}

func ExampleClient_GetProductList() {
	c := myjvn.NewClient(nil)
	params := &myjvn.Parameter{}
	p := myjvn.NewParamsGetProductList(params)
	productList, err := c.GetProductList(context.Background(), p)
	if err != nil {
		fmt.Println(err)
	}

	for _, product := range productList.VendorInfo.Vendors[100].Products {
		fmt.Println(*product)
	}
}

func ExampleClient_GetVulnOverviewList() {
	c := myjvn.NewClient(nil)
	params := &myjvn.Parameter{}
	p := myjvn.NewParamsGetVulnOverviewList(params)
	vulnOverviewList, err := c.GetVulnOverviewList(context.Background(), p)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(vulnOverviewList.Item[0].Title)
}

func ExampleClient_GetVulnDetailInfo() {
	c := myjvn.NewClient(nil)
	params := &myjvn.Parameter{
		VulnID: "JVNDB-2020-006469",
	}
	p := myjvn.NewParamsGetVulnDetailInfo(params)
	vulnDetailInfo, err := c.GetVulnDetailInfo(context.Background(), p)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(vulnDetailInfo.VulInfo.VulInfoID)
}

func ExampleClient_GetStatisticsHND() {
	c := myjvn.NewClient(nil)
	params := &myjvn.Parameter{
		Theme:            "SumCvss",
		CWEID:            "CEW-20",
		DatePublicStartY: 2015,
	}
	p := myjvn.NewParamsGetStatisticsHND(params)
	statisticsHND, err := c.GetStatisticsHND(context.Background(), p)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(statisticsHND.SumCVSS.Title)
}

func ExampleClient_GetStatisticsITM() {
	c := myjvn.NewClient(nil)
	params := &myjvn.Parameter{
		Theme:            "sumCvss",
		CWEID:            "CWE-20",
		DatePublicStartY: 2015,
	}
	p := myjvn.NewParamsGetStatisticsITM(params)
	statisticsITM, err := c.GetStatisticsITM(context.Background(), p)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(statisticsITM.SumCVSS.Title)
}
