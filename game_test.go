package game_test

import (
	"testing"

	"github.com/kukichalang/game"
)

func TestOverlaps(t *testing.T) {
	cases := []struct {
		name string
		a, b game.Rect
		want bool
	}{
		{"overlapping", game.Rect{X: 0, Y: 0, Width: 10, Height: 10}, game.Rect{X: 5, Y: 5, Width: 10, Height: 10}, true},
		{"not overlapping", game.Rect{X: 0, Y: 0, Width: 10, Height: 10}, game.Rect{X: 20, Y: 20, Width: 10, Height: 10}, false},
		{"touching edges", game.Rect{X: 0, Y: 0, Width: 10, Height: 10}, game.Rect{X: 10, Y: 0, Width: 10, Height: 10}, false},
		{"contained", game.Rect{X: 0, Y: 0, Width: 20, Height: 20}, game.Rect{X: 5, Y: 5, Width: 5, Height: 5}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := game.Overlaps(tc.a, tc.b); got != tc.want {
				t.Errorf("Overlaps() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestOverlapsCircle(t *testing.T) {
	cases := []struct {
		name string
		a, b game.Circle
		want bool
	}{
		{"overlapping", game.Circle{X: 0, Y: 0, Radius: 10}, game.Circle{X: 15, Y: 0, Radius: 10}, true},
		{"not overlapping", game.Circle{X: 0, Y: 0, Radius: 5}, game.Circle{X: 20, Y: 0, Radius: 5}, false},
		{"concentric", game.Circle{X: 5, Y: 5, Radius: 10}, game.Circle{X: 5, Y: 5, Radius: 3}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := game.OverlapsCircle(tc.a, tc.b); got != tc.want {
				t.Errorf("OverlapsCircle() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestCircleOverlapsRect(t *testing.T) {
	cases := []struct {
		name string
		c    game.Circle
		r    game.Rect
		want bool
	}{
		{"overlapping", game.Circle{X: 5, Y: 5, Radius: 5}, game.Rect{X: 0, Y: 0, Width: 10, Height: 10}, true},
		{"not overlapping", game.Circle{X: 50, Y: 50, Radius: 5}, game.Rect{X: 0, Y: 0, Width: 10, Height: 10}, false},
		{"circle inside rect", game.Circle{X: 5, Y: 5, Radius: 2}, game.Rect{X: 0, Y: 0, Width: 20, Height: 20}, true},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			if got := game.CircleOverlapsRect(tc.c, tc.r); got != tc.want {
				t.Errorf("CircleOverlapsRect() = %v, want %v", got, tc.want)
			}
		})
	}
}

func TestMakeColor(t *testing.T) {
	c := game.MakeColor(255, 0, 0, 255)
	if c.R != 255 || c.G != 0 || c.B != 0 || c.A != 255 {
		t.Errorf("MakeColor(255,0,0,255) = %+v", c)
	}

	c = game.MakeColor(0, 0, 0, 0)
	if c.R != 0 || c.G != 0 || c.B != 0 || c.A != 0 {
		t.Errorf("MakeColor(0,0,0,0) = %+v", c)
	}
}

func TestRandom(t *testing.T) {
	r := game.Random(5, 10)
	if r < 5 || r >= 10 {
		t.Errorf("Random(5, 10) = %d, want [5, 10)", r)
	}

	same := game.Random(7, 7)
	if same != 7 {
		t.Errorf("Random(7, 7) = %d, want 7", same)
	}
}
