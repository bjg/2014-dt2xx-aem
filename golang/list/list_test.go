package list

import (
    "testing"
)

func TestNew(t *testing.T) {
    if New().head.next != nil {
        t.Fatal("Expected a new empty list to be created")
    }
}

