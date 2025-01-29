package board

import (
	"fmt"
	"gosweeper/logger"
	"math/rand"
	"slices"
)

var BoardRows = 8
var BoardCols = 16
var BoardMines = 20

type BoardError string

func (e BoardError) Error() string {
	return string(e)
}

const (
	ErrOutOfBounds       = BoardError("invalid address, out of bounds")
	ErrIsMine            = BoardError("cell is a mine, cannot populate adjacents")
	ErrMoreMines         = BoardError("more mines than cells")
	ErrIsPopulated       = BoardError("board is already populated")
	ErrIsAlreadyRevealed = BoardError("cell is already revealed")
	ErrIsFlagged         = BoardError("cell is flagged")
)

type GameBoard struct {
	Rows      int
	Columns   int
	Mines     int
	Cells     [][]*Cell
	Revealed  int
	Flags     int
	populated bool
}

type Cell struct {
	IsMine     bool
	IsFlag     bool
	IsRevealed bool
	Adj        int
}

func (c Cell) String() string {
	if c.IsMine {
		return "[*]"
	}
	return fmt.Sprintf("[%d]", c.Adj)
}

func (c Cell) StringVal() string {
	if c.IsFlag {
		return "[⚑]"
	}
	if !c.IsRevealed {
		return "[ ]"
	}
	return c.String()
}

func (c *Cell) Reveal() (string, error) {
	if c.IsRevealed {
		logger.DebugLogf("Cell is already revealed, with value %s", c.String())
		return "", ErrIsAlreadyRevealed
	}
	if c.IsFlag {
		return "", ErrIsFlagged
	}
	c.IsRevealed = true
	if c.IsMine {
		return "*", nil
	}
	return fmt.Sprint(c.Adj), nil
}

func (c *Cell) Flag() {
	if c.IsRevealed {
		return
	}
	c.IsFlag = !c.IsFlag
}

func (c *Cell) RevealifZero() bool {
	if c.Adj == 0 {
		c.IsRevealed = true
		return true
	}
	return false
}

// NewBoard creates a new board with the given dimensions and number of mines
func NewBoard(rows, columns, mines int) *GameBoard {
	board := &GameBoard{
		Rows:    rows,
		Columns: columns,
		Mines:   0,
	}

	board.Cells = make([][]*Cell, rows)
	for i := range board.Cells {
		board.Cells[i] = make([]*Cell, columns)
		for j := range board.Cells[i] {
			board.Cells[i][j] = &Cell{}
		}
	}
	board.populateMines(mines)
	board.PopulateAdjacents()
	logger.DebugLogf("Board populated with %d mines", board.Mines)
	logger.DebugLogf("Board: \n%v", board)
	board.populated = true

	return board
}

func (b GameBoard) String() string {
	s := ""
	for _, row := range b.Cells {
		for _, cell := range row {
			s += cell.String()
		}
		s += "\n"
	}
	return s
}

func (b GameBoard) GetCell(row, col int) (*Cell, error) {
	if col < 0 || col >= b.Columns {
		return nil, ErrOutOfBounds
	}
	if row < 0 || row >= b.Rows {
		return nil, ErrOutOfBounds
	}
	return b.Cells[row][col], nil
}

func (b GameBoard) IsMine(row, col int) (bool, error) {
	cell, err := b.GetCell(row, col)
	if err != nil {
		return false, err
	}
	return cell.IsMine, nil
}

func (b *GameBoard) AddMine(row, col int) (bool, error) {
	if b.Mines > b.Rows*b.Columns {
		return false, ErrMoreMines
	}
	cell, err := b.GetCell(row, col)
	if err != nil {
		return false, err
	}
	if cell.IsMine {
		logger.DebugLogf("Cell %d,%d is already a mine", row, col)
		return false, nil
	} else {
		logger.DebugLogf("Adding mine to cell %d,%d", row, col)
		cell.IsMine = true
		b.Mines++
		return true, nil
	}
}

func (b *GameBoard) addRandomMine() (err error) {
	success := false
	for !success {
		col := rand.Intn(b.Columns)
		row := rand.Intn(b.Rows)
		success, err = b.AddMine(row, col)
		if err != nil {
			return err
		}
	}
	return nil
}

func (b *GameBoard) populateMines(mines int) (err error) {
	logger.DebugLogf("Populating %d mines", mines)
	if b.Mines > b.Rows*b.Columns {
		return ErrMoreMines
	}
	for b.Mines < mines {
		logger.DebugLogf("Adding random mine, current count %d/%d", b.Mines, mines)
		err = b.addRandomMine()
		if err != nil {
			break
		}
	}
	return err
}

func (b GameBoard) getAdjacents(row, col int) (int, error) {
	c, err := b.GetCell(row, col)
	if err != nil {
		return 0, err
	}
	if c.IsMine {
		return 0, ErrIsMine
	}
	adjacents := 0
	for i := row - 1; i <= row+1; i++ {
		for j := col - 1; j <= col+1; j++ {
			if i == row && j == col {
				continue
			}
			cell, err := b.GetCell(i, j)
			if err != nil {
				continue
			}
			if cell.IsMine {
				adjacents++
			}
		}
	}
	return adjacents, nil
}

func (b *GameBoard) PopulateAdjacents() {
	for row := 0; row < b.Rows; row++ {
		for col := 0; col < b.Columns; col++ {
			if b.Cells[row][col].IsMine {
				continue
			}
			adjacents, _ := b.getAdjacents(row, col)
			b.Cells[row][col].Adj = adjacents
		}
	}
}

func (b *GameBoard) Flag(row, col int) string {
	cell, _ := b.GetCell(row, col)
	cell.Flag()
	logger.DebugLogf("Flagged cell %d,%d: %v", row, col, cell.IsFlag)
	if cell.IsFlag {
		b.Flags++
	} else {
		b.Flags--
	}
	return cell.StringVal()
}

func (b *GameBoard) Reveal(row, col int) (string, error) {
	cell, err := b.GetCell(row, col)
	if err != nil {
		return "", err
	}
	s, err := cell.Reveal()
	if err == nil {
		b.Revealed++
	}
	return s, err
}

func (b *GameBoard) RevealAdjacent(row, col int, visited []string, first bool) []string {
	posStr := fmt.Sprintf("%d,%d", row, col)
	if slices.Contains(visited, posStr) {
		return visited
	} else {
		visited = append(visited, posStr)
	}

	logger.DebugLogf("Revealing adjacent cells for %d,%d", row, col)
	cell, err := b.GetCell(row, col)
	if err != nil {
		return visited
	}
	if (cell.IsRevealed || cell.IsMine || cell.IsFlag) && !first {
		return visited
	}
	b.Reveal(row, col)
	if cell.Adj == 0 {
		for i := row - 1; i <= row+1; i++ {
			for j := col - 1; j <= col+1; j++ {
				if i == row && j == col {
					continue
				}
				visited = b.RevealAdjacent(i, j, visited, false)
			}
		}
	}
	return visited
}
