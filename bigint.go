package bigint

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"math/big"
)

type Bigint big.Int

func NewBigint(x *big.Int) *Bigint {
	return (*Bigint)(x)
}

func FromInt64(x int64) *Bigint {
	return new(Bigint).FromBigInt(big.NewInt(x))
}

func FromString(x string) (*Bigint, error) {
	if x == "" {
		x = "0"
	}
	a := big.NewInt(0)
	b, ok := a.SetString(x, 10)

	if !ok {
		return nil, fmt.Errorf("cannot create Bigint from string")
	}

	return NewBigint(b), nil
}

func (b *Bigint) Value() (driver.Value, error) {
	return (*big.Int)(b).String(), nil
}

func (b *Bigint) Scan(value interface{}) error {

	var i sql.NullString

	if err := i.Scan(value); err != nil {
		return err
	}

	if _, ok := (*big.Int)(b).SetString(i.String, 10); ok {
		return nil
	}

	return fmt.Errorf("Error converting type %T into Bigint", value)
}

func (b *Bigint) toBigInt() *big.Int {
	return (*big.Int)(b)
}

func (b *Bigint) Sub(x *Bigint) *Bigint {
	return (*Bigint)(big.NewInt(0).Sub(b.toBigInt(), x.toBigInt()))
}

func (b *Bigint) Add(x *Bigint) *Bigint {
	return (*Bigint)(big.NewInt(0).Add(b.toBigInt(), x.toBigInt()))
}

func (b *Bigint) Mul(x *Bigint) *Bigint {
	return (*Bigint)(big.NewInt(0).Mul(b.toBigInt(), x.toBigInt()))
}

func (b *Bigint) Div(x *Bigint) *Bigint {
	return (*Bigint)(big.NewInt(0).Div(b.toBigInt(), x.toBigInt()))
}

func (b *Bigint) Neg() *Bigint {
	return (*Bigint)(big.NewInt(0).Neg(b.toBigInt()))
}

func (b *Bigint) ToUInt64() uint64 {
	return b.toBigInt().Uint64()
}

func (b *Bigint) ToInt64() int64 {
	return b.toBigInt().Int64()
}

// same as NewBigint()
func (b *Bigint) FromBigInt(x *big.Int) *Bigint {
	return (*Bigint)(x)
}

func (b *Bigint) String() string {
	return b.toBigInt().String()
}

func (b *Bigint) Cmp(target *Bigint) Cmp {
	return &cmp{r: b.toBigInt().Cmp(target.toBigInt())}
}

func (b *Bigint) Abs() *Bigint {
	return (*Bigint)(new(big.Int).Abs(b.toBigInt()))
}
