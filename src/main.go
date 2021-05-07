package main

import (
	database "3network-backend/src/model"
	"fmt"
	//_ "3network-backend/src/router"
)

func main() {
	defer database.Disconnect()
	fmt.Printf("%v", database.GetNodeAreaStatistics())
}
