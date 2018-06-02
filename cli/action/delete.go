package action

import "fmt"

func DeleteKey(key string) error {
	fmt.Printf("Delete Key %s\n", key)
	return nil
}
