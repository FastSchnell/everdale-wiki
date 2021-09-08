package specialties

import (
	"encoding/csv"
	"everdale-wiki/pkg/config"
	"fmt"
	"io"
	"os"
)

func GetDetail(name string) ([]*SpecialtiesDetail, error) {
	csvFile, err := os.Open(fmt.Sprintf("%s/logic/specialties.csv", config.CfgInstance.Everdale.Path))
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)
	cName := ""
	cLevel := 1

	bds := make([]*SpecialtiesDetail, 0)

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


		//fmt.Println(cName, cLevel, row[7], row[8], row[9], row[10], row[11], row[12])
		if cName == name {
			sd := new(SpecialtiesDetail)
			sd.Name = cName
			sd.Level = cLevel
			sd.BoostIncrease = row[5]
			sd.StudySeconds = row[6]
			sd.GoldCost = row[7]
			sd.XPNeededMinutes = row[8]
			bds = append(bds, sd)
		}

	}

	return bds, nil
}

func GetByLevel(name string, level int) ([]*SpecialtiesDetail, error) {
	nbds := make([]*SpecialtiesDetail, 0)
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



type SpecialtiesDetail struct {
	Name string
	Level int
	BoostIncrease string
	StudySeconds string
	GoldCost string
	XPNeededMinutes string

}

