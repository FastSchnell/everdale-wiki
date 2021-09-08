package challenges


import (
	"fmt"
)

var BaseStr = `
    <tr class="table-active">
      <th scope="col">%s</th>
      <th scope="col">%s</th>
      <th scope="col">%s</th>
      <th scope="col">%s</th>
      <th scope="col">%s</th>
      <th scope="col">%s</th>
      <th scope="col">%s</th>
    </tr>
`

func FakeBds(name string) (string, error) {
	bds, err := GetDetail(name)
	if err != nil {
		return "", err
	}

	bdsStr := ""
	for _, bd := range bds {
		bdsStr += fmt.Sprintf(BaseStr, bd.Name, bd.Npc, bd.MinNationLevel, bd.BuildingNeeded, bd.BuildingLevelNeeded, bd.OtherChallengeNeeded, bd.RewardPerk)
	}

	return bdsStr, nil
}



