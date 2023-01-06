package main

import (
	"canvas"
	"fmt"
	"log"
	"math/rand"
	"os"
	"runtime"
	"strconv"
	"time"
)

func main() {
	// Set parameter from command line
	// Arg[1]: for board size
	size, err := strconv.Atoi(os.Args[1])
	if err != nil {
		panic(err)
	} else if size <= 0 {
		// Size should be positive
		fmt.Println("Please enter positive integers")
	}
	// Arg[2]: for coins number
	pile, err := strconv.Atoi(os.Args[2])
	if err != nil {
		panic(err)
	} else if pile <= 0 {
		// Coins should be positive
		fmt.Println("Please enter positive integers")
	}
	// Arg[3]: for different initial board
	placement := os.Args[3]
	if placement != "central" && placement != "random" {
		// Could only be "central", "random"
		fmt.Println("Please enter 'central' or 'random'")
	}

	// Create initial board for Serial
	s := EmptyBoard(size)
	if placement == "central" {
		s.CreateCentralBoard(pile)
	} else {
		s.CreateRandomBoard(pile)
	}
	// Create initial board for Parallel
	p := EmptyBoard(size)
	p.CopyBoard(s)

	////////// For Serial program
	fmt.Println("Start serial program")
	//// Starting time
	startS := time.Now()
	// Running Game
	s.RunningSerialGame()
	//// Stop time
	elapsedS := time.Since(startS)
	log.Printf("Simulating in serial took %s", elapsedS)
	fmt.Println("Start Drawing serial picture")
	s.BoardToImage("s", size)

	////////// For Parallel program
	fmt.Println("Start parallel program")
	// Get the number of CPU
	numProcs := runtime.NumCPU()
	//// Starting time
	start := time.Now()
	// Running Game
	p.RunningPar(numProcs)
	// Stop time
	elapsed := time.Since(start)
	log.Printf("Simulating in parallel took %s", elapsed)
	fmt.Println("Start Drawing parallel picture")
	p.BoardToImage("p", size)

}

// Set Board data type
type Board [][]int

// For channel
type AddBoard struct {
	coins          Board
	topRowIndex    []int
	bottomRowIndex []int
	index          int
	run            bool
}

// Input: size of board
// Output: empty board
// Create a empty board
func EmptyBoard(size int) Board {
	b := make(Board, size)
	// Make all slices inside the cells
	for i := range b {
		b[i] = make([]int, size)
	}
	return b
}

// Input: board, size of board
// Output: board
// Copy board
func (b2 Board) CopyBoard(b1 Board) {
	for i := range b1 {
		for j := range b1[i] {
			b2[i][j] = b1[i][j]
		}
	}
}

// Board Method; Receiver: Board; Input: pile
// Create initial Board for Central System
func (b Board) CreateCentralBoard(pile int) {
	size := len(b)
	// Set all coins in the center
	b[size/2][size/2] = pile
}

// Board Method; Receiver: Board; Input: pile
// Create initial Board for Random System
func (b Board) CreateRandomBoard(pile int) {
	// count for how many coins have been put into board
	var count int
	size := len(b)
	// For 1~99 squares
	rand.Seed(time.Now().UTC().UnixNano())
	for i := 0; i < 99; i++ {
		// If is still has coins can put into board
		if count <= pile {
			// Random choose row, column
			row := rand.Intn(size)
			col := rand.Intn(size)
			// how many coins need to be put
			rand := rand.Intn(pile - count) // Pile-count is remaining coins
			// Add coins to square and count
			b[row][col] += rand
			count += rand
		} else {
			// If coins is used up, break the loop
			break
		}
	}
	// If it still has coins after 99 times
	// Last one square will equal to remaining coins
	if count <= pile {
		row := rand.Intn(size)
		col := rand.Intn(size)
		b[row][col] = pile - count + b[row][col]
	}
}

// Board method; Receiver: Board
// Running game until it is stable
func (b Board) RunningSerialGame() {
	run := b.UpdateSerialBoard()
	// If it's not stable, continuously running
	if run {
		b.RunningSerialGame()
	}
}

// Board Method; Receiver: Board
// Output: bool
// Update Board by using CalculateBoardSerial, and check if board is stable
func (b Board) UpdateSerialBoard() bool {
	// Get the adding board and if it is stable
	add, run := CalculateBoardSerial(b)
	// Add adding board to board
	b.AddTwoBoard(add)
	return run
}

// Board Method; Receiver: Board
// Input: suqure posistion
// Calculate square for total changing coins
func CalculateBoardSerial(b Board) (Board, bool) {
	var count int
	// Set initial run status as false
	run := false
	row := len(b)
	// Create a addboard for each square should add how many coins to board
	addboard := EmptyBoard(row)
	for i := 0; i < row; i++ {
		for j := 0; j < row; j++ {
			// Calculate any condition on b[i][j]
			// If this square is larger than 4, the -4
			if b[i][j] >= 4 {
				count++
				addboard[i][j] = addboard[i][j] - 4
			}
			// Check if north, south, east and west is larger than 4
			// If it is larger than 4, then add 1 to b[i][j]
			if i < row-1 {
				if b[i+1][j] >= 4 {
					addboard[i][j]++
				}
			}
			if i != 0 {
				if b[i-1][j] >= 4 {
					addboard[i][j]++
				}
			}
			if j < row-1 {
				if b[i][j+1] >= 4 {
					addboard[i][j]++
				}
			}
			if j != 0 {
				if b[i][j-1] >= 4 {
					addboard[i][j]++
				}
			}
		}
	}
	if count > 0 {
		// Change run to true if board is unstable
		run = true
	}
	return addboard, run
}

// Board method; Receiver: Board
// Add the coins in two boards together
func (b1 Board) AddTwoBoard(b2 Board) {
	row := len(b1)
	// Loop through boards
	for i := 0; i < row; i++ {
		for j := 0; j < row; j++ {
			b1[i][j] = b1[i][j] + b2[i][j]

		}
	}
}

// For Parallel Running
// Board method; Receiver: Board; Input: number of CPU
// Running game until it is stable
func (b Board) RunningPar(numProcs int) {
	// Use SandpileMultiprocs to calculate how many coins should add to board
	// And check if board is stable
	add, run := SandpileMultiprocs(b, numProcs)
	b.AddTwoBoard(add) // Add boards
	// If it's not stable, continuously running
	if run {
		b.RunningPar(numProcs)
	}
}

// For Parallel Running
// Input: Board, numProcs
// Output: addBoard, bool (if still needs to run)
// Multiprocs; Divide board to several pieces and assign them to each numProcs
// Calculate how much coins change in each squares, and check if board is stable
func SandpileMultiprocs(b Board, numProcs int) (Board, bool) {
	// Create a empty channel
	boardList := make([]AddBoard, numProcs)
	// Set run to false
	run := false
	row := len(b)
	c := make(chan AddBoard, numProcs) // Make channel which len is equals to numProcs
	//Divide board to numProcs
	for i := 0; i < numProcs; i++ {
		// Set Star index and end index
		startIndex := i * (row / numProcs)
		endIndex := (i + 1) * (row / numProcs)
		// Normal, for 1~numProcs-1
		if i < numProcs-1 {
			smallBoard := b[startIndex:endIndex][:]
			// Store index into smallBoard for further arrange small boards
			go SandpileSingleproc(smallBoard, c, i)
		} else { // For last numProcs
			// From start index to end of the board
			smallBoard := b[startIndex:][:]
			go SandpileSingleproc(smallBoard, c, i)
		}
	}
	for i := 0; i < numProcs; i++ {
		// Take out all addBoard from channel
		chanBoard := <-c
		// Arrange boards using index
		boardList[chanBoard.index] = chanBoard
		// If any small boards is still unstable, then run whole board again
		if chanBoard.run == true {
			run = true
		}
	}
	// Handle the boundry (Top, bottom) from each boards on add board
	for i := 0; i < numProcs; i++ {
		if i != 0 {
			// Use the index store in the list of channel to get information
			// that which square should add 1
			for j := 0; j < len(boardList[i].topRowIndex); j++ {
				previousLastRow := len(boardList[i-1].coins) - 1
				boardList[i-1].coins[previousLastRow][boardList[i].topRowIndex[j]]++
			}
		}
		if i < numProcs-1 {
			for j := 0; j < len(boardList[i].bottomRowIndex); j++ {
				boardList[i+1].coins[0][boardList[i].bottomRowIndex[j]]++
			}
		}
	}
	// Combine all small add boards
	var addBoard Board
	for i := 0; i < numProcs; i++ {
		for j := 0; j < len(boardList[i].coins); j++ {
			addBoard = append(addBoard, boardList[i].coins[j])
		}
	}
	return addBoard, run
}

// For Parallel Running
// Input: Board, c channel
// Single procs; each CPU will run this process to calculate small boards
func SandpileSingleproc(b Board, c chan AddBoard, index int) {
	row := len(b)
	col := len(b[0])
	var count int
	var addBoard AddBoard
	// Create an empty board
	addBoard.coins = make([][]int, row)
	for i := 0; i < row; i++ {
		addBoard.coins[i] = make([]int, col)
	}
	// If this square is larger than 4, the -4
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if b[i][j] >= 4 {
				count++
				addBoard.coins[i][j] = addBoard.coins[i][j] - 4
				if i == 0 {
					// If it is first row of small board, then add index to list
					// Because the same index of last row of previous board should add 1
					addBoard.topRowIndex = append(addBoard.topRowIndex, j)

				} else if i == row-1 {
					// Because the same index of first row of next board should add 1
					addBoard.bottomRowIndex = append(addBoard.bottomRowIndex, j)
				}
			}
			// Calculate N, W, E, S of b[i][j], if they are larger than 4 then add 1
			if i != 0 {
				if b[i-1][j] >= 4 {
					addBoard.coins[i][j]++
				}
			}
			if i < row-1 {
				if b[i+1][j] >= 4 {
					addBoard.coins[i][j]++
				}
			}
			if j != 0 {
				if b[i][j-1] >= 4 {
					addBoard.coins[i][j]++
				}
			}
			if j < col-1 {
				if b[i][j+1] >= 4 {
					addBoard.coins[i][j]++
				}
			}
		}
	}
	// Set the fields
	if count == 0 {
		addBoard.run = false
	} else if count > 0 {
		addBoard.run = true
	}
	addBoard.index = index
	// Send to Channel
	c <- addBoard
}

// Reference: from recitation (spatial)
// Create PNG
func (b Board) BoardToImage(name string, size int) {
	row := len(b)
	col := len(b[0])
	// Create a canvas
	c := canvas.CreateNewCanvas(row, col)
	// For each pixel, set color according to coins number
	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			if b[i][j] == 3 {
				c.SetFillColor(canvas.MakeColor(255, 255, 255))

			} else if b[i][j] == 2 {
				c.SetFillColor(canvas.MakeColor(170, 170, 170))

			} else if b[i][j] == 1 {
				c.SetFillColor(canvas.MakeColor(85, 85, 85))

			} else if b[i][j] == 0 {
				c.SetFillColor(canvas.MakeColor(0, 0, 0))

			}
			// Draw
			x1, y1 := j, i
			x2, y2 := (j + 1), (i + 1)
			c.ClearRect(x1, y1, x2, y2)
			c.Fill()
		}
	}
	if name == "p" {
		c.SaveToPNG("Parallel.png")
	} else if name == "s" {
		c.SaveToPNG("Serial.png")
	}
}
