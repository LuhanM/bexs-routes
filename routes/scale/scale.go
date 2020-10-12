package scale

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"
)

type TypeScale struct {
	Origin      string `json:"origin"`
	Destination string `json:"destination"`
	Cost        int    `json:"cost"`
}

//Scales cache com a lista de escalas
var Scales []TypeScale

var pathCSV string

//AddScale Adiciona uma escala na lista de escalas disponíveis para consulta
func AddScale(scale TypeScale) error {
	Scales = append(Scales, scale)

	if err := saveInFile(scale); err != nil {
		return err
	}

	return nil
}

func saveInFile(scale TypeScale) error {
	var err error
	csvFile, err := os.OpenFile(pathCSV, os.O_RDWR|os.O_APPEND|os.O_CREATE, os.ModeAppend)
	if err != nil {
		return err
	}
	defer csvFile.Close()
	csvWriter := csv.NewWriter(csvFile)
	record := []string{scale.Origin, scale.Destination, strconv.Itoa(scale.Cost)}

	err = csvWriter.Write(record)
	if err != nil {
		return err
	}
	csvWriter.Flush()

	return csvWriter.Error()
}

//GetScale Retorna a escata se existir na lista de escalas. Caso não exista será retornado erro
func GetScale(origin, destination string) (TypeScale, error) {
	for _, scale := range Scales {
		if scale.Origin == origin && scale.Destination == destination {
			return scale, nil
		}
	}
	return TypeScale{}, fmt.Errorf("scale not found")
}

//LoadScalesFile Le o arquivo de input e adiciona na lista de escalas na memória
func LoadScalesFile(path string) {
	pathCSV = path
	csvFile, err := os.OpenFile(path, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	csvReader := csv.NewReader(csvFile)
	defer csvFile.Close()

	for {
		record, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}

		var scale TypeScale
		scale.Origin = record[0]
		scale.Destination = record[1]
		scale.Cost, _ = strconv.Atoi(record[2])
		Scales = append(Scales, scale)
	}

}
