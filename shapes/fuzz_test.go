package shapes

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Fuzz_Between(t *testing.T) {
	cases := []struct {
		name string
		fuzz Fuzz
		v    float64
		a    float64
		b    float64
		want bool
	}{
		{
			name: "non-fuzz trivial case",
			fuzz: Fuzz(0.1),
			v:    2.0,
			a:    1.0,
			b:    3.0,
			want: true,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.fuzz.Between(tc.v, tc.a, tc.b)
			t.Logf(
				"want: %v, got: %v, v: %f, a: %f, b: %f, fuzz: %v",
				tc.want, got, tc.v, tc.a, tc.b, tc.fuzz,
			)
			assert.Equal(t, tc.want, got)
		})
	}
}

func Test_Fuzz_Eq(t *testing.T) {
	cases := []struct {
		name string
		fuzz Fuzz
		a    float64
		b    float64
		want bool
	}{
		{
			name: "with fuzz 0.00001 then 1.0 eq 1.000001 is true",
			fuzz: Fuzz(0.00001), a: 1.0, b: 1.000001, want: true,
		},
		{
			name: "with fuzz 0.00001 then 1.0 eq 1.001 is false",
			fuzz: Fuzz(0.00001), a: 1.0, b: 1.001, want: false,
		},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.fuzz.Eq(tc.a, tc.b)
			t.Logf(
				"want: %v, got: %v, a: %f, b: %f, fuzz: %v",
				tc.want, got, tc.a, tc.b, tc.fuzz,
			)
			assert.Equal(t, tc.want, got)
		})
	}
}
