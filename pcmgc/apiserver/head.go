package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"os"

	restful "github.com/emicklei/go-restful/v3"
)

const (
	port           = ":80"
	domain         = "http://fileservice/db/"
	uploadPath     = "upload"
	effectEndpoint = "http://effectservice/effect/random"
	MIME_IMG       = "image/jpeg"
	MIME_MULTIPART = "multipart/form-data"
	logfile        = "text.log"
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
	ws.Consumes(restful.MIME_JSON, MIME_IMG, MIME_MULTIPART)
	ws.Produces(restful.MIME_JSON, MIME_IMG)

	ws.Route(ws.GET("/list").To(p.getList).
		Doc("get all metadata"))

	ws.Route(ws.GET("/random.jpg").To(p.getRandom).
		Doc("get random mutation"))

	ws.Route(ws.GET("/{id}.jpg").To(p.getPhoto).
		Doc("get the product by its id").
		Param(ws.PathParameter("id", "identifier of the product").DataType("integer")))

	ws.Route(ws.POST("/upload").To(p.addPhoto).
		Doc("create a photo"))

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
	f, err := os.OpenFile(logfile,
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

func (p *APIResource) addPhoto(req *restful.Request, resp *restful.Response) {
	//forward to fileservice a then parse there with bodyparams?
	logger.Print("uploading a photo...")
	defer req.Request.Body.Close()
	dbresp, err := http.Post(domain+uploadPath, MIME_IMG, req.Request.Body)
	if err != nil {
		logger.Print(err)
		return
	}
	data, err := ioutil.ReadAll(dbresp.Body)
	if err != nil {
		logger.Print(err)
		return
	}
	//resp.Header(). = req.Request.Header()
	resp.Write(data)
	logger.Print("upload success!")
}

func (p *APIResource) getPhoto(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("id")
	logger.Print("req for " + id)
	p.forwardPhoto(req, resp, domain+id)
}

func (p *APIResource) getRandom(req *restful.Request, resp *restful.Response) {
	logger.Print("random effect")
	p.forwardPhoto(req, resp, effectEndpoint)
}

func (p *APIResource) forwardPhoto(req *restful.Request, resp *restful.Response, endpoint string) {

	photo, err := http.Get(endpoint)
	if err != nil {
		logger.Print(err)
		return
	}

	defer photo.Body.Close()

	data, err := ioutil.ReadAll(photo.Body)
	if err != nil {
		logger.Print(err)
		return
	}
	resp.Header().Set("content-type", MIME_IMG)
	resp.Write(data)
}

func (p *HealthResource) returnOK(req *restful.Request, resp *restful.Response) {
	resp.WriteEntity("OK")
}
