package tvm

import (
	"math"
	"testing"

	"gonum.org/v1/gonum/mat"
)

// A test to check whether the present value at time 0 of 3 cashflows of 100 at times 1, 2 and 3 is 248.68520...
func TestTimeValue(t *testing.T) {
	tp, i := 0.0, 0.1
	var T, Ct mat.Vector = mat.NewVecDense(3, []float64{1, 2, 3}), mat.NewVecDense(3, []float64{100, 100, 100})
	pv, pvAns := TimeValue(tp, i, T, Ct), 248.68520
	pvDiff := pv - pvAns

	if math.Round(pv*100000) != pvAns*100000 {
		t.Errorf("TimeValue failed (%v, %v) with error of %v", pv, pvAns, pvDiff)
	}
}
