package page

import (
	"everdale-wiki/pkg/logger"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var (
	HtmlStr = `
<html>
<head>
<title>
Everdale Wiki - Buildings | Postbook
</title>
    <link rel="stylesheet" href="/static/bootstrap.css" media="screen">

</head>

<body>

%s

<div class="form-group" align="center">
      <label for="Building Name" class="form-label mt-4">Building Name</label>
      <select class="form-select" id="BuildingName">

        <option value="Village Kitchen">Village Kitchen</option>
        <option value="WoodStorage">WoodStorage</option>
        <option value="House">House</option>
        <option value="PumpkinField">PumpkinField</option>
        <option value="StoneStorage">StoneStorage</option>
<option value="Study">Study</option>
<option value="Sawmill">Sawmill</option>
<option value="StoneMine">StoneMine</option>
<option value="ChickenCoop">ChickenCoop</option>
<option value="BrickYard">BrickYard</option>
<option value="ClayPit">ClayPit</option>
<option value="ClayStorage">ClayStorage</option>
<option value="Warehouse">Warehouse</option>
<option value="Treasury">Treasury</option>
<option value="Fountain">Fountain</option>
<option value="VillageTrader">VillageTrader</option>
<option value="FigurineWorkshop">FigurineWorkshop</option>
<option value="PotteryWorkshop">PotteryWorkshop</option>
<option value="StatueWorkshop">StatueWorkshop</option>
<option value="JuicePress">JuicePress</option>
<option value="Windmill">Windmill</option>
<option value="Beehive">Beehive</option>
<option value="AppleTreeGrove">AppleTreeGrove</option>
<option value="VillageTraderCart">VillageTraderCart</option>


      </select>
    </div>

<div class="form-group" align="center">
      <label for="Building Level" class="form-label mt-4">Building Level</label>
      <select class="form-select" id="BuildingLevel">
        <option>ALL</option>
        <option>1</option>
        <option>2</option>
        <option>3</option>
        <option>4</option>
        <option>5</option>
        <option>6</option>
        <option>7</option>
        <option>8</option>
        <option>9</option>
        <option>10</option>
        <option>11</option>
        <option>12</option>
        <option>13</option>
        <option>14</option>
        <option>15</option>
        <option>16</option>
      </select>
    </div>

</br></br>

<div align="center">
<button type="button" class="btn btn-primary btn-lg" onclick="gotoBuildings()">Search</button>
</div>
</br></br>

<table class="table table-hover">
  <thead>
    <tr>
      <th scope="col">Name</th>
      <th scope="col">Level</th>
      <th scope="col">Coins</th>
      <th scope="col">Wood</th>
      <th scope="col">Clay</th>
      <th scope="col">Stone</th>
      <th scope="col">Plank</th>

    </tr>
  </thead>
  <tbody>
    %s

  </tbody>
</table>
</br></br>


       <footer id="footer">
        <div class="row" align="center">
          <div class="col-lg-12">
            <ul class="list-unstyled">
              <li><a href="https://everdale.postbook.xyz/read/dad2a14c-7c7c-4937-800d-5fe6b115c44b">Feedback</a></li>
              <li><a href="https://github.com/FastSchnell/everdale-wiki">GitHub</a></li>

            </ul>
            <p>Made by <a>Frank & Luis Tam</a>.</p>
            <p>Code released under the <a href="https://github.com/FastSchnell/everdale-wiki/blob/master/LICENSE">MIT License</a>.</p>

          </div>
        </div>
      </footer>

<script type="text/javascript">
  window.onload = function se() {
      let nameIdx = getQueryVariable("name_idx")
      let levelIdx = getQueryVariable("level_idx")

      if (nameIdx != '') {

          document.getElementById("BuildingName")[nameIdx].selected=true;
      }
      if (levelIdx != '') {
          document.getElementById("BuildingLevel")[levelIdx].selected=true;
      }

  }

</script>


<script>
function gotoBuildings() {
  var nameS = document.getElementById("BuildingName");
  var nameIdx = nameS.selectedIndex;
  var name = nameS.value;
  var levelS = document.getElementById("BuildingLevel");
  var levelIdx = levelS.selectedIndex;
  var level = levelS.value;
  window.location.href="/buildings?name=" + name + "&level=" + level + "&name_idx=" + nameIdx + "&level_idx=" + levelIdx ;
}

function getQueryVariable(variable)
{
       var query = window.location.search.substring(1);
       var vars = query.split("&");
       for (var i=0;i<vars.length;i++) {
               var pair = vars[i].split("=");
               if(pair[0] == variable){return pair[1];}
       }
       return(false);
}

</script>


    <script src="/static/bootstrap.bundle.min.js"></script>


</body>
</html>



`
)


func fakeName(name string) string {
	switch name {
	case "Village Kitchen":
		name = "Town Hall"
	case "Study":
		name = "Library"

	}
	return name
}

func outPutName(name string) string {
	switch name {
	case "Town Hall":
		name = "Village Kitchen"
	case "Library":
		name = "Study"

	}
	return name
}



func Page(ctx *gin.Context) {
	name, _ := ctx.GetQuery("name")
	name = fakeName(name)
	level, _ := ctx.GetQuery("level")

	if name == "" {
		name = "Town Hall"
		level = "0"
	}

	if level == "ALL" {
		level = "0"
	}

	levelInt, _ := strconv.Atoi(level)
	tableStr, err := FakeBds(name, levelInt)
	if err != nil {
		logger.Error.Json(map[string]interface{}{
			"flag": "FakeBds error",
			"err": err.Error(),
		})
	}


	ctx.Data(http.StatusOK, "text/html", []byte(fmt.Sprintf(HtmlStr, NavStr, tableStr)))
}
