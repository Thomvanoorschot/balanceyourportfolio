//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

import (
	"github.com/google/uuid"
	"time"
)

type Fund struct {
	ID                 uuid.UUID `sql:"primary_key"`
	Name               *string
	Currency           *string
	Isin               *string
	TotalHoldings      *float64
	Price              *float64
	Provider           *string
	ExternalIdentifier *string
	OutstandingShares  *float64
	EffectiveDate      *time.Time
}
