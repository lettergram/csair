package main

import (
	"fmt"
	"os/exec"
	"graph"
	"data"
)

/**
 * The options panel, calls the function selected
 * @param database, data base which stores all flight information
 * @param i - the selected user input
 */
func options(m *data.Map, g *graph.Graph, i rune) bool  {
	fmt.Print("\nOptions: \n" +
			"a. Its code\n" +
			"b. Its name\n" +
			"c. All City Names\n" +
			"d. All Possible Routes\n" +
			"e. Statistics\n" +
			"f. Generate Map\n" +
			"g. Edit Map\n" +
			"h. Make Flight Plans\n" +
			"z. TO EXIT\n>> ")
		_, err := fmt.Scanf("%c ", &i)
		if err != nil { fmt.Scanf("%c\n", &i) }
	if(i == 'a'){
		fmt.Print("Please enter the city code: ")
		data.CityCode(g, lookup())
	}else if(i == 'b'){
		fmt.Print("Please enter the city name: ")
		data.CityName(g, lookup())
	}else if(i == 'c'){
		data.CityAll(g)
	}else if(i == 'd'){
		data.RoutesAll(m)
	}else if(i == 'e'){
		data.AirlineStatistics(m, g)
	}else if(i == 'f'){
		openBrowser(*m)
	}else if(i == 'g'){
		return EditMap(m, g, i)
	}else if(i == 'h'){
		data.MakeFlight(g)
	}else if(i == 'z'){
		return false;
	}
	return true;
}

/**
 * This is the menu used to display editing options
 */
func EditMap(m *data.Map, g *graph.Graph, i rune) bool {

	fmt.Print("\nEditing Options: \n" +
				"a. Add a city to the map\n" +
				"b. Edit an existing city\n" +
				"c. Remove a city\n" +
				"d. Remove a Route\n" +
				"e. Save freshly edited Graph\n" +
				"z. TO EXIT\n>> ")
	_, err := fmt.Scanf("%c ", &i)

	if err != nil { fmt.Scanf("%c\n", &i) }
	if(i == 'a'){
		data.CityAdd(m, g)
	}else if(i == 'b'){
		data.CityEdit(m, g)
	}else if(i == 'c'){
		data.CityDelete(m, g)
	}else if(i == 'd'){
		data.RemoveSpecifiedRoute(m, g)
	}else if(i == 'e'){
		data.SaveGraph(m, g)
	}else if(i == 'z'){
		return false
	}
	return true
}

/**
 * Scans for users input and returns users input
 */
func lookup() string {
	s := ""
	_, err := fmt.Scanf("%s\n", &s)
	if err != nil {
		fmt.Println(err)
	}
	return s
}

/**
 * @param database which scores all of the flight information
 * Opens a web browser and displays all of the routes
 */
func openBrowser(database data.Map) {
	str := "http://www.gcmap.com/mapui?P=" + data.AllRoutes(&database)
	err := exec.Command("open", str).Start()
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	m, g := data.ReturnGraph()
	i := 'q'
	for options(&m, &g, i) {}

}
