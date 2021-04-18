package main

import (
	database "3network-backend/src/model"
	//_ "3network-backend/src/router"
)

func main() {
	defer database.Disconnect()
	database.GetSituationHandleInfo()
}
