package main

import (
	"image"

	"github.com/anthonynsimon/bild/effect"
)

func MDilate(image image.Image) *image.RGBA {
	return effect.Dilate(image, 3)
}

func MEdgeDetection(image image.Image) *image.RGBA {
	return effect.EdgeDetection(image, 1.0)
}

func MMedian(image image.Image) *image.RGBA {
	return effect.Median(image, 10.0)
}

func MErode(image image.Image) *image.RGBA {
	return effect.Erode(image, 3)
}
