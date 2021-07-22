package util

import "github.com/ktnyt/go-moji"

func NormalizeName(name string) string {
	// zenkaku english -> hankaku english
	name = moji.Convert(name, moji.ZE, moji.HE)
	name = moji.Convert(name, moji.ZS, moji.HS)
	return name
}
