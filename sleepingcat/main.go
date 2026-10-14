package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/signal"
	"strings"
	"syscall"
	"time"
)

type Zzz struct {
	x    int
	y    int
	life int
	size int
}

const (
	width  = 60
	height = 18
)

func main() {
	fmt.Print("\033[?25l")
	defer fmt.Print("\033[?25h")

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sig
		fmt.Print("\033[?25h")
		fmt.Print("\033[H\033[2J")
		os.Exit(0)
	}()

	cat := []string{
		`⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⣠⡤⠶⠒⠛⠉⠙⠛⠒⠶⢤⣀⠀⠀⠀⠀⠀⠀`,
		`⠀⠀⠀⠀⠀⠀⠀⠀⠀⢀⡴⠋⠁⠀⠀⠀⠀⠀⠀⠁⠀⠀⠀⠈⠙⢦⡀⠀⠀⠀`,
		`⠀⠀⠀⠀⣀⣠⠤⠤⡴⠻⢓⣶⠦⠤⣄⡀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠙⢦⡀⠀`,
		`⣀⡠⠴⠊⠁⠀⠀⠀⠀⠀⠒⠽⠀⠀⠀⠉⢙⠒⢢⡄⠀⠀⠀⠀⠀⠀⠀⠈⢷⡀`,
		`⠘⢆⠉⠑⡆⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢆⠁⣠⠎⠀⠀⠀⠀⠀⠀⠀⠀⠀⠈⣇`,
		`⠀⠈⠓⠦⠁⢀⣀⠀⠀⠀⠀⠀⠀⣀⣀⢸⡊⠁⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⠀⢻`,
		`⠀⠀⠀⡸⠀⠉⠀⠙⠆⠀⠀⠀⠏⠀⠈⠀⢇⠀⠀⠀⠀⠀⠀⠀⢀⠔⠀⠀⠀⣼`,
		`⠀⠀⠀⠹⣄⠀⠠⣤⡡⠪⡭⠃⡤⢤⡄⡰⠋⠀⠀⠀⣀⡠⠴⠊⠁⠀⠀⠀⢠⠇`,
		`⠀⠀⢠⢿⣠⠟⠓⠛⠉⠛⡟⠛⠛⠛⠛⠒⠒⠚⠉⠉⠁⠀⠀⠀⠀⠀⢀⡴⠋⠀`,
		`⠀⠀⠈⠛⢯⣀⣀⣀⡤⠤⠤⠤⢤⣤⣀⣀⣀⣀⣀⣀⣀⣤⠤⠴⠒⠋⠁⠀⠀⠀`,
	}

	catW := 0
	for _, row := range cat {
		if r := len([]rune(row)); r > catW {
			catW = r
		}
	}

	catX := (width - catW) / 2
	catY := height - len(cat)

	headX := catX + catW*40/100
	headY := catY

	rng := rand.New(rand.NewSource(time.Now().UnixNano()))

	zs := []Zzz{}

	for {

		for i := range zs {
			zs[i].y--
			zs[i].x++
			zs[i].life--
			if zs[i].life < 15 && zs[i].size < 3 {
				zs[i].size++
			}
		}

		alive := zs[:0]
		for _, z := range zs {
			if z.y+1 >= 0 && z.life > 0 && z.x < width {
				alive = append(alive, z)
			}
		}
		zs = alive

		if rng.Intn(100) < 12 {
			zs = append(zs, Zzz{
				x:    headX,
				y:    headY - 1,
				life: 25 + rng.Intn(10),
				size: 1,
			})
		}

		fmt.Print("\033[H\033[2J")

		canvas := make([][]rune, height)
		for i := range canvas {
			canvas[i] = make([]rune, width)
			for j := range canvas[i] {
				canvas[i][j] = ' '
			}
		}

		for i, row := range cat {
			cy := catY + i
			if cy < 0 || cy >= height {
				continue
			}
			for j, ch := range []rune(row) {
				cx := catX + j
				if cx >= 0 && cx < width {
					canvas[cy][cx] = ch
				}
			}
		}

		for _, z := range zs {
			for k := 0; k < z.size; k++ {
				cx := z.x + k
				if z.y >= 0 && z.y < height && cx >= 0 && cx < width {
					canvas[z.y][cx] = 'Z'
				}
			}
		}

		fmt.Println("  ┌" + strings.Repeat("─", width) + "┐")
		for _, row := range canvas {
			fmt.Printf("  │%s│\n", string(row))
		}
		fmt.Println("  └" + strings.Repeat("─", width) + "┘")

		time.Sleep(90 * time.Millisecond)
	}
}
