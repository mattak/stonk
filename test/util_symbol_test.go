package test

import (
	"github.com/mattak/stonk/pkg/util"
	"testing"
)

func TestIsTokyoNoiseSymbol(t *testing.T) {
	if util.IsTokyoNoiseSymbol("7731.T") {
		t.Fatal("expect to be false")
	}
	if !util.IsTokyoNoiseSymbol("1649910D.T") {
		t.Fatal("expect to be true")
	}
}
