// Copyright (c) 2013 Melvin Tercan, https://github.com/melvinmt
//
// Permission is hereby granted, free of charge, to any person obtaining a copy of this 
// software and associated documentation files (the "Software"), to deal in the Software 
// without restriction, including without limitation the rights to use, copy, modify, 
// merge, publish, distribute, sublicense, and/or sell copies of the Software, and to permit 
// persons to whom the Software is furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in all copies or 
// substantial portions of the Software.
// 
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, 
// INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR 
// PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE 
// FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR 
// OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER 
// DEALINGS IN THE SOFTWARE.

package gt

import (
	"testing"
)

func TestBuild(t *testing.T) {
	g := &Build{
		Index: Strings{
			"homepage-greeting": {
				"en":    "Welcome to %s#title, %s#name!",
				"es-LA": "¡Bienvenido a %s#title, %s#title!",
			},
		},
		Origin: "en",
		Target: "es-LA",
	}
	_ = g
}

func TestNoTargetString(t *testing.T) {
	g := &Build{
		Index: Strings{
			"homepage-greeting": {
				"en": "Welcome to %s#title, %s#name!",
			},
		},
		Origin: "en",
		Target: "es-LA",
	}
	func() {
		t.Log("Test key")
		s := "homepage-greeting"
		tr, err := g.Translate(s)
		if s != tr {
			t.Errorf("Returned string '%s' is not the same as input string '%s'", tr, s)
		}
		if err == nil {
			t.Error("Should return: could not find error")
		}
	}()
	func() {
		t.Log("Test string")
		s := "Welcome to %s#title, %s#name!"
		tr, err := g.Translate(s)
		if s != tr {
			t.Errorf("Returned string '%s' is not the same as input string '%s'", tr, s)
		}
		if err == nil {
			t.Error("Should return: could not find error")
		}
	}()
}

func TestNoOriginString(t *testing.T) {
	g := &Build{
		Index: Strings{
			"homepage-greeting": {
				"es-LA": "¡Bienvenido a %s#title, %s#title!",
			},
		},
		Origin: "en",
		Target: "es-LA",
	}
	func() {
		t.Log("Test key")
		s := "homepage-greeting"
		tr, err := g.Translate(s)
		if s != tr {
			t.Errorf("Returned string '%s' is not the same as input string '%s'", tr, s)
		}
		if err == nil {
			t.Error("Should return: could not find error")
		}
	}()
	func() {
		t.Log("Test string")
		s := "Welcome to %s#title, %s#name!"
		tr, err := g.Translate(s)
		if s != tr {
			t.Errorf("Returned string '%s' is not the same as input string '%s'", tr, s)
		}
		if err == nil {
			t.Error("Should return: could not find error")
		}
	}()
}

func TestNoOrigin(t *testing.T) {
	g := &Build{
		Index: Strings{
			"homepage-greeting": {
				"en":    "Welcome to %s#title, %s#name!",
				"es-LA": "¡Bienvenido a %s#title, %s#title!",
			},
		},
		Target: "es-LA",
	}
	func() {
		t.Log("Test key")
		s := "homepage-greeting"
		tr, err := g.Translate(s)
		if s != tr {
			t.Errorf("Returned string '%s' is not the same as input string '%s'", tr, s)
		}
		if err == nil {
			t.Error("Should return: target or origin is not set")
		}
	}()
	func() {
		t.Log("Test string")
		s := "Welcome to %s#title, %s#name!"
		tr, err := g.Translate(s)
		if s != tr {
			t.Errorf("Returned string '%s' is not the same as input string '%s'", tr, s)
		}
		if err == nil {
			t.Error("Should return: target or origin is not set")
		}
	}()
}

func TestNoTarget(t *testing.T) {
	g := &Build{
		Index: Strings{
			"homepage-greeting": {
				"en":    "Welcome to %s#title, %s#name!",
				"es-LA": "¡Bienvenido a %s#title, %s#title!",
			},
		},
		Origin: "en",
	}
	func() {
		t.Log("Test key")
		s := "homepage-greeting"
		tr, err := g.Translate(s)
		if s != tr {
			t.Errorf("Returned string '%s' is not the same as input string '%s'", tr, s)
		}
		if err == nil {
			t.Error("Should return: target or origin is not set")
		}
	}()
	func() {
		t.Log("Test string")
		s := "Welcome to %s#title, %s#name!"
		tr, err := g.Translate(s)
		if s != tr {
			t.Errorf("Returned string '%s' is not the same as input string '%s'", tr, s)
		}
		if err == nil {
			t.Error("Should return: target or origin is not set")
		}
	}()
}
