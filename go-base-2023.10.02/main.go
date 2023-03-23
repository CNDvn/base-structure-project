package main

import (
	"fmt"
	"gobase/pkg/helpers"
	"gobase/pkg/integrates"
	"gobase/pkg/routes"
)

func init() {
	if err := helpers.LoadENV(); err != nil {
		panic(err)
	}

	if err := integrates.InitAWS(); err != nil {
		panic(err)
	}
}

func main() {
	routes := routes.Route()
	if err := routes.Run(":" + helpers.GetENV().API_PORT); err != nil {
		fmt.Println("run routes back end service error: ", err)
		return
	}

}
