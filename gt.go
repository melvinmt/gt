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
	"errors"
	"fmt"
	"regexp"
)

type Strings map[string]map[string]string // map[key]map[language][translation]

// Set up Build environment first before starting translations.
type Build struct {
	Origin     string         // the origin env
	Target     string         // the target env
	Index      Strings        // the index which contains all keys and strings
	regexVerbs *regexp.Regexp // caching regex Compiles
	regexTags  *regexp.Regexp
}

// T() is a shorthand method for Translate. Ignores errors and strictly returns strings.
func (b *Build) T(s string, a ...interface{}) (t string) {
	t, _ = b.Translate(s, a...)
	return t
}

// Translate() translates a key or string from origin to target.
// Parses augmented sprintf format when additional arguments are given.
func (b *Build) Translate(str string, args ...interface{}) (t string, err error) {

	if b.Origin == "" || b.Target == "" {
		return str, errors.New("Origin or target is not set.")
	}

	// Try to find origin string by key or key[:2]
	var o string // origin string
	key := str   // key can differ from str

	if b.Index[key][b.Origin] != "" {
		o = b.Index[key][b.Origin]
	} else if b.Index[key][b.Origin[:2]] != "" {
		o = b.Index[key][b.Origin[:2]]
	} else {
		// If key is not found, try matching strings in origin.
		for k, v := range b.Index {
			if key == v[b.Origin] {
				o, key = key, k
				break
			}
		}
	}

	// Try to find target string by key or key[:2]
	if b.Index[key][b.Target] != "" {
		t = b.Index[key][b.Target]
	} else if b.Index[key][b.Target[:2]] != "" {
		t = b.Index[key][b.Target[:2]]
	}

	if o == "" || t == "" {
		return str, errors.New("Couldn't find origin or target string.")
	}

	// When no additional arguments are given, there's nothing left to do.
	if len(args) == 0 {
		return t, err
	}

	// Find verbs in both strings.
	if b.regexVerbs == nil {
		b.regexVerbs, _ = regexp.Compile(`%(?:\d+\$)?[+-]?(?:[ 0]|'.{1})?-?\d*(?:\.\d+)?#?[bcdeEfFgGopqstTuUvxX%]?[#[\w0-9-_]+]?`)
	}
	oVerbs := b.regexVerbs.FindAllStringSubmatch(o, -1)
	tVerbs := b.regexVerbs.FindAllStringSubmatch(t, -1)

	if len(oVerbs) != len(args) || len(tVerbs) != len(args) {
		return str, errors.New("Arguments count is different than verbs count.")
	}

	// Check if verbs are unique.
	uniqueVerbs := true
	q := make(map[string]int, len(tVerbs))
	for _, v := range tVerbs {
		if len(v) > 0 {
			q[v[0]] = q[v[0]] + 1
			if q[v[0]] > 1 {
				uniqueVerbs = false
				break
			}
		}
	}

	// Swap argument positions when swapping is necessary and verbs are unique.
	if uniqueVerbs == true {
		for i, v1 := range tVerbs {
			for j, v2 := range oVerbs {
				if len(args)-1 == i {
					break
				}
				if v1[0] == v2[0] && args[i] != args[j] {
					args[j], args[i] = args[i], args[j]
					break
				}
			}
		}
	} else if h1, h2 := fmt.Sprintf("%v", oVerbs), fmt.Sprintf("%v", tVerbs); h1 != h2 {
		return str, errors.New("Verbs have to be swapped but are not unique.")
	}

	// Clean up tags
	if b.regexTags == nil {
		b.regexTags, _ = regexp.Compile(`#[\w0-9-_]+`)
	}
	t = b.regexTags.ReplaceAllLiteralString(t, "")

	// Parse arguments into string.
	t = fmt.Sprintf(t, args...)
	return t, err
}
