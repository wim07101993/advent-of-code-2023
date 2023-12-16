package shared

type Coordinate struct {
	X int
	Y int
}

func ListGalaxies(rs [][]rune) []Coordinate {
	var gs []Coordinate
	for y := range rs {
		for x := range rs[y] {
			if rs[y][x] == '#' {
				gs = append(gs, Coordinate{x, y})
			}
		}
	}
	return gs
}

func IsEmptyRow(rs [][]rune, y int) bool {
	for x := range rs[y] {
		if rs[y][x] != rs[y][0] {
			return false
		}
	}
	return true
}

func IsEmptyColumn(rs [][]rune, x int) bool {
	for y := range rs {
		if rs[y][x] != rs[0][x] {
			return false
		}
	}
	return true
}

func AbsDiff(i1, i2 int) int {
	if i1 > i2 {
		return i1 - i2
	}
	return i2 - i1
}
