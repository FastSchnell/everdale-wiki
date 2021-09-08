package nation_buildings


import (
"encoding/csv"
"everdale-wiki/pkg/config"
"everdale-wiki/pkg/gin_util"
"fmt"
"github.com/gin-gonic/gin"
"github.com/gin-gonic/gin/binding"
"io"
"os"
)

func GetNationBuildings(ctx *gin.Context) {
	req := new(GetNationBuildingsRequest)
	err := ctx.ShouldBindWith(req, binding.JSON)
	if err != nil {
		data := map[string]interface{}{
			"msg": "InvalidParameter",
		}
		gin_util.RespStatusOK(ctx, &data)
		return
	}

	bds, err := GetByLevel(req.Name, req.Level)
	if err != nil {
		data := map[string]interface{}{
			"msg": "UnknownError",
		}
		gin_util.RespStatusOK(ctx, &data)
		return
	}

	gin_util.RespStatusOK(ctx, bds)

}


func GetDetail(name string) ([]*BuildDetail, error) {
	csvFile, err := os.Open(fmt.Sprintf("%s/logic/nation_buildings.csv", config.CfgInstance.Everdale.Path))
	if err != nil {
		return nil, err
	}
	defer csvFile.Close()
	csvReader := csv.NewReader(csvFile)
	cName := ""
	cLevel := 1

	bds := make([]*BuildDetail, 0)

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
			bd := new(BuildDetail)
			bd.Name = cName
			bd.Level = cLevel
			setVal(row[21], row[22], bd)
			setVal(row[23], row[24], bd)
			setVal(row[25], row[26], bd)
			setVal(row[27], row[28], bd)
			setVal(row[29], row[30], bd)
			setVal(row[31], row[32], bd)
			bds = append(bds, bd)
		}

	}

	return bds, nil
}

func GetByLevel(name string, level int) ([]*BuildDetail, error) {
	nbds := make([]*BuildDetail, 0)
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

func setVal(res, cost string, bd *BuildDetail) {
	switch res {
	case "Coins":
		bd.Coins = cost
	case "Wood":
		bd.Wood = cost
	case "Clay":
		bd.Clay = cost
	case "Plank":
		bd.Plank = cost
	case "Stone":
		bd.Stone = cost
	case "Brick":
		bd.Brick = cost

	}
}


type BuildDetail struct {
	Name string
	Level int
	Coins string
	Wood string
	Clay string
	Plank string
	Stone string
	Brick string
}