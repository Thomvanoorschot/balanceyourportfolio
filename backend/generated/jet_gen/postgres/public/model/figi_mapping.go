//
// Code generated by go-jet DO NOT EDIT.
//
// WARNING: Changes to this file may cause incorrect behavior
// and will be lost if the code is regenerated
//

package model

type FigiMapping struct {
	Figi   string `sql:"primary_key"`
	Ticker *string
	Name   *string
	Isin   *string
	Sedol  *string
	Cusip  *string
}
