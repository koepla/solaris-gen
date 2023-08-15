package gen

import (
	"fmt"
	"github.com/agnivade/levenshtein"
	"solaris-gen/model"
	"solaris-gen/utility"
	"strings"
)

// EnumEntry represents an enum entry of generated code
type EnumEntry string

// String expands the SolarisEnum to meet the string spec
func (e EnumEntry) String() string {
	parts := strings.Split(string(e), "_")
	return strings.Join(utility.Transform(parts, utility.PascalCase), " ")
}

// Enum represents an enum type of generated code
type Enum struct {
	Entries []EnumEntry
}

// NewEnum creates a new SolarisEnum instance
func NewEnum() Enum {
	return Enum{
		Entries: []EnumEntry{},
	}
}

// Append tries to append the enum entry to the enum
func (e *Enum) Append(entry EnumEntry) {
	var closeCandidate *EnumEntry
	for _, enumEntry := range e.Entries {
		if levenshtein.ComputeDistance(string(enumEntry), string(entry)) < 2 {
			closeCandidate = &enumEntry
			break
		}
	}
	if closeCandidate != nil {
		return
	}
	e.Entries = append(e.Entries, entry)
}

// CodeWriter generates code
type CodeWriter struct {
	ObjectTypes    Enum
	CatalogTypes   Enum
	Constellations Enum
}

// NewCodeWriter creates a SolarisWriter instance
func NewCodeWriter() *CodeWriter {
	return &CodeWriter{
		ObjectTypes:    NewEnum(),
		CatalogTypes:   NewEnum(),
		Constellations: NewEnum(),
	}
}

// buildConfig builds the writer config to the enum symbols
func (writer *CodeWriter) buildConfig(config *model.ObjectConfig) {
	for _, entry := range config.Objects {
		writer.ObjectTypes.Append(EnumEntry(entry.ObjectType))
		writer.CatalogTypes.Append(EnumEntry(entry.Designation.Catalog))
		writer.Constellations.Append(EnumEntry(entry.Constellation))
	}
}

// GenerateCode produces the source and header files that are used by solaris
func (writer *CodeWriter) GenerateCode(config *model.ObjectConfig) (source string, header string, err error) {
	writer.buildConfig(config)
	fmt.Println("ObjectTypes:")
	for _, objectType := range writer.ObjectTypes.Entries {
		fmt.Println(objectType.String())
	}
	fmt.Println("CatalogTypes:")
	for _, catalogType := range writer.CatalogTypes.Entries {
		fmt.Println(catalogType.String())
	}
	fmt.Println("Constellations:")
	for _, constellation := range writer.Constellations.Entries {
		fmt.Println(constellation.String())
	}
	return source, header, nil
}
