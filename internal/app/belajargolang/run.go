package belajargolang

import (
	"fmt"

	"belajar-golang/internal/pkg/appinfo"
	"belajar-golang/pkg/greeting"
)

// Run menjalankan aplikasi utama.
func Run() error {
	fmt.Printf("%s\n", appinfo.Banner())
	fmt.Println(greeting.Build("Gopher"))
	return nil
}
