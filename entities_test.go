package database

import (
    "testing"
)

func TestTypeIDGeneration(t *testing.T) {
    type first struct {a, b int}
    type second struct {a, b int}

    w := NewWorld()

    one := w.AddType(first{})
    two := w.AddType(second{})

    if one == two {
        t.Fatalf("expected two different TypeIDs")
    }

    if one != 0 {
        t.Fatalf("expected TypeID to be %d", 0)
    }

    if two != 1 {
        t.Fatalf("expected TypeID to be %d", 1)
    }
}

