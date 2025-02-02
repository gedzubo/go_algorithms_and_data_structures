package grid

type Grid struct {
	rows    int
	columns int
	grid    [][]string
	values  []string
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
		g.values = append(g.values, value)
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

func (g *Grid) FillAll(value string) {
	for i := 0; i < g.rows; i++ {
		g.FillRow(i, value)
	}
}

func (g *Grid) FillDiagonal(value string) {
	for i := 0; i < g.rows; i++ {
		g.Set(i, i, value)
	}
}

func (g *Grid) FillReverseDiagonal(value string) {
	for i := 0; i < g.rows; i++ {
		g.Set(i, g.columns-i-1, value)
	}
}

func (g *Grid) FillBorder(value string) {
	g.FillRow(0, value)
	g.FillRow(g.rows-1, value)
	g.FillColumn(0, value)
	g.FillColumn(g.columns-1, value)
}

func (g *Grid) CheckIfValueIsPresent(value string) bool {
	for _, v := range g.values {
		if v == value {
			return true
		}
	}
	return false
}

func (g *Grid) CheckIfValueIsPresentInRow(row int, value string) bool {
	for i := 0; i < g.columns; i++ {
		if g.Get(row, i) == value {
			return true
		}
	}
	return false
}

func (g *Grid) CheckIfValueIsPresentInColumn(column int, value string) bool {
	for i := 0; i < g.rows; i++ {
		if g.Get(i, column) == value {
			return true
		}
	}
	return false
}
