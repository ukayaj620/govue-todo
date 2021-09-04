package main

import (
	"github.com/ukayaj620/govue-todo/api/db"
	"github.com/ukayaj620/govue-todo/api/routes"
)

func main() {
	db.Init()
	routes.InitRouter()
}
