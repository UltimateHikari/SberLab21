package main

import (
	"io"
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
)

const (
	port     = ":8000"
	database = "./photos.json"
)

func main() {
	ws := new(restful.WebService)

	ws.Route(ws.GET("/random").To(hello))
	ws.Route(ws.GET("/list").To(serveJson))
	restful.Add(ws)

	log.Fatal(http.ListenAndServe(port, nil))
}

func serveJson(req *restful.Request, resp *restful.Response) {
	//dummy, change to database server
	http.ServeFile(resp.ResponseWriter, req.Request, database)
}

func hello(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "world")
}
