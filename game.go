package main

import "math/rand"

const (
	alivePercent = 50
)

// State 生命的状态
type State struct {
	state     bool
	nextState bool
}

// Game 结构体
type Game struct {
	row  int
	col  int
	grid [][]State
}

// NewGame 新游戏
func NewGame(row, col int) *Game {
	game := Game{
		row: row,
		col: col,
	}
	game.grid = grid(row, col)
	game.initLife(alivePercent)
	return &game
}

//随机初始化生命，按百分比随机生成生命数量
func (game *Game) initLife(alivePercent int) {
	for i := 0; i < game.row; i++ {
		for j := 0; j < game.col; j++ {
			if rand.Intn(100) < alivePercent {
				game.grid[i][j] = State{state: true, nextState: true}
			} else {
				game.grid[i][j] = State{state: false, nextState: false}
			}
		}
	}
}

// 生成grid
func grid(row, col int) [][]State {
	g := make([][]State, row)
	for i := 0; i < row; i++ {
		g[i] = make([]State, col)
	}

	return g
}

//左右边界的映射，超出左边界则认为是右边界关联，如-1会映射为是最右侧,这样会让游戏的宽度是无限延展的
func (game *Game) mapX(x int) int {
	if x >= game.row || x < 0 {
		return (x%game.row + game.row) % game.row
	}
	return x
}

//上下边界的映射，参见mapX
func (game *Game) mapY(y int) int {
	if y >= game.col || y < 0 {
		return (y%game.col + game.col) % game.col
	}
	return y
}

//计算某个生命的邻居生存个数
func (game *Game) aliveCountAround(x, y int) int {
	around := [8][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	count := 0
	for i := 0; i < 8; i++ {
		if game.grid[game.mapX(x+around[i][0])][game.mapY(y+around[i][1])].state {
			count++
		}
	}
	return count
}

//计算某个生命的下一回合的生存状态
func (game *Game) nextState(x, y int) bool {
	aliveCountAround := game.aliveCountAround(x, y)
	if aliveCountAround >= 4 {
		return false
	} else if aliveCountAround == 3 {
		return true
	} else if aliveCountAround >= 2 {
		return game.grid[x][y].nextState
	} else {
		return false
	}
}

//计算所有生命的下一回合的生存状态
func (game *Game) calcNextState() {
	for i := 0; i < game.row; i++ {
		for j := 0; j < game.col; j++ {
			game.grid[i][j].nextState = game.nextState(i, j)
		}
	}
}

//转换到下一回合的生存状态
func (game *Game) switchNextState() {
	for i := 0; i < game.row; i++ {
		for j := 0; j < game.col; j++ {
			game.grid[i][j].state = game.grid[i][j].nextState
		}
	}
}

// IsAlive 是否处于生存状态
func (game *Game) IsAlive(row, col int) bool {
	return game.grid[row][col].state
}

//NextAround 下一回合
func (game *Game) NextAround() {
	game.calcNextState()
	game.switchNextState()
}
