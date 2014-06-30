package main

import (
	"fmt"
	m "math"
)

type TimeSeriesData []float64

type TimeSeries struct {
	size int
	data TimeSeriesData
}

// for funsies
func sum(data []float64) (val float64) {
	for i := 0; i < len(data); i++ {
		val += data[i]
	}
	return
}

//also for funsies
func (t *TimeSeries) mean() (mean float64) {
	for i := 0; i < t.size; i++ {
		mean += t.data[i] / float64(t.size)
	}
	return
}

//weighted mean
// this is a normalized weighted mean, so the weights must sum to 1
// returns false if the weights length and the data length don't match, or if sum of weights > 1
func (t *TimeSeries) weightedMean(weights []float64) (float64, bool) {
	//make sure length of weights checks out
	if len(weights) != t.size {
		fmt.Println("Weights must be same length as data")
		return 0.0, false
	}
	//make sure sum of weights is 1.0
	if sum(weights) != 1.0 {
		return 0.0, false
	}

	mean := 0.0
	for i := 0; i < t.size; i++ {
		mean += t.data[i] * weights[i]
	}
	return mean, true

}

//variance in array of floats
func (t *TimeSeries) variance() float64 {
	average := t.mean()
	variance := 0.0
	for i := 0; i < t.size; i++ {
		diff := t.data[i] - average
		variance += (diff * diff)
	}
	return variance / float64(t.size)
}

//weighted variance - weights should sum to 1
func (t *TimeSeries) weightedVariance(weights []float64) float64 {
	average := t.mean()
	variance := 0.0
	for i := 0; i < t.size; i++ {
		diff := t.data[i] - average
		variance += (diff * diff) * weights[i]
	}
	return variance
}

// covariance & Correlation
// i got too lazy to make them separate functions
func (x *TimeSeries) covCorr(y *TimeSeries) (float64, float64, bool) {
	if x.size == y.size {
		cov := 0.0
		for i := 0; i < x.size; i++ {
			xVar := x.data[i] - x.mean()
			yVar := y.data[i] - y.mean()
			cov += (xVar * yVar) / float64(x.size)
		}

		return cov, cov / (x.sd() * y.sd()), true
	}
	return 0.0, 0.0, false
}

// standard deviation in array of floats
func (t *TimeSeries) sd() float64 {
	return m.Sqrt(t.variance())
}

//autocovariance function
func (t *TimeSeries) autoCov(lag int) (acf float64) {
	tbar := t.mean()
	for i := 0; i < (t.size - lag); i++ {
		acf += (t.data[i] - tbar) * (t.data[i+lag] - tbar)
	}
	return acf / float64(t.size)
}

//autocorrelation function
func (t *TimeSeries) autoCorr(lag int) (acf float64) {
	tbar := t.mean()
	for i := 0; i < (t.size - lag); i++ {
		acf += (t.data[i] - tbar) * (t.data[i+lag] - tbar)
	}
	return acf / (t.variance() * float64(t.size))
}

//remove value at lag k
func (t *TimeSeries) removeLag(lag int) *TimeSeries {
	if lag > t.size {
		return nil
	}
	newvals := make([]float64, t.size-1)
	for i := 0; i < t.size; i++ {
		if i != lag {
			newvals = append(newvals, t.data[i])
		}
	}
	return &TimeSeries{size: t.size - 1, data: newvals}
}

// function to calculate the partial auto-correlation
// otherwise known as the PACF
func (t *TimeSeries) pacf(lag int) float64 {
	newTS := t.removeLag(lag)
	return newTS.autoCorr(lag - 1)
}
