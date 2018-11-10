package models

import (
	"sort"
)

type Engine struct {
	Id int
	Code string
	Diameter float64
	Length float64
	Delay int
	PropellantWeight float64
	EngineWeight float64
	Manufacturer string
	BurnTime float64
	ThrustCurvePoints []Datapoint
	MassCurvePoints []Datapoint
}

func (e Engine) force_value_at(time float64) float64 {
	if time >= e.BurnTime {
		return 0.0
	}

	idx_start := sort.Search(len(e.ThrustCurvePoints), func(i int) bool {
		return (e.ThrustCurvePoints[i].TimeStamp) >= time	
	})

	idx_end := idx_start + 1

	return interpolate_segment(e.ThrustCurvePoints[idx_start].TimeStamp,
							   e.ThrustCurvePoints[idx_start].Value,
							   e.ThrustCurvePoints[idx_end].TimeStamp, 
							   e.ThrustCurvePoints[idx_end].Value, time)

}

func (e Engine) mass_value_at(time float64) float64 {
	if time >= e.BurnTime {
		return (e.EngineWeight - e.PropellantWeight) 
	}

	idx_start := sort.Search(len(e.MassCurvePoints), func(i int) bool {
		return (e.MassCurvePoints[i].TimeStamp) >= time	
	})

	idx_end := idx_start + 1

	return interpolate_segment(e.ThrustCurvePoints[idx_start].TimeStamp,
							   e.ThrustCurvePoints[idx_start].Value,
							   e.ThrustCurvePoints[idx_end].TimeStamp, 
							   e.ThrustCurvePoints[idx_end].Value, time)


}



// linear interpolation.   Interpolates y value from x value.  In this case x is time
// return value is thrust or mass
// Saturates to y0 or y1 if x outside interval [x0, x1].
func interpolate_segment(x0 float64, y0 float64, x1 float64, y1 float64, x float64) float64 {

	t := 0.0;

    if (x <= x0) { return y0 }
	if (x >= x1) { return y1 }
	if (x1 == x0) { return y0 }

    t =  (x-x0)
    t = t/(x1-x0)

    return y0 + t*(y1-y0)

}



