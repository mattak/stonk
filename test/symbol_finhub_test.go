package test

import (
	"github.com/mattak/stonk/cmd"
	"testing"
)

func TestFilterFinhubSymbols(t *testing.T) {
	dic := map[string]cmd.SymbolInfo{
		"12345678.T": {"12345678.T", "Name1"},
		"1234.T":     {"1234.T", "Name2"},
	}
	cmd.FilterFinhubSymbols("T", &dic)
	_, containsKey := dic["1234.T"]
	if !containsKey {
		t.Fatal("should be contains key: 1234.T")
	}
	_, containsKey = dic["12345678.T"]
	if containsKey {
		t.Fatal("should be contains key: 12345678.T")
	}
}
