package counter

import (
	"testing"
)

func testIsSocialDistancing(expectation bool, source dataSource, minDistance int, t *testing.T) {
	if expectation != IsSocialDistancing(source, minDistance) {
		t.Error("wrong expection", expectation)
		return
	}
}

func TestIsSocialDistancing1(t *testing.T) {
	expectation := true
	source := dataSource{{1, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 1}}
	testIsSocialDistancing(expectation, source, 6, t)
}
func TestIsSocialDistancing2(t *testing.T) {
	expectation := true
	source := dataSource{{1, 0, 0, 0, 0, 0, 0, 1}}
	testIsSocialDistancing(expectation, source, 6, t)
}

func TestIsSocialDistancing3(t *testing.T) {
	expectation := false
	source := dataSource{{1, 0, 0, 0, 1}}
	testIsSocialDistancing(expectation, source, 6, t)
}
func TestIsSocialDistancing4(t *testing.T) {
	expectation := false
	source := dataSource{{1, 0, 0, 1}}
	testIsSocialDistancing(expectation, source, 6, t)
}

func TestIsSocialDistancing5(t *testing.T) {
	expectation := false
	source := dataSource{
		{1},
		{0},
		{0},
		{0},
		{0},
		{1}}
	testIsSocialDistancing(expectation, source, 6, t)
}
func TestIsSocialDistancing6(t *testing.T) {
	expectation := true
	source := dataSource{
		{1},
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
		{1}}
	testIsSocialDistancing(expectation, source, 6, t)
}
func TestIsSocialDistancing7(t *testing.T) {
	expectation := false
	source := dataSource{
		{1, 0, 0, 0, 0, 0, 1, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0, 1}}
	testIsSocialDistancing(expectation, source, 6, t)
}
func TestIsSocialDistancing8(t *testing.T) {
	expectation := true
	source := dataSource{
		{1, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0}}
	testIsSocialDistancing(expectation, source, 6, t)
}
func TestIsSocialDistancing9(t *testing.T) {
	expectation := true
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
	testIsSocialDistancing(expectation, source, 6, t)
}
func TestIsSocialDistancing10(t *testing.T) {
	expectation := false
	source := dataSource{
		{1, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 1, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 1, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0},
		{1, 0, 0, 0, 0, 0, 0, 0, 1}}
	testIsSocialDistancing(expectation, source, 6, t)
}

func TestIsSocialDistancing11(t *testing.T) {
	expectation := true
	source := dataSource{}
	testIsSocialDistancing(expectation, source, 6, t)
}

func TestIsSocialDistancing12(t *testing.T) {
	expectation := true
	source := dataSource{{}, {}, {}}
	testIsSocialDistancing(expectation, source, 6, t)
}

func TestIsSocialDistancing13(t *testing.T) {
	expectation := true
	source := dataSource{
		{1, 0, 0, 0, 0, 0, 0, 1},
		{0},
		{0},
		{0},
		{0},
		{0},
		{0},
		{1},
	}
	testIsSocialDistancing(expectation, source, 6, t)
}

func TestIsSocialDistancing14(t *testing.T) {
	expectation := false
	source := dataSource{
		{1, 0, 0, 0, 0, 0, 0, 1},
		{0},
		{0},
		{0},
		{0},
		{0},
		{1},
		{0},
	}
	testIsSocialDistancing(expectation, source, 6, t)
}

func TestIsSocialDistancing15(t *testing.T) {
	expectation := true
	source := dataSource{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 1, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	testIsSocialDistancing(expectation, source, 3, t)
}

func TestIsSocialDistancing16(t *testing.T) {
	expectation := false
	source := dataSource{
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0},
		{0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0},
		{0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		{0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
	}
	testIsSocialDistancing(expectation, source, 6, t)
}
func TestIsSocialDistancing17(t *testing.T) {
	expectation := true
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
	testIsSocialDistancing(expectation, source, distance, t)
}
