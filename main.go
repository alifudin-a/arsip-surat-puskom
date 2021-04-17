package main

import (
	database "github.com/alifudin-a/arsip-surat-puskom/database/psql"
	"github.com/alifudin-a/arsip-surat-puskom/route"
)

func main() {
	database.OpenDB()
	route.InitRoute()
}
