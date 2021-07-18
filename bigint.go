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
	a := big.NewInt(0)
	b, ok := a.SetString(x, 10)

	if !ok {
		return nil, fmt.Errorf("cannot create Bigint from string")
	}

	return NewBigint(b), nil
}

func (i *Bigint) Value() (driver.Value, error) {
	return (*big.Int)(i).String(), nil
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

func (i *Bigint) toBigInt() *big.Int {
	return (*big.Int)(i)
}

func (i *Bigint) Sub(x *Bigint) *Bigint {
	return (*Bigint)(big.NewInt(0).Sub(i.toBigInt(), x.toBigInt()))
}

func (i *Bigint) Add(x *Bigint) *Bigint {
	return (*Bigint)(big.NewInt(0).Add(i.toBigInt(), x.toBigInt()))
}

func (i *Bigint) Mul(x *Bigint) *Bigint {
	return (*Bigint)(big.NewInt(0).Mul(i.toBigInt(), x.toBigInt()))
}

func (i *Bigint) Div(x *Bigint) *Bigint {
	return (*Bigint)(big.NewInt(0).Div(i.toBigInt(), x.toBigInt()))
}

func (i *Bigint) Neg() *Bigint {
	return (*Bigint)(big.NewInt(0).Neg(i.toBigInt()))
}

func (i *Bigint) ToUInt64() uint64 {
	return i.toBigInt().Uint64()
}

func (i *Bigint) ToInt64() int64 {
	return i.toBigInt().Int64()
}

func (i *Bigint) FromBigInt(x *big.Int) *Bigint {
	return (*Bigint)(x)
}

func (i *Bigint) String() string {
	return i.toBigInt().String()
}
