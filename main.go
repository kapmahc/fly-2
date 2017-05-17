package main

import (
	"log"

	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	_ "github.com/kapmahc/fly/plugins/erp"
	_ "github.com/kapmahc/fly/plugins/forms"
	_ "github.com/kapmahc/fly/plugins/forum"
	_ "github.com/kapmahc/fly/plugins/mall"
	_ "github.com/kapmahc/fly/plugins/ops/mail"
	_ "github.com/kapmahc/fly/plugins/ops/vpn"
	_ "github.com/kapmahc/fly/plugins/pos"
	_ "github.com/kapmahc/fly/plugins/reading"
	_ "github.com/kapmahc/fly/plugins/shop"
	_ "github.com/kapmahc/fly/plugins/site"
	"github.com/kapmahc/fly/web"
)

func main() {
	if err := web.Main(); err != nil {
		log.Fatal(err)
	}
}
