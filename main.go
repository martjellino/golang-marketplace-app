package main

import (
	"golang-marketplace-app/router"
)

var (
	PORT = ":8000"
)

func main() {
	r := router.StartApp()
	r.Run(PORT)
}
