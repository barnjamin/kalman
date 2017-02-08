package particle

import (
	"log"
	"math/rand"
	"testing"
)

func TestNew(t *testing.T) {
	u := []float64{1, 1}
	std := []float64{.1, .1}

	pf := New(5000, [][]float64{{0, .2}, {0, .2}}, Stratified)

	for x := 1; x < 20; x++ {
		X := (rand.Float64() * 0.1) + float64(x)
		Y := (rand.Float64() * 0.1) + float64(x)
		measurements := []float64{X, Y}

		pf.Predict(u, std)

		pf.Update(measurements, 0.1)

		pf.Resample()

		mu, std := pf.Estimate()

		log.Printf("%+v %+v", mu, std)
	}

}
