package streams

import (
	"reflect"
	"testing"
)

func TestNewEntry(t *testing.T) {
	e1 := NewEntry(10, 20)
	expect := &Entry[int, int]{K: 10, V: 20}
	if !reflect.DeepEqual(e1, expect) {
		t.Fatal("test failed!")
	}
}

func TestMapToEntry(t *testing.T) {
	m := map[int]string{1: "1"}
	expect := []*Entry[int, string]{{1, "1"}}
	if !reflect.DeepEqual(MapToEntry[int, string](m), expect) {
		t.Fatal("test failed!")
	}
}

func TestEntryToMap(t *testing.T) {
	es := []*Entry[int, string]{{1, "1"}}
	expect := map[int]string{1: "1"}
	if !reflect.DeepEqual(EntryToMap[int, string](es), expect) {
		t.Fatal("test failed!")
	}
}
