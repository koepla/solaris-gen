package main

const HeaderTemplateDefinition = `
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
	{{- range .ObjectTypes.Entries }}
	{{ .Definition }},
	{{- end }}
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
    {{- range .Catalogs.Entries }}
	{{ .Definition }},
	{{- end }}
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
	{{- range .Constellations.Entries }}
	{{ .Definition }},
	{{- end }}
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

`

const SourceTemplateDefinition = `
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

#include "{{ .HeaderFile.Name }}"

const char *solaris_object_type_to_string(enum solaris_object_type type) {
	switch (type) {
		{{- range .ObjectTypes.Entries }}
	    case {{ .Definition }}:
	        return "{{ .Text }}";
		{{- end }}
		default:
			return "unknown";
	}
}

const char *solaris_catalog_type_to_string(enum solaris_catalog_type type) {
	switch (type) {
		{{- range .Catalogs.Entries }}
	    case {{ .Definition }}:
	        return "{{ .Text }}";
		{{- end }}
		default:
			return "unknown";
	}
}

const char *solaris_constellation_to_string(enum solaris_constellation constellation) {
	switch (constellation) {
		{{- range .Constellations.Entries }}
	    case {{ .Definition }}:
	        return "{{ .Text }}";
		{{- end }}
		default:
			return "unknown";
	}
}

static struct solaris_object generated_objects[] = {
	{{- range .Objects }}
	{{ . }},
	{{- end }}
};

static struct solaris_object_list internal_object_list = {
	generated_objects,
	{{ .ObjectCount }}
};

struct solaris_object_list *solaris_objects() {
	return &internal_object_list;
}

`
