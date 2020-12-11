package main

import (
	"fmt"
	"log"

	"github.com/casbin/casbin/v2"
)

func main() {
	e, err := casbin.NewEnforcer("model.conf", "policy.csv")
	if err != nil {
		log.Fatalf("unable to create Casbin enforcer: %v", err)
	}

	result, err := e.Enforce("alice", "data1", "read")
	if err != nil {
		log.Fatalf("Enforce error: %v", err)
	}

	fmt.Printf("alice, data1, read: %v\n", result)

	result, err = e.Enforce("alice", "data1", "write")
	if err != nil {
		log.Fatalf("Enforce error: %v", err)
	}

	fmt.Printf("alice, data1, write: %v\n", result)

	result, err = e.Enforce("alice", "data2", "read")
	if err != nil {
		log.Fatalf("Enforce error: %v", err)
	}

	fmt.Printf("alice, data2, read: %v\n", result)

	result, err = e.Enforce("alice", "data2", "write")
	if err != nil {
		log.Fatalf("Enforce error: %v", err)
	}

	fmt.Printf("alice, data2, write: %v\n", result)

}
