// Copyright 2017 The Bazel Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package pkgscript_test

// This file defines tests of the Value API.

import (
	"fmt"
	"testing"

	"github.com/andrewchambers/pkgscript/pkgscript"
)

func TestStringMethod(t *testing.T) {
	s := pkgscript.String("hello")
	for i, test := range [][2]string{
		// quoted string:
		{s.String(), `"hello"`},
		{fmt.Sprintf("%s", s), `"hello"`},
		{fmt.Sprintf("%+s", s), `"hello"`},
		{fmt.Sprintf("%v", s), `"hello"`},
		{fmt.Sprintf("%+v", s), `"hello"`},
		// unquoted:
		{s.GoString(), `hello`},
		{fmt.Sprintf("%#v", s), `hello`},
	} {
		got, want := test[0], test[1]
		if got != want {
			t.Errorf("#%d: got <<%s>>, want <<%s>>", i, got, want)
		}
	}
}

func TestListAppend(t *testing.T) {
	l := pkgscript.NewList(nil)
	l.Append(pkgscript.String("hello"))
	res, ok := pkgscript.AsString(l.Index(0))
	if !ok {
		t.Errorf("failed list.Append() got: %s, want: pkgscript.String", l.Index(0).Type())
	}
	if res != "hello" {
		t.Errorf("failed list.Append() got: %+v, want: hello", res)
	}
}
