package main

import "fmt"
import "math"
import "math/rand"

var r = rand.New(rand.NewSource(1111))

type Point struct {
	x int
	y int
}

type GradientVector struct {
	x   int
	y   int
	mag int
}

func perlin(x float64, y float64) float64 {

	ul := Point{int(math.Floor(x)), int(math.Ceil(y))}
	ur := Point{int(math.Ceil(x)), int(math.Floor(y))}
	bl := Point{int(math.Floor(x)), int(math.Ceil(y))}
	br := Point{int(math.Ceil(x)), int(math.Ceil(y))}

	fmt.Printf("ul: %v, %v \n", ul.x, ul.y)
	fmt.Printf("ur: %v, %v \n", ur.x, ur.y)
	fmt.Printf("bl: %v, %v \n", bl.x, bl.y)
	fmt.Printf("br: %v, %v \n", br.x, br.y)

	return 0.0
}

func createGradientVectorGrid(size int) [7][7]GradientVector {
	var grid [7][7]GradientVector
	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			grid[i][j] = GradientVector{r.Int(), r.Int(), 1}
			fmt.Printf(" | %v,%v | ", grid[i][j].x, grid[i][j].y)
		}
		fmt.Println()
	}

	return grid
}

func main() {
	createGradientVectorGrid(7)
	fmt.Println("Hello!")
	perlin(2.43, 3.21)
}
