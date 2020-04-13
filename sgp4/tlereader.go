package sgp4

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

type TLEReader struct {
	*bufio.Scanner
}

func NewTLEReader(r io.Reader) *TLEReader {
	return &TLEReader{Scanner: bufio.NewScanner(r)}
}

func (r *TLEReader) ReadAllTLE() ([]*TLE, error) {
	var tles []*TLE
	count := 0
	lines := make([]string, 3)
	for r.Scan() {
		lines[count] = r.Text()
		if err := r.Err(); err != nil {
			return nil, err
		}
		if count >= 2 {
			tle, err := NewTLE(strings.TrimSpace(lines[0]), lines[1], lines[2])
			if err != nil {
				fmt.Printf("error tle [%s]\n[%s]", lines[1], lines[2])
				return nil, err
			}
			tles = append(tles, tle)
			count = 0
			continue
		}
		count++
	}

	return tles, nil
}
