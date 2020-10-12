// Copyright © 2020 Pieter07. All rights reserved.
// Use of this source code is governed by a GNU General Public License that can be found in the LICENSE file.
package rates

import (
	"errors"
	"math"
)

// Takes in the nominal interest rate and the number of compoundings per period and returns the effective annual interest rate.
// Uses the formula: i = (1 + ip / p) ^ p - 1.
func Effective(ip float64, p int) (float64, error) {
	if p < 0 {
		return 0, errors.New("The number of compoundings per period must be greater than 0")
	}
	return math.Pow(1+ip/float64(p), float64(p)) - 1, nil
}

// Takes in the annual effective interest rate and the number of compoundings per period and returns the nominal interest rate per period.
// Uses the formula: ip = p * ((1 + i)^(1/p) - 1).
func Nominal(i float64, p int) (float64, error) {
	if p < 0 {
		return 0, errors.New("The number of compoundings per period must be greater than 0")
	}
	return float64(p) * (math.Pow(1+i, 1/float64(p)) - 1), nil
}

// Takes in the annual effective interest rate and converts it into the continuously compounded force of interest.
// Uses the formula: delta = ln(1 + i).
func EffectiveToContinuousForce(i float64) float64 {
	return math.Log(1 + i)
}

// Takes in the continuously compounded force of interest and converts it into the annual effective interest rate.
// Uses the formula: i = exp(delta) - 1.
func ContinuousForceToEffective(delta float64) float64 {
	return math.Exp(delta) - 1
}
