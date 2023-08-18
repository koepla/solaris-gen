
// Copyright (c) 2023 koepla
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all
// copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
// SOFTWARE.

#ifndef SOLARIS_GENERATED_OBJECTS_H
#define SOLARIS_GENERATED_OBJECTS_H

// type of celestial object
// -------------------------------------------------------------
// visit https://cdsarc.cds.unistra.fr/ftp/VII/118/ReadMe
// for full list of objects
// -------------------------------------------------------------
enum solaris_object_type {
	OBJECT_TYPE_GALAXY,
	OBJECT_TYPE_UNIDENTIFIED,
	OBJECT_TYPE_NONEXISTENT,
	OBJECT_TYPE_REFLECTION_NEBULA,
	OBJECT_TYPE_SINGLE_STAR,
	OBJECT_TYPE_DOUBLE_STAR,
	OBJECT_TYPE_ASTERISM,
	OBJECT_TYPE_PLANETARY_NEBULA,
	OBJECT_TYPE_UNCERTAIN,
	OBJECT_TYPE_GLOBULAR_STAR_CLUSTER,
	OBJECT_TYPE_OPEN_STAR_CLUSTER,
	OBJECT_TYPE_CLUSTER,
	OBJECT_TYPE_TRIPLE_STAR,
	OBJECT_TYPE_KNOT,
	OBJECT_TYPE_PHOTOGRAPHIC_PLATE_DEFECT,
};

// string representation of the specified object type
// --------------------------------------------------
// @param type -> object type
// --------------------------------------------------
SOLARIS_API const char *solaris_object_type_to_string(enum solaris_object_type type);

// catalog type
// -------------------------------------------------------------
// visit https://cdsarc.cds.unistra.fr/ftp/VII/118/ReadMe
// -------------------------------------------------------------
enum solaris_catalog_type {
	CATALOG_INDEX_CATALOG,
	CATALOG_NEW_GENERAL_CATALOG,
};

// string representation of the specified catalog type
// --------------------------------------------------
// @param type -> catalog type
// --------------------------------------------------
SOLARIS_API const char *solaris_catalog_type_to_string(enum solaris_catalog_type type);

// constellation type
// -------------------------------------------------------------
// visit https://cdsarc.cds.unistra.fr/ftp/VII/118/ReadMe
// -------------------------------------------------------------
enum solaris_constellation {
	CONSTELLATION_ANDROMEDA,
	CONSTELLATION_CASSIOPEIA,
	CONSTELLATION_PISCES,
	CONSTELLATION_PEGASUS,
	CONSTELLATION_TUCANA,
	CONSTELLATION_SCULPTOR,
	CONSTELLATION_CETUS,
	CONSTELLATION_CEPHEUS,
	CONSTELLATION_PHOENIX,
	CONSTELLATION_HYDRUS,
	CONSTELLATION_TRIANGULUM,
	CONSTELLATION_OCTANS,
	CONSTELLATION_PERSEUS,
	CONSTELLATION_ARIES,
	CONSTELLATION_FORNAX,
	CONSTELLATION_ERIDANUS,
	CONSTELLATION_HOROLOGIUM,
	CONSTELLATION_RETICULUM,
	CONSTELLATION_TAURUS,
	CONSTELLATION_CAMELOPARDALIS,
	CONSTELLATION_MENSA,
	CONSTELLATION_DORADO,
	CONSTELLATION_CAELUM,
	CONSTELLATION_ORION,
	CONSTELLATION_PICTOR,
	CONSTELLATION_AURIGA,
	CONSTELLATION_LEPUS,
	CONSTELLATION_COLUMBA,
	CONSTELLATION_GEMINI,
	CONSTELLATION_MONOCEROS,
	CONSTELLATION_CARINA,
	CONSTELLATION_PUPPIS,
	CONSTELLATION_CANIS_MAJOR,
	CONSTELLATION_LYNX,
	CONSTELLATION_VOLANS,
	CONSTELLATION_CANIS_MINOR,
	CONSTELLATION_CANCER,
	CONSTELLATION_VELA,
	CONSTELLATION_HYDRA,
	CONSTELLATION_PYXIS,
	CONSTELLATION_URSA_MAJOR,
	CONSTELLATION_LEO,
	CONSTELLATION_LEO_MINOR,
	CONSTELLATION_CHAMAELEON,
	CONSTELLATION_ANTLIA,
	CONSTELLATION_DRACO,
	CONSTELLATION_SEXTANS,
	CONSTELLATION_CRATER,
	CONSTELLATION_CENTAURUS,
	CONSTELLATION_VIRGO,
	CONSTELLATION_URSA_MINOR,
	CONSTELLATION_MUSCA,
	CONSTELLATION_CORVUS,
	CONSTELLATION_COMA_BERENICES,
	CONSTELLATION_CRUX,
	CONSTELLATION_CANES_VENATICI,
	CONSTELLATION_BO�TES,
	CONSTELLATION_CIRCINUS,
	CONSTELLATION_APUS,
	CONSTELLATION_LIBRA,
	CONSTELLATION_TRIANGULUM_AUSTRALE,
	CONSTELLATION_SERPENS,
	CONSTELLATION_CORONA_BOREALIS,
	CONSTELLATION_NORMA,
	CONSTELLATION_SCORPIUS,
	CONSTELLATION_HERCULES,
	CONSTELLATION_OPHIUCHUS,
	CONSTELLATION_ARA,
	CONSTELLATION_PAVO,
	CONSTELLATION_SAGITTARIUS,
	CONSTELLATION_CORONA_AUSTRALIS,
	CONSTELLATION_TELESCOPIUM,
	CONSTELLATION_LYRA,
	CONSTELLATION_SCUTUM,
	CONSTELLATION_AQUILA,
	CONSTELLATION_VULPECULA,
	CONSTELLATION_CYGNUS,
	CONSTELLATION_SAGITTA,
	CONSTELLATION_CAPRICORNUS,
	CONSTELLATION_DELPHINUS,
	CONSTELLATION_MICROSCOPIUM,
	CONSTELLATION_INDUS,
	CONSTELLATION_AQUARIUS,
	CONSTELLATION_EQUULEUS,
	CONSTELLATION_GRUS,
	CONSTELLATION_PISCIS_AUSTRINUS,
	CONSTELLATION_LACERTA,
};

// string representation of the specified constellation
// ----------------------------------------------------
// @param type -> object type
// ----------------------------------------------------
SOLARIS_API const char *solaris_constellation_to_string(enum solaris_constellation constellation);

// celestial object
// ------------------------------------------------------------------
// @member designation   -> NGC or IC designation
// @member position      -> equatorial position
// @member type          -> see @solaris_object_type
// @member constellation -> constellation where the object resides in
// @member dimension     -> dimension in minutes of arc
// @member magnitude     -> apparent magnitude
// ------------------------------------------------------------------
struct solaris_object {
	struct solaris_designation designation;
	struct solaris_equatorial position;
	enum solaris_object_type type;
	enum solaris_constellation constellation;
	f32 dimension;
	f32 magnitude;
};

// celestial object list
// ------------------------------------------------------------------
// @member objects -> array of objects
// @member size    -> list size
// ------------------------------------------------------------------
struct solaris_object_list {
	struct solaris_object *objects;
	usize size;
};

// get all objects
// ----------------------------------------------------
// @param type -> object type
// ----------------------------------------------------
SOLARIS_API struct solaris_object_list *solaris_objects();

#endif// SOLARIS_GENERATED_OBJECTS_H

