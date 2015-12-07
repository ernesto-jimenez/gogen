/*
* CODE GENERATED AUTOMATICALLY WITH github.com/ernesto-jimenez/gogen/unmarshalmap
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package testpkg

import (
	"fmt"
)

// UnmarshalMap takes a map and unmarshals the fieds into the struct
func (s *SimpleStruct) UnmarshalMap(m map[string]interface{}) error {

	if v, ok := m["SimpleField"].(string); ok {
		s.SimpleField = v
	} else if v, exists := m["SimpleField"]; exists && v != nil {
		return fmt.Errorf("expected field SimpleField to be string but got %T", m["SimpleField"])
	}

	if v, ok := m["field2"].(string); ok {
		s.SimpleJSONTagged = v
	} else if v, exists := m["field2"]; exists && v != nil {
		return fmt.Errorf("expected field field2 to be string but got %T", m["field2"])
	}

	if v, ok := m["field3"].(string); ok {
		s.SimpleJSONTaggedOmitted = v
	} else if v, exists := m["field3"]; exists && v != nil {
		return fmt.Errorf("expected field field3 to be string but got %T", m["field3"])
	}

	// Pointer SimplePointer
	if p, ok := m["pointer"]; ok {

		if m, ok := p.(string); ok {
			s.SimplePointer = &m
		}

	}

	return nil
}
