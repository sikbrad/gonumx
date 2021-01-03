package gony

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
)

func Zeros(row, col int) *mat.Dense{
	arr := make([]float64, row*col)
	for idx, elem := range arr{
		arr[idx] = elem
	}
	gd := mat.NewDense(row, col, arr)
	return gd
}

func PrintMatrix(mtx mat.Matrix, name string) {
	PrintMatrixWithFormatter(mtx, name, "%4.2f")
}

func PrintMatrixWithFormatter(mtx mat.Matrix, name, numFormat string) {
	rows, cols := mtx.Dims()
	fmt.Printf("prints matrix [%s], dim[%v,%v]\n", name, rows, cols)
	fa := mat.Formatted(mtx, mat.Prefix("    "), mat.Squeeze())
	fmtStr := fmt.Sprintf("%%s = %s\n", numFormat)
	fmt.Printf(fmtStr, name, fa)
}



//type Gmx interface{
//	mat.Matrix
//}
//type Gdense struct{
//	mat.Dense
//}
//func NewGdense(dense *mat.Dense) *Gdense{
//	return &Gdense{
//		*dense,
//	}
//}
//
//func Array(arr []float64) *mat.VecDense{
//	return mat.NewVecDense(len(arr), arr)
//}
//
////func (gd* Gdense) Array(arr []float64) *Gdense{
//func (gx* Gdense) Zeros(row, col int) *Gdense{
//	arr := make([]float64, row*col)
//	for idx, elem := range arr{
//		arr[idx] = elem
//	}
//
//	gd := NewGdense(mat.NewDense(row, col, arr))
//	return gd
//}
