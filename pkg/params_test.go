package pkg

import (
	"reflect"
	"testing"
)

func TestSingleLevel(t *testing.T) {
	var p Params

	if err := p.Set("foo="); err != nil {
		t.Fatal(err)
	}

	if err := p.Set("bar=4"); err != nil {
		t.Fatal(err)
	}

	if err := p.Set("foo=12"); err != nil {
		t.Fatal(err)
	}

	exp := map[string]interface{}{
		"foo": []string{"", "12"},
		"bar": "4",
	}
	if !reflect.DeepEqual(p.Values, exp) {
		t.Fatalf("expected %#v got %#v", exp, p.Values)
	}
}

func TestMultiLevel(t *testing.T) {
	var p Params

	if err := p.Set("a.b=100"); err != nil {
		t.Fatal(err)
	}

	if err := p.Set("a.c=200"); err != nil {
		t.Fatal(err)
	}

	if err := p.Set("a.a.a="); err != nil {
		t.Fatal(err)
	}

	if err := p.Set("a.d=hello"); err != nil {
		t.Fatal(err)
	}

	if err := p.Set("a.c="); err != nil {
		t.Fatal(err)
	}

	exp := map[string]interface{}{
		"a": map[string]interface{}{
			"b": "100",
			"c": []string{"200", ""},
			"a": map[string]interface{}{
				"a": "",
			},
			"d": "hello",
		},
	}
	if !reflect.DeepEqual(p.Values, exp) {
		t.Fatalf("expected %#v got %#v", exp, p.Values)
	}
}

func TestConflicts(t *testing.T) {
	var p Params

	// long key first
	if err := p.Set("a.b.c=22"); err != nil {
		t.Fatal(err)
	}

	if ok, keys := IsConflictErr(p.Set("a.b=33")); ok {
		exp := []string{"a.b", "a.b.c"}
		if !reflect.DeepEqual(keys, exp) {
			t.Fatalf("expected %#v got %#v", exp, keys)
		}
	} else {
		t.Fatal("expected conflict")
	}

	// short key first
	if err := p.Set("b.a=1"); err != nil {
		t.Fatal(err)
	}

	if ok, keys := IsConflictErr(p.Set("b.a.c=1")); ok {
		exp := []string{"b.a.c", "b.a"}
		if !reflect.DeepEqual(keys, exp) {
			t.Fatalf("expected %#v got %#v", exp, keys)
		}
	} else {
		t.Fatal("expected conflict")
	}
}
