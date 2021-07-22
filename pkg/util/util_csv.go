package util

import (
	"encoding/csv"
	"log"
	"os"
	"os/exec"
)

func ReadCsv(filePath string) [][]string {
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	cells, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	return cells
}

func ConvertXlsFileToCsvFile(filePathXls, filePathCsv string) {
	_, err := exec.LookPath("ssconvert")
	if err != nil {
		log.Fatal("Command `ssconvert` not found. please install gnumeric")
	}
	err = exec.Command("ssconvert", filePathXls, filePathCsv).Run()
	if err != nil {
		log.Fatal(err)
	}
}
