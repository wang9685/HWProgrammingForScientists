package main
import (
	"math"
)

//BarnesHut is our highest level function.
//Input: initial Universe object, a number of generations, and a time interval.
//Output: collection of Universe objects corresponding to updating the system
//over indicated number of generations every given time interval.
func BarnesHut(initialUniverse *Universe, numGens int, time, theta float64) []*Universe {
	timePoints := make([]*Universe, numGens+1)
	// Your code goes here. Use subroutines! :)
	// index 0 equals to initialUniverse which will be create in the main
	timePoints[0] = initialUniverse
	// Update each universe by previous
	for i := 1; i <= numGens; i++ {
		timePoints[i] = UpdateUniverse(timePoints[i-1], time, theta)
	}
	// Return the collection of each universes
	return timePoints
}
// Input: previous universe, time
// Update: newUniverse
// Update universe from previous one
func UpdateUniverse(previousUniverse *Universe, time float64, theta float64) *Universe {
	// Create a variable to store new universe
	newUniverse := CopyUniverse(previousUniverse)
	tree := CreateTree(previousUniverse.stars, previousUniverse.width)
	// Loop through universe.stars to update each stars in universe
	for i := range newUniverse.stars {
		newUniverse.stars[i].acceleration = UpdateAcceleration(tree, previousUniverse.stars[i], theta)
		newUniverse.stars[i].UpdateVelocity(time)
		newUniverse.stars[i].UpdatePosition(time)
	}
	return newUniverse
}
// Input: previous universe
// Update: point to newUniverse
// Copy Universe from  previous one
func CopyUniverse(previousUniverse *Universe) *Universe {
	// Create a new universe
	var newUniverse Universe
	// Copy width
	newUniverse.width = previousUniverse.width
	// Get length of stars 
	numStars := len(previousUniverse.stars)
	// Make a list to store stars addresses
	newUniverse.stars = make([]*Star, numStars)
	// Loop through stars, and copy each fields in star
	for i := range previousUniverse.stars {
		//Deep Copy to create star for new universe
		var star Star
		star.mass = previousUniverse.stars[i].mass
		star.radius = previousUniverse.stars[i].radius
		star.position.x = previousUniverse.stars[i].position.x
		star.position.y = previousUniverse.stars[i].position.y
		star.velocity.x = previousUniverse.stars[i].velocity.x
		star.velocity.y = previousUniverse.stars[i].velocity.y
		star.acceleration.x = previousUniverse.stars[i].acceleration.x
		star.acceleration.y = previousUniverse.stars[i].acceleration.y
		star.red = previousUniverse.stars[i].red
		star.green = previousUniverse.stars[i].green
		star.blue = previousUniverse.stars[i].blue
		// Refer new star to star list in newUniverse
		newUniverse.stars[i] = &star
	}
	return &newUniverse
}
// Input: QuadTree,theta, star
// Output: Updated acceleration
// Use ComputeNetForce function to search through tree to get force
func UpdateAcceleration(tree *QuadTree, s *Star, theta float64) OrderedPair {
	// Set accel and force to store values
	var accel OrderedPair
	var netForce OrderedPair
	// ComputeNetForce through tree
	for i := 0; i < len(tree.root.children); i++ {
		if tree.root.children[i].star != nil {
			force := ComputeNetForce(s, tree.root.children[i], theta)
			netForce.x += force.x
			netForce.y += force.y
		}	
	}
	// Calculate accel
	accel.x = netForce.x / s.mass
	accel.y = netForce.y / s.mass
	return accel
}

// Input: current star, Node, theta
// Output: Calculate NetForce for current star
// Use ComputeNetForce function to search through tree to get force
func ComputeNetForce(s1 *Star, n *Node, theta float64) OrderedPair {
	var netForce OrderedPair
	// Distance between current star and star in node
	d := Distance(s1.position, n.star.position)
	// Node's sector width
	s := n.sector.width
	// If star is not in same id as node
	// If s/d > theta and Node is a leaf
	// or if it is dummy star and s/d < theta
	if !IsInSector(s1, n.sector) && s1 != n.star && n.star != nil {
		if (s/d > theta && len(n.children) == 0) || (s/d < theta) {
			// Compute force of node's star to star
			computeForce := ComputeForce(s1, n)
			netForce.x += computeForce.x
			netForce.y += computeForce.y
		// If Node is inner node & s/d > theta, continue go through its children
		} else {
			for i := 0; i < len(n.children); i++ {
				// If child doesn't contain star then skip
				if n.children[i].star != nil && s1 != n.children[i].star {	
					// If it contains star, then use ComputeNetForce functio to check again		
					computeForce := ComputeNetForce(s1, n.children[i], theta)
					netForce.x += computeForce.x
					netForce.y += computeForce.y
				}
			}
		}
	// If star is at same side as node (Discussion with Tao Luo)
	} else if IsInSector(s1, n.sector) && s1 != n.star && n.star != nil {
		// If it is a leaf, then count force directly
		if len(n.children) == 0 {
			computeForce := ComputeForce(s1, n)
			netForce.x += computeForce.x
			netForce.y += computeForce.y
		// Else, go through ComputeNetForce again
		// Because if dummy stars is same side with star, they contain the mass of star,
		// we can't directly calculate all the node in this side 
		} else {
			for i := 0; i < len(n.children); i++ {
				// If child doesn't contain star than skip
				if n.children[i].star != nil && s1 != n.children[i].star {	
					// If it contains star, then use ComputeNetForce functio to check again		
					computeForce := ComputeNetForce(s1, n.children[i], theta)
					netForce.x += computeForce.x
					netForce.y += computeForce.y
				}
			}
		}
	}
	return netForce
}
// Input: current star, node's sector
// Output: true or false
// To check if star is in the same sector with node
func IsInSector(s *Star, ns Quadrant) bool{
	if s.position.x >= ns.x && s.position.x < (ns.x + ns.width) && s.position.y < ns.y && s.position.y >= (ns.y - ns.width) {
		return true
	} 
	return false
}
// Input: current star, node
// Ouput: force
// Compute force from node to star
func ComputeForce(s1 *Star, n *Node) OrderedPair {
	var force OrderedPair
	// Distance between node and star
	dist := Distance(s1.position, n.star.position)
	// magnitude of gravity
	F := G * s1.mass * n.star.mass / (dist * dist) 
	deltaX := n.star.position.x - s1.position.x
	deltaY := n.star.position.y - s1.position.y
	//deltaX/dist = cos theta
	force.x = F * deltaX / dist 
	//deltaY/dist = sin theta
	force.y = F * deltaY / dist 
	return force
}

// Input: Total stars in universe, width
// Output: pointer of Tree
// Create Tree for all the stars
func CreateTree(starList []*Star, width float64) *QuadTree {
	// Create tree, root
	var tree QuadTree
	var root Node
	// Set root sector at (0, width)
	// Width equals to width
	root.sector.x = 0
	root.sector.y = width
	root.sector.width = width
	// Asign root to tree.root
	tree.root = &root
	// Create children nodes form root
	tree.root = DivideList(starList, tree.root)
	return &tree
}

// Input: Total stars in universe, root node
// Output: node
// Use position to divide stars to 4 directions
// At the same time, also divides node to 4 children
// And continuely divide the stars (and node) to 4 directions using recursion
// until each lists only has one star
// Then, assign star to node
func DivideList(starList []*Star, n *Node) *Node{
	// Create children node
	n = InitializeNodeChildren(n)
	// Set node as dummy star
	n.star = SetDummyStar(starList)
	// Make lists to store stars from four directions 
	nw := make([]*Star, 0)
	ne := make([]*Star, 0)
	sw := make([]*Star, 0)
	se := make([]*Star, 0)
	// Loop through each lists
	// Check direction of each star, and divide them into four lists
	for i := range starList {
		d  := CheckDirection(starList[i], n.sector)
		if d == 0 {
			nw = append(nw, starList[i])
		} else if d == 1 {
			ne = append(ne, starList[i])
		} else if d == 2 {
			sw = append(sw, starList[i])
		} else if d == 3 {
			se = append(se, starList[i])
		}
	}
	// If list only has one star, then assign it to children node
	if len(nw) == 1 {
		n.children[0].star = nw[0]
	// If not, then execute the function to continuously divide the list
	} else if len(nw) > 1 {
		n.children[0] = DivideList(nw, n.children[0])
	}
	if len(ne) == 1 {
		n.children[1].star = ne[0]
	} else if len(ne) > 1 {
		n.children[1] = DivideList(ne, n.children[1])
	}
	if len(sw) == 1 {
		n.children[2].star = sw[0]
	} else if len(sw) > 1 {
		n.children[2] = DivideList(sw, n.children[2])
	}
	if len(se) == 1 {
		n.children[3].star = se[0]
	} else if len(se) > 1 {
		n.children[3] = DivideList(se, n.children[3])
	}
	return n
}

// Input: Stars under this dummy star
// Output: pointer of star
// Set the dummy star and calculate its position and mass using all of its children stars
func SetDummyStar(stars []*Star) *Star {
	var dummyStar Star
	var mass float64
	var position OrderedPair
	// Loop through children
	for i := 0; i < len(stars); i++ {
		// Add mass together
		mass += stars[i].mass
		// Add x.position
		position.x += (stars[i].position.x) * stars[i].mass
		// Add y.position
		position.y += (stars[i].position.y) * stars[i].mass
	} 
	// Calculate for center of those children
	dummyStar.position.x = position.x / mass
	dummyStar.position.y = position.y / mass
	dummyStar.mass = mass
	return &dummyStar
}

// Input: Node
// Output: Node with children list
// Create a list of node for node.children, and set initial parameters
func InitializeNodeChildren(n *Node) *Node {
	n.children = make([]*Node, 4)
	// Create sector
	boundary := CreatSquare(n.sector.x, n.sector.y, n.sector.width)
	// Loop through four children
	for i := 0; i < len(n.children); i++ {
		// Create a node in it
		var newNode Node
		// Set each child sector, and it would be calculated from mother node
		newNode.sector.x = boundary[i].x
		newNode.sector.y = boundary[i].y
		newNode.sector.width = boundary[i].width
		// newNode refer to n.children[i]
		n.children[i] = &newNode
	}
	return n
}

// Input: x, y of sector and width
// Output: Quadrant
// Calculate sector from mother parameters for children
func CreatSquare(x, y, width float64) []Quadrant{
	// Different direction has different x, y
	var nw, ne, sw, se Quadrant
	nw.x = x
	nw.y = y - width/2
	nw.width = width/2
	ne.x = x + width/2
	ne.y = y - width/2
	ne.width = width/2
	sw.x = x
	sw.y = y
	sw.width = width/2
	se.x = x + width/2
	se.y = y
	se.width = width/2
	// Create a list to store
	boundry := 	[]Quadrant{nw, ne, sw, se}
	return boundry
}

// Input: current star, Quandrant
// Output: int for which direction
// Check current star should be in which child
func CheckDirection(s *Star, q Quadrant) int {
	// ne
	if s.position.x >= q.x + q.width/2 && s.position.x < q.x + q.width && s.position.y < q.y - q.width/2 && s.position.y >= q.y - q.width {
		return 1
		// sw
	} else if s.position.x < q.x + q.width/2 && s.position.x >= q.x && s.position.y >= q.y - q.width/2 && s.position.y < q.y {
		return 2
		// se
	} else if s.position.x >= q.x + q.width/2 && s.position.x < q.x + q.width && s.position.y >= q.y - q.width/2 && s.position.y < q.y {
		return 3
		// nw
	} else if s.position.x < q.x + q.width/2 && s.position.x >= q.x && s.position.y < q.y - q.width/2 && s.position.y >= q.y - q.width {
		return 0
	}
	return 4
}

// Star method
// Input: time
// Output: The new velocity of the star
// Calculate th velocity by time and acceleration
func (s *Star) UpdateVelocity (time float64) {
	s.velocity.x = s.velocity.x + s.acceleration.x*time
	s.velocity.y = s.velocity.y + s.acceleration.y*time
}

// Star method
// Input: time
// Output: updated position of star
func (s *Star) UpdatePosition (time float64) {
	// Calculate
	s.position.x = 0.5*s.acceleration.x*time*time + s.velocity.x*time + s.position.x
	s.position.y = 0.5*s.acceleration.y*time*time + s.velocity.y*time + s.position.y
}

// Input: Position of two stars
// Ouput: The distance between two stars
// Calculate distance between two stars
func Distance(p1, p2 OrderedPair) float64 {
	deltaX := p1.x - p2.x
	deltaY := p1.y - p2.y
	return math.Sqrt(deltaX*deltaX + deltaY*deltaY)
}