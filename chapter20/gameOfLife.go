package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const (
	width  = 25
	height = 50
) //设置长宽

type Universe [][]bool

var clean = func() {
	cmd := exec.Command("cmd", "/c", "cls") // Windows系统下清屏命令
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func NewUniverse() Universe { //新建世界
	universe := make(Universe, width)

	for i := range universe {
		universe[i] = make([]bool, height)
	}

	return universe
}

func (u Universe) Show() { //展示世界
	for _, u1 := range u {
		for _, i := range u1 {
			if i {
				fmt.Print("*")
			} else {
				fmt.Printf(".")
			}
		}
		fmt.Printf("\n")
	}
}

func (u Universe) Seed() { //激活世界
	count := 0
	for {
		i := rand.Intn(width)
		j := rand.Intn(height)
		if u[i][j] {
			continue
		} else {
			u[i][j] = true
			count++
		}
		if count >= height*width/4 {
			break
		}
	}
}

func (u Universe) Alive(x, y int) bool { //判断细胞是否存活
	x = (x + width) % width
	y = (y + height) % height
	if u[x][y] {
		return true
	} else {
		return false
	}
}

func (u Universe) Neighbors(x, y int) int { //统计临近的细胞
	count := 0
	for i := 0; i < 8; i++ {
		switch i {
		case 0:
			if u.Alive(x+1, y+1) {
				count++
			}
		case 1:
			if u.Alive(x+1, y-1) {
				count++
			}
		case 2:
			if u.Alive(x-1, y+1) {
				count++
			}
		case 3:
			if u.Alive(x-1, y-1) {
				count++
			}
		case 4:
			if u.Alive(x, y+1) {
				count++
			}
		case 5:
			if u.Alive(x, y-1) {
				count++
			}
		case 6:
			if u.Alive(x+1, y) {
				count++
			}
		case 7:
			if u.Alive(x-1, y) {
				count++
			}
		}
	}
	return count
}

func (u Universe) Next(x, y int) bool { //下一个世代
	count := u.Neighbors(x, y)
	if count == 2 || count == 3 {
		return true
	} else {
		return false
	}
}

func Step(a, b Universe) { //世代更新
	for i := range a {
		for j := range a[i] {
			b[i][j] = a.Next(i, j)
		}
	}
	for i := range a {
		for j := range a[i] {
			a[i][j], b[i][j] = b[i][j], a[i][j]
		}
	}
}

func main() {
	universe_1 := NewUniverse()
	universe_2 := NewUniverse()
	universe_1.Seed()

	for {
		universe_1.Show()
		Step(universe_1, universe_2)
		time.Sleep(1 * time.Second)
		clean()
	}
}
