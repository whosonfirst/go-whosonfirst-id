package id

import (
	"context"
	"fmt"
	"testing"
)

func TestNextInt(t *testing.T) {

	ctx := context.Background()
	cl, err := NewIdClient(ctx)

	if err != nil {
		t.Fatal(err)
	}

	id, err := cl.NextInt()

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(id)
}

func TestNextIntWithURI(t *testing.T) {

	uri := "artisanal:///?client=brooklynintegers%3A%2F%2F&minimum=5&pool=memory%3A%2F%2F"
	
	ctx := context.Background()
	cl, err := NewIdClientWithURI(ctx, uri)

	if err != nil {
		t.Fatal(err)
	}

	id, err := cl.NextInt()

	if err != nil {
		t.Fatal(err)
	}

	fmt.Println(id)
}
