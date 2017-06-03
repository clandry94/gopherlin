package main

import (
	"encoding/csv"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
)

var r = rand.New(rand.NewSource(219381908309281094))

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

	// Interpolation weights
	sX := x - math.Floor(x)
	sY := y - math.Floor(y)

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

	ix0 := linearInterpolate(dotUl, dotUr, sX)
	ix1 := linearInterpolate(dotBl, dotBr, sX)

	return linearInterpolate(ix0, ix1, sY)
}

func linearInterpolate(a0 float64, a1 float64, w float64) float64 {
	return ((1.0 - w) * a0) + (w * a1)
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
	testPoint := Point{0.0, 0.0}
	file, err := os.Create("perlin2.csv")
	if err != nil {
		fmt.Println("bad")
	}
	defer file.Close()
	w := csv.NewWriter(file)

	for i := 0; i < 600; i++ {
		if testPoint.x >= 6.0 {
			testPoint.x = 0
		}
		testPoint.y = 0.0
		for j := 0; j < 600; j++ {
			if testPoint.y >= 6.0 {
				break
			}
			//fmt.Printf("TESTING: %v, %v with i,j = %v,%v\n", testPoint.x, testPoint.y, i, j)
			value := perlin(testPoint, vectorGrid)
			fmt.Printf("Perlin(%v, %v) = %v\n", testPoint.x, testPoint.y, value)
			x := strconv.FormatFloat(testPoint.x, 'f', 6, 64)
			y := strconv.FormatFloat(testPoint.y, 'f', 6, 64)
			val := strconv.FormatFloat(value, 'f', 6, 64)

			record := []string{x, y, val}
			w.Write(record)
			testPoint.y += 0.01
		}
		testPoint.x += 0.01
	}
	w.Flush()
}
