package main

import (
	"github.com/gdamore/tcell"
)

// Render 渲染
type Render struct {
	game   *Game
	screen tcell.Screen
}

// NewRender 生成新对象
func NewRender(screen tcell.Screen) *Render {
	w, h := screen.Size()
	game := NewGame(h, w/2)
	return &Render{game: game, screen: screen}
}

func (render *Render) rend() {
	st := tcell.StyleDefault.Background(tcell.ColorWhite)
	bst := tcell.StyleDefault.Background(tcell.ColorBlack)
	for i := 0; i < render.game.row; i++ {
		for j := 0; j < render.game.col; j++ {
			if render.game.IsAlive(i, j) {
				render.screen.SetCell(j*2, i, st, ' ')
				render.screen.SetCell(j*2+1, i, st, ' ')
			} else {
				render.screen.SetCell(j*2, i, bst, ' ')
				render.screen.SetCell(j*2+1, i, bst, ' ')
			}
		}
	}
	render.game.NextAround()
	render.screen.Show()
}
