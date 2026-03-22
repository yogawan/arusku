package greeting

import "fmt"

// Build membuat sapaan sederhana untuk package publik.
func Build(name string) string {
	return fmt.Sprintf("Hello, %s!", name)
}
