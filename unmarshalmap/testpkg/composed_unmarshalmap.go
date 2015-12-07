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

	// Struct Embedded
	if m, ok := m["Embedded"].(map[string]interface{}); ok {
		s := &s.Embedded
		// Fill object

		if v, ok := m["Field"].(string); ok {
			s.Field = v
		} else if v, exists := m["Field"]; exists && v != nil {
			return fmt.Errorf("expected field Field to be string but got %T", m["Field"])
		}

	} else if v, exists := m["Embedded"]; exists && v != nil {
		return fmt.Errorf("expected field Embedded to be map[string]interface{} but got %T", m["Embedded"])
	}

	if v, ok := m["Base"].(string); ok {
		s.Base = v
	} else if v, exists := m["Base"]; exists && v != nil {
		return fmt.Errorf("expected field Base to be string but got %T", m["Base"])
	}

	return nil
}
