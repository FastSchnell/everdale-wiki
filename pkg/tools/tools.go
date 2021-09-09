package tools


import (
	"encoding/csv"
	"everdale-wiki/pkg/config"
	"fmt"
	"io"
	"os"
	"strings"
)

func GetDetail(name string) ([]*ToolsDetail, error) {
	csvFile, err := os.Open(fmt.Sprintf("%s/logic/tools.csv", config.CfgInstance.Everdale.Path))
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)
	cName := ""
	cLevel := 1

	bds := make([]*ToolsDetail, 0)

	for {
		row, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		if row[0] != "" {
			cName = row[0]
			cLevel = 1
		} else {
			cLevel += 1
		}
		//
		//fmt.Println(cName, name)

		//fmt.Println(cName, cLevel, row[7], row[8], row[9], row[10], row[11], row[12])
		if strings.HasSuffix(cName, name) {
			sd := new(ToolsDetail)
			sd.Name = cName
			sd.Level = cLevel
			sd.BuildingLevel = row[11]
			sd.BoostIncrease = row[12]
			sd.HungerPercentage = row[14]
			sd.SpeedIncrease = row[15]
			sd.DurabilitySec = row[17]
			sd.CreationPhaseSeconds = row[18]

			bds = append(bds, sd)

		}

	}

	return bds, nil
}

func GetByLevel(name string, level int) ([]*ToolsDetail, error) {
	nbds := make([]*ToolsDetail, 0)
	bds, err := GetDetail(name)
	if err != nil {
		return nil, err
	}

	if level == 0 {
		return bds, nil
	}


	for _, v := range bds {
		if v.Level == level {
			nbds = append(nbds, v)
		}
	}

	return nbds, nil
}



type ToolsDetail struct {
	Name string
	Level int
	BuildingLevel string
	BoostIncrease string
	HungerPercentage string
	SpeedIncrease string
	DurabilitySec string
	CreationPhaseSeconds string
}