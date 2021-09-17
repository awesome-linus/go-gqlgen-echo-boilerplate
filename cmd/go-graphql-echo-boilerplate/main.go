package main

import (
	// 注意: MySQL用ドライバは削除すると接続できなくなるので注意
	"github.com/awesome-linus/go-graphql-echo-boilerplate/pkg/external"
	"github.com/awesome-linus/go-graphql-echo-boilerplate/pkg/external/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	defer mysql.Close()

	external.StartServer()
}
