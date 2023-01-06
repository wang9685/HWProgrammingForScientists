package main

import (
	"fmt"
	"gifhelper"
	"math"
	"math/rand"
	"os"
	"strconv"
	"time"
)
func main() {
	// Take command line argument
	// os.Args[1] is num of num of boids
	numBoids, err1 := strconv.Atoi(os.Args[1])
	if err1!= nil {
        fmt.Println(err1)
	}
	// Number of boids can't be negative
	if numBoids < 0 {
		panic("numBoids must be greater or equal to 0")
	}
	// os.Args[2] is sky width
	skyWidth, err2 := strconv.ParseFloat(os.Args[2], 64)
	if err2!= nil {
        fmt.Println(err2)
    }
	// os.Args[3] initialSpeed for boid
	initialSpeed, err3 := strconv.ParseFloat(os.Args[3], 64)
	if err3!= nil {
        fmt.Println(err3)
    }
	// os.Args[4] max speed for boid
	maxBoidSpeed, err4 := strconv.ParseFloat(os.Args[4], 64)
	if err4!= nil {
        fmt.Println(err4)
    }
	// os.Args[5] num of generation
	numGens, err5 := strconv.Atoi(os.Args[5])
	if err5!= nil {
        fmt.Println(err5)
    }
	// num of generation can't be negative
	if numGens <= 0 {
		panic("numGens must be greater to 0")
	}
	// os.Args[6] is threshold distance between boids to generate forces
	proximity, err6 := strconv.ParseFloat(os.Args[6], 64)
	if err6!= nil {
        fmt.Println(err6)
    }
	// os.Args[7] to os.Args[9] are factors
	separationFactor, err7 := strconv.ParseFloat(os.Args[7], 64)
	if err7!= nil {
        fmt.Println(err7)
    }
	alignmentFactor, err8 := strconv.ParseFloat(os.Args[8], 64)
	if err8!= nil {
        fmt.Println(err8)
    }
	cohesionFactor, err9 := strconv.ParseFloat(os.Args[9], 64)
	if err9!= nil {
        fmt.Println(err9)
    }
	// os.Args[10] is time of per generation
	timeStep, err10 := strconv.ParseFloat(os.Args[10], 64)
	if err10!= nil {
        fmt.Println(err10)
    }
	// time can't be negative
	if timeStep <= 0 {
		panic("timeStep must be greater than 0")
	}
	// os.Args[11] is canvas width
	canvasWidth, err11 := strconv.Atoi(os.Args[11])
	if err11!= nil {
        fmt.Println(err11)
    }
	// os.Args[12] is frequence to catch image
	imageFrequency, err12 := strconv.Atoi(os.Args[12])
	if err12!= nil {
        fmt.Println(err12)
    }
	// Use random function to set initial velocity and position
	rand.Seed(time.Now().UTC().UnixNano())
	// Set list that length equals to num of Boids
	boidsList := make([]Boid, numBoids)
	// Set initila velocity and position of each Boids and store in the list
	for i := 0; i < numBoids; i++ {
		// Random position that is in sky
		boidsList[i].position.x = float64(rand.Intn(int(skyWidth)))
		boidsList[i].position.y = float64(rand.Intn(int(skyWidth)))
		// Set random float for unitVector X and calculate unitVector Y
		var unitVector OrderedPair
		unitVector.x = rand.Float64()
		unitVector.y = math.Sqrt(1-unitVector.x*unitVector.x)
		// intial speed will equal to initialSpeed that user input
		// Random assign direction for velocity
		boidsList[i].velocity.x = unitVector.x * initialSpeed * math.Pow(-1, float64(rand.Intn(100)))
		boidsList[i].velocity.y = unitVector.y * initialSpeed * math.Pow(-1, float64(rand.Intn(100)))
		// initial acceleration equals to 0
		boidsList[i].acceleration.x = 0
		boidsList[i].acceleration.y = 0
	}
	// Set initial sky
	var initialSky Sky
	// fields of sky equal to argument that user input
	initialSky.width = skyWidth
	initialSky.boids = boidsList
	initialSky.maxBoidSpeed = maxBoidSpeed
	initialSky.proximity = proximity
    initialSky.alignmentFactor = alignmentFactor
	initialSky.separationFactor = separationFactor
	initialSky.cohesionFactor = cohesionFactor

	fmt.Println("Successfully read command line argument!")
	fmt.Println("Simulating system.")
	// Starts to simulate sky from initial sky
	simulateSky := SimulateBoids(initialSky, numGens, timeStep)

	fmt.Println("Simulating system created!")
	// Starts to create inages
	images := AnimateSystem(simulateSky, canvasWidth, imageFrequency)

	fmt.Println("Images have been created!")
	fmt.Println("Drawing...")
	// Generate gif
	gifhelper.ImagesToGIF(images, "boids")

	fmt.Println("Done!")
	fmt.Println("Mission successfully completed!")

}	
