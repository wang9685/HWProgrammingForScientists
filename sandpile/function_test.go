package main
import (
	"fmt"
	"testing"
	"runtime"
)

func TestEmptyBoard(t *testing.T) {
	type test struct {
		size int
	}
	// Test 1
	var t1 test
	t1.size = 10
	outcome := EmptyBoard(t1.size)
	// Check row number
	if len(outcome) != t1.size {
		t.Errorf("Test EmptyBoard: Expected number of row %d, got %d", t1.size, len(outcome))
	} else {
		fmt.Println("Test EmptyBoard: Correct row number!")
	}
	// Check col number
	if len(outcome[0]) != t1.size {
		t.Errorf("Test EmptyBoard: Expected number of col %d, got %d", t1.size, len(outcome[0]))
	} else {
		fmt.Println("Test EmptyBoard: Correct col number!")
	}
	// Test 2
	var t2 test
	t2.size = 100
	outcome2 := EmptyBoard(t2.size)
	if len(outcome2) != t2.size {
		t.Errorf("Test EmptyBoard: Expected number of row %d, got %d", t2.size, len(outcome2))
	} else {
		fmt.Println("Test EmptyBoard: Correct row number!")
	}
	if len(outcome2[0]) != t2.size {
		t.Errorf("Test EmptyBoard: Expected number of col %d, got %d", t2.size, len(outcome2[0]))
	} else {
		fmt.Println("Test EmptyBoard: Correct col number!")
	}
}
func TestCreateCentralBoard(t *testing.T) {
	type test struct {
		pile int
		size int
	}
	// Test 1
	var t1 test
    t1.pile = 1000
	t1.size = 100
	outcome := EmptyBoard(t1.size)
	outcomePile := 0
	outcome.CreateCentralBoard(t1.pile)
	// Check total pile
	for i := 0; i < len(outcome); i++ {
		for j := 0; j < len(outcome[0]); j++ {
			outcomePile += outcome[i][j]
		}
	}
	if outcomePile != t1.pile {
		t.Errorf("Test CreateCentralBoard: Expected number of pile %d, got %d", t1.pile, outcomePile)
	} else {
		fmt.Println("Test CreateCentralBoard: Correct pile number!")
	}
	// Check pile position
	if outcome[t1.size/2][t1.size/2] != t1.pile {
		t.Errorf("Test CreateCentralBoard: Expected number of pile in central %d , got %d", t1.pile, outcome[t1.size/2][t1.size/2])
	} else {
		fmt.Println("Test CreateCentralBoard: Correct pile position!")
	}
	// Test 2
	var t2 test
    t2.pile = 300
	t2.size = 10
	outcome2 := EmptyBoard(t2.size)
	outcomePile2 := 0
	outcome2.CreateCentralBoard(t2.pile)
	// Check total pile
	for i := 0; i < len(outcome2); i++ {
		for j := 0; j < len(outcome2[0]); j++ {
			outcomePile2 += outcome2[i][j]
		}
	}
	if outcomePile2 != t2.pile {
		t.Errorf("Test CreateCentralBoard: Expected number of pile %d, got %d", t2.pile, outcomePile2)
	} else {
		fmt.Println("Test CreateCentralBoard: Correct pile number!")
	}
	// Check pile position
	if outcome2[t2.size/2][t2.size/2] != t2.pile {
		t.Errorf("Test CreateCentralBoard: Expected number of pile in central %d , got %d", t2.pile, outcome2[t2.size/2][t2.size/2])
	} else {
		fmt.Println("Test CreateCentralBoard: Correct pile position!")
	}

}

func TestCreateRandomBoard(t *testing.T) {
	type test struct {
		pile int
		size int
	}
	// Test 1
	var t1 test
    t1.pile = 1000
	t1.size = 100
	outcome := EmptyBoard(t1.size)
	outcomePile := 0
	outcome.CreateRandomBoard(t1.pile)
	for i := 0; i < len(outcome); i++ {
		for j := 0; j < len(outcome[0]); j++ {
			outcomePile += outcome[i][j]
		}
	}
	// Check if piles are the same as input
	if outcomePile != t1.pile {
		t.Errorf("Test CreateRandomBoard: Expected number of pile %d, got %d", t1.pile, outcomePile)
	} else {
		fmt.Println("Test CreateRandomBoard: Correct pile number!")
	}
	// Test 2
	var t2 test
    t2.pile = 4000
	t2.size = 10
	outcome2 := EmptyBoard(t2.size)
	outcomePile2 := 0
	outcome2.CreateRandomBoard(t2.pile)
	for i := 0; i < len(outcome2); i++ {
		for j := 0; j < len(outcome2[0]); j++ {
			outcomePile2 += outcome2[i][j]
		}
	}
	if outcomePile2 != t2.pile {
		t.Errorf("Test CreateRandomBoard: Expected number of pile %d, got %d", t2.pile, outcomePile2)
	} else {
		fmt.Println("Test CreateRandomBoard: Correct pile number!")
	}
}
func TestAddTowBoard(t *testing.T) {
	type test struct {
		firstBord Board
		secondBord Board
		answer Board
    }
	// Test 1
	var t1 test
    t1.firstBord = EmptyBoard(5)
	t1.secondBord = EmptyBoard(5)
	t1.answer = EmptyBoard(5)
	for i := 0; i < len(t1.firstBord); i++ {
		for j := 0; j < len(t1.firstBord); j++ {
			t1.firstBord[i][j] = 1
			t1.secondBord[i][j] = 2
			t1.answer[i][j] = 3
		}
	}
	t1.firstBord.AddTwoBoard(t1.secondBord)
	for i := 0; i < len(t1.firstBord); i++ {
		for j := 0; j < len(t1.firstBord[0]); j++ {
			// Check if successfully add secondBoard to firstBord
			if t1.firstBord[i][j] != t1.answer[i][j] {
				t.Errorf("Test AddTowBoard: (%d, %d)Expected %d, got %d", i, j, t1.answer[i][j], t1.firstBord[i][j])
			}
		}
	}
	// Test 2
	var t2 test
    t2.firstBord = EmptyBoard(10)
	t2.secondBord = EmptyBoard(10)
	t2.answer = EmptyBoard(10)
	for i := 0; i < len(t1.firstBord); i++ {
		for j := 0; j < len(t1.firstBord); j++ {
			t1.firstBord[i][j] = 20
			t1.secondBord[i][j] = 30
			t1.answer[i][j] = 50
		}
	}
	t2.firstBord.AddTwoBoard(t2.secondBord)
	for i := 0; i < len(t2.firstBord); i++ {
		for j := 0; j < len(t2.firstBord[0]); j++ {
			if t2.firstBord[i][j] != t2.answer[i][j] {
				t.Errorf("Test AddTowBoard: (%d, %d)Expected %d, got %d", i, j, t2.answer[i][j], t2.firstBord[i][j])
			}
		}
	}

}
func TestRunningSerialGame(t *testing.T) {
	type test struct {
        answer Board
		outcomeBoard Board
    }
	// Test 1
	var t1 test
	t1.answer = EmptyBoard(3)
	t1.answer[0][1] = 1
	t1.answer[1][0] = 1
	t1.answer[1][1] = 1
	t1.answer[1][2] = 1
	t1.answer[2][1] = 1
	t1.outcomeBoard = EmptyBoard(3)
	t1.outcomeBoard.CreateCentralBoard(5)
	t1.outcomeBoard.RunningSerialGame()
	// check if outcomeBoard has same size as expected
	if len(t1.outcomeBoard) != len(t1.answer) {
		t.Errorf("Test RunningSerialGame: Expected number of row %d, got %d", len(t1.answer), len(t1.outcomeBoard))
	} else {
		fmt.Println("Test RunningSerialGame: Correct row number!")
	}
	if len(t1.outcomeBoard[0]) != len(t1.answer[0]) {
		t.Errorf("Test RunningSerialGame: Expected number of col %d, got %d", len(t1.answer[0]), len(t1.outcomeBoard[0]))
	} else {
		fmt.Println("Test RunningSerialGame: Correct col number!")
	}
	// Check if final Board is as expected
	for i := 0; i < len(t1.outcomeBoard); i++ {
		for j := 0; j < len(t1.outcomeBoard[0]); j++ {
			if t1.outcomeBoard[i][j] != t1.answer[i][j] {
				t.Errorf("Test RunningSerialGame: (%d, %d)Expected %d, got %d", i, j, t1.answer[i][j], t1.outcomeBoard[i][j])
			}
		}
	}
	// Test 2
	var t2 test
	t2.answer = EmptyBoard(3)
	t2.answer[0][0] = 2
	t2.answer[0][1] = 1
	t2.answer[0][2] = 2
	t2.answer[1][0] = 1
	t2.answer[1][1] = 1
	t2.answer[1][2] = 1
	t2.answer[2][0] = 2
	t2.answer[2][1] = 1
	t2.answer[2][2] = 2
	t2.outcomeBoard = EmptyBoard(3)
	t2.outcomeBoard.CreateCentralBoard(17)
	t2.outcomeBoard.RunningSerialGame()
	if len(t2.outcomeBoard) != len(t2.answer) {
		t.Errorf("Test RunningSerialGame: Expected number of row %d, got %d", len(t2.answer), len(t2.outcomeBoard))
	} else {
		fmt.Println("Test RunningSerialGame: Correct row number!")
	}
	if len(t2.outcomeBoard[0]) != len(t2.answer[0]) {
		t.Errorf("Test RunningSerialGame: Expected number of col %d, got %d", len(t2.answer[0]), len(t2.outcomeBoard[0]))
	} else {
		fmt.Println("Test RunningSerialGame: Correct col number!")
	}
	for i := 0; i < len(t2.outcomeBoard); i++ {
		for j := 0; j < len(t2.outcomeBoard[0]); j++ {
			if t2.outcomeBoard[i][j] != t2.answer[i][j] {
				t.Errorf("Test RunningSerialGame: (%d, %d)Expected %d, got %d", i, j, t2.answer[i][j], t2.outcomeBoard[i][j])
			}
		}
	}
}

func TestCalculateBoardSerial(t *testing.T) {
	type test struct {
        answer Board
		outcomeBoard Board
		answerRun bool
    }
	// Test 1
	var t1 test
	t1.answerRun = true
    t1.answer = EmptyBoard(3)
	t1.answer[0][0] = 0
	t1.answer[0][1] = 1
	t1.answer[0][2] = 0
	t1.answer[1][0] = 1
	t1.answer[1][1] = -4
	t1.answer[1][2] = 1
	t1.answer[2][0] = 0
	t1.answer[2][1] = 1
	t1.answer[2][2] = 0
	t1.outcomeBoard = EmptyBoard(3)
	t1.outcomeBoard.CreateCentralBoard(9)
	b1, run1 := CalculateBoardSerial(t1.outcomeBoard)
	// Check if run status is as expected
	if run1 != t1.answerRun {
		t.Errorf("Test CalculateBoardSerial: Expected status of run %t, got %t", t1.answerRun, run1)
	} else {
		fmt.Println("Test CalculateBoardSerial: Correct run status!")

	}
	// Check if adding Board is as expected
	for i := 0; i < len(b1); i++ {
		for j := 0; j < len(b1); j++ {
			if b1[i][j] != t1.answer[i][j] {
				t.Errorf("Test CalculateBoardSerial: (%d, %d)Expected %d, got %d", i, j, t1.answer[i][j], b1[i][j])
			}
		}
	}
	// Test 2
	var t2 test
	t2.answerRun = false
    t2.answer = EmptyBoard(3)
	t2.answer[0][0] = 0
	t2.answer[0][1] = 0
	t2.answer[0][2] = 0
	t2.answer[1][0] = 0
	t2.answer[1][1] = 0
	t2.answer[1][2] = 0
	t2.answer[2][0] = 0
	t2.answer[2][1] = 0
	t2.answer[2][2] = 0
	t2.outcomeBoard = EmptyBoard(3)
	t2.outcomeBoard.CreateCentralBoard(3)
	b2, run2 := CalculateBoardSerial(t2.outcomeBoard)
	if run2 != t2.answerRun {
		t.Errorf("Test CalculateBoardSerial: Expected status of run %t, got %t", t2.answerRun, run2)
	} else {
		fmt.Println("Test CalculateBoardSerial: Correct run status!")

	}
	for i := 0; i < len(b2); i++ {
		for j := 0; j < len(b2[0]); j++ {
			if b2[i][j] != t2.answer[i][j] {
				t.Errorf("Test CalculateBoardSerial: (%d, %d)Expected %d, got %d", i, j, t2.answer[i][j], b2[i][j])
			}
		}
	}
}

// Compare the outcomes between Serial and Parallel
func TestCompareSerialParallel(t *testing.T) {
	type test struct {
		serial Board
		parallel Board
	}
	// Test 1
	var t1 test
	t1.serial = EmptyBoard(200)
    t1.serial.CreateCentralBoard(1000)
	t1.serial.RunningSerialGame()
	numProcs := runtime.NumCPU()
	t1.parallel = EmptyBoard(200)
	t1.parallel.CreateCentralBoard(1000)
	t1.parallel.RunningPar(numProcs)
	// Check if each square has same value
	for i := 0; i < len(t1.serial); i++ {
		for j := 0; j < len(t1.serial[0]); j++ {
			if t1.serial[i][j] != t1.parallel[i][j] {
				t.Errorf("Test CompareSerialParallel: (%d, %d)Expected serial %d, got %d in parallel", i, j, t1.serial[i][j], t1.parallel[i][j])
			}
		}
	}
	// Test 2
	var t2 test
	t2.serial = EmptyBoard(300)
    t2.serial.CreateCentralBoard(5000)
	t2.serial.RunningSerialGame()
	t2.parallel = EmptyBoard(300)
	t2.parallel.CreateCentralBoard(5000)
	t2.parallel.RunningPar(numProcs)
	for i := 0; i < len(t2.serial); i++ {
		for j := 0; j < len(t2.serial[0]); j++ {
			if t2.serial[i][j] != t2.parallel[i][j] {
				t.Errorf("Test CompareSerialParallel: (%d, %d)Expected serial %d, got %d in parallel", i, j, t2.serial[i][j], t2.parallel[i][j])
			}
		}
	}
}