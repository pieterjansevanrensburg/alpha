// Copyright © 2020 Pieter07. All rights reserved.
// Use of this source code is governed by a GNU General Public License that can be found in the LICENSE file.
package tvm

import (
	"fmt"
	"math"

	"gonum.org/v1/gonum/mat"
)

func TimeValue(tp float64, i float64, T mat.Vector, Ct mat.Vector) float64 {
	v := math.Pow(1+i, -1)

	n, _ := Ct.Dims()

	fmt.Println("n=", n)

	vArray := make([]float64, n)

	for i := 0; i < n; i++ {
		vArray[i] = math.Pow(v, T.At(i, 0))
	}

	V, pvZero := mat.NewVecDense(n, vArray), mat.NewVecDense(1, nil)

	pvZero.MulVec(Ct.T(), V)

	answer := pvZero.At(0, 0) * math.Pow(1+i, tp)

	return answer
}
