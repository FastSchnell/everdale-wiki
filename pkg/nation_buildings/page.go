package nation_buildings

import (
	"everdale-wiki/pkg/logger"
	"everdale-wiki/pkg/page"
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
Everdale Wiki - Nation Buildings | Postbook
</title>
    <link rel="stylesheet" href="/static/bootstrap.css" media="screen">

</head>

<body>
%s

<div class="form-group" align="center">
      <label for="Building Name" class="form-label mt-4">Nation Building Name</label>
      <select class="form-select" id="BuildingName">

        <option>Castle</option>
        <option>Research Guild</option>
        <option>ConstructionAcademy</option>
        <option>LumberjackAcademy</option>
        <option>WheatField</option>
        <option>ClaydiggerAcademy</option>
        <option>StonemasonAcademy</option>
        <option>FarmerAcademy</option>
        <option>NationBakery</option>
        <option>NationTailor</option>
        <option>NationPotionMaker</option>
        <option>NationCannery</option>
        <option>NationDyeShop</option>
        <option>NationToyWorkshop</option>
        <option>NationFishingSpot</option>
        <option>NationCottonField</option>
        <option>NationIndigoField</option>
        <option>NationSugarField</option>
        <option>NationSaltField</option>
        <option>NationMonument01</option>


      </select>
    </div>

<div class="form-group" align="center">
      <label for="Building Level" class="form-label mt-4">Nation Building Level</label>
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
            <th scope="col">Brick</th>

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
  window.location.href="/nation_buildings?name=" + name + "&level=" + level + "&name_idx=" + nameIdx + "&level_idx=" + levelIdx ;
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
	case "Research Guild":
		name = "ResearchAcademy"

	}
	return name
}

func outPutName(name string) string {
	switch name {
	case "ResearchAcademy":
		name = "Research Guild"

	}
	return name
}


func NationBuildingsPage(ctx *gin.Context) {
	name, _ := ctx.GetQuery("name")
	name = fakeName(name)
	level, _ := ctx.GetQuery("level")



	if name == "" {
		name = "Castle"
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


	ctx.Data(http.StatusOK, "text/html", []byte(fmt.Sprintf(HtmlStr, page.NavStr, tableStr)))
}
