package main

import (
	"fmt"
	"testing"
)

// Test IsInSector
func TestIsInSector(t *testing.T) {
	type test struct {
		s      *Star
		ns     Quadrant
		answer bool
	}
	var t1 test
	var s1 Star
	var ns Quadrant
	t1.ns = ns
	ns.x, ns.y, ns.width = 0, 10, 10
	t1.s = &s1
	t1.s.position.x, t1.s.position.y = 0, 10
	t1.answer = false
	outcome := IsInSector(t1.s, t1.ns)
	if outcome != t1.answer {
		t.Errorf("Expected %t, got %t", t1.answer, outcome)
	} else {
		fmt.Println("Correct!")
	}
	var t2 test
	var s2 Star
	t2.ns = ns
	t2.s = &s2
	t2.s.position.x, t2.s.position.y = 0, 0
	t2.answer = true
	outcome2 := IsInSector(t2.s, t2.ns)
	if outcome2 != t2.answer {
		t.Errorf("Expected %t, got %t", t2.answer, outcome2)
	} else {
		fmt.Println("Correct!")
	}
}

// Test ComputeForce
func TestComputeForce(t *testing.T) {
	type test struct {
		s      *Star
		n      *Node
		answer OrderedPair
	}
	var s Star
	var sNode Star
	var n Node
	var t1 test
	n.star = &sNode
	t1.s, t1.n = &s, &n
	t1.s.position.x, t1.s.position.y, t1.s.mass = 1, 1, 10
	t1.n.star.position.x, t1.n.star.position.y, t1.n.star.mass = 4, 5, 20
	F := G * 10 * 20 / 25
	t1.answer.x, t1.answer.y = (F*3.0)/5.0, (F*4.0)/5.0
	outcome := ComputeForce(t1.s, t1.n)
	if outcome != t1.answer {
		t.Errorf("Expected %f, got %f", t1.answer, outcome)
	} else {
		fmt.Println("Correct!")
	}
}

// Test CreatSquare
func TestCreatSquare(t *testing.T) {
	type test struct {
		x, y, width float64
		answer      []Quadrant
	}
	var t1 test
	t1.x, t1.y, t1.width = 0, 10, 10
	var nwTest, neTest, swTest, seTest Quadrant
	nwTest.x, nwTest.y, nwTest.width = 0, 5, 5
	neTest.x, neTest.y, neTest.width = 5, 5, 5
	swTest.x, swTest.y, swTest.width = 0, 10, 5
	seTest.x, seTest.y, seTest.width = 5, 10, 5
	t1.answer = append(t1.answer, nwTest, neTest, swTest, seTest)
	outcome1 := CreatSquare(t1.x, t1.y, t1.width)

	if len(outcome1) != len(t1.answer) {
		t.Errorf("Expected %d, got %d", len(t1.answer), len(outcome1))
	} else {
		fmt.Println("Correct!")
	}

	for i := 0; i < len(outcome1); i++ {
		if outcome1[i].x != t1.answer[i].x {
			t.Errorf("Expected %f, got %f", t1.answer[i].x, outcome1[i].x)
		} else {
			fmt.Println("Correct!")
		}
		if outcome1[i].y != t1.answer[i].y {
			t.Errorf("Expected %f, got %f", t1.answer[i].y, outcome1[i].y)
		} else {
			fmt.Println("Correct!")
		}
		if outcome1[i].width != t1.answer[i].width {
			t.Errorf("Expected %f, got %f", t1.answer[i].width, outcome1[i].width)
		} else {
			fmt.Println("Correct!")
		}
	}
}

// Test CheckDirection
func TestCheckDirection(t *testing.T) {
	type test struct {
		s      *Star
		q      Quadrant
		answer int
	}
	var t1 test
	var qTest Quadrant
	qTest.width, qTest.x, qTest.y = 10, 0, 10
	var s1Test Star
	s1Test.position.x, s1Test.position.y = 5, 5
	t1.q, t1.s, t1.answer = qTest, &s1Test, 3
	outcome1 := CheckDirection(t1.s, t1.q)
	if outcome1 != t1.answer {
		t.Errorf("Expected %d, got %d", t1.answer, outcome1)
	} else {
		fmt.Println("Correct!")
	}
	var t2 test
	var s2Test Star
	s2Test.position.x, s2Test.position.y = 10, 10
	t2.q, t2.s, t2.answer = qTest, &s2Test, 4
	outcome2 := CheckDirection(t2.s, t2.q)
	if outcome2 != t2.answer {
		t.Errorf("Expected %d, got %d", t2.answer, outcome2)
	} else {
		fmt.Println("Correct!")
	}
	var t3 test
	var s3Test Star
	s3Test.position.x, s1Test.position.y = 0, 0
	t3.q, t3.s, t3.answer = qTest, &s3Test, 0
	outcome3 := CheckDirection(t3.s, t3.q)
	if outcome3 != t3.answer {
		t.Errorf("Expected %d, got %d", t3.answer, outcome3)
	} else {
		fmt.Println("Correct!")
	}

}

// Test UpdateVelocity
func TestUpdateVelocity(t *testing.T) {
	// Construct test
	type test struct {
		s      *Star
		time   float64
		answer OrderedPair
	}
	var t1 test
	var testStar Star
	t1.s = &testStar
	t1.time = 1
	t1.s.velocity.x, t1.s.velocity.y = -5, 2
	t1.s.acceleration.x, t1.s.acceleration.y = 1, 1
	t1.answer.x, t1.answer.y  = -4, 3
	t1.s.UpdateVelocity(t1.time)
	if t1.s.velocity != t1.answer {
		t.Errorf("Expected %f, got %f", t1.answer, t1.s.velocity)
	} else {
		fmt.Println("Correct!")
	}
	var t2 test
	var test2Star Star
	t2.s = &test2Star
	t2.time = 1
	t2.s.velocity.x, t2.s.velocity.y = -5, 2
	t2.s.acceleration.x, t2.s.acceleration.y = 1, -1
	t2.answer.x, t2.answer.y = -4, 1
	t2.s.UpdateVelocity(t2.time)
	if t2.s.velocity != t2.answer {
		t.Errorf("Expected %f, got %f", t2.answer, t2.s.velocity)
	} else {
		fmt.Println("Correct!")
	}
}

// Test UpdatePosition
func TestUpdatePosition(t *testing.T) {
	// Construct test
	type test struct {
		s      *Star
		time   float64
		answer OrderedPair
	}
	var t1 test
	// timeStep won't less than 0 because have already avoid this situation in main.go
	var testStar Star
	t1.s = &testStar
	t1.time = 1
	t1.s.acceleration.x, t1.s.acceleration.y = 0, -5
	t1.s.velocity.x, t1.s.velocity.y = -10, 0
	t1.s.position.x, t1.s.position.y = -20, 5
	t1.answer.x, t1.answer.y = (0.5*0*1*1 + -10*1 + -20), (0.5*-5*1*1 + 0*1 + 5)
	t1.s.UpdatePosition(t1.time)
	if t1.answer != t1.s.position {
		t.Errorf("Expected %f, got %f", t1.answer, t1.s.position)
	} else {
		fmt.Println("Correct!")
	}
}

// Test Distance
func TestDistance(t *testing.T) {
	type test struct {
		p2     OrderedPair
		answer float64
	}
	// Test different values
	tests := make([]test, 3)
	tests[0].p2.x, tests[0].p2.y, tests[0].answer = 0, 0, 0
	tests[1].p2.x, tests[1].p2.y, tests[1].answer = -3, -4, 5
	tests[2].p2.x, tests[2].p2.y, tests[2].answer = 3, -4, 5

	var p1 OrderedPair
	p1.x, p1.y = 0, 0
	for i := range tests {
		outcome := Distance(p1, tests[i].p2)
		if outcome != tests[i].answer {
			t.Errorf("Error! Expected %f, got %f", tests[i].answer, outcome)
		} else {
			fmt.Println("Correct! When position is", tests[i].p2, "distance is", outcome)
		}
	}
}
