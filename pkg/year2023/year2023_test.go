// Code generated by aocgen; DO NOT EDIT.
package year2023

import (
	"testing"

	"AdventOfCode/pkg/aoc"
)

func Benchmark2023Day01(b *testing.B) {
	Init()
	input := aoc.TestInput(2023, 1)
	p := aoc.NewPuzzle(2023, 1)
	b.Run("PartA", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			p.PartA(input)
		}
	})
	b.Run("PartB", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			p.PartB(input)
		}
	})
}
func Benchmark2023Day02(b *testing.B) {
	Init()
	input := aoc.TestInput(2023, 2)
	p := aoc.NewPuzzle(2023, 2)
	b.Run("PartA", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			p.PartA(input)
		}
	})
	b.Run("PartB", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			p.PartB(input)
		}
	})
}
func Benchmark2023Day03(b *testing.B) {
	Init()
	input := aoc.TestInput(2023, 3)
	p := aoc.NewPuzzle(2023, 3)
	b.Run("PartA", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			p.PartA(input)
		}
	})
	b.Run("PartB", func(b *testing.B) {
		b.ResetTimer()
		for i := 0; i < b.N; i++ {
			p.PartB(input)
		}
	})
}
