package fluidcurrency

import (
  "fmt"
  "time"
	"testing"
)

func TestRates(t *testing.T) {
  unixMilli := time.Now().UnixNano() / 1000000
  var rates = Rates{
    Base: "EUR",
    Millitimestamp: unixMilli,
  }

  if rates.Base != "EUR" {
    t.Errorf("rates.Base == %q, want %q", rates.Base, "EUR")
  }
  if rates.Millitimestamp != unixMilli {
    t.Errorf("rates.Millitimestamp == %q, want %q", rates.Millitimestamp, unixMilli)
  }
}
//
func TestNewFluidCurrency(t *testing.T) {
  fluidCurrency := FluidCurrency{ "APIFixer" }
  pms := Parameters{
    From: "EUR",
    To: []string{"USD"},
  }
  r, e := fluidCurrency.GetRates(&pms)
  //
  if e != nil {
    t.Errorf("error == %q, want %q", e, nil)
  }
  if r.Base != "EUR" {
    t.Errorf("r.Base == %q, want %q", r.Base, "EUR")
  }
}
