package main

import (
	"fmt"

	"github.com/manifoldco/promptui"
)

func main() {
	defer func() {
		fmt.Println("Program sonlandırılıyor...")

	}()
	prompt := promptui.Select{
		Label: "Select Day",
		Items: []string{"Çalıştır", "Temizle"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return
	}
	switch result {
	case "Çalıştır":
	}
}
