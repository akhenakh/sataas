package sgp4

import (
	"testing"
)

const (
	tle1 = "1 25544U 98067A   20099.88933163  .00000859  00000-0  23991-4 0  9991"
	tle2 = "2 25544  51.6455 331.3904 0003968  99.3691 357.2762 15.48691931221277"
)

func TestNewTLE(t *testing.T) {
	tests := []struct {
		name    string
		tle1    string
		tle2    string
		notNil  bool
		wantErr bool
	}{
		{name: "empty", tle1: "", tle2: "", notNil: false, wantErr: true},
		{name: "bogus", tle1: "234324 234 234324", tle2: "234234 234432", notNil: false, wantErr: true},
		{name: "valid", tle1: tle1, tle2: tle2, notNil: true, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt := tt
			gotTle, err := NewTLE(tt.tle1, tt.tle2)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTLE() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if (gotTle != nil) != tt.notNil {
				t.Errorf("NewTLE() tle = %v, want %v", gotTle, tt.notNil)
			}
		})
	}
}
