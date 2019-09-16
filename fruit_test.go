package main

import (
	"testing"

	"github.com/pangpanglabs/goutils/test"
)

func TestJsonCR(t *testing.T) {
	colors := map[string]string{"red": "1"}
	stores := map[string]*Store{"001": &Store{Name: "果想你"},
		"002": &Store{Name: "许仙网"}}
	var ID int64
	t.Run("Create", func(t *testing.T) {
		f := &Fruit{
			Colors: colors,
			Stores: stores,
		}
		affectedRow, err := f.Create()
		test.Ok(t, err)
		test.Equals(t, int64(1), affectedRow)
		ID = f.Id
		test.Assert(t, ID != int64(0), "create failure")
	})
	t.Run("Get", func(t *testing.T) {
		has, v, err := Fruit{}.GetById(ID)
		test.Ok(t, err)
		test.Equals(t, true, has)
		test.Equals(t, colors, v.Colors)
		test.Equals(t, stores, v.Stores)
	})
}
