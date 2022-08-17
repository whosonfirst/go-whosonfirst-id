package id

import (
	"testing"
)

func TestContants(t *testing.T) {

	tests := map[int64]int64{
		-4: MULTIPLE_PARENTS,
		-2: ITS_COMPLICATED,
		-1: UNKNOWN,
		0:  EARTH,
		1:  NULL_ISLAND,
	}

	for id, expected := range tests {

		if id != expected {
			t.Fatalf("Unexpected value for %d: %d", id, expected)
		}
	}
}
