
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

#include "data/object_generated_min.h"

const char *solaris_object_type_to_string(enum solaris_object_type type) {
	switch (type) {
	    case OBJECT_TYPE_SUPERNOVA_REMNANT:
	        return "Supernova Remnant";
	    case OBJECT_TYPE_SPIRAL_GALAXY:
	        return "Spiral Galaxy";
		default:
			return "unknown";
	}
}

const char *solaris_catalog_type_to_string(enum solaris_catalog_type type) {
	switch (type) {
	    case CATALOG_MESSIER:
	        return "Messier";
		default:
			return "unknown";
	}
}

const char *solaris_constellation_to_string(enum solaris_constellation constellation) {
	switch (constellation) {
	    case CONSTELLATION_TAURUS:
	        return "Taurus";
	    case CONSTELLATION_ANDROMEDA:
	        return "Andromeda";
		default:
			return "unknown";
	}
}

static struct solaris_object generated_objects[] = {
	{ { CATALOG_MESSIER, 1 }, { 5.500000, 22.100000, 1.000000 }, OBJECT_TYPE_SUPERNOVA_REMNANT, CONSTELLATION_TAURUS, 6.000000, 8.400000 },
	{ { CATALOG_MESSIER, 31 }, { 0.750000, 41.600000, 2904.000000 }, OBJECT_TYPE_SPIRAL_GALAXY, CONSTELLATION_ANDROMEDA, 178.000000, 3.400000 },
};

static struct solaris_object_list internal_object_list = {
	generated_objects,
	2
};

struct solaris_object_list *solaris_objects() {
	return &internal_object_list;
}

