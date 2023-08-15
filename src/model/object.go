package model

import "fmt"

// SolarisEquatorial represents equatorial coordinates of
// - right ascension
// - declination
// - distance
type SolarisEquatorial struct {
	RightAscension float64 `json:"RightAscension"`
	Declination    float64 `json:"Declination"`
	Distance       float64 `json:"Distance"`
}

// String retrieves a string representation of equatorial coordinates
func (e SolarisEquatorial) String() string {
	return fmt.Sprintf("eq{ %f, %f, %f }", e.RightAscension, e.Declination, e.Distance)
}

// NewSolarisEquatorial creates a new SolarisEquatorial instance
func NewSolarisEquatorial() SolarisEquatorial {
	return SolarisEquatorial{
		RightAscension: 0,
		Declination:    0,
		Distance:       0,
	}
}

// SolarisDesignation represents a designation in a
// specific catalog, e.g. M107
// - catalog (Messier)
// - index (107)
type SolarisDesignation struct {
	Catalog string `json:"Catalog"`
	Index   uint64 `json:"Index"`
}

// NewSolarisDesignation creates a new designation instance
func NewSolarisDesignation() SolarisDesignation {
	return SolarisDesignation{
		Catalog: "",
		Index:   0,
	}
}

// String retrieves a string representation of a designation
func (d SolarisDesignation) String() string {
	return fmt.Sprintf("designation{ %s [%d] }", d.Catalog, d.Index)
}

// SolarisObject represents a celestial object
// seen from the prime meridian of planet earth
// with properties of
// - designation (e.g. I474)
// - objectType (e.g. planetary cluster)
// - position (equatorial)
// - dimension (angle size) [']
// - magnitude (brightness) [1]
type SolarisObject struct {
	Designation   SolarisDesignation `json:"Designation"`
	ObjectType    string             `json:"ObjectType"`
	Constellation string             `json:"Constellation"`
	Position      SolarisEquatorial  `json:"Position"`
	Dimension     float64            `json:"Dimension"`
	Magnitude     float64            `json:"Magnitude"`
}

// NewSolarisObject creates a new SolarisObject instance
func NewSolarisObject() SolarisObject {
	return SolarisObject{
		Designation:   NewSolarisDesignation(),
		ObjectType:    "unidentified",
		Constellation: "",
		Position:      NewSolarisEquatorial(),
		Dimension:     0,
		Magnitude:     0,
	}
}

// String retrieves a string representation of a SolarisObject
func (o SolarisObject) String() string {
	return fmt.Sprintf("%s - %s (%s): %s dim{ %f' } mag{ %f }",
		o.Designation.String(),
		o.ObjectType,
		o.Constellation,
		o.Position.String(),
		o.Dimension,
		o.Magnitude,
	)
}
