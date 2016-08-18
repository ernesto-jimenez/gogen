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
		var s *Embedded = &s.First
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
			if s.Second == nil {
				s.Second = &Embedded{}
			}
			s := s.Second

			if v, ok := m["Field"].(string); ok {
				s.Field = v

			} else if v, exists := m["Field"]; exists && v != nil {
				return fmt.Errorf("expected field Field to be string but got %T", m["Field"])
			}

		} else if p == nil {
			s.Second = nil
		} else {
			return fmt.Errorf("expected field Second to be map[string]interface{} but got %T", p)
		}

	}

	// ArrayOrSlice Third

	if v, ok := m["Third"].([]interface{}); ok {

		s.Third = make([]Embedded, len(v))

		prev := s
		for i, el := range v {
			var s *Embedded

			s = &prev.Third[i]

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

	// ArrayOrSlice Fourth

	if v, ok := m["Fourth"].([]interface{}); ok {

		s.Fourth = make([]*Embedded, len(v))

		prev := s
		for i, el := range v {
			var s *Embedded

			if el == nil {
				continue
			}
			prev.Fourth[i] = &Embedded{}
			s = prev.Fourth[i]

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

	// ArrayOrSlice Fifth

	if v, ok := m["Fifth"].([]interface{}); ok {

		if len(s.Fifth) < len(v) {
			return fmt.Errorf("expected field Fifth to be an array with %d elements, but got an array with %d", len(s.Fifth), len(v))
		}

		prev := s
		for i, el := range v {
			var s *Embedded

			s = &prev.Fifth[i]

			if m, ok := el.(map[string]interface{}); ok {
				// Fill object

				if v, ok := m["Field"].(string); ok {
					s.Field = v

				} else if v, exists := m["Field"]; exists && v != nil {
					return fmt.Errorf("expected field Field to be string but got %T", m["Field"])
				}

			}
		}
	} else if v, exists := m["Fifth"]; exists && v != nil {
		return fmt.Errorf("expected field Fifth to be []interface{} but got %T", m["Fifth"])
	}

	// ArrayOrSlice Sixth

	if v, ok := m["Sixth"].([]interface{}); ok {

		if len(s.Sixth) < len(v) {
			return fmt.Errorf("expected field Sixth to be an array with %d elements, but got an array with %d", len(s.Sixth), len(v))
		}

		prev := s
		for i, el := range v {
			var s *Embedded

			if el == nil {
				continue
			}
			prev.Sixth[i] = &Embedded{}
			s = prev.Sixth[i]

			if m, ok := el.(map[string]interface{}); ok {
				// Fill object

				if v, ok := m["Field"].(string); ok {
					s.Field = v

				} else if v, exists := m["Field"]; exists && v != nil {
					return fmt.Errorf("expected field Field to be string but got %T", m["Field"])
				}

			}
		}
	} else if v, exists := m["Sixth"]; exists && v != nil {
		return fmt.Errorf("expected field Sixth to be []interface{} but got %T", m["Sixth"])
	}

	return nil
}
