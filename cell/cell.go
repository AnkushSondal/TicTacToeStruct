package cell

type Cell struct {
	mark string
}

func NewCell() *Cell {
	return &Cell{
		mark: "-",
	}
}

func (c *Cell) IsCellMarked() bool {
	return c.mark != "-"
}

func (c *Cell) MarkCell(mark string) {
	c.mark = mark
}
func (c *Cell) ReturnMark() string {
	return c.mark
}
