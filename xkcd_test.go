package xkcd

import (
	"testing"
)

func Test_getFilename(t *testing.T) {
	type fields struct {
		ID    int
		Year  string
		Month string
		Day   string
	}

	// test template
	tt := []struct {
		name   string
		fields fields
		expect string
	}{
		{"comic #1", fields{ID: 1, Year: "2006", Month: "1", Day: "1"}, "1_2006_1_1.png"},
		{"comic #2", fields{ID: 2, Year: "2006", Month: "1", Day: "1"}, "2_2006_1_1.png"},
		{"comic #1000", fields{ID: 1000, Year: "2012", Month: "1", Day: "6"}, "1000_2012_1_6.png"},
	}
	for _, tc := range tt {
		t.Run(tc.name, func(t *testing.T) {
			c := Comic{ID: tc.fields.ID, Year: tc.fields.Year, Month: tc.fields.Month, Day: tc.fields.Day}
			if c.getFilename() != tc.expect {
				t.Fatalf("sum of %v should be %v; got: %v", tc.name, tc.expect, c)
			}
		})
	}
}
