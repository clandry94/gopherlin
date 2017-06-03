package main

import "fmt"
import "math"
import "math/rand"

var r = rand.New(rand.NewSource(831994))

type Point struct {
	x float64
	y float64
}

type GradientVector struct {
	x float64
	y float64
}

func perlin(curPoint Point, vectorGrid [7][7]GradientVector) float64 {
	x := curPoint.x
	y := curPoint.y

	ul := Point{math.Floor(x), math.Ceil(y)}
	ur := Point{math.Ceil(x), math.Floor(y)}
	bl := Point{math.Floor(x), math.Ceil(y)}
	br := Point{math.Ceil(x), math.Ceil(y)}

	// Calculate the distance vectors
	distToUl := distVector(ul, curPoint)
	distToUr := distVector(ur, curPoint)
	distToBl := distVector(bl, curPoint)
	distToBr := distVector(br, curPoint)

	// Calculate dot products
	dotUl := dotProd(distToUl, vectorGrid[int(ul.x)][int(ul.y)])
	dotUr := dotProd(distToUr, vectorGrid[int(ur.x)][int(ur.y)])
	dotBl := dotProd(distToBl, vectorGrid[int(bl.x)][int(bl.y)])
	dotBr := dotProd(distToBr, vectorGrid[int(br.x)][int(br.y)])

	fmt.Printf("ul: %v, %v \n", ul.x, ul.y)
	fmt.Printf("ur: %v, %v \n", ur.x, ur.y)
	fmt.Printf("bl: %v, %v \n", bl.x, bl.y)
	fmt.Printf("br: %v, %v \n", br.x, br.y)

	fmt.Println("Dot products: ")
	fmt.Println(dotUl)
	fmt.Println(dotUr)
	fmt.Println(dotBl)
	fmt.Println(dotBr)

	return 0.0
}

func dotProd(distVec Point, gradientVec GradientVector) float64 {
	return (distVec.x * gradientVec.x) + (distVec.y * gradientVec.y)
}

func distVector(node Point, curPoint Point) Point {
	var dist Point
	dist.x = curPoint.x - node.x
	dist.y = curPoint.y - node.y
	return dist
}

func createGradientVectorGrid(size int) [7][7]GradientVector {
	var grid [7][7]GradientVector
	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			randRad := r.Float64() * (2 * math.Pi)
			grid[i][j] = GradientVector{math.Cos(randRad), math.Sin(randRad)}
			fmt.Printf("Position %v, %v = %v,%v \n", i, j, grid[i][j].x, grid[i][j].y)
		}
		fmt.Println()
	}

	return grid
}

func main() {
	vectorGrid := createGradientVectorGrid(7)
	fmt.Println("Hello!")
	testPoint := Point{2.43, 3.21}
	perlin(testPoint, vectorGrid)
}
