package main

import (
	"fmt"
	"solaris-gen/model"
	"solaris-gen/utility"
	"strings"
)

// NgcEater parses gen files of the ngc specification,
// see https://cdsarc.cds.unistra.fr/ftp/VII/118/ReadMe
type NgcEater struct {
	objects []model.SolarisObject
}

// NewNgcEater creates a new NgcEater instance
func NewNgcEater() *NgcEater {
	return &NgcEater{
		objects: []model.SolarisObject{},
	}
}

// Objects returns the objects parsed by the ngc parser
func (eater *NgcEater) Objects() []model.SolarisObject {
	return eater.objects
}

// AsNewConfig creates a SolarisConfig from the eaten objects
func (eater *NgcEater) AsNewConfig() *model.ObjectConfig {
	config := model.NewSolarisConfig()
	config.Objects = eater.objects
	return config
}

// EatLine takes a line extracts object properties from it,
// then appends to NgcParser objects
func (eater *NgcEater) EatLine(line string) {
	if len(line) < 46 {
		fmt.Println("[error] skipping line: invalid length")
		return
	}

	var err error
	object := model.NewSolarisObject()
	if object.Designation, err = eatDesignation(line[0:5]); err != nil {
		fmt.Println("[error] skipping line: unable to retrieve designation")
		return
	}

	if object.ObjectType, err = eatObjectType(line[6:9]); err != nil {
		fmt.Println("[error] skipping line: unable to retrieve object type")
		return
	}

	if object.Position, err = eatEquatorial(line[10:]); err != nil {
		fmt.Println("[error] skipping line: unable to retrieve position")
		return
	}

	// Abbreviated constellation
	if object.Constellation, err = eatConstellation(line[29:32]); err != nil {
		fmt.Println("[error] skipping line: unable to retrieve constellation")
		return
	}

	// Try to parse the dimension
	object.Dimension, err = utility.StringToFloat64(line[33:38])
	if err != nil {
		fmt.Println("[error] skipping line: unable to retrieve dimension")
		return
	}

	// Try to parse the magnitude
	object.Magnitude, err = utility.StringToFloat64(line[40:44])
	if err != nil {
		fmt.Println("[error] skipping line: unable to retrieve magnitude")
		return
	}

	eater.objects = append(eater.objects, object)
}

// eatDesignation retrieves a designation
// from a given input string
func eatDesignation(line string) (model.SolarisDesignation, error) {
	designation := model.NewSolarisDesignation()
	if line[0] == 'I' {
		designation.Catalog = "Index Catalog"
	} else {
		designation.Catalog = "New General Catalog"
	}
	var err error
	if designation.Index, err = utility.StringToUint64(line[1:5]); err != nil {
		return designation, fmt.Errorf("unable to retrieve designation index: %s", err)
	}
	return designation, nil
}

// eatObjectType expands the abbreviated NGC object type
func eatObjectType(line string) (string, error) {
	objectTypeTable := map[string]string{
		" Gx": "Galaxy",
		" OC": "Open Star Cluster",
		" Gb": "Globular Star Cluster",
		" Nb": "Reflection Nebula",
		" Pl": "Planetary Nebula",
		"C+N": "Cluster",
		"Ast": "Asterism",
		" Kt": "Knot",
		"***": "Triple Star",
		" D*": "Double Star",
		"  *": "Single Star",
		"  ?": "Uncertain",
		"   ": "Unidentified",
		"  -": "Nonexistent",
		" PD": "Photographic Plate Defect",
	}
	if strings.Contains(line, "?") {
		return "Uncertain", nil
	}
	if expanded, ok := objectTypeTable[line]; ok {
		return expanded, nil
	}
	return line, fmt.Errorf("unable to expand abbreviated objectType: %s", line)
}

// eatConstellation expands the abbreviated NGC constellation names
func eatConstellation(line string) (string, error) {
	constellationTable := map[string]string{
		"And": "Andromeda",
		"Ant": "Antlia",
		"Aps": "Apus",
		"Aqr": "Aquarius",
		"Aql": "Aquila",
		"Ara": "Ara",
		"Ari": "Aries",
		"Aur": "Auriga",
		"Boo": "Bo�tes",
		"Cae": "Caelum",
		"Cam": "Camelopardalis",
		"Cnc": "Cancer",
		"CVn": "Canes Venatici",
		"CMa": "Canis Major",
		"CMi": "Canis Minor",
		"Cap": "Capricornus",
		"Car": "Carina",
		"Cas": "Cassiopeia",
		"Cen": "Centaurus",
		"Cep": "Cepheus",
		"Cet": "Cetus",
		"Cha": "Chamaeleon",
		"Cir": "Circinus",
		"Col": "Columba",
		"Com": "Coma Berenices",
		"CrA": "Corona Australis",
		"CrB": "Corona Borealis",
		"Crv": "Corvus",
		"Crt": "Crater",
		"Cru": "Crux",
		"Cyg": "Cygnus",
		"Del": "Delphinus",
		"Dor": "Dorado",
		"Dra": "Draco",
		"Equ": "Equuleus",
		"Eri": "Eridanus",
		"For": "Fornax",
		"Gem": "Gemini",
		"Gru": "Grus",
		"Her": "Hercules",
		"Hor": "Horologium",
		"Hya": "Hydra",
		"Hyi": "Hydrus",
		"Ind": "Indus",
		"Lac": "Lacerta",
		"Leo": "Leo",
		"Lib": "Libra",
		"LMi": "Leo Minor",
		"Lep": "Lepus",
		"Lup": "Lupus",
		"Lyn": "Lynx",
		"Lyr": "Lyra",
		"Men": "Mensa",
		"Mic": "Microscopium",
		"Mon": "Monoceros",
		"Mus": "Musca",
		"Nor": "Norma",
		"Oct": "Octans",
		"Oph": "Ophiuchus",
		"Ori": "Orion",
		"Pav": "Pavo",
		"Peg": "Pegasus",
		"Per": "Perseus",
		"Phe": "Phoenix",
		"Pic": "Pictor",
		"Psc": "Pisces",
		"PsA": "Piscis Austrinus",
		"Pup": "Puppis",
		"Pyx": "Pyxis",
		"Ret": "Reticulum",
		"Sge": "Sagitta",
		"Sgr": "Sagittarius",
		"Sco": "Scorpius",
		"Scl": "Sculptor",
		"Sct": "Scutum",
		"Ser": "Serpens",
		"Sex": "Sextans",
		"Tau": "Taurus",
		"Tel": "Telescopium",
		"Tri": "Triangulum",
		"TrA": "Triangulum Australe",
		"Tuc": "Tucana",
		"UMa": "Ursa Major",
		"UMi": "Ursa Minor",
		"Vel": "Vela",
		"Vir": "Virgo",
		"Vol": "Volans",
		"Vul": "Vulpecula",
	}
	if expanded, ok := constellationTable[line]; ok {
		return expanded, nil
	}
	return line, fmt.Errorf("unable to expand abbreviated constellation: %s", line)
}

// eatEquatorial retrieves equatorial position information
// from a given input string
func eatEquatorial(line string) (model.SolarisEquatorial, error) {
	var position model.SolarisEquatorial
	rightAscensionHours, err := utility.StringToFloat64(line[0:2])
	if err != nil {
		return position, fmt.Errorf("unable to retrieve hours of right ascension: %s", err)
	}

	rightAscensionMinutes, err := utility.StringToFloat64(line[3:7])
	if err != nil {
		return position, fmt.Errorf("unable to retrieve minutes of right ascension: %s", err)
	}

	declinationSign := line[9]
	declinationDegrees, err := utility.StringToFloat64(line[10:12])
	if err != nil {
		return position, fmt.Errorf("unable to retrieve degrees of declination: %s", err)
	}

	declinationMinutes, err := utility.StringToFloat64(line[13:15])
	if err != nil {
		return position, fmt.Errorf("unable to retrieve minutes of declination: %s", err)
	}

	var declination float64
	if declinationSign == '-' {
		declination = -1 * (declinationDegrees + declinationMinutes/60)
	} else {
		declination = declinationDegrees + declinationMinutes/60
	}

	position.RightAscension = 15 * (rightAscensionHours + rightAscensionMinutes/60)
	position.Declination = declination
	position.Distance = 0
	return position, nil
}
