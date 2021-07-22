package test

import (
	"github.com/mattak/stonk/pkg/util"
	"testing"
)

func TestNormalizeName(t *testing.T) {
	realNames := []string{
		"ＮＥＸＴ　ＦＵＮＤＳ　ＴＯＰＩＸ連動型上場投信",
		"ｉシェアーズ　ＪＰＸ日経４００",
		"Ｓ＆Ｐ　ＧＳＣＩ商品指数",
		"Ｓ＆Ｐ　ＧＳＣＩ商品指数　エネルギー＆メタル・キャップド・コンポーネント３５／２０・ＴＨＥＡＭ・イージーＵＣＩＴＳ・ＥＴＦクラスＡ米ドル建受益証券",
		"市場第一部（内国株）",
		"JASDAQ(グロース・内国株）",
	}
	expectNames := []string{
		"NEXT FUNDS TOPIX連動型上場投信",
		"iシェアーズ JPX日経400",
		"S&P GSCI商品指数",
		"S&P GSCI商品指数 エネルギー&メタル・キャップド・コンポーネント35/20・THEAM・イージーUCITS・ETFクラスA米ドル建受益証券",
		"市場第一部(内国株)",
		"JASDAQ(グロース・内国株)",
	}
	for i := 0; i < len(realNames); i++ {
		normalizedName := util.NormalizeName(realNames[i])
		if normalizedName != expectNames[i] {
			t.Fatalf("normalize name failed: %s <=> %s\n", normalizedName, expectNames[i])
		}
	}
}
