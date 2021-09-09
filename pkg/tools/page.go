package tools


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
Everdale Wiki - Tools | Postbook
</title>
    <link rel="stylesheet" href="/static/bootstrap.css" media="screen">

</head>

<body>
%s

<div class="form-group" align="center">
      <label for="Building Name" class="form-label mt-4">Specialties Name</label>
      <select class="form-select" id="BuildingName">

        <option>SpeedLumberjack</option>
        <option>SpeedFarmer</option>
        <option>SpeedResearcher</option>
        <option>SpeedStonemason</option>
        <option>SpeedSawyer</option>
        <option>SpeedClaydigger</option>
        <option>SpeedBrickmaker</option>
        <option>NoFood</option>




      </select>
    </div>

<div class="form-group" align="center">
      <label for="Building Level" class="form-label mt-4">Specialties Level</label>
      <select class="form-select" id="BuildingLevel">
        <option>ALL</option>
        <option>1</option>
        <option>2</option>
        <option>3</option>

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
      <th scope="col">BuildingLevel</th>
      <th scope="col">BoostIncrease</th>
      <th scope="col">HungerPercentage</th>
      <th scope="col">SpeedIncrease</th>
      <th scope="col">DurabilitySec</th>
      <th scope="col">CreationPhaseSeconds</th>


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
  window.location.href="/tools?name=" + name + "&level=" + level + "&name_idx=" + nameIdx + "&level_idx=" + levelIdx ;
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

func ToolsPage(ctx *gin.Context) {
	name, _ := ctx.GetQuery("name")
	level, _ := ctx.GetQuery("level")

	if name == "" {
		name = "SpeedLumberjack"
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
