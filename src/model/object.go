package model

import (
	"fmt"
	"solaris-gen/utility"
)

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
	return fmt.Sprintf("{ %f, %f, %f }", e.RightAscension, e.Declination, e.Distance)
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

// CatalogText retrieves the text representation of the catalog
func (d SolarisDesignation) CatalogText() string {
	return d.Catalog
}

// CatalogDefinition retrieves the enum representation of the catalog
func (d SolarisDesignation) CatalogDefinition() string {
	return fmt.Sprintf("CATALOG_%s", utility.AsEnumDefinition(d.Catalog))
}

// String retrieves the object definition representation of the designation
func (d SolarisDesignation) String() string {
	return fmt.Sprintf("{ %s, %d }", d.CatalogDefinition(), d.Index)
}

// SolarisObject represents a celestial object
// seen from the prime meridian of planet earth
// with properties of
// - designation (e.g. I474)
// - position (equatorial)
// - objectType (e.g. planetary cluster)
// - dimension (angle size) [']
// - magnitude (brightness) [1]
type SolarisObject struct {
	Designation   SolarisDesignation `json:"Designation"`
	Position      SolarisEquatorial  `json:"Position"`
	ObjectType    string             `json:"ObjectType"`
	Constellation string             `json:"Constellation"`
	Dimension     float64            `json:"Dimension"`
	Magnitude     float64            `json:"Magnitude"`
}

// ObjectTypeText retrieves the text representation of the object type
func (o SolarisObject) ObjectTypeText() string {
	return o.ObjectType
}

// ObjectTypeDefinition retrieves the enum representation of the object type
func (o SolarisObject) ObjectTypeDefinition() string {
	return fmt.Sprintf("OBJECT_TYPE_%s", utility.AsEnumDefinition(o.ObjectType))
}

// ConstellationText retrieves the text representation of the constellation
func (o SolarisObject) ConstellationText() string {
	return o.Constellation
}

// ConstellationDefinition retrieves the enum representation of the constellation
func (o SolarisObject) ConstellationDefinition() string {
	return fmt.Sprintf("CONSTELLATION_%s", utility.AsEnumDefinition(o.Constellation))
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
	return fmt.Sprintf("{ %s, %s, %s, %s, %f, %f }",
		o.Designation.String(),
		o.Position.String(),
		o.ObjectTypeDefinition(),
		o.ConstellationDefinition(),
		o.Dimension,
		o.Magnitude,
	)
}
