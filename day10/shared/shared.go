package shared

func FindStartingPoint(rs [][]rune) (x, y int) {
	for y := range rs {
		for x := range rs[y] {
			if rs[y][x] == 'S' {
				return x, y
			}
		}
	}
	panic("S not found")
}

func CanGoNorth(x, y int, rs [][]rune) bool {
	if y == 0 {
		return false
	}
	north := rs[y-1][x]
	current := rs[y][x]
	return (current == '|' || current == 'L' || current == 'J' || current == 'S') &&
		(north == '|' || north == 'F' || north == '7' || north == 'S')
}

func CanGoSouth(x, y int, rs [][]rune) bool {
	if y >= len(rs)-1 {
		return false
	}
	south := rs[y+1][x]
	current := rs[y][x]
	return (current == '|' || current == 'F' || current == '7' || current == 'S') &&
		(south == '|' || south == 'L' || south == 'J' || south == 'S')
}

func CanGoWest(x, y int, rs [][]rune) bool {
	if x == 0 {
		return false
	}
	west := rs[y][x-1]
	current := rs[y][x]
	return (current == '-' || current == '7' || current == 'J' || current == 'S') &&
		(west == '-' || west == 'F' || west == 'L' || west == 'S')
}

func CanGoEast(x, y int, rs [][]rune) bool {
	if x >= len(rs[0])-1 {
		return false
	}
	east := rs[y][x+1]
	current := rs[y][x]
	return (current == '-' || current == 'F' || current == 'L' || current == 'S') &&
		(east == '-' || east == '7' || east == 'J' || east == 'S')
}
