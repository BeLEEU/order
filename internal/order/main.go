package main

import (
	"github.com/BeLEEU/order/common/config"
	"github.com/spf13/viper"
	"log"
)

func init() {
	if err := config.NewViperConfig(); err != nil {
		log.Fatal(err)
	}
}

func main() {

	log.Printf("%v", viper.Get("order"))
}
