package parallel

import (
	"time"
	"math/rand"
)

var MatrixA = [][]int{}
var MatrixB = [][]int{}
var MatrixC = [][]int{}

func Setup(size int) {
	rand.Seed(time.Now().UTC().UnixNano())
	MatrixA = [][]int{}
	MatrixB = [][]int{}
	MatrixC = [][]int{}

	for i := 0; i < size; i++ {
		rowA := make([]int, size)
		rowB := make([]int, size)
		//rowC := make([]int, size)
		for j:=0; j < size; j++ {
			rowA[j] = rand.Intn(100)
			rowB[j] = rand.Intn(100)
		}
		MatrixA = append(MatrixA, rowA)
		MatrixB = append(MatrixB, rowB)
		MatrixC = append(MatrixC, make([]int, size))
	}
}

func Run(size int, parallelExecutions int) {
	MatrixC = [][]int{}
	for i := 0; i < size; i++ {
		MatrixC = append(MatrixC, make([]int, size))
	}
	c := make(chan int)
	for i := 0; i < parallelExecutions; i++ {
		go multiply(i*size/parallelExecutions, (i+1)*size/parallelExecutions, size, &MatrixA, &MatrixB, &MatrixC, c)
	}
	for i := 0; i < parallelExecutions; i++ {
		<-c
	}
}

func multiply(indexNum int, indexEnd int, size int, MatrixA_ptr *[][]int, MatrixB_ptr *[][]int, MatrixC_ptr *[][]int, ch chan<- int){
	for i := indexNum; i < indexEnd; i++ {
		for j:=0; j < size; j++ {
			var sum int=0
			for k:=0; k<size; k++ {
				sum = sum + (*MatrixA_ptr)[i][k] * (*MatrixB_ptr)[k][j]
			}
			(*MatrixC_ptr)[i][j] = sum
		}
	}
	ch <- 0
}