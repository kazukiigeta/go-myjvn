# go-myjvn
go-myjvn is a Go client library for accessing the [MyJVN API](https://jvndb.jvn.jp/apis/index.html).

MyJVN API is provided by [IPA](https://www.ipa.go.jp/index-e.html) to offer to create calls to get the data of vulnerabilities in Japanese products. 

[![Golang CI](https://github.com/kazukiigeta/go-myjvn/workflows/Golang%20CI/badge.svg)](https://github.com/kazukiigeta/go-myjvn/actions?query=workflow%3A%22Golang+CI%22)
[![GoDoc](https://pkg.go.dev/badge/github.com/kazukiigeta/go-myjvn?status.svg)](https://pkg.go.dev/github.com/kazukiigeta/go-myjvn)
[![GitHub](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/kazukiigeta/go-myjvn/blob/master/LICENSE)
[![Go Report Card](https://goreportcard.com/badge/github.com/kazukiigeta/go-myjvn)](https://goreportcard.com/report/github.com/kazukiigeta/go-myjvn)

## Usage
```golang
import "github.com/kazukiigeta/go-myjvn"
```

Construct a new MyJVN client, then prepare an instance of parameters for a method which you want to use. You can get a result of calling a method as an instance of JSON struct unmarshalled from a HTTP response.

For example:

```golang
c := myjvn.NewClient(nil)
alertList, err := c.GetAlertList(context.Background(),
	SetKeyword("android"),
	SetRangeDatePublic("n"),
	SetRangeDatePublished("n"),
	SetRangeDateFirstPublished("n"),
)
if err != nil {
	fmt.Println(err)
}
fmt.Println(alertList.Title)
```
## Trying examples
Working examples are available in [examples/ directory](./examples).
You can try them just to execute the following commands.

```sh
# Example of getAlertList
cd examples/get-alert-list
go run main.go

# Example of getVendorList
cd examples/get-vendor-list
go run main.go

# Example of getProductList
cd examples/get-product-list
go run main.go -venorID 4499

# Example of getVulnOverviewList
cd examples/get-vuln-overview-list
go run main.go

# Example of getVulnDetailInfo
cd examples/get-vuln-detail-info
go run main.go -vulnid JVNDB-2020-007528

# Example of getStatics ver HND
cd examples/get-statistics-hnd
go run main.go -cweID CWE-20 -datePublicStartY 2019

# Example of getStatistics ver ITM
cd examples/get-statistics-itm
go run main.go -theme sumCvss -cweID CWE-20 -datePublicStartY 2019
```


## Implemented API
| Version | Method              | Supported | Notes       |
|---------|---------------------|-----------|-------------|
| HND     | getAlertList        | Yes       |             |
|         | getVendorList       | Yes       |             |
|         | getProductList      | Yes       |             |
|         | getVulnOverviewList | Yes       |             |
|         | getVulnDetailInfo   | Yes       |             |
|         | getStatistics       | Yes       |             |
| ITM     | getStatistics       | Yes       |             |
| 3.1     | getOvalList         |           | not planned |
|         | getOvalData         |           | not planned |
|         | getXccdfList        |           | not planned |
|         | getXccdfData        |           | not planned |

## LICENSE

[MIT](https://github.com/kazukiigeta/go-myjvn/blob/master/LICENSE)
