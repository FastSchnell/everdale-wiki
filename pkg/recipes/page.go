package recipes



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
Everdale Wiki - Recipes | Postbook
</title>
    <link rel="stylesheet" href="/static/bootstrap.css" media="screen">

</head>

<body>
%s

<div class="form-group" align="center">
      <label for="Building Name" class="form-label mt-4">Recipes Name</label>
      <select class="form-select" id="BuildingName">
  <option>ALL</option>
        <option>Bread</option>
        <option>BerryPie</option>
        <option>BerryCake</option>
        <option>SugarApple</option>
        <option>HoneyCookie</option>
        <option>FishPie</option>
        <option>ApplePie</option>
        <option>CakeRoyale</option>
        <option>SaltedFish</option>
        <option>SaltPot</option>
        <option>HoneyPot</option>
        <option>AppleJam</option>
        <option>SweetHerring</option>
        <option>DyeRed</option>
        <option>DyeYellow</option>
        <option>DyeBlue</option>
        <option>DyeGreen</option>
        <option>DyeOrange</option>
        <option>DyePurple</option>
        <option>DyeRainbow</option>
        <option>Plushie</option>
        <option>RuneGame</option>
        <option>WoodenDuck</option>
        <option>Marbles</option>
        <option>WoodenDoll</option>
        <option>Socks</option>
        <option>Trousers</option>
        <option>Hat</option>
        <option>Tunic</option>
        <option>Shirt</option>
        <option>Dress</option>
        <option>Jacket</option>
        <option>RainbowSuit</option>

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
      <th scope="col">NationBuilding</th>
      <th scope="col">BuildingLevel</th>
      <th scope="col">DurationSeconds</th>
      <th scope="col">InputResource1</th>
      <th scope="col">InputResource2</th>
      <th scope="col">InputResource3</th>
      <th scope="col">InputResource4</th>


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

  window.location.href="/recipes?name=" + name +  "&name_idx=" + nameIdx ;
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

func RecipesPage(ctx *gin.Context) {
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




