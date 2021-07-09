package main

import (
	"log"
	"net/http"
	"os"

	restful "github.com/emicklei/go-restful/v3"
)

const (
	port     = ":8000"
	database = "./photos.json"
)

type APIResource struct{}

func (p APIResource) RegisterTo(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/photos")
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON, "image/jpeg")

	ws.Route(ws.GET("list").To(p.getList).
		Doc("get all todos"))

	ws.Route(ws.GET("/{id}").To(p.getPhoto).
		Doc("get the product by its id").
		Param(ws.PathParameter("id", "identifier of the product").DataType("integer")))

	// ws.Route(ws.POST("/").To(p.addToDo).
	// 	Doc("update or create a product").
	// 	Param(ws.BodyParameter("ToDo", "a ToDo (JSON)").DataType("main.ToDo")))

	// ws.Route(ws.PUT("/{id}").To(p.updateTodo).
	// 	Doc("get the product by its id").
	// 	Param(ws.PathParameter("id", "identifier of the product").DataType("integer")).
	// 	Param(ws.BodyParameter("ToDo", "a ToDo (JSON)").DataType("main.ToDo")))

	// ws.Route(ws.DELETE("/{id}").To(p.deleteTodo).
	// 	Doc("update or create a product").
	//Param(ws.PathParameter("id", "identifier of the product").DataType("integer")))

	container.Add(ws)
}

func main() {
	port := ":8000"
	if len(os.Args) == 1 {
		log.Print("no port arg, using default")
	} else {
		port = ":" + os.Args[1]
	}

	wsContainer := restful.NewContainer()
	t := APIResource{}
	t.RegisterTo(wsContainer)

	// Add container filter to enable CORS
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept"},
		AllowedMethods: []string{"PUT", "POST", "GET", "DELETE"},
		AllowedDomains: []string{"*"},
		CookiesAllowed: false,
		Container:      wsContainer}
	wsContainer.Filter(cors.Filter)

	// Add container filter to respond to OPTIONS
	wsContainer.Filter(wsContainer.OPTIONSFilter)
	wsContainer.Filter(CORSFilter)

	log.Print("start listening on localhost:" + port)
	log.Fatal(http.ListenAndServe(port, wsContainer))
}

func CORSFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	resp.AddHeader(restful.HEADER_AccessControlAllowOrigin, "*")
	chain.ProcessFilter(req, resp)
}

func (p *APIResource) getList(req *restful.Request, resp *restful.Response) {
	//dummy, change to database server
	http.ServeFile(resp.ResponseWriter, req.Request, database)
}

func (p *APIResource) getPhoto(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("id")
	log.Print("req for " + id)
}
