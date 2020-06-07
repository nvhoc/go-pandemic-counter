package counter

type rowSource []uint8

type dataSource []rowSource

func abs(num int) int {
	if num < 0 {
		return num * -1
	}
	return num
}

func (d dataSource) check(minDistance, px, py int) bool {
	maxY := len(d)
	for increaseY := -minDistance; increaseY <= minDistance; increaseY++ {
		y := py + increaseY
		if y < 0 || y >= maxY {
			continue
		}
		deltaX := minDistance - abs(increaseY)
		maxX := len(d[y])
		for increaseX := -deltaX; increaseX <= deltaX; increaseX++ {
			x := px + increaseX
			if x < 0 || x >= maxX {
				continue
			}

			if d[y][x] == 0 {
				continue
			}
			if x == px && y == py {
				continue
			}

			return false
		}
	}
	return true
}

// IsSocialDistancing main function of the solution
func IsSocialDistancing(data dataSource, minDistance int) bool {
	for i, vi := range data {
		for j, vij := range vi {
			if vij == 0 {
				continue
			}
			isNoHumanInMalArea := data.check(minDistance, j, i)
			if !isNoHumanInMalArea {
				return false
			}
		}
	}
	return true
}
