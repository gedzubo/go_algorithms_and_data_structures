package grid

type Grid struct {
	rows    int
	columns int
	grid    [][]string
}

func NewGrid(rows, columns int) *Grid {
	grid := make([][]string, rows)
	for i := range grid {
		grid[i] = make([]string, columns)
	}
	return &Grid{
		rows:    rows,
		columns: columns,
		grid:    grid,
	}
}

func (g *Grid) Set(row, column int, value string) {
	if row >= 0 && row < g.rows && column >= 0 && column < g.columns {
		g.grid[row][column] = value
	}
}

func (g *Grid) Get(row, column int) string {
	if row >= 0 && row < g.rows && column >= 0 && column < g.columns {
		return g.grid[row][column]
	} else {
		return ""
	}
}

func (g *Grid) Print() string {
	var result string
	for _, row := range g.grid {
		for _, cell := range row {
			result += cell + " "
		}
		result += "\n"
	}
	return result
}

func (g *Grid) FillRow(row int, value string) {
	for i := 0; i < g.columns; i++ {
		g.Set(row, i, value)
	}
}

func (g *Grid) FillColumn(column int, value string) {
	for i := 0; i < g.rows; i++ {
		g.Set(i, column, value)
	}
}
