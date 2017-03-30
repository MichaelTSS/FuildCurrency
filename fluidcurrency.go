package fluidcurrency

import (
//   "time"
  "errors"
)

// Rates is the object returned by the library
type Rates struct {
	Base           string // EUR
	Millitimestamp int64 // 1490882558385
  Currencies     map[string]float64 // { "USD": 1.64526, "CAN": 0.83645 }
}

// Parameters are passed in the GetRates method
type Parameters struct {
  From           string // EUR
	Millitimestamp int64 // 1490882558385
	To             []string // { "USD", "CAN" }
}

// FluidCurrency is an exported struct with the same name as the package. Can we do better?
type FluidCurrency struct {
  ProviderName string
}

// NewFluidCurrency is the helper function to initialize the module
func NewFluidCurrency(providerName string) *FluidCurrency {
  if providerName == "" {
    providerName = "APIFixer"
  }
  return &FluidCurrency{
    ProviderName: providerName,
  }
}

// GetRates is the main exported function
func (fluidC FluidCurrency) GetRates(pms *Parameters) (*Rates, error) {
  // set defaults parameters if not set
  // return normalized value
  switch fluidC.ProviderName {
    case "APIFixer":
      p := APIFixerProvider{}
      return p.GetRatesFromAPI(pms)
  }
  return nil, errors.New("Provider not supported")
}
