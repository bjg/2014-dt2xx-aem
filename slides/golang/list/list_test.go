package list

import (
    "testing"
)

func TestNew(t *testing.T) {
      if New() == nil {
          t.Fatal("Expected a new list to be created")
      }
  }

