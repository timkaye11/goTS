## Time Series in GO!
I had some trouble falling asleep tonight (6/29/2014) so I decided to write some basic time series functions in Go. So far this supports autocovariance, autocorrelation, and partial auto-correlation. I'll be adding some more functionality soon!

For an example, try :
```
	data := []float64{1.33, 2.33, 3.22, 4.55}
	data2 := []float64{7.73, 1.38, 2.22, 5.55, }

	ts := TimeSeries{size: 4, data: data}
	ts2 := TimeSeries{size: 4, data: data2}

	// covariacne and correlation
	fmt.Println(ts.covCorr(&ts2))
	//partial autocorrelation at lag 1
	fmt.Println(ts.pacf(1))
```

more to come soon...
