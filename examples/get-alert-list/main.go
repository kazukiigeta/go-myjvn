// Copyright 2020 go-myjvn authors. All rights reserved.
// Use of this source code is governed by a MIT-style license that can be
// found in the LICENSE file.

package main

import (
	"context"
	"fmt"

	"github.com/kazukiigeta/go-myjvn"
)

func main() {
	c := myjvn.NewClient(nil)
	params := &myjvn.Parameter{}
	p := myjvn.NewParamsGetAlertList(params)
	alertList, err := c.GetAlertList(context.Background(), p)
	if err != nil {
		fmt.Println(err)
	}

	n := len(alertList.Entries)

	fmt.Println("---------------------------------------")
	fmt.Println(alertList.Title)
	fmt.Println("---------------------------------------")
	fmt.Printf("%d alerts are found.\n\n", n)
	fmt.Printf("Here is the latest one.\n\n")
	fmt.Printf("%+v\n", *alertList.Entries[n-1])
	fmt.Println("---------------------------------------")
}
