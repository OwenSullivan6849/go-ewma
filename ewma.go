// Package ewma is a streaming exponential moving average, dependency-free.
package ewma

type EWMA struct {
	alpha float64
	value float64
	set   bool
}

func New(alpha float64) *EWMA { return &EWMA{alpha: alpha} }

func (e *EWMA) Update(x float64) float64 {
	if !e.set {
		e.value, e.set = x, true
	} else {
		e.value = e.alpha*x + (1-e.alpha)*e.value
	}
	return e.value
}
