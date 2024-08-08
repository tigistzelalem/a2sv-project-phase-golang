package main

import (
	"task_manager/router"
)

func main() {
	r := router.SetUpRouter()
	r.Run(":8080")
}
