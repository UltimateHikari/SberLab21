package main

import (
	"io"
	"log"
	"net/http"

	restful "github.com/emicklei/go-restful/v3"
)

func main() {
	ws := new(restful.WebService)
	ws.Route(ws.GET("/random").To(hello))
	restful.Add(ws)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func hello(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "world")
}
