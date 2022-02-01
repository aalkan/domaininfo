package pkg

import (
	"errors"
	"fmt"
	"os"

	"github.com/fatih/color"
)

func GetWr() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Çalışma dizini bulunumadı, izinleri kontrol ediniz.")
		return ""
	}
	return dir
}

func CheckFileExists(file string) error {
	_, err := os.Stat(file)
	if errors.Is(err, os.ErrNotExist) {
		Danger("Çalışma dizininde domains.txt bulunamadı.")
		if _, err := os.OpenFile(file, os.O_CREATE, 0666); err == nil {
			Primary("Dosya sizin için oluşturuldu ama program sonlandırıldı. Lütfen içeriğini doldurunuz v eprogramı tekrar çalıştırınız.")
			return errors.New("Dosya oluşturuldu.")
		}
		return err
	}
	return nil
}

func Danger(message string) {
	danger := color.New(color.FgRed).PrintlnFunc()
	danger(message)
}	

func Warning(message string) {
	warning := color.New(color.FgYellow).PrintlnFunc()
	warning(message)
}

func Success(message string) {
	success := color.New(color.FgGreen).PrintlnFunc()
	success(message)
}

func White(message string) {
	white := color.New(color.FgWhite).PrintlnFunc()
	white(message)
}

func Primary(message string) {
	primary := color.New(color.FgBlue).PrintlnFunc()
	primary(message)
}
