package main

import (
	"math"
)

//place your non-drawing functions here.
// Input: initialSky, numGens, timeStep
// Output: a list of sky
// Simulate Boids
func SimulateBoids(initialSky Sky, numGens int, timeStep float64) []Sky {
	// Create a list to store skys
	// The length is equal to numGen+1(include initial Sky)
	skyList := make([]Sky, numGens+1)
	// index 0 stores initial Sky
	skyList[0] = initialSky
	// Others store sky that be updated by previous one 
	for i := 1; i <= numGens; i++ {
		skyList[i] = UpdateSky(skyList[i-1], timeStep)
	}
	// Return skyList
	return skyList
}

// Input: CurrentSky, time
// Output: New Sky
// Update Sky with new acceleration
func UpdateSky(currentSky Sky, timeStep float64) Sky {
	// Copy current ky to new sky
	newSky := CopySky(currentSky)
	// Loopthrough boids to update factors
	for i := 0 ; i < len(newSky.boids); i++ {
		// Update acceleration
		newSky.boids[i].acceleration = UpdateAcceleration(currentSky, newSky.boids[i])
        // Update velocity
		newSky.boids[i].velocity = UpdateVelocity(newSky.boids[i], timeStep, currentSky.maxBoidSpeed)
		// Update position
		newSky.boids[i].position = UpdatePosition(newSky.boids[i], timeStep)
		// If boids totally run out the sky, they will return into the sky from the other side
		if newSky.boids[i].position.x > newSky.width {
			newSky.boids[i].position.x = newSky.boids[i].position.x - newSky.width 
		} 
		if newSky.boids[i].position.x < 0 {
			newSky.boids[i].position.x = newSky.boids[i].position.x + newSky.width
		}
		if newSky.boids[i].position.y > newSky.width {
			newSky.boids[i].position.y = newSky.boids[i].position.y - newSky.width
		} 
		if newSky.boids[i].position.y < 0 {
			newSky.boids[i].position.y = newSky.boids[i].position.y + newSky.width
		}
	}
	// Return newSky
	return newSky
}

// Input: currentSky, current boid
// Output: new acceleration
// Assume boid's mass equals to 1, Acceleration = Netforce
func UpdateAcceleration(currentSky Sky, b Boid) OrderedPair {
	// Set acceleration variable
	var accel OrderedPair
	// Calculate netForce for current boid
	accel = ComputeNetForce(b, currentSky.boids, currentSky.separationFactor, currentSky.alignmentFactor, currentSky.cohesionFactor, currentSky.proximity)
	return accel
}

// Input: the boid, boids list, factors, threshold
// Output: the total forces of the boid
// include separation, alignment, cohesion
// Calculate all the force for the boid
func ComputeNetForce(b Boid, boids []Boid, sepFactor, aliFactor, cohFactor float64, proximity float64) OrderedPair {
	// Set variable to store forces and num of boids that are in threshold distance
	var totalForce  OrderedPair
	var count int
	// Calculate distance
	for i :=0; i < len(boids); i++ {
		 dist := Distance(b.position, boids[i].position)
		 // Won't calculate boids which are too far away, 
		 // and also won't calculate boids which at same position
		if dist <= proximity && dist != 0 {
		// Calculate all the forces
			sepForce := SeparationForce(b, boids[i], sepFactor, dist)
			aliForce := AlignmentForce(boids[i], aliFactor, dist)
			cohForce := CohesionForce(b, boids[i], cohFactor, dist)
			// Add forces together
			totalForce.x += sepForce.x + aliForce.x +cohForce.x
			totalForce.y += sepForce.y + aliForce.y +cohForce.y
			// Count the number of nearby boids
			count ++
		}
	} 
	// Calculate average of force
	// If there is no nearby boid then total force equals to 0
	if count == 0 {
		totalForce.x = 0
		totalForce.y = 0
	} else {
		// Calculate the average force cause by each boids
		totalForce.x = totalForce.x/float64(count)
		totalForce.y = totalForce.y/float64(count)
	}
	// Return totalForce for the boid
	return totalForce
}

// Input: current Sky
// Output: the copy of current Sky
// To copy current sky
func CopySky(currentSky Sky) Sky {
	// Set the newSky type
	var newSky Sky
	// Copy factors in currentSky to newSky
	newSky.width = currentSky.width
	newSky.maxBoidSpeed = currentSky.maxBoidSpeed
	newSky.proximity = currentSky.proximity
	newSky.separationFactor = currentSky.separationFactor
	newSky.alignmentFactor = currentSky.alignmentFactor
	newSky.cohesionFactor = currentSky.cohesionFactor
	
	// Create list to store boid
	newSky.boids = make([]Boid, len(currentSky.boids))

	// Loopthrough bodies from currentSky
	// Copy fields in bodies to newSky
	for i := 0; i < len(currentSky.boids); i++ {
		newSky.boids[i].acceleration = currentSky.boids[i].acceleration
		newSky.boids[i].velocity = currentSky.boids[i].velocity
		newSky.boids[i].position = currentSky.boids[i].position
	}
	return newSky
}

// Input: the boid, time
// OutOut: update velocity
// Conculate update velocity through acceleation
func UpdateVelocity(b Boid, timeStep float64, maxSpeed float64) OrderedPair {
	// Set OrderedPair to store new velocity
	var vel OrderedPair
	// Conculate new velocity
	vel.x = b.velocity.x + b.acceleration.x*timeStep
	vel.y = b.velocity.y + b.acceleration.y*timeStep
	// Calculate speed from new velocity
	speedQ := vel.x*vel.x + vel.y*vel.y
	// Because the speed can't exceed the max speed
	if  math.Sqrt(speedQ) > maxSpeed {
		// If exceed, stay at the max speed and same dirction
		var velMax OrderedPair
		// Use unit vector to calculate the velocity of x,y
		// , so the speed won't exceed the max speed after multiplied by maxSpeed
		velMax.x = math.Sqrt(vel.x*vel.x / speedQ)
        velMax.y = math.Sqrt(vel.y*vel.y / speedQ)
		// Because the direction won't change, if original velocity is negative,
		// then new velocity multiplied by -1
		if vel.x < 0 {
			vel.x = velMax.x * maxSpeed * -1.0
		} else if vel.x > 0{

			vel.x = velMax.x * maxSpeed
		} else {
			vel.x = 0.0
		}
		if vel.y < 0 {
			vel.y = velMax.y * maxSpeed * -1.0
		} else if vel.y > 0{
			vel.y = velMax.y * maxSpeed
		} else {
			vel.y = 0.0
		}
	}
	// return update velocity
	return vel
}

// Input: boid, time
// OutOut: current poistion
// Conculate position through acceleation and velocity
func UpdatePosition(b Boid, timeStep float64) OrderedPair {
	// Set OrderedPair to store new position
	var pos OrderedPair
	timeSqu := timeStep*timeStep
	// Calculate update position
	pos.x = 0.5*b.acceleration.x*timeSqu + b.velocity.x*timeStep + b.position.x
	pos.y = 0.5*b.acceleration.y*timeSqu + b.velocity.y*timeStep + b.position.y
	// Return position
	return pos
}

// Distance between to position
func Distance(p1, p2 OrderedPair) float64 {
	distX := p1.x - p2.x
	distY := p1.y - p2.y
	// Calculate distance
	dist := math.Sqrt(distX*distX + distY*distY)
	// Return distance
	return dist
}

// Input: two boids, time, factor, distance
// OutOut: a SeparationForce between two boids
// Conculate SeparationForce between two Boids
func SeparationForce(b1, b2 Boid, separationFactor float64, d float64) OrderedPair {
	// Set variable to store force in OrderedPair
	var sep OrderedPair
	// Calculate separation force separately 
	sep.x = separationFactor * (b1.position.x - b2.position.x)/(d*d)
	sep.y = separationFactor * (b1.position.y - b2.position.y)/(d*d)
	// Return separation force
	return sep
}

// Input: two boids, time, factor, distance
// OutOut: a AlognmentForce between two boids
// Conculate AlignmentForce between two Boids
func AlignmentForce(b2 Boid, alignmentFactor float64, d float64) OrderedPair {
	// Set variable to store force in OrderedPair
	var ali OrderedPair
	// Calculate alignment force separately 
	ali.x = alignmentFactor * (b2.velocity.x)/d
	ali.y = alignmentFactor * (b2.velocity.y)/d
	// Return alignment force
	return ali
}

// Input: two boids, time, factor, distance
// OutOut: a CohesionForce between two boids
// Conculate CohesionForce between two Boids
func CohesionForce(b1, b2 Boid, cohesionFactor float64, d float64) OrderedPair {
	// Set variable to store force in OrderedPair
	var coh OrderedPair
	// Calculate cohesion force separately 
	coh.x = cohesionFactor * (b2.position.x - b1.position.x)/d
	coh.y = cohesionFactor * (b2.position.y - b1.position.y)/d
	// return cohesiom force
	return coh
}