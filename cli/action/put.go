package action

import "fmt"

func PutKey(key, value string) error {
	fmt.Printf("Put %s:%s\n", key, value)
	return nil
}
