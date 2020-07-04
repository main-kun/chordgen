package main

import "github.com/main-kun/chordgen/pkg/picgen"

func main() {
	arr := []int{9, 13, 16, 20}
	picgen.DrawChord(arr, "test.png")
}
