package main

import (
	"sharehouse1/library"
	_ "sharehouse1/router"
)

func main(){

	library.Run(":8080")
}