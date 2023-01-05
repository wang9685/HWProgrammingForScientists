package main

import (
	"fmt"
	"gifhelper"
	"math"
	"os"
)

func main() {
	// you probably want to apply a "push" function at this point to these galaxies to move
	// them toward each other to collide.
	// be careful: if you push them too fast, they'll just fly through each other.
	// too slow and the black holes at the center collide and hilarity ensues.
	width := 1.0e23
	// now evolve the universe: feel free to adjust the following parameters.
	numGens := 100000
	time := 2e14
	theta := 2.0
	canvasWidth := 1000
	frequency := 1000
	scalingFactor := 1e11 // a scaling factor is needed to inflate size of stars when drawn because galaxies are very sparse
	system := os.Args[1]
	if system == "jupiter" {
		var jupiter, io, europa, ganymede, callisto Star
		jupiter.red, jupiter.green, jupiter.blue = 223, 227, 202
		io.red, io.green, io.blue = 249, 249, 165
		europa.red, europa.green, europa.blue = 132, 83, 52
		ganymede.red, ganymede.green, ganymede.blue = 76, 0, 153
		callisto.red, callisto.green, callisto.blue = 0, 153, 76
		jupiter.mass = 1.898 * math.Pow(10, 27)
		io.mass = 8.9319 * math.Pow(10, 22)
		europa.mass = 4.7998 * math.Pow(10, 22)
		ganymede.mass = 1.4819 * math.Pow(10, 23)
		callisto.mass = 1.0759 * math.Pow(10, 23)

		jupiter.radius = 71000000
		// 10 times the radius for other stars
		io.radius = 18210000
		europa.radius = 15690000
		ganymede.radius = 26310000
		callisto.radius = 24100000

		jupiter.position.x, jupiter.position.y = 2000000000, 2000000000
		io.position.x, io.position.y = 2000000000-421600000, 2000000000
		europa.position.x, europa.position.y = 2000000000, 2000000000+670900000
		ganymede.position.x, ganymede.position.y = 2000000000+1070400000, 2000000000
		callisto.position.x, callisto.position.y = 2000000000, 2000000000-1882700000

		jupiter.velocity.x, jupiter.velocity.y = 0, 0
		io.velocity.x, io.velocity.y = 0, -17320
		europa.velocity.x, europa.velocity.y = -13740, 0
		ganymede.velocity.x, ganymede.velocity.y = 0, 10870
		callisto.velocity.x, callisto.velocity.y = 8200, 0
		// Set jupiter using previous parameters
		jupiterWidth := 4000000000.0
		numGensJupiter := 1000000
		timeJupiter := 30.0
		canvasWidthJupiter := 1000
		frequencyJupiter := 300
		scalingFactorJupiter := 1.0

		g0 := []*Star{&jupiter, &io, &europa, &ganymede, &callisto}
		galaxies := []Galaxy{g0}
		fmt.Println("Running for jupiter")
		initialUniverse := InitializeUniverse(galaxies, jupiterWidth)
		timePoints := BarnesHut(initialUniverse, numGensJupiter, timeJupiter, theta)
		imageList := AnimateSystem(timePoints, canvasWidthJupiter, frequencyJupiter, scalingFactorJupiter)
		fmt.Println("Simulation run. Now drawing images.")
		fmt.Println("Images drawn. Now generating GIF.")
		gifhelper.ImagesToGIF(imageList, "galaxy")
		fmt.Println("GIF drawn.")
	} else if system == "galaxy" {
		g0 := InitializeGalaxy(500, 4e21, 7e22, 2e22)
		galaxies := []Galaxy{g0}
		fmt.Println("Running for galaxy")
		initialUniverse := InitializeUniverse(galaxies, width)
		timePoints := BarnesHut(initialUniverse, numGens, time, theta)
		imageList := AnimateSystem(timePoints, canvasWidth, frequency, scalingFactor)
		fmt.Println("Simulation run. Now drawing images.")
		fmt.Println("Images drawn. Now generating GIF.")
		gifhelper.ImagesToGIF(imageList, "galaxy")
		fmt.Println("GIF drawn.")

	} else if system == "collision" {
		// the following sample parameters may be helpful for the "collide" command
		// all units are in SI (meters, kg, etc.)
		// but feel free to change the positions of the galaxies.
		g0 := InitializeGalaxy(500, 4e21, 7e22, 2e22)
		g1 := InitializeGalaxy(500, 4e21, 3e22, 7e22)
		blackHole0 := g0[len(g0)-1]
		blackHole1 := g1[len(g1)-1]
		d := Distance(blackHole0.position, blackHole1.position)
		// Add a push force to galaxies to make them closer
		cohForce0 := CohesionForce(blackHole0, blackHole1, 2000.0, d)
		cohForce1 := CohesionForce(blackHole1, blackHole0, 2000.0, d)
		for star := range g0 {
			g0[star].velocity.x += cohForce0.x 
			g0[star].velocity.y += cohForce0.y // Can *factor, so galaxies won't have collision directly
		}
		for star := range g1 {
			g1[star].velocity.x += cohForce1.x 
			g1[star].velocity.y += cohForce1.y
		}
		galaxies := []Galaxy{g0, g1}
		fmt.Println("Running for collision")
		initialUniverse := InitializeUniverse(galaxies, width)
		timePoints := BarnesHut(initialUniverse, numGens, time, theta)
		imageList := AnimateSystem(timePoints, canvasWidth, frequency, scalingFactor)
		fmt.Println("Simulation run. Now drawing images.")
		fmt.Println("Images drawn. Now generating GIF.")
		gifhelper.ImagesToGIF(imageList, "galaxy")
		fmt.Println("GIF drawn.")

	}
}
// Input: *Star, Factor, distance between stars
// Output: force
// To help galaxies get closer
// From boids
func CohesionForce(b1, b2 *Star, cohesionFactor float64, d float64) OrderedPair {
	// Set variable to store force in OrderedPair
	var coh OrderedPair
	// Calculate cohesion force separately
	coh.x = cohesionFactor * (b2.position.x - b1.position.x) / d
	coh.y = cohesionFactor * (b2.position.y - b1.position.y) / d
	// return cohesiom force
	return coh
}
