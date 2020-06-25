package main

import (
	"github.com/nsf/termbox-go"
	"github.com/viile/games/common"
	"github.com/viile/games/snake"
	//"github.com/viile/games/tetris"
	"math/rand"
	"time"
)

func inputFromTermbox(g common.Game) (err error) {
	err = termbox.Init()
	if err != nil {
		return
	}
	defer termbox.Close()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			switch ev.Key {
			case 0:
				switch ev.Ch {
				case 'w','W':
					g.Input(65517)
				case 's','S':
					g.Input(65516)
				case 'a','A':
					g.Input(65515)
				case 'd','D':
					g.Input(65514)
				}
			case termbox.KeyEsc:
				g.Stop()
				return
			case termbox.KeyArrowUp,termbox.KeyCtrlW:
				g.Input(65517)
			case termbox.KeyArrowDown,termbox.KeyCtrlS:
				g.Input(65516)
			case termbox.KeyArrowLeft,termbox.KeyCtrlA:
				g.Input(65515)
			case termbox.KeyArrowRight,termbox.KeyCtrlD:
				g.Input(65514)
			default:
				continue
			}
		}
	}
}


func main() {
	rand.Seed(time.Now().UnixNano())
	defer func() {
		if r := recover(); r != nil {
			println(r)
		}
	}()

	g := snake.NewGame()
	go g.Run()

	println(inputFromTermbox(g))

}
