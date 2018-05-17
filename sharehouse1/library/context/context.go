package context

import (
	"net/http"
)

type CtxInterface interface{
	construct()
	response()
}

type Ctx struct{
	req *http.Request
}

func(c *Ctx)construct(){

}

func(c *Ctx)response(data []byte){
	
}