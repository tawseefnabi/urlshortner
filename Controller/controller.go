package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	model "github.com/tawseefnabi/urlshortner/Model"
	service "github.com/tawseefnabi/urlshortner/Service"
)

var (
	defaultUrl string = "http://localhost:8080/home/"
)

type Controller struct {
	ser *service.Service
}

func NewController(ser *service.Service) Controller {
	return Controller{
		ser: ser,
	}
}
func (c *Controller) GenerateTinyUrl(rw http.ResponseWriter, req *http.Request) {
	var url model.UrlModel
	err := json.NewDecoder(req.Body).Decode(&url)
	rw.Header().Set("Content-Type", "application/json")
	if err != nil {
		log.Fatalln(err)
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{"status":"failed","message": "failed to decode"}`))
		return
	}
	resp := c.ser.GenerateTinyUrl(url)
	jsonBody, err := json.Marshal(resp)

	if err != nil {
		log.Println("error: ", err.Error())
		rw.WriteHeader(http.StatusInternalServerError)
		rw.Write([]byte(`{"status":"failed","message": "failed to unmarshal"}`))
		return
	}
	rw.WriteHeader(http.StatusOK)
	rw.Write(jsonBody)

}

func (c *Controller) RedirectTinyUrl(rw http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	hashCode := params["hashCode"]
	url := c.ser.RedirectTinyUrl(hashCode)
	if url != "" {
		log.Println("redirecting to : ", url)
		http.Redirect(rw, req, url, http.StatusSeeOther)
	} else {
		log.Println("url not found: ", params)
		log.Println("redirecting to default url: ", defaultUrl)
		http.Redirect(rw, req, defaultUrl, http.StatusSeeOther)
	}

}
func (c *Controller) HomePage(rw http.ResponseWriter, req *http.Request) {
	rw.Header().Set("Content-Type", "application/json")
	rw.WriteHeader(http.StatusNotFound)
	rw.Write([]byte(`{"status" : "failed","message" : "generate tiny url please"}`))
	return
}
