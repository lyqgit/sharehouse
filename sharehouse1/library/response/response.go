package response

import (
	"encoding/json"
)

func ErrMsg(code int)string{
	
	switch code{
		case 0:
		return "访问正确"


		case 500:
		return "服务器错误"
		
		case 404:
		return "无服务"

		default:
		return "未知错误"
	}

}

type ResponseMsg struct{
	Errcode int
	Msg string
}

func(this *ResponseMsg)Response(errcode int)([]byte,error){
	this.Errcode = errcode
	this.Msg = ErrMsg(errcode)
	return json.Marshal(this)
}


type ResponseDataPage struct{
	Errcode int
	Msg string
	Data interface{}
	Curpage int
	Allpage int
}

func(this *ResponseDataPage)Response(errcode int,data interface{},curpage int,allpage int)([]byte,error){
	this.Errcode = errcode
	this.Msg = ErrMsg(errcode)
	this.Data = data
	this.Curpage = curpage
	this.Allpage = allpage
	return json.Marshal(this)
}

type ResponseData struct{
	Errcode int
	Msg string
	Data interface{}
}

func(this *ResponseData)Response(errcode int,data interface{})([]byte,error){
	this.Errcode = errcode
	this.Msg = ErrMsg(errcode)
	this.Data = data
	return json.Marshal(this)
}
