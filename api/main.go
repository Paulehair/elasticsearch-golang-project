package main

import (
	"fmt"

	"github.com/google/uuid"
)

func main() {
	uuid := uuid.NewString()
	fmt.Println(uuid)
}
