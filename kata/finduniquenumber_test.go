package kata_test

import (
	. "github.com/AlKoFDC/codewars/kata"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

var _ = Describe("FindUniq", func() {
	It("should work for some basic cases", func() {
		Expect(FindUniq([]float32{1.0, 1.0, 1.0, 2.0, 1.0, 1.0})).To(Equal(float32(2)))
		Expect(FindUniq([]float32{0, 0, 0.55, 0, 0})).To(Equal(float32(0.55)))
	})
})

func TestFindUniq(t *testing.T) {
	for desc, testIO := range map[string]struct {
		in  []float32
		out float32
	}{
		"2":      {in: []float32{1.0, 1.0, 1.0, 2.0, 1.0, 1.0}, out: float32(2)},
		"0.55":   {in: []float32{0, 0, 0.55, 0, 0}, out: float32(0.55)},
		"first":  {in: []float32{1, 2, 2}, out: float32(1)},
		"second": {in: []float32{2, 1, 2}, out: float32(1)},
		"third":  {in: []float32{2, 2, 1}, out: float32(1)},
	} {
		t.Run(desc, func(t *testing.T) {
			got := FindUniq(testIO.in)
			if got != testIO.out {
				t.Fatal("in", testIO.in, "expect", testIO.out, "got", got)
			}
		})
	}
}

func BenchmarkFindUniq(b *testing.B) {
	const (
		equal     float32 = 0
		different float32 = 2
	)
	in := []float32{}

	in = append(in, equal)
	in = append(in, equal)
	for n := 0; n < b.N; n++ {
		in = append(in, equal)
	}
	in = append(in, different)

	got := FindUniq(in)
	if got != different {
		b.Fatal("in", in, "expect", different, "got", got)
	}
}
