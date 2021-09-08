package challenges

import (
	"everdale-wiki/pkg/logger"
	"everdale-wiki/pkg/page"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var (
	HtmlStr = `
<html>
<head>
<title>
Everdale Wiki - Challenges | Postbook
</title>
    <link rel="stylesheet" href="/static/bootstrap.css" media="screen">

</head>

<body>
%s

<div class="form-group" align="center">
      <label for="Building Name" class="form-label mt-4">Challenges Name</label>
      <select class="form-select" id="BuildingName">
  <option>ALL</option>
        <option>ValleyFeast</option>
        <option>Shipwreck</option>
        <option>SheepEscape</option>
        <option>PlaturtlePancake</option>
        <option>RoadBuilding</option>
        <option>BrokenCart</option>
        <option>Ruins01</option>
        <option>Ruins02</option>
        <option>FishingFrenzy</option>
        <option>MechaScarecrow</option>
        <option>PlaturtlePants</option>
        <option>BiggestCake</option>
        <option>StuckinJam</option>
        <option>PsychicOctopus</option>
        <option>UnityQuilt</option>
        <option>SlingshotShenanigans</option>
        <option>IndigoCloudSpirit</option>
        <option>ExoticHoney</option>
        <option>FoodFight</option>
        <option>HotAirBalloon</option>
        <option>CloudFlowerJourney</option>
        <option>BayCleanup</option>
        <option>SpiritParade</option>
        <option>BreadPupLost</option>
        <option>SaltySealife</option>
        <option>PlaturtleCoaster</option>
        <option>SweetToothMole</option>
        <option>SaveEarpuffs</option>
        <option>SandcastleSymposium</option>
        <option>DyeshopExplosion</option>

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
      <th scope="col">Npc</th>
      <th scope="col">MinNationLevel</th>
      <th scope="col">BuildingNeeded</th>
      <th scope="col">BuildingLevelNeeded</th>
      <th scope="col">OtherChallengeNeeded</th>
      <th scope="col">RewardPerk</th>
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

      if (nameIdx != '') {

          document.getElementById("BuildingName")[nameIdx].selected=true;
      }

  }

</script>


<script>
function gotoBuildings() {
  var nameS = document.getElementById("BuildingName");
  var nameIdx = nameS.selectedIndex;
  var name = nameS.value;

  window.location.href="/challenges?name=" + name +  "&name_idx=" + nameIdx ;
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

func ChallengesPage(ctx *gin.Context) {
	name, _ := ctx.GetQuery("name")


	if name == "" {
		name = "ALL"

	}




	tableStr, err := FakeBds(name)
	if err != nil {
		logger.Error.Json(map[string]interface{}{
			"flag": "FakeBds error",
			"err": err.Error(),
		})
	}


	ctx.Data(http.StatusOK, "text/html", []byte(fmt.Sprintf(HtmlStr, page.NavStr, tableStr)))
}


