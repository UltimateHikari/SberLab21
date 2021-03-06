package main

import (
	"fmt"
	"image"
	"io"
	"log"
	"math/rand"
	"net/http"
	"os"
	"strconv"

	"github.com/anthonynsimon/bild/effect"
	"github.com/anthonynsimon/bild/imgio"
	"github.com/anthonynsimon/bild/transform"
	"github.com/emicklei/go-restful"
	"github.com/google/uuid"
)

const (
	port       = ":80"
	in_suffix  = ".jpg"
	out_suffix = "_out.jpg"
	domain     = "http://fileservice/db/"
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

	firstmagic := applyEffect()(img)
	secondmagic := applyEffect()(firstmagic)
	resized := transform.Resize(secondmagic, 800, 800, transform.Linear)
	rotated := transform.Rotate(resized, float64(rand.Intn(360)), nil)

	if err := imgio.Save(imageUuid+out_suffix, rotated, imgio.JPEGEncoder(100)); err != nil {
		fmt.Println(err)
		return
	}
}

func applyEffect() (f func(image image.Image) *image.RGBA) {
	Functions := []func(image.Image) *image.RGBA{
		effect.Invert,
		effect.Emboss,
		effect.Sepia,
		effect.Sharpen,
		effect.Sobel,
		MDilate,
		MEdgeDetection,
		MErode,
		MMedian,
	}
	return Functions[rand.Intn(len(Functions))]
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
	var err error
	logger.Print("req for effect for")
	imageUuid := uuid.NewString()
	logger.Print("spawning " + imageUuid + " guy")

	fetchImage(imageUuid)
	process(imageUuid)

	http.ServeFile(resp.ResponseWriter, req.Request, imageUuid+out_suffix)
	err = os.Remove(imageUuid + in_suffix)
	if err != nil {
		logger.Print(err)
	}
	err = os.Remove(imageUuid + out_suffix)
	if err != nil {
		logger.Print(err)
	}
}

func fetchImage(imageUuid string) {
	photo, err := http.Get(domain + strconv.Itoa(rand.Intn(9)))
	logger.Print(domain + strconv.Itoa(rand.Intn(9)))
	if err != nil {
		logger.Print(err)
		return
	}
	defer photo.Body.Close()

	new, err := os.Create(imageUuid + in_suffix)
	if err != nil {
		logger.Print(err)
	}
	defer new.Close()

	//This will copy
	_, ok := io.Copy(new, photo.Body)
	if ok != nil {
		logger.Print(err)
	}
}

func (p *HealthResource) returnOK(req *restful.Request, resp *restful.Response) {
	resp.WriteEntity("OK")
}
