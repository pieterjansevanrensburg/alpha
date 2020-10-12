// Copyright © 2020 Pieter07. All rights reserved.
// Use of this source code is governed by a GNU General Public License that can be found in the LICENSE file.
package rates

import (
	"math"
	"testing"
)

// A test to check whether a nominal interest rate of 10% compounded 12 times per period converts into an annual effective interest rate of 10.471...%.
func TestEffective(t *testing.T) {
	t.Parallel()
	ip, p, iAns := 0.1, 12, 0.10471
	i, _ := Effective(ip, p)
	iDiff := i - iAns
	if math.Round(i*100000) != iAns*100000 {
		t.Errorf("Effective failed (%v, %v) with error of %v", i, iAns, iDiff)
	}
}

// A test to check whether an effective interest rate of 10.471...% converts into a nominal effective interest rate of 10% compounded 12 times per period.
func TestNominal(t *testing.T) {
	t.Parallel()
	i, p, ipAns := 0.10471306744129724159057263529753, 12, 0.10000
	ip, _ := Nominal(i, p)
	ipDiff := ip - ipAns
	if math.Round(ip*100000) != ipAns*100000 {
		t.Errorf("Nominal failed (%v, %v) with error of %v", ip, ipAns, ipDiff)
	}
}

// A test to check whether an effective interest rate of 10% converts into a continuously compounded force of interest of 9.531...%.
func TestEffectiveToContinuousForce(t *testing.T) {
	i, deltaAns := 0.1, 0.09531
	delta := EffectiveToContinuousForce(i)
	deltaDiff := delta - deltaAns
	if math.Round(delta*100000) != deltaAns*100000 {
		t.Errorf("EffectiveToContinuousForce failed (%v, %v) with error of %v", delta, deltaAns, deltaDiff)
	}
}

// A test to check whether a continuously compounded force of interest of 9.531...% converts into an effective interest rate of 10%.
func TestContinuousForceToEffective(t *testing.T) {
	delta, iAns := 0.09531017980432486004395212328077, 0.10000
	i := ContinuousForceToEffective(delta)
	iDiff := i - iAns
	if math.Round(i*100000) != iAns*100000 {
		t.Errorf("ContinuousForceToEffective failed (%v, %v) with error of %v", i, iAns, iDiff)
	}
}
