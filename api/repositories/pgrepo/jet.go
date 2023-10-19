package pgrepo

import (
	"fmt"

	"github.com/go-jet/jet/v2/postgres"
)

func ILike(column postgres.ColumnString, value string) postgres.BoolExpression {
	return postgres.BoolExp(
		postgres.Raw(
			fmt.Sprintf("%s.%s ILIKE :value", column.TableName(), column.Name()),
			postgres.RawArgs{":value": fmt.Sprintf("%%%s%%", value)},
		),
	)
}
