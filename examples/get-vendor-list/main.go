// Copyright 2020 go-myjvn authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"fmt"
	"net/url"
	"os"
	"strconv"

	"github.com/kazukiigeta/go-myjvn"
)

func main() {
	var (
		keyword = flag.String("keyword", "", "Keyword for vendor search")
	)
	flag.Parse()

	c := myjvn.NewClient(nil)
	params := &myjvn.Parameter{Keyword: url.QueryEscape(*keyword)}
	p := myjvn.NewParamsGetVendorList(params)
	vendorList, err := c.GetVendorList(context.Background(), p)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("---------------------------------------")
	fmt.Println("Result of the getVendorList command")
	fmt.Println("---------------------------------------")

	n := len(vendorList.VendorInfo.Vendors)
	var s string = strconv.Itoa(n)
	switch n {
	case 0:
		fmt.Println("vendorList.VendorInfo.Vendors is nil")
		os.Exit(1)
	case 1:
		s += " vendor is found."
	case 10000:
		s += " vendors are found. (Reached upper limit of displaying)"
	default:
		s += " vendors are found."
	}

	fmt.Printf("%s\n\n", s)
	fmt.Printf("Here is the results.\n\n")
	for _, v := range vendorList.VendorInfo.Vendors {
		fmt.Printf("%+v\n", *v)
	}
	fmt.Println("---------------------------------------")
}
