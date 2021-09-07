package page

import (
	"everdale-wiki/pkg/buildings"
	"fmt"
)

var BaseStr = `
    <tr class="table-active">
      <th scope="col">%s</th>
      <th scope="col">%d</th>
      <th scope="col">%s</th>
      <th scope="col">%s</th>
      <th scope="col">%s</th>
      <th scope="col">%s</th>
      <th scope="col">%s</th>
    </tr>
`

func FakeBds(name string, level int) (string, error) {
	bds, err := buildings.GetByLevel(name, level)
	if err != nil {
		return "", err
	}

	bdsStr := ""
	for _, bd := range bds {
		bdsStr += fmt.Sprintf(BaseStr, bd.Name, bd.Level, bd.Coins, bd.Wood, bd.Clay, bd.Stone, bd.Plank)
	}

    return bdsStr, nil
}
