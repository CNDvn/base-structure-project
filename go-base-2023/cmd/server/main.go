package main

import (
	"basego/pkg/helpers"
	"basego/pkg/routes"
	"basego/pkg/utils"
	"fmt"
)

func main() {
	if err := helpers.LoadENV(); err != nil {
		fmt.Println("Error::", err)
		return
	}

	// connect mysql
	if db, err := helpers.MysqlConnect(); err != nil {
		fmt.Printf("Connect mysql error::%s\n", err)
		return
	} else {
		utils.MySqlDB = db
		helpers.MysqlAutoMigrate(utils.MySqlDB)
	}

	// initial firebase
	credentialsJSON := helpers.GetENV().FIREBASE_SERVICE_ACCOUNT
	if firebaseInstance, err := helpers.InitialFirebase([]byte(credentialsJSON)); err != nil {
		fmt.Printf("Firebase initial error: %s\n", err)
		return
	} else {
		utils.FirebaseApp = firebaseInstance
	}

	routes := routes.Routes()
	if err := routes.Run(helpers.GetENV().PORT); err != nil {
		fmt.Println("run routes error: ", err)
		return
	}
}
