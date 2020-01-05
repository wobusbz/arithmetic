package test

import "testing"

func TestIsJudegeSymbol(t *testing.T) {
	symbol := "asd(asdasd)asd)asasda(asdasd)(()123()asd(asd((asd(sd()("
	IsJudegeSymbol(symbol)
}
