package router

import (
	"sharehouse1/library/route"
	"sharehouse1/controller"
)

func init(){
	route.AddRoute("/login",controller.Login)
}