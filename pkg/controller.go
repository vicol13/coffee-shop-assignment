package pkg

import (
	"io"
	"net/http"
	"errors"
	"github.com/julienschmidt/httprouter"
	log "github.com/sirupsen/logrus"
	"fmt"
)


var PATHS = map[string]bool {
	AMERICANO:true,
	CAPPUCCINO:true,
	ESPRESSO :true,
}



//struct which contains logic related to http interaction with system
type CoffeRequestHandler struct {
	Coffeeservice *CoffeeService
}


func (crh *CoffeRequestHandler) GetCoffe(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	w.Header().Add("Content-Type", "application/json")
	
	if _ ,ok := PATHS[params.ByName("name")]; !ok { //wrong path
		w.WriteHeader(400)
		io.WriteString(w, jsonMessage(errors.New("No such route: " + params.ByName("name"))))	
		log.Debug(fmt.Sprintf("Rejected request to  [/coffe/%s]", params.ByName("name") ))
		return 
	}
	
	
	quota,userId,err := crh.loadAndValidateHeaders(w,r) 
	if err != nil { //request wihtou headers
		w.WriteHeader(400)
		io.WriteString(w, jsonMessage(err))	
		return 
	}

	
	if res,err := crh.Coffeeservice.ProcessOrder(quota,userId,params.ByName("name")); err!= nil { //not valid quota
		w.WriteHeader(429) 
		io.WriteString(w, jsonMessage(err))	
		log.Debug(fmt.Sprintf("Rejected request to [%s] as he exceed hist limit of requests", userId))
	}else{
		w.WriteHeader(200)
		io.WriteString(w, fmt.Sprintf("{\"Message\": \"%s\"}",res))	
	}
		
}



// This function is responsible for for validating headers 
func  (h *CoffeRequestHandler) loadAndValidateHeaders(w http.ResponseWriter, r *http.Request) (string,string,error) {
	quota := r.Header.Get("quota")
	if quota == "" {
		log.Debug("Rejected request header [quota] is missing")
		return "","", errors.New("header [quota] is missing or is blank")
	}
	userId := r.Header.Get("user-id")
	if userId == ""  {
		log.Debug("Rejected request header [user-id] is missing")
		return "","", errors.New("header [user-id] is missing or is blank")
	}
	return quota,userId,nil
}



