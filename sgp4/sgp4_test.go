package sgp4_test

import (
	"testing"
	"time"

	"github.com/stretchr/testify/require"

	"github.com/akhenakh/sataas/sgp4"
)

const (
	tle1 = "1 25544U 98067A   20101.90690972 -.00000449  00000-0  00000+0 0  9993"
	tle2 = "2 25544  51.6446 321.4198 0003848 108.5166  84.0719 15.48680394221581"
)

func TestSGP4_FindPosition(t *testing.T) {
	tle, err := sgp4.NewTLE("ISS", tle1, tle2)
	if err != nil {
		t.Fatal(err)
	}
	p, err := sgp4.NewSGP4(tle)
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
			require.NoError(t, err)
			gotLat, gotLng, gotAlt, err := p.Position(t1)
			require.NoError(t, err)

			require.InDeltaf(t, tt.wantLat, gotLat, 0.0000001, "invalid lat not in deltra")
			require.InDeltaf(t, tt.wantLng, gotLng, 0.0000001, "invalid lng not in deltra")
			require.InDeltaf(t, tt.wantAlt, gotAlt, 0.0000001, "invalid alt not in deltra")
		})
	}
}
