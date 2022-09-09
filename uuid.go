package uuid

import (
	"fmt"
	"github.com/google/uuid"
)

func GenerateUUID() {
	id := uuid.New().String()
	fmt.Println("UUID: ", id)
}
