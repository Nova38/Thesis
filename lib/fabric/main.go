package main

import (
	"fmt"
	"testing"

	"github.com/bufbuild/protovalidate-go"
	pb "github.com/nova38/thesis/lib/gen/go/rbac"
)

func TestV(t *testing.T) {
}

func main() {
	println(msg)

	if err = v.Validate(msg); err != nil {
		fmt.Println("validation failed:", err)
	} else {
		fmt.Println("validation succeeded")
	}
}
