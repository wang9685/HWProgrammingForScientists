package main

import (
	"canvas"
	"image"
)

//place your drawing code here.
func AnimateSystem(simulateSky []Sky, canvasWidth, imageFrequency int) []image.Image {
	// Create list to store generated images
	images := make([]image.Image, 0)
	// Range through simulateSky to draw images
	for i := range simulateSky {
		// Because program set imageFrequency, only draw the specific skys
		if i % imageFrequency == 0 {
			images = append(images, DrawToCanvas((simulateSky[i]), canvasWidth))
		}
	}
	// Return all the images
	return images
}

func DrawToCanvas(sky Sky, canvasWidth int) image.Image {
	// Create canvas and use command line argument to set width
	c := canvas.CreateNewCanvas(canvasWidth, canvasWidth)
	// Black background
	c.SetFillColor(canvas.MakeColor(0, 0, 0))
	c.ClearRect(0, 0, canvasWidth, canvasWidth)
	c.Fill()

	// Range over bodies and draw them to canvas
	for i, b := range sky.boids {
		// Set bodies color
		if i % 3 == 0 {
			c.SetFillColor(canvas.MakeColor(235, 206, 250))
		} else if i % 3 == 1 {
			c.SetFillColor(canvas.MakeColor(240, 240, 0))
		} else {
			c.SetFillColor(canvas.MakeColor(204, 255, 204))
		}
		// Draw boids
		// Position of boid on canvas
		cx := (b.position.x / sky.width) * float64(canvasWidth)
		cy := (b.position.y / sky.width) * float64(canvasWidth)
		// Radius of boid
		r := 5.0
		// Draw boid as circle
		c.Circle(cx, cy, r)
		c.Fill()
	}
	return c.GetImage()
}
