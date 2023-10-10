//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package table

// UseSchema sets a new schema name for all generated table SQL builder types. It is recommended to invoke
// this method only once at the beginning of the program.
func UseSchema(schema string) {
	Fund = Fund.FromSchema(schema)
	FundHolding = FundHolding.FromSchema(schema)
	FundListing = FundListing.FromSchema(schema)
	Holding = Holding.FromSchema(schema)
	Portfolio = Portfolio.FromSchema(schema)
	PortfolioFund = PortfolioFund.FromSchema(schema)
	User = User.FromSchema(schema)
}
