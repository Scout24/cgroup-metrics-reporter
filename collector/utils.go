package collector

import "fmt"

func CreateTag(name, value string) string {
	return fmt.Sprintf("%s:%s", name, value)
}
