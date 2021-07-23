# bigint

bigint is a wrapper around math/big package to let us use big.int type in postgresql

In order to use big.Int structure into postgres database, one may thinks this is not straightforward and even painful.
In this package, we are going to solve it.

## math/big

This package uses math/big in its heart and extends its usefulnell even into postgres.

## Example use with go-pg

**go-pg** is an amazing orm for gophers to utilize postgres. This package is used to help **go-pg** users implement **math/big** functionalities.

```
import (
	"net"

	"github.com/d-fal/bigint"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
)

func main() {

	type TableWithBigint struct {
		Id        uint64
		tableName struct{} `pg:"table_with_bigint"`
		Name      string
		Deposit   *bigint.Bigint
	}

	db := pg.Connect(&pg.Options{
		User:     "user",
		Password: "password",
		Database: "testdb",
		Addr: net.JoinHostPort(
			"postgres-url",
			"postgres-port",
		),
	})
	err := db.Model((*TableWithBigint)(nil)).CreateTable(&orm.CreateTableOptions{
		Temp:          true,
		FKConstraints: true,
		IfNotExists:   true,
	})

	if err != nil {
		println(err)
	}

}


```

# Support for [bun](https://github.com/uptrace/bun)

**bun** will be the successor of **go-pg** and this package is under development to cover it as well.