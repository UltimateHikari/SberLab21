package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	restful "github.com/emicklei/go-restful/v3"
)

const (
	port   = ":80"
	domain = "http://fileservice/db/"
)

var logger *log.Logger

type APIResource struct{}
type HealthResource struct{}

func (p HealthResource) RegisterTo(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/")
	ws.Consumes(restful.MIME_JSON)
	ws.Produces(restful.MIME_JSON)

	ws.Route(ws.GET("health").To(p.returnOK).
		Doc("healthcheck endpoint for kube"))
	container.Add(ws)
}

func (p APIResource) RegisterTo(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/photos")
	ws.Consumes(restful.MIME_JSON, "image/jpeg")
	ws.Produces(restful.MIME_JSON, "image/jpeg")

	ws.Route(ws.GET("/list").To(p.getList).
		Doc("get all metadata"))

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
	f, err := os.OpenFile("text.log",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()

	logger = log.New(f, "prefix", log.LstdFlags)

	wsContainer := restful.NewContainer()
	t := APIResource{}
	t.RegisterTo(wsContainer)
	h := HealthResource{}
	h.RegisterTo(wsContainer)

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

	logger.Print("start listening on localhost" + port)
	logger.Fatal(http.ListenAndServe(port, wsContainer))
}

func CORSFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	resp.AddHeader(restful.HEADER_AccessControlAllowOrigin, "*")
	chain.ProcessFilter(req, resp)
}

func (p *APIResource) getList(req *restful.Request, resp *restful.Response) {

	logger.Print("getting list")
	list, err := http.Get(domain + "list")
	if err != nil {
		logger.Print(err)
		return
	}

	defer list.Body.Close()

	data, err := ioutil.ReadAll(list.Body)
	if err != nil {
		logger.Print(err)
		return
	}

	resp.Header().Set("content-type", restful.MIME_JSON)
	resp.Write(data)
}

func (p *APIResource) getPhoto(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("id")
	//logger.Print("got content-type from client:" + req.HeaderParameter("Content-Type"))
	logger.Print("req for " + id)
	photo, err := http.Get(domain + id)
	if err != nil {
		logger.Print(err)
		return
	}
	//logger.Print("got content-type from fileserver:" + photo.Header.Get("Content-Type"))

	defer photo.Body.Close()

	data, err := ioutil.ReadAll(photo.Body)
	if err != nil {
		logger.Print(err)
		return
	}
	//logger.Print("header was " + strings.Join(resp.Header().Values("Content-Type"), " "))
	resp.Header().Set("content-type", "image/jpeg")
	//logger.Print("header now " + strings.Join(resp.Header().Values("Content-Type"), " "))
	resp.Write(data)
}

func (p *HealthResource) returnOK(req *restful.Request, resp *restful.Response) {
	logger.Print("healthcheck")
	resp.WriteEntity("OK")
}
