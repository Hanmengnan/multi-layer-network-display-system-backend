package main

import (
	"fmt"
	database "multi-layer-network-display-system-backend/src/model"
	_ "multi-layer-network-display-system-backend/src/router"
)

func main() {
	defer database.Disconnect()
	fmt.Printf("%v", database.GetNodeAreaStatistics())
}
