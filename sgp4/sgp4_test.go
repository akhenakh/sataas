package sgp4

import (
	"testing"
	"time"
)

const (
	tle1 = "1 25544U 98067A   20101.90690972 -.00000449  00000-0  00000+0 0  9993"
	tle2 = "2 25544  51.6446 321.4198 0003848 108.5166  84.0719 15.48680394221581"
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
		{name: "valid", tle1: tle1, tle2: tle2, notNil: true, wantErr: false, wantNoradNumber: 25544},
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
		})
	}
}

func TestSGP4_FindPosition(t *testing.T) {
	tle, err := NewTLE(tle1, tle2)
	if err != nil {
		t.Fatal(err)
	}
	p, err := NewSGP4(tle)
	if err != nil {
		t.Fatal(err)
	}

	tests := []struct {
		name    string
		tString string
		wantLat float64
		wantLng float64
		wantAlt float64
	}{
		{
			name:    "ISS",
			tString: "2020-04-10T22:19:26-04:00",
			wantLat: 6.091596525354912,
			wantLng: -98.96289626192333,
			wantAlt: 419.552968612672,
		},
		{
			name:    "ISS",
			tString: "2020-04-10T22:21:31-04:00",
			wantLat: -0.2721344053329822,
			wantLng: -94.46077407870163,
			wantAlt: 420.51309857543765,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t1, err := time.Parse(time.RFC3339, tt.tString)
			if err != nil {
				t.Error(err)
			}
			gotLat, gotLng, gotAlt := p.FindPosition(t1)
			if gotLat != tt.wantLat {
				t.Errorf("FindPosition() gotLat = %v, want %v", gotLat, tt.wantLat)
			}
			if gotLng != tt.wantLng {
				t.Errorf("FindPosition() gotLng = %v, want %v", gotLng, tt.wantLng)
			}
			if gotAlt != tt.wantAlt {
				t.Errorf("FindPosition() gotAlt = %v, want %v", gotAlt, tt.wantAlt)
			}
		})
	}
}
