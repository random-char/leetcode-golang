package solutions

func islandPerimeter(grid [][]int) int {
	perimeter := 0

	for i := range grid {
		for j := range grid[i] {
			if grid[i][j] == 0 {
				continue
			}

			if i-1 < 0 || grid[i-1][j] == 0 {
				perimeter += 1
			}

			if i+1 >= len(grid) || grid[i+1][j] == 0 {
				perimeter += 1
			}

			if j-1 < 0 || grid[i][j-1] == 0 {
				perimeter += 1
			}

			if j+1 >= len(grid[i]) || grid[i][j+1] == 0 {
				perimeter += 1
			}
		}
	}

	return perimeter
}
