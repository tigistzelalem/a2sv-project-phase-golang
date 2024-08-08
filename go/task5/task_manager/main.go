package main

import (
	"task_manager/data"
	"task_manager/router"
)

func main() {
	data.INitMongoDB()
	r := router.SetUpRouter()
	r.Run(":8080")
}
