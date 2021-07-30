package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"

	restful "github.com/emicklei/go-restful/v3"
)

const (
	port           = ":80"
	prefix_debug   = "./resources/thumbnails/"
	database_debug = "./resources/photos.json"
	prefix         = "./thumbnails/"
	database       = "./photos.json"
)

var logger *log.Logger

type APIResource struct{}
type HealthResource struct{}

type Source struct {
	Name string
	Url  string
}

type Photographer struct {
	Name string
	Url  string
}

type PhotoEntry struct {
	Id           int
	Title        string
	Filename     string
	Location     string
	Source       Source       `json:"Source"`
	Photographer Photographer `json:"Photographer"`
}

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
	ws.Path("/db")
	ws.Consumes(restful.MIME_JSON)
	ws.Produces("image/jpeg")

	ws.Route(ws.GET("/{id}").To(p.getPhoto).
		Doc("get the photo by its id").
		Param(ws.PathParameter("id", "identifier of photo").DataType("integer")))

	ws.Route(ws.GET("/list").To(p.getList).
		Doc("get all metadata"))

	ws.Route(ws.POST("/upload").To(p.addPhoto).
		Doc("create a photo"))

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

	logger.Print("start listening on localhost:" + port)
	logger.Fatal(http.ListenAndServe(port, wsContainer))
}

func CORSFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	resp.AddHeader(restful.HEADER_AccessControlAllowOrigin, "*")
	chain.ProcessFilter(req, resp)
}
func (p *APIResource) getList(req *restful.Request, resp *restful.Response) {
	http.ServeFile(resp.ResponseWriter, req.Request, database)
}

func (p *APIResource) getPhoto(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("id")
	logger.Print("req for " + id)
	data, err := ioutil.ReadFile(database)
	if err != nil {
		logger.Print(err)
		return
	}
	var photos []PhotoEntry
	err = json.Unmarshal(data, &photos)
	if err != nil {
		logger.Print(err)
		return
	}
	idInt, _ := strconv.Atoi(id)
	location := prefix + photos[idInt].Filename
	logger.Print(location)
	http.ServeFile(resp.ResponseWriter, req.Request, location)
}

func (p *APIResource) addPhoto(req *restful.Request, resp *restful.Response) {
	logger.Print("got photo")
	r := req.Request
	err := r.ParseForm()
	if err != nil {
		logger.Print(err)
		return
	}
	populateDb(r.PostForm.Get("name"), r.PostForm.Get("description"))
	//r.PostForm.Get("file")
}

func populateDb(filename string, description string) {
	logger.Print("populating..")
	var db []PhotoEntry

	data, err := ioutil.ReadFile(database)
	if err != nil {
		log.Print(err)
		return
	}
	err = json.Unmarshal([]byte(data), &db)
	if err != nil {
		log.Print(err)
		return
	}
	db = append(db, PhotoEntry{Id: len(db), Filename: filename, Location: description})
	logger.Print("populated")
	result, err := json.Marshal(db)
	if err != nil {
		log.Print(err)
		return
	}
	f, err := os.Open(database)
	if err != nil {
		log.Print(err)
		return
	}
	defer f.Close()

	// write bytes to the file
	_, err = f.Write(result)
	if err != nil {
		log.Print(err)
		return
	}
	logger.Print("wrote")
}

func (p *HealthResource) returnOK(req *restful.Request, resp *restful.Response) {
	resp.WriteEntity("OK")
}
