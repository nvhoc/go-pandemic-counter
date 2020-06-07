package counter

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
	"testing"
)

func bmIsSocialDistancing(source dataSource, minDistance int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		IsSocialDistancing(source, minDistance)
	}
}

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

func BenchmarkIsSocialDistancing17(t *testing.B) {
	M := 10000
	source := make(dataSource, M)
	distance := 6
	for i := 0; i < M; i++ {
		source[i] = make(rowSource, M)
		if i%(distance+1) != 0 {
			continue
		}
		for j := 0; j < M; j++ {
			if j%(distance+1) == 0 {
				source[i][j] = 1
			}
		}
	}
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}

	bmIsSocialDistancing(source, distance, t)
}

func BenchmarkIsSocialDistancing(b *testing.B) {
	source := dataSource{
		{1, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 1}}
	bmIsSocialDistancing(source, 6, b)
}
