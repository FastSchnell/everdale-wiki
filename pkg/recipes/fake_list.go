package recipes


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
		bdsStr += fmt.Sprintf(BaseStr, bd.Name, bd.NationBuilding, bd.BuildingLevel, bd.DurationSeconds, bd.InputResource1, bd.InputResource2, bd.InputResource3, bd.InputResource4)
	}

	return bdsStr, nil
}
