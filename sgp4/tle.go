package sgp4

import (
	"github.com/akhenakh/sataas/cppsgp4"
)

type TLE struct {
	ctle cppsgp4.Tle
}

func NewTLE(tle1, tle2 string) (tle *TLE, err error) {
	defer catch(&err)
	ctle := cppsgp4.NewTle(tle1, tle2)

	tle = &TLE{ctle: ctle}
	return tle, nil
}

func (tle *TLE) NoradNumber() int {
	return int(tle.ctle.NoradNumber())
}

func (tle *TLE) Line1() string {
	return tle.ctle.Line1()
}

func (tle *TLE) Line2() string {
	return tle.ctle.Line2()
}
