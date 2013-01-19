package gt

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
)

// var g = &gt.Build{
//  Index: gt.Keys{
//      "homepage-greeting": gt.Strings{
//          "en": "Welcome to %s#title, %s#name!"
//          "es-LA": "¡Bienvenido a %s#title, %s#titlename!",
//          "nl": "Welkom bij %s#title, %s#name!",
//          "tr": "%s#name, %s#title'ya hoşgeldiniz!", // tag notation proves useful in SOV languages where subject comes before object/verb
//          "zh-CN": "欢迎%s#title, %s#name!", 
//      },
//  }, 
//  Origin: "en",
// }
// t.Target = "es"
// 
// keyStr := g.T("homepage-greeting")
// 
// literalStr := g.T("Welcome to %s#title, %s#name!")
// 
type Strings map[string]string
type Keys map[string]Strings

type Build struct {
	Origin string // the originating env
	Target string // the target env
	Index  Keys   // the index which contains all keys and strings
}

// Translate finds a string for key in target env and transliterates it.
// Will return key if string or target env is not found.
// 
func (b *Build) Translate(key string, a ...interface{}) (t string, err error) {

	if len(b.Index) == 0 {
		return key, errors.New("Index needs to be set first")
	}

	var o string // origin string

	// try to find origin string by key
	if b.Index[key][b.Origin] != "" {
		o = b.Index[key][b.Origin]
	} else if b.Index[key][b.Origin[:2]] != "" {
		o = b.Index[key][b.Origin[:2]]
	}

	if o == "" { // no key found? try matching strings
		for k, v := range b.Index {
			if v[b.Origin] == key {
				o = key // key is origin string!
				key = k // found key in k
				break
			}
		}
	}

	// try to find target string
	if b.Index[key][b.Target] != "" {
		t = b.Index[key][b.Target]
	} else if b.Index[key][b.Target[:2]] != "" {
		t = b.Index[key][b.Target[:2]]
	}

	if o == "" || t == "" || b.Target == b.Origin {
		return t, errors.New("Couldn't find key/string or target is the same as origin")
	}

	if len(a) == 0 { // no arguments? return string
		return t, err
	}

	oVerbs, err := ParseStr(o)

	if err != nil {
		return t, errors.New("Couldn't parse origin string")
	}

	tVerbs, err := ParseStr(t)

	if err != nil {
		return t, errors.New("Couldn't parse target string")
	}

	var cleanVerb string

	// time to switch it up!
	if len(oVerbs) == len(tVerbs) { // check if both verbs arrays are the same length
		var r *regexp.Regexp
		r, err = regexp.Compile(`(#[\w0-9-_]+)`) // compile regex to match tags

		if err != nil {
			return t, errors.New("Couldn't compile regex to match tags")
		}

		newArgs := make([]interface{}, len(oVerbs)) // create new args slice

		for ti, dirtyVerb := range tVerbs {
			for oi, ov := range oVerbs {
				if ov == dirtyVerb && a[oi] != nil { // find original argument 
					newArgs[ti] = a[oi]                                  // assign argument to current position
					a[oi] = nil                                          // unset arg
					cleanVerb = r.ReplaceAllLiteralString(dirtyVerb, "") // remove tags
					strings.Replace(t, dirtyVerb, cleanVerb, -1)         // replace dirty verbs with clean verbs
				}
			}
		}

		if len(newArgs) == len(a) {
			t = fmt.Sprintf(t, newArgs...) // replace verbs with args
		} else {
			return t, errors.New("Impossible to assign arguments to string")
		}

	} else {
		return t, errors.New("Number of printf verbs in origin and target string do not match")
	}

	return t, err
}

// T is a shorthand function for Translate, ignores errors and strictly returns strings
func (b *Build) T(key string, a ...interface{}) (t string) {
	t, _ = b.Translate(key, a...)
	return t
}

// ParseStr returns an array of parsed verbs with optional tags
func ParseStr(str string) (verbs []string, err error) {
	r, err := regexp.Compile(`(%(?:\d+\$)?[+-]?(?:[ 0]|'.{1})?-?\d*(?:\.\d+)?#?[bcdeEfFgGopqstTuUvxX%]?)(#[\w0-9-_]+)?`)

	if err != nil {
		return verbs, errors.New("Couldn't compile regex to parse strings")
	}

	m := r.FindAllStringSubmatch(str, -1)

	if len(m) > 0 {
		verbs = make([]string, len(m))
		for i, v := range m[0] {
			verbs[i] = v
		}
	} else {
		return verbs, errors.New("No matches found")
	}

	return verbs, err
}

// t.Target = "es"

// str := g.T("homepage-greeting")

// t.Target = "nl"

// str = g.T("homepage-greeting")
// Example:
//     fmt.Printf(t.S("Welcome to %s#title, %s!#name!"), "gotr", "Melvin")
//     // outputs: "¡Bienvenido a gotr, Melvin!"
//     
// You can add meta data to the printf verbs by using the tag notation. Not only will translators
// understand more about the context that they're translating in, but it also prevents order issues 
// in certain languages. (The returning string is always correctly parsed).
// 
// Tags are recommended but you can also use regular printf strings. But then the
// order can be wrong in certain SOV languages (Turkish, Japanese):
//
//     fmt.Printf(t.S("Welcome to %s, %s"), "gotr", "Melvin")
//     // outputs: "gotr, Melvin'ya hoşgeldiniz!" in Turkish. 
//     // This roughly translates to: "Welcome to Melvin, gotr!" which is not what you want!
//     
// Tags are normally ended by a single space but you can also end tags with a plus sign:
// "somethingbefore%s#name+somethingafter" 
// 
