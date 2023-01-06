package main
import(
	"fmt"
	"testing"
)

// Test UpdateVelocity
func TestUpdateVelocity(t *testing.T) {
	// Construct test
	type test struct {
		b Boid
		timeStep float64
		maxSpeed float64
		answer OrderedPair
	}
	var t1 test
	// give values to test
	t1.maxSpeed = 3
	t1.timeStep = 1
	t1.b.velocity.x = -5
	t1.b.velocity.y = 2
	t1.b.acceleration.x = 1
	t1.b.acceleration.y = 1
	// Because exceeds max speed
	t1.answer.x = (4.0/5.0) * t1.maxSpeed * -1
	t1.answer.y = (3.0/5.0) * t1.maxSpeed
	outcome := UpdateVelocity(t1.b, t1.timeStep, t1.maxSpeed)
	if outcome != t1.answer {
		if t1.answer != outcome {
			t.Errorf("Expected %f, got %f", t1.answer, outcome)
		} else {
			fmt.Println("Correct!")
		}
	}
	// Velocity not exceeds max speed
	var t2 test
	// give values to test
	t2.maxSpeed = 50
	t2.timeStep = 1
	t2.b.velocity.x = -5
	t2.b.velocity.y = 2
	t2.b.acceleration.x = 1
	t2.b.acceleration.y = -1
	// Because exceeds max speed
	t2.answer.x = -4
	t2.answer.y = 1
	outcome2 := UpdateVelocity(t2.b, t2.timeStep, t2.maxSpeed)
	if outcome2 != t2.answer {
		if t2.answer != outcome2 {
			t.Errorf("Expected %f, got %f", t2.answer, outcome2)
		} else {
			fmt.Println("Correct!")
		}
	}
}

// Test UpdatePosition
func TestUpdatePosition(t *testing.T) {
	// Construct test
	type test struct {
        b Boid
		timeStep float64
		answer OrderedPair
    }
	var t1 test
	// timeStep won't less than 0 because have already avoid this situation in main.go
	t1.timeStep = 1
	t1.b.acceleration.x = 0
	t1.b.acceleration.y = -5
	t1.b.velocity.x = -10
	t1.b.velocity.y = 0
	t1.b.position.x = -20
	t1.b.position.y = 5
	t1.answer.x = 0.5 * 0 * 1 * 1 + -10 * 1 + -20
	t1.answer.y = 0.5 * -5 * 1 * 1 + 0 * 1 + 5
	outcome := UpdatePosition(t1.b, t1.timeStep)
	if t1.answer != outcome {
		t.Errorf("Expected %f, got %f", t1.answer, outcome)
	} else {
		fmt.Println("Correct!")
	}
}

// Test Distance
func TestDistance(t *testing.T) {
	type test struct {
		p2 OrderedPair
		answer float64
	}
	// Test different values
	tests := make([]test, 3)
	tests[0].p2.x = 0
	tests[0].p2.y = 0
	tests[0].answer = 0
	tests[1].p2.x = -3
	tests[1].p2.y = -4
	tests[1].answer = 5
	tests[2].p2.x = 3
	tests[2].p2.y = -4
	tests[2].answer = 5
	var p1 OrderedPair
	p1.x = 0
    p1.y = 0
	for i := range tests {
		outcome := Distance(p1, tests[i].p2)
		if outcome != tests[i].answer {
			t.Errorf("Error! Expected %f, got %f", tests[i].answer, outcome)
		} else {
			fmt.Println("Correct! When position is", tests[i].p2, "distance is", outcome)
		}
	}
}

// Test SeparationForce
func TestSeparationForce(t *testing.T) {
	// Construct test
	type test struct {
        b1, b2 Boid
		separationFactor float64
		d float64
		answer OrderedPair
    }
	var t1 test
	t1.b1.position.x = 0
	t1.b1.position.y = 0
	t1.b2.position.x = -4
    t1.b2.position.y = 3
	t1.separationFactor = 1
	t1.d = Distance(t1.b1.position, t1.b2.position)
	t1.answer.x = 4.0 / 25.0
	t1.answer.y = -3.0 / 25.0
    outcome := SeparationForce(t1.b1, t1.b2, t1.separationFactor, t1.d)
	if t1.answer != outcome {
			t.Errorf("Error! Expected (%f, %f), got (%f, %f)", t1.answer.x, t1.answer.y, outcome.x, outcome.y)
	} else {
			fmt.Printf("Correct!")
	}

}

// Test AlignmentForce
func TestAlignmentForce(t *testing.T) {
	// Construct test
	type test struct {
        b1, b2 Boid
		alignmentFactor float64
		d float64
		answer OrderedPair
    }
	var t1 test
	t1.b1.position.x = 0
	t1.b1.position.y = 0
	t1.b2.position.x = -4
    t1.b2.position.y = 3
	t1.b2.velocity.x = -5
    t1.b2.velocity.y = 0
	t1.alignmentFactor = -5
	t1.d = Distance(t1.b1.position, t1.b2.position)
	t1.answer.x = -5.0 * -5 / 5.0
	t1.answer.y = 0 / 5.0
	outcome := AlignmentForce(t1.b2, t1.alignmentFactor, t1.d)
	if t1.answer != outcome {
			t.Errorf("Error! Expected (%f, %f), got (%f, %f)", t1.answer.x, t1.answer.y, outcome.x, outcome.y)
	} else {
			fmt.Println("Correct!")
	}

}

// Test CohesionForce
func TestCohesionForce(t *testing.T) {
	type test struct {
        b1, b2 Boid
		cohesionFactor float64
		d float64
		answer OrderedPair
    }
	var t1 test
	t1.b1.position.x = 0
	t1.b1.position.y = 0
	t1.b2.position.x = -4
    t1.b2.position.y = 3
	t1.cohesionFactor = -5
	t1.d = Distance(t1.b1.position, t1.b2.position)
	t1.answer.x = -5.0 * -4.0 / 5.0
	t1.answer.y = -5.0 * 3.0 / 5.0
	outcome := CohesionForce(t1.b1, t1.b2, t1.cohesionFactor, t1.d)
	if t1.answer != outcome {
			t.Errorf("Error! Expected (%f, %f), got (%f, %f)", t1.answer.x, t1.answer.y, outcome.x, outcome.y)
	} else {
			fmt.Println("Correct!")
	}
}
