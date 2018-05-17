package controller

import (
	"sharehouse1/library/context"
)

type Api struct{
	context.Ctx
	username string
	password string
}


