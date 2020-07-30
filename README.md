# go-myjvn
go-myjvn is a Go client library for accessing the [MyJVN API](https://jvndb.jvn.jp/apis/index.html).

MyJVN API is provided by [IPA](https://www.ipa.go.jp/index-e.html) to offer to create calls to get the data of vulnerabilities in Japanese products. 

[![Golang CI](https://github.com/kazukiigeta/go-myjvn/workflows/Golang%20CI/badge.svg)](https://github.com/kazukiigeta/go-myjvn/actions?query=workflow%3A%22Golang+CI%22)
[![GoDoc](https://godoc.org/github.com/kazukiigeta/go-myjvn?status.svg)](https://godoc.org/github.com/kazukiigeta/go-myjvn)
[![GitHub](https://img.shields.io/github/license/mashape/apistatus.svg)](https://github.com/kazukiigeta/go-myjvn/blob/master/LICENSE)
[![Maintainability](https://api.codeclimate.com/v1/badges/4ad5c3ade8eb39cec428/maintainability)](https://codeclimate.com/github/kazukiigeta/go-myjvn/maintainability)

## Usage
```golang
import "github.com/kazukiigeta/go-myjvn"
```

Construct a new MyJVN client, then prepare an instance of parameters for a method which you want to use. You can get a result of calling a method as an instance of JSON struct unmarshalled from a HTTP response.

For example:

```golang
c := myjvn.NewClient(nil)
params := &myjvn.Parameter{}
p := myjvn.NewParamsGetAlertList(params)
alertList, err := c.GetAlertList(context.Background(), p)
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
cd examples/get-alert-list
go run main.go
```


## Implemented API
| Version | Method              | Supported | Notes       |
|---------|---------------------|-----------|-------------|
| HND     | getAlertList        | Yes       |             |
|         | getVendorList       | Yes       |             |
|         | getProductList      | Yes       |             |
|         | getVulnOverviewList |           |             |
|         | getVulnDetailInfo   |           |             |
|         | getStatistics       |           |             |
| ITM     | getStatistics       |           |             |
| 3.1     | getOvalList         |           | not planned |
|         | getOvalData         |           | not planned |
|         | getXccdfList        |           | not planned |
|         | getXccdfData        |           | not planned |

## LICENSE

[MIT](https://github.com/kazukiigeta/go-myjvn/blob/master/LICENSE)
