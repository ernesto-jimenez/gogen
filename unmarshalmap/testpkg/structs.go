package testpkg

//go:generate go run ../../cmd/gounmarshalmap/main.go -o simple_struct_unmarshalmap.go SimpleStruct

type SimpleStruct struct {
	SimpleField             string
	SimpleJSONTagged        string  `json:"field2"`
	SimpleJSONTaggedOmitted string  `json:"field3,omitempty"`
	SimpleOmitEmptyNoName   string  `json:",omitempty"`
	SimpleSkipped           string  `json:"-,omitempty"`
	SimplePointer           *string `json:"pointer"`
	SimpleInteger           int     `json:"integer"`
	SimpleIntegerPtr        *int    `json:"integer_ptr"`
	Ignored                 string  `json:"-"`
}

//go:generate go run ../../cmd/gounmarshalmap/main.go -o array_unmarshalmap.go Array

type Array struct {
	List []string
	//ListPointers []*string
	// PointerToList *[]string
}

type Embedded struct {
	Field string
}

//go:generate go run ../../cmd/gounmarshalmap/main.go -o composed_unmarshalmap.go Composed

type Composed struct {
	Embedded
	Base string
}

//go:generate go run ../../cmd/gounmarshalmap/main.go -o nested_unmarshalmap.go Nested

type Nested struct {
	First  Embedded
	Second *Embedded
	Third  []Embedded
	Fourth []*Embedded
	Fifth  [3]Embedded
	Sixth  [3]*Embedded
	// TODO: Implement map support
	// Seventh map[string]Embedded
	// Eight   map[string]*Embedded
	// Nineth  map[int]Embedded
	// Tenth   map[int]*Embedded
}
