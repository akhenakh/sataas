package sgp4_test

import (
	"os"
	"testing"

	"github.com/akhenakh/sataas/sgp4"
)

func TestNewTLE(t *testing.T) {
	tests := []struct {
		name            string
		tle1            string
		tle2            string
		notNil          bool
		wantErr         bool
		wantNoradNumber int
	}{
		{name: "empty", tle1: "", tle2: "", notNil: false, wantErr: true},
		{name: "bogus", tle1: "234324 234 234324", tle2: "234234 234432", notNil: false, wantErr: true},
		{name: "iss", tle1: tle1, tle2: tle2, notNil: true, wantErr: false, wantNoradNumber: 25544},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			gotTle, err := sgp4.NewTLE(tt.name, tt.tle1, tt.tle2)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTLE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (gotTle != nil) != tt.notNil {
				t.Errorf("NewTLE() tle = %v, want %v", gotTle, tt.notNil)
				return
			}

			if (gotTle != nil) && tt.wantNoradNumber != gotTle.NoradNumber() {
				t.Errorf("NoradNumber() = %v, wantCatalogNumber %v", gotTle.NoradNumber(), tt.wantNoradNumber)
				return
			}

			if (gotTle != nil) && tt.tle1 != gotTle.Line1() {
				t.Errorf("Line1() = %v, tle1 %v", gotTle.Line1(), tt.tle1)
				return
			}

			if (gotTle != nil) && tt.tle2 != gotTle.Line2() {
				t.Errorf("Line2() = %v, tle2 %v", gotTle.Line2(), tt.tle2)
				return
			}

			if (gotTle != nil) && tt.name != gotTle.Name() {
				t.Errorf("Name() = %v, name %v", gotTle.Name(), tt.name)
				return
			}
		})
	}
}

func TestActives(t *testing.T) {
	f, err := os.Open("testdata/active.txt")
	if err != nil {
		t.Fatal(err)
	}

	r := sgp4.NewTLEReader(f)
	tles, err := r.ReadAllTLE()
	if err != nil {
		t.Fatal(err)
	}

	// 8232 lines
	if len(tles) != 8226/3 {
		t.Fatal("invalid sats count")
	}
}
