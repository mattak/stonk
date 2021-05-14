package test

import (
	"github.com/mattak/stonk/cmd"
	"testing"
)

func TestIsTokyoNoiseSymbol(t *testing.T) {
	if cmd.IsTokyoNoiseSymbol("7731.T") {
		t.Fatal("expect to be false")
	}
	if !cmd.IsTokyoNoiseSymbol("1649910D.T") {
		t.Fatal("expect to be true")
	}
}
