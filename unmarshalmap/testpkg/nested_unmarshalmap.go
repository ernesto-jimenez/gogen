/*
* CODE GENERATED AUTOMATICALLY WITH github.com/ernesto-jimenez/gogen/unmarshalmap
* THIS FILE SHOULD NOT BE EDITED BY HAND
 */

package testpkg

import (
	"fmt"
)

// UnmarshalMap takes a map and unmarshals the fieds into the struct
func (s *Nested) UnmarshalMap(m map[string]interface{}) error {

	// Struct First
	if m, ok := m["First"].(map[string]interface{}); ok {
		s := &s.First
		// Fill object

		if v, ok := m["Field"].(string); ok {
			s.Field = v
		} else if v, exists := m["Field"]; exists && v != nil {
			return fmt.Errorf("expected field Field to be string but got %T", m["Field"])
		}

	} else if v, exists := m["First"]; exists && v != nil {
		return fmt.Errorf("expected field First to be map[string]interface{} but got %T", m["First"])
	}

	// Pointer Second
	if p, ok := m["Second"]; ok {

		if m, ok := p.(map[string]interface{}); ok {
			s.Second = &Embedded{}
			s := s.Second

			if v, ok := m["Field"].(string); ok {
				s.Field = v
			} else if v, exists := m["Field"]; exists && v != nil {
				return fmt.Errorf("expected field Field to be string but got %T", m["Field"])
			}

		}

	}

	// Array Third

	if v, ok := m["Third"].([]interface{}); ok {
		s.Third = make([]Embedded, len(v))
		prev := s
		for i, el := range v {

			s := &prev.Third[i]

			if m, ok := el.(map[string]interface{}); ok {
				// Fill object

				if v, ok := m["Field"].(string); ok {
					s.Field = v
				} else if v, exists := m["Field"]; exists && v != nil {
					return fmt.Errorf("expected field Field to be string but got %T", m["Field"])
				}

			}
		}
	} else if v, exists := m["Third"]; exists && v != nil {
		return fmt.Errorf("expected field Third to be []interface{} but got %T", m["Third"])
	}

	// Array Fourth

	if v, ok := m["Fourth"].([]interface{}); ok {
		s.Fourth = make([]*Embedded, len(v))
		prev := s
		for i, el := range v {

			prev.Fourth[i] = &Embedded{}
			s := prev.Fourth[i]

			if m, ok := el.(map[string]interface{}); ok {
				// Fill object

				if v, ok := m["Field"].(string); ok {
					s.Field = v
				} else if v, exists := m["Field"]; exists && v != nil {
					return fmt.Errorf("expected field Field to be string but got %T", m["Field"])
				}

			}
		}
	} else if v, exists := m["Fourth"]; exists && v != nil {
		return fmt.Errorf("expected field Fourth to be []interface{} but got %T", m["Fourth"])
	}

	return nil
}
