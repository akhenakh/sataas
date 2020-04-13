package sgp4

import (
	"testing"
	"time"
)

const (
	tle1 = "1 25544U 98067A   20101.90690972 -.00000449  00000-0  00000+0 0  9993"
	tle2 = "2 25544  51.6446 321.4198 0003848 108.5166  84.0719 15.48680394221581"
)

func TestSGP4_FindPosition(t *testing.T) {
	tle, err := NewTLE("ISS", tle1, tle2)
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
			gotLat, gotLng, gotAlt, err := p.FindPosition(t1)
			if err != nil {
				t.Error(err)
			}
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
