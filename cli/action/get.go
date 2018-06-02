package action

import "fmt"

func GetKey(key string) error {
	if key == "" {
		return fmt.Errorf("must specify key.\n")
	}

	fmt.Printf("Get Key %s\n", key)
	return nil
}
