package main

import (
	"github.com/Hilst/tuirest/screen"
)

func main() {
	s := screen.MakeScreen()
	if err := s.Run(); err != nil {
		println(err.Error())
		panic(err)
	}
}
