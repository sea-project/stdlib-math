package math

import (
	"math/big"
	"testing"
)

func Test_FirstBitSet(t *testing.T) {

	temp := new(big.Int).SetBytes([]byte("abc"))
	t.Log(temp)
}
