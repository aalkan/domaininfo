package main

import (
	"fmt"
	"io/ioutil"
	"strings"
	"time"

	"github.com/aalkan/domaininfo/pkg"
	"github.com/manifoldco/promptui"
	whois "github.com/undiabler/golang-whois"
	"github.com/xuri/excelize/v2"
)

func main() {
	
	defer func() {
		time.Sleep(time.Second * 1)
		pkg.White("Program sonlandırılıyor...")
		time.Sleep(time.Second * 1)
		pkg.White("3...")
		time.Sleep(time.Second * 1)
		pkg.White("2...")
		time.Sleep(time.Second * 1)
		pkg.White("1...")
	}()
	pkg.Primary("Domain listesi dosya dizini: " + pkg.GetWr()+"/domains.txt")
	time.Sleep(time.Second * 1)
	fmt.Println("")
	pkg.White("Yapmak istediğiniz işlemi seçiniz:")
	prompt := promptui.Select{
		Label: "Select Day",
		Items: []string{"Çalıştır", "Temizle"},
	}

	_, result, err := prompt.Run()

	if err != nil {
		pkg.Danger(fmt.Sprintf("Prompt failed %v\n", err))
		return
	}
	switch result {
	case "Çalıştır":
		workDirectory := pkg.GetWr()
		filename := workDirectory + "/domains.txt"
		if err := pkg.CheckFileExists(filename); err != nil {
			pkg.Danger(err.Error())
			return
		}

		pkg.Success("Excel dosyası okunuyor...")
		time.Sleep(time.Second * 1)

		fileByte, err := ioutil.ReadFile(filename)
		if err != nil {
			pkg.Danger("domains.txt okunurken hata meydana geldi..")
			return
		}
		domains := strings.Split(string(fileByte), "\n")

		if len(domains) == 0 {
			pkg.Danger("Yeterli veri okunamadı...")
			return
		}

		f := excelize.NewFile()
		// Create a new sheet.
		style, _ := f.NewStyle(`{"alignment":{"vertical":"top"}}`)
		styleLast, _ := f.NewStyle(`{"fill":{"type":"pattern","color":["#9ACD32"],"pattern":4}}`)
		bold, _ := f.NewStyle(`{"font":{"bold":true}}`)

		index := f.NewSheet("Sheet1")
		// Set value of a cell.

		f.SetColWidth("Sheet1", "A", "H", 20)
		f.SetColStyle("Sheet1", "A:H", style)
		f.SetColStyle("Sheet1", "E", styleLast)
		f.SetColStyle("Sheet1", "A", bold)

		f.SetCellValue("Sheet1", "A1", "Domain")
		f.SetCellValue("Sheet1", "B1", "Kayıtlı Kuruluş")
		f.SetCellValue("Sheet1", "C1", "Kayıt Tarihi")
		f.SetCellValue("Sheet1", "D1", "Son Güncelleme Tarihi")
		f.SetCellValue("Sheet1", "E1", "Sonlanacağı Tarih")
		f.SetCellValue("Sheet1", "F1", "NS1")
		f.SetCellValue("Sheet1", "G1", "NS2")
		f.SetCellValue("Sheet1", "H1", "Ülke")

		// Set active sheet of the workbook.
		f.SetActiveSheet(index)
		// Save spreadsheet by the given path.
		pkg.Success("Dosyaya yazılıyor..")
		time.Sleep(time.Second * 2)

		for i, domain := range domains {
			result, _ := whois.GetWhois(domain)
			lines := strings.Split(result, "\n")
			f.SetCellValue("Sheet1", fmt.Sprintf("A%d", i+2), domain)
			ns1Flasg := false
			for _, line := range lines {
				if strings.Contains(line, "Registrar WHOIS Server") {
					value := strings.Split(line, ":")[1]
					f.SetCellValue("Sheet1", fmt.Sprintf("B%d", i+2), value)
				} else if strings.Contains(line, "Creation Date") {
					value := strings.Split(line, ":")[1]
					f.SetCellValue("Sheet1", fmt.Sprintf("C%d", i+2), value)
				} else if strings.Contains(line, "Updated Date") {
					value := strings.Split(line, ":")[1]
					f.SetCellValue("Sheet1", fmt.Sprintf("D%d", i+2), value)
				} else if strings.Contains(line, "Registry Expiry Date") {
					value := strings.Split(line, ":")[1]
					f.SetCellValue("Sheet1", fmt.Sprintf("E%d", i+2), value)
				} else if strings.Contains(line, "Name Server") && !ns1Flasg {
					value := strings.Split(line, ":")[1]
					f.SetCellValue("Sheet1", fmt.Sprintf("F%d", i+2), value)
					ns1Flasg = true
				} else if strings.Contains(line, "Name Server") && ns1Flasg {
					value := strings.Split(line, ":")[1]
					f.SetCellValue("Sheet1", fmt.Sprintf("G%d", i+2), value)
					ns1Flasg = true
				} else if strings.Contains(line, "Registrant Country") {
					value := strings.Split(line, ":")[1]
					f.SetCellValue("Sheet1", fmt.Sprintf("H%d", i+2), value)
				}
			}
		}

		randomName := pkg.GetWr() + "/" + "results.xlsx"
		if err := f.SaveAs(randomName); err != nil {
			pkg.Danger(err.Error() + " Dosya kayıt edilemedi")
		}
		
		
		pkg.Success("İşlem başarılı...")
	

	}
}
