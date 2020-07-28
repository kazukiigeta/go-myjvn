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
