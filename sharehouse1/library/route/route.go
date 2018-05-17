package route

import (
	"net/http"
	"sharehouse1/library/response"
	"sharehouse1/library/context"
	"log"
)

func init(){
	Router = new(Mux)
}

type controller func(c context.CtxInterface)


type routeInfo struct{
	path string
	controllerFunc controller
	c context.CtxInterface
	data []byte
}

func AddRoute(path string,f controller,c context.CtxInterface){
	var temp = routeInfo{path,f,c,[]byte("")}
	Router.route = append(Router.route,temp)

}

type Mux struct{
	route []routeInfo
	routeStatus bool
}

var Router *Mux

func(m *Mux)ServeHTTP(w http.ResponseWriter,r *http.Request){
	m.routeStatus = false
	log.Println(r.URL.Path+"1")
	defer func(){
		if !m.routeStatus{
			responseErr := new(response.ResponseMsg)
			resbyteErr,_:= responseErr.Response(500)
			w.Write(resbyteErr)
		}
	}()
	for _,v := range m.route{
		if v.path == r.URL.Path{
			
			v.controllerFunc(v.c)
			w.Write(v.data)
			log.Println(v.path+"2")
			m.routeStatus = true
		}
	}

	if !m.routeStatus {
		for _,v := range m.route{
			if v.path == "*"{
				v.controllerFunc(v.c)
				w.Write(v.data)
				m.routeStatus = true
			}
		}

		if !m.routeStatus{
			response := new(response.ResponseMsg)
			resbyte,err := response.Response(404)
			if err != nil{
				panic(err)
			}
			w.Write(resbyte)
			m.routeStatus = true
		}
	}
}