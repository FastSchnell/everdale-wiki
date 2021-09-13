package recipes



import (
	"encoding/csv"
	"everdale-wiki/pkg/config"
	"fmt"
	"io"
	"os"
	"strings"
)

func GetDetail(name string) ([]*RecipesDetail, error) {
	csvFile, err := os.Open(fmt.Sprintf("%s/logic/resource_recipes.csv", config.CfgInstance.Everdale.Path))
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)
	cName := ""


	bds := make([]*RecipesDetail, 0)

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

		if cName != "" && strings.HasPrefix(cName, "Recipe") && (row[3] == name || name == "ALL") {
			//ocn := row[16]
			//if strings.HasPrefix(ocn, "Challenge") {
			//	ocn = strings.Split(row[16], "_")[1]
			//}
			sd := new(RecipesDetail)
			sd.Name = row[3]
			sd.NationBuilding = row[1]
			sd.BuildingLevel = row[2]
			sd.DurationSeconds = row[9]
			sd.InputResource1 = fakeInputResource(row[10], row[11])
			sd.InputAmount1 = row[11]
			sd.InputResource2 = fakeInputResource(row[12], row[13])
			sd.InputAmount2 = row[13]
			sd.InputResource3 = fakeInputResource(row[14], row[15])
			sd.InputAmount3 = row[15]
			sd.InputResource4 = fakeInputResource(row[16], row[17])
			sd.InputAmount4 = row[17]

			bds = append(bds, sd)
		}

	}

	return bds, nil
}


func fakeInputResource(a, b string) string {
	if a == "" {
		return ""
	}

	return fmt.Sprintf("%s x %s", a, b)
}


type RecipesDetail struct {
	Name string
	NationBuilding string
	BuildingLevel string
	DurationSeconds string
	InputResource1 string
	InputAmount1 string
	InputResource2 string
	InputAmount2 string
	InputResource3 string
	InputAmount3 string
	InputResource4 string
	InputAmount4 string

}

