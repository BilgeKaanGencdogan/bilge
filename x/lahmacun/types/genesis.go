package types

import fmt "fmt"

// this line is used by starport scaffolding # genesis/types/import

// DefaultIndex is the default global index
const DefaultIndex uint64 = 1

// DefaultGenesis returns the default genesis state
func DefaultGenesis() *GenesisState {
	return &GenesisState{
		LahmacunList: []*Lahmacun{
			{
				Creator: "admin",
				Id:      "1",
				Heat:    "medium",
				Dough:   "thin",
				Meat:    "beef",
				Tomato:  "fresh",
				Onion:   "sliced",
			},
			{
				Creator: "admin",
				Id:      "2",
				Heat:    "hot",
				Dough:   "thick",
				Meat:    "lamb",
				Tomato:  "diced",
				Onion:   "chopped",
			},
		},
		// this line is used by starport scaffolding # genesis/types/default
		Params: DefaultParams(),
	}
}

// Validate performs basic genesis state validation returning an error upon any
// failure.
func (gs GenesisState) Validate() error {
	for _, elem := range gs.LahmacunList {
		if err := elem.Validate(); err != nil {
			return fmt.Errorf("invalid lahmacun: %w", err)
		}
	}
	// this line is used by starport scaffolding # genesis/types/validate

	return gs.Params.Validate()
}

// Validate performs basic validation on the Lahmacun object.
func (l Lahmacun) Validate() error {
	if l.Heat == "" {
		return fmt.Errorf("heat cannot be empty")
	}
	if l.Dough == "" {
		return fmt.Errorf("dough cannot be empty")
	}
	if l.Meat == "" {
		return fmt.Errorf("meat cannot be empty")
	}
	if l.Tomato == "" {
		return fmt.Errorf("tomato cannot be empty")
	}
	if l.Onion == "" {
		return fmt.Errorf("onion cannot be empty")
	}
	return nil
}
