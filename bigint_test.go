package bigint_test

import (
	"database/sql/driver"
	"math/big"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/d-fal/bigint"
)

func TestAll(t *testing.T) {
	a := big.NewInt(100)
	b := big.NewInt(200)

	t.Run("multiply", func(t *testing.T) {
		x := bigint.NewBigint(a)
		y := bigint.NewBigint(b)
		// 100 * 200 = 20000
		assert.Equal(t, bigint.NewBigint(big.NewInt(20000)), x.Mul(y))
	})
	t.Run("add", func(t *testing.T) {
		x := bigint.NewBigint(a)
		y := bigint.NewBigint(b)
		// 100 + 200 = 300
		assert.Equal(t, bigint.NewBigint(big.NewInt(300)), x.Add(y))
	})

	t.Run("sub", func(t *testing.T) {
		x := bigint.NewBigint(a)
		y := bigint.NewBigint(b)
		// 100 -200 = -100
		assert.Equal(t, bigint.NewBigint(big.NewInt(-100)), x.Sub(y))
	})

	t.Run("div", func(t *testing.T) {
		x := bigint.NewBigint(a)
		y := bigint.NewBigint(b)
		// 200 / 100 = 2
		assert.Equal(t, bigint.NewBigint(big.NewInt(2)), y.Div(x))
	})

	t.Run("negation", func(t *testing.T) {
		x := bigint.NewBigint(a)
		assert.Equal(t, bigint.NewBigint(big.NewInt(-100)), x.Neg())
	})

	t.Run("int64", func(t *testing.T) {
		x := bigint.NewBigint(a)
		assert.Equal(t, int64(-100), x.Neg().ToInt64())
	})

	t.Run("uint64", func(t *testing.T) {
		x := bigint.NewBigint(a)
		assert.Equal(t, uint64(100), x.ToUInt64())
	})
	t.Run("toString", func(t *testing.T) {
		x := bigint.NewBigint(a)
		assert.Equal(t, "100", x.String())
	})
	t.Run("fromString", func(t *testing.T) {
		x, err := bigint.FromString("100")
		assert.Nil(t, err)
		assert.Equal(t, "100", x.String())
	})
	t.Run("fromInt64", func(t *testing.T) {
		x := bigint.FromInt64(100000000)
		assert.Equal(t, int64(100000000), x.ToInt64())
	})

	t.Run("value", func(t *testing.T) {
		x := bigint.FromInt64(10000000000)
		v, err := x.Value()
		assert.Nil(t, err)
		assert.True(t, true, driver.IsValue(v))
	})

	t.Run("from-bigint", func(t *testing.T) {
		x := bigint.NewBigint(a)
		y := bigint.NewBigint(b)
		d := x.FromBigInt(b)
		// 100 * 200 = 20000
		assert.Equal(t, y, d)
	})
	t.Run("scan", func(t *testing.T) {
		x := bigint.NewBigint(a)
		s, err := x.Value()
		assert.NoError(t, err)
		assert.NotEmpty(t, s)
	})

}
