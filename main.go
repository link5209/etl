package main

import (
	_ "play/etl/config"
	"play/etl/controllers"
)

func main() {
	etl := new(controllers.ETL)
	etl.Run()
}
