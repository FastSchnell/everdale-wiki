package challenges


import (
	"encoding/csv"
	"everdale-wiki/pkg/config"
	"fmt"
	"io"
	"os"
	"strings"
)

func GetDetail(name string) ([]*ChallengesDetail, error) {
	csvFile, err := os.Open(fmt.Sprintf("%s/logic/challenges.csv", config.CfgInstance.Everdale.Path))
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)
	cName := ""


	bds := make([]*ChallengesDetail, 0)

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

		} else {
            continue
		}


		//fmt.Println(cName, cLevel, row[7], row[8], row[9], row[10], row[11], row[12])

		if cName != "" && strings.HasPrefix(cName, "Challenge") && (strings.Split(cName, "_")[1] == name || name == "ALL") {
			ocn := row[16]
			if strings.HasPrefix(ocn, "Challenge") {
				ocn = strings.Split(row[16], "_")[1]
			}
			sd := new(ChallengesDetail)
			sd.Name = strings.Split(cName, "_")[1]
	        sd.Npc = row[2]
	        sd.MinNationLevel = row[11]
	        sd.BuildingNeeded = row[13]
	        sd.BuildingLevelNeeded = row[14]
	        sd.OtherChallengeNeeded = ocn
	        sd.RewardPerk = row[62]

			bds = append(bds, sd)
		}

	}

	return bds, nil
}



type ChallengesDetail struct {
	Name string
	Npc string
	MinNationLevel string
	BuildingNeeded string
	BuildingLevelNeeded string
	OtherChallengeNeeded string
	RewardPerk string
}
