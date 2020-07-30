// Copyright 2020 go-myjvn authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"flag"
	"fmt"
	"os"
	"strconv"

	"github.com/kazukiigeta/go-myjvn"
)

func main() {
	var (
		vendorid = flag.String("vendorid", "", "Vendor ID for listing products")
	)
	flag.Parse()
	if *vendorid == "" {
		fmt.Println("vendorid must be specified.")
		os.Exit(1)
	} else if _, err := strconv.Atoi(*vendorid); err != nil {
		fmt.Printf("Invalid vendor ID: %s\n", err)
		os.Exit(1)
	}

	c := myjvn.NewClient(nil)
	params := &myjvn.Parameter{VendorID: *vendorid}
	p := myjvn.NewParamsGetProductList(params)
	productList, err := c.GetProductList(context.Background(), p)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("---------------------------------------")
	fmt.Println("Result of the getProductList command")
	fmt.Println("---------------------------------------")

	n := len(productList.VendorInfo.Vendors[0].Products)
	var s string = strconv.Itoa(n)
	switch n {
	case 0:
		fmt.Println("Can't find any products")
		os.Exit(1)
	case 1:
		s += " product is found."
	default:
		s += " products are found."
	}

	fmt.Printf("%s\n\n", s)
	fmt.Printf("Here is the results.\n\n")
	for _, v := range productList.VendorInfo.Vendors[0].Products {
		fmt.Printf("%+v\n", *v)
	}
	fmt.Println("---------------------------------------")
}
