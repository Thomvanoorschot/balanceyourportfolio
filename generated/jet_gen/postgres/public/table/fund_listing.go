//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

import (
	"github.com/go-jet/jet/v2/postgres"
)

var FundListing = newFundListingTable("public", "fund_listing", "")

type fundListingTable struct {
	postgres.Table

	// Columns
	FundID postgres.ColumnString
	Ticker postgres.ColumnString

	AllColumns     postgres.ColumnList
	MutableColumns postgres.ColumnList
}

type FundListingTable struct {
	fundListingTable

	EXCLUDED fundListingTable
}

// AS creates new FundListingTable with assigned alias
func (a FundListingTable) AS(alias string) *FundListingTable {
	return newFundListingTable(a.SchemaName(), a.TableName(), alias)
}

// Schema creates new FundListingTable with assigned schema name
func (a FundListingTable) FromSchema(schemaName string) *FundListingTable {
	return newFundListingTable(schemaName, a.TableName(), a.Alias())
}

// WithPrefix creates new FundListingTable with assigned table prefix
func (a FundListingTable) WithPrefix(prefix string) *FundListingTable {
	return newFundListingTable(a.SchemaName(), prefix+a.TableName(), a.TableName())
}

// WithSuffix creates new FundListingTable with assigned table suffix
func (a FundListingTable) WithSuffix(suffix string) *FundListingTable {
	return newFundListingTable(a.SchemaName(), a.TableName()+suffix, a.TableName())
}

func newFundListingTable(schemaName, tableName, alias string) *FundListingTable {
	return &FundListingTable{
		fundListingTable: newFundListingTableImpl(schemaName, tableName, alias),
		EXCLUDED:         newFundListingTableImpl("", "excluded", ""),
	}
}

func newFundListingTableImpl(schemaName, tableName, alias string) fundListingTable {
	var (
		FundIDColumn   = postgres.StringColumn("fund_id")
		TickerColumn   = postgres.StringColumn("ticker")
		allColumns     = postgres.ColumnList{FundIDColumn, TickerColumn}
		mutableColumns = postgres.ColumnList{FundIDColumn, TickerColumn}
	)

	return fundListingTable{
		Table: postgres.NewTable(schemaName, tableName, alias, allColumns...),

		//Columns
		FundID: FundIDColumn,
		Ticker: TickerColumn,

		AllColumns:     allColumns,
		MutableColumns: mutableColumns,
	}
}
