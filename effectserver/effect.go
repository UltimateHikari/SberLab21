package main

import (
	"fmt"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"github.com/emicklei/go-restful"
	"github.com/google/uuid"
)

const (
	port       = ":80"
	out_suffix = "_out.jpg"
)

var logger *log.Logger

type EffectResource struct{}
type HealthResource struct{}

func process(imageUuid string) {
	img, err := imgio.Open(imageUuid + ".jpg")
	if err != nil {
		fmt.Println(err)
		return
	}

	inverted := effect.Invert(img)
	resized := transform.Resize(inverted, 800, 800, transform.Linear)
	rotated := transform.Rotate(resized, float64(rand.Intn(360)), nil)

	if err := imgio.Save(imageUuid+out_suffix, rotated, imgio.JPEGEncoder(100)); err != nil {
		fmt.Println(err)
		return
	}
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

func (p EffectResource) RegisterTo(container *restful.Container) {
	ws := new(restful.WebService)
	ws.Path("/effect")
	ws.Consumes(restful.MIME_JSON)
	ws.Produces("image/jpeg")

	ws.Route(ws.GET("/random").To(p.randomEffect).
		Doc("get all metadata"))
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
	t := EffectResource{}
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

func (p *EffectResource) randomEffect(req *restful.Request, resp *restful.Response) {
	id := req.PathParameter("id")
	logger.Print("req for effect for " + id)
	imageUuid := uuid.NewString()
	logger.Print("spawning " + imageUuid + " guy")

	fetchImage(id, imageUuid)
	process(imageUuid)

	http.ServeFile(resp.ResponseWriter, req.Request, imageUuid+out_suffix)
}

func fetchImage(id string, imageUuid string) {
	original, err := os.Open("input.jpg")
	if err != nil {
		logger.Print(err)
	}
	defer original.Close()

	new, err := os.Create(imageUuid + ".jpg")
	if err != nil {
		logger.Print(err)
	}
	defer new.Close()

	//This will copy
	_, ok := io.Copy(new, original)
	if ok != nil {
		logger.Print(err)
	}
}

func (p *HealthResource) returnOK(req *restful.Request, resp *restful.Response) {
	logger.Print("healthcheck")
	resp.WriteEntity("OK")
}
