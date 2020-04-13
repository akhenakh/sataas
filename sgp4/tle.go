package sgp4

import (
	"fmt"

	"github.com/akhenakh/sataas/cppsgp4"
)

type TLE struct {
	ctle cppsgp4.Tle
}

func NewTLE(name, tle1, tle2 string) (tle *TLE, err error) {
	defer catch(&err)
	ctle := cppsgp4.NewTle(name, tle1, tle2)

	tle = &TLE{ctle: ctle}
	return tle, nil
}

func (tle *TLE) NoradNumber() int {
	return int(tle.ctle.NoradNumber())
}

func (tle *TLE) String() string {
	return fmt.Sprintf("<TLE: %d %s>", tle.NoradNumber(), tle.Name())
}

func (tle *TLE) Line1() string {
	return tle.ctle.Line1()
}

func (tle *TLE) Line2() string {
	return tle.ctle.Line2()
}

func (tle *TLE) Name() string {
	return tle.ctle.Name()
}
