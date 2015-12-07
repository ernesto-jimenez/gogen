/*
* CODE GENERATED AUTOMATICALLY WITH github.com/ernesto-jimenez/gogen/unmarshalmap
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package testpkg

import (
	"fmt"
)

// UnmarshalMap takes a map and unmarshals the fieds into the struct
func (s *Composed) UnmarshalMap(m map[string]interface{}) error {

	// Anonymous Embedded
	if scoped := true; scoped {
		var s *Embedded = &s.Embedded
		// Fill object

		if v, ok := m["Field"].(string); ok {
			s.Field = v

		} else if v, exists := m["Field"]; exists && v != nil {
			return fmt.Errorf("expected field Field to be string but got %T", m["Field"])
		}

	}

	if v, ok := m["Base"].(string); ok {
		s.Base = v

	} else if v, exists := m["Base"]; exists && v != nil {
		return fmt.Errorf("expected field Base to be string but got %T", m["Base"])
	}

	return nil
}
