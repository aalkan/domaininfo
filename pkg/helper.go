package pkg

import (
	"errors"
	"fmt"
	"os"
)

func GetWr() string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println("Çalışma dizini bulunumadı, izinleri kontrol ediniz.")
		return ""
	}
	fmt.Println(dir)
	return dir
}

func CheckFileExists(file string) error {
	_, err := os.Stat(file)
	if errors.Is(err, os.ErrNotExist) {
		fmt.Println("Çalışma dizininde domains.txt bulunamadı.")
		if _, err := os.OpenFile(file, os.O_CREATE, 0666); err == nil {
			fmt.Println("Dosya sizin için oluşturuldu ama program sonlandırıldı. Lütfen içeriğini doldurunuz v eprogramı tekrar çalıştırınız.")
			return errors.New("Dosya oluşturuldu.")
		}
		return err
	}
	return nil
}
