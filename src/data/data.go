package data

import(

	"graph"
	"flight"

	"fmt"
	"os"
	"encoding/json"
	"io/ioutil"
	"strconv"
)

type Map struct {
	Metro	[]Metro		`json:"metros"`
	Routes	[]Routes	`json:"routes"`
}

type Metro struct {
	Code   		string	`json:"code"`
	Name   		string 	`json:"name"`
	Country   	string 	`json:"country"`
	Continent   string 	`json:"continent"`
	Timezone   	float64	`json:"timezone"`
	Coordinates map[string]int`json:"coordinates"`
	Population  int		`json:"population"`
	Region  	int		`json:"region"`
}

type Routes struct {
	Nodes 	[]string  `json:"ports"`
	Dist	int		  `json:"distance"`
}

/**
 * @Return m - the json structure of the map
 * @Return g - a graph from a given json
 * Converts a json file into memory
 */
func ReturnGraph() (Map, graph.Graph) {

	filename := "/Users/austin/IdeaProjects/CSAir2.1/json/"

	g := graph.Graph{}
	m := Map{}
	new_m := Map{}
	s := "a_route_map.json"

	original, err := ioutil.ReadFile(filename + s)

	fmt.Print("Default route map opened, if you wish to over ride that type 'overide', else hit enter'\n>> ")
	_, err = fmt.Scanf("%s\n", &s)
	if s == "overide" {
		fmt.Print("Please enter the json file name: ")
		_, err = fmt.Scanf("%s\n", &s)
		original, err = ioutil.ReadFile(filename + s)
	}

	err = json.Unmarshal(original, &m)

	for {

		fmt.Print("Would you like to add any additional files (y or n)?\n>>> ")
		_, err = fmt.Scanf("%s\n", &s)

		if s == "y" {

			fmt.Print("What is the file name?\n>>> ")
			_, err = fmt.Scanf("%s\n", &s)
			newfile := "/Users/austin/IdeaProjects/CSAir2.1/json/" + s

			file, err := ioutil.ReadFile(newfile)
			if err != nil { fmt.Println("error2:", err) }

			err = json.Unmarshal(file, &new_m)
			if err != nil { fmt.Println("error2:", err) }

			for i := 0; i < len(new_m.Metro); i++ {
				m.Metro = append(m.Metro, new_m.Metro[i])
			}
			for i := 0; i < len(new_m.Routes); i++ {
				m.Routes = append(m.Routes, new_m.Routes[i])
			}
		}else{ break }
	}

	if err != nil { fmt.Println("error:", err) }

	BuildGraph(&m, &g)
	return m, g
}

/**
 * This function saves the  network (map) to a json
 */
func SaveGraph(m *Map, g *graph.Graph) {

	b, err := json.Marshal(m)
	if err != nil { fmt.Println("error:", err) }
	err = ioutil.WriteFile("/Users/austin/IdeaProjects/CSAir2.1/json/a_route_map.json", b, 0777)
	if err != nil {
		fmt.Printf("File error in opening: %v\n", err)
		os.Exit(1)
	}
}

/**
 * Function used to build a graph from a jsonmap
 * @param m - the mapped json
 * @param grapn - the graph to map the json to
 */
func BuildGraph(m *Map, g *graph.Graph){

	g.Node = make(map[string]graph.Node)

	for i := 0; i < len(m.Metro); i++ {
		node := graph.Node{m.Metro[i].Code, m.Metro[i], nil}
		g.Node[m.Metro[i].Code] = node
	}

	for i := 0; i < len(m.Routes); i++ {
		AddEdge(g, m.Routes[i])
	}
}

/**
 * @param graph - graph of flights/routes/destinations
 * @param route - a destination, source, and time
 * Adds a edge to a given node
 */
func AddEdge(g *graph.Graph, route Routes) {

	edgeone := graph.Edge{}
	edgetwo := graph.Edge{}

	nodeone := g.Node[route.Nodes[0]]
	nodetwo := g.Node[route.Nodes[1]]

	edgeone.Destination = &nodetwo
	edgetwo.Destination = &nodeone

	edgeone.Len = route.Dist
	edgetwo.Len = route.Dist


	nodeone.OutEdges = append(nodeone.OutEdges, edgeone)

	graph.AddOutEdge(&nodetwo, edgetwo)
	g.Node[route.Nodes[0]] = nodeone
	g.Node[route.Nodes[1]] = nodetwo
}

/**
 * @Param graph *Graph - A graph created from a json
 * Prints all city names
 */
func CityAll(g *graph.Graph) {
	for _, node := range g.Node {
		metro := node.Data.(Metro)
		fmt.Println(metro.Name)
	}
}

/**
 * @Param graph *Graph - A graph created from a json
 * @Param code string - The code for a particular cities airport
 * Prints all city graph for given airport
 */
func CityCode(g *graph.Graph, code string) {
	node := g.Node[code] // TODO
	printNode(&node)
}

/**
 * @Param graph *Graph - A graph created from a json
 * @Param name string - The name for a particular city
 * Prints all city graph for given city name
 */
func CityName(g *graph.Graph, name string) {
	for _, node := range g.Node {
		metro := node.Data.(Metro)
		if metro.Name == name {
			printNode(&node)
		}
	}
}

/**
 * Returns all of the routes in the format used for the URL
 * @param m - Map used for the given flights
 */
func AllRoutes(m *Map) string{
	str := ""
	for i := 0; i < len(m.Routes); i++ {
		str += m.Routes[i].Nodes[0] + "-" + m.Routes[i].Nodes[1] + ",+"
	}
	return str
}

/**
 * Prints all of the routes given a graph
 * @param m - Map used for the given flights
 */
func RoutesAll(m *Map){
	for i := 0; i < len(m.Routes); i++ {
		fmt.Println(m.Routes[i].Nodes[0] + "-" + m.Routes[i].Nodes[1])
	}
}

/**
 * Adds a city to the graph/map
 * @param m - Map used for the given flights
 * @param g - graph used to represent flights
 */
func CityAdd(m *Map, g *graph.Graph) {

	newmetro := Metro{}

	fmt.Println("Adding a city to the network!")

	setCode(&newmetro.Code)
	setCity(&newmetro.Name)
	setCountry(&newmetro.Country)
	setContinent(&newmetro.Continent)
	setTimezone(&newmetro.Timezone)

	setPopulation(&newmetro.Population)
	setRegion(&newmetro.Region)

	m.Metro = append(m.Metro, newmetro)
	g.Node[newmetro.Code] = graph.Node{newmetro.Code, newmetro, nil}

	setEdges(g, m, &newmetro)

}

/**
 * Edits a city already belonging to the grpah/map
 * @param m - Map used for the given flights
 * @param g - graph used to represent flights
 */
func CityEdit(m *Map, g *graph.Graph) {

	s := ""

	fmt.Print("What city would you like to edit (please type city code): ")
	_, err := fmt.Scanf("%s\n", &s)
	if err != nil { fmt.Println(err) }
	if g.Node[s].Data == nil {
		fmt.Println("City does not exist")
		return
	}

	metro := g.Node[s].Data.(Metro)
	outedges := g.Node[s].OutEdges

	selection := 'q'
	fmt.Println("What about the city would you like to edit: ")
	fmt.Println("a. Code")
	fmt.Println("b. Name")
	fmt.Println("c. Country")
	fmt.Println("d. Continent")
	fmt.Println("e. Timezone")
	fmt.Println("f. Coordinates")
	fmt.Println("g. Population")
	fmt.Println("h. Region")
	fmt.Println("i. New Route")
	fmt.Print(">> ")

	_, err = fmt.Scanf("%c\n", &selection)
	if err != nil { fmt.Println(err) }

	delete(g.Node, metro.Code)

	if(selection == 'a'){
		setCode(&metro.Code)
	}else if(selection == 'b'){
		setCity(&metro.Name)
	}else if(selection == 'c'){
		setCountry(&metro.Country)
	}else if(selection == 'd'){
		setContinent(&metro.Continent)
	}else if(selection == 'e'){
		setTimezone(&metro.Timezone)
	}else if(selection == 'f'){
		setCoordinates(&metro)
	}else if(selection == 'g'){
		setPopulation(&metro.Population)
	}else if(selection == 'h'){
		setRegion(&metro.Region)
	}else if(selection == 'i'){
		setEdges(g, m, &metro)
	}

	g.Node[metro.Code] = graph.Node{metro.Code, metro, outedges}
}

/**
 * Removes a city from the map
 */
func CityDelete(m *Map, g *graph.Graph) {

	code := ""
	setCode(&code)

	rmnode := g.Node[code]
	if rmnode.Data == nil {
		fmt.Println("Error node to remove does not exist")
		return
	}

	removeOutEdge(g, code)
	removeRoute(m, code)

	var metro []Metro

	for i := 0; i < len(m.Metro); i++ {
		if( m.Metro[i].Code != code ) { metro = append(metro, m.Metro[i]) }
	}
	m.Metro = metro
	delete(g.Node, code)
}

/**
 * Iterates over all routes and removes any route with the given code
 */
func removeRoute(m *Map, code string) {

	var routes []Routes
	for i := 0; i < len(m.Routes); i++ {
		if m.Routes[i].Nodes[0] != code && m.Routes[i].Nodes[1] != code {
			routes = append(routes, m.Routes[i])
		}
	}
	m.Routes = routes
}

/**
 * Removes a specific route from the map given user input
 */
func RemoveSpecifiedRoute(m *Map, g *graph.Graph) {

	src := ""
	dest := ""

	fmt.Println("Source City: ")
	setCode(&src)
	fmt.Println("Destination City: ")
	setCode(&dest)

	var routes []Routes
	for i := 0; i < len(m.Routes); i++ {
		if m.Routes[i].Nodes[0] != src && m.Routes[i].Nodes[1] != dest {
			routes = append(routes, m.Routes[i])
		}
	}
	m.Routes = routes

	removeEdge(g, src, dest)
	removeEdge(g, dest, src)

}

/**
 * Iterates over an entire graph and removes any outbound edges to the
 * Node that is to be removed
 * @param g - graph representing the possible flights
 * @param code - code of the graph to remove
 */
func removeOutEdge(g *graph.Graph, code string) {
	for s, _ := range g.Node {
		removeEdge(g, s, code)
	}
}

/**
 * Helper function, which removes a single edge, given a source and destination
 * @param g - grpah representing flight paths
 * @param src - source node-code of flightpath
 * @param dest - destination node-code of flightpath
 */
func removeEdge(g *graph.Graph, src string, dest string){

	node := g.Node[src]

	var outedges []graph.Edge
	metro := g.Node[src].Data.(Metro)
	delete(g.Node, metro.Code)

	for i := 0; i < len(node.OutEdges); i++ {
		test := *node.OutEdges[i].Destination
		if dest != test.Data.(Metro).Code {
			outedges = append(outedges, node.OutEdges[i])
		}
	}
	g.Node[metro.Code] = graph.Node{metro.Code, metro, outedges}
}


/************************************************************************
 * From this point on, it is just "Getters" and "Setters",
 * Mostly simply Println statements.
 *************************************************************************/


/**
 * Used to set a code given a string
 * s - used for meta.Code adding a code to a node
 */
func setCode(s *string) {

	fmt.Print("Please enter the city code: ")
	_, err := fmt.Scanf("%s\n", s)
	if err != nil { fmt.Println(err) }
}

/**
 * Used to set a city given a string
 * s - used for meta.City adding a city to a node
 */
func setCity(s *string) {

	fmt.Print("Please enter the city name: ")
	_, err := fmt.Scanf("%s\n", s)
	if err != nil { fmt.Println(err) }
}

/**
 * Used to set a country given a string
 * s - used for meta.Country adding a country to a node
 */
func setCountry(s *string) {

	fmt.Print("Please enter the cities country: ")
	_, err := fmt.Scanf("%s\n", s)
	if err != nil { fmt.Println(err) }
}

/**
 * Used to set a continent given a string
 * s - used for meta.Continent adding a continent to a node
 */
func setContinent(s *string) {

	fmt.Print("Please enter the cities continent: ")
	_, err := fmt.Scanf("%s\n", s)
	if err != nil { fmt.Println(err) }
}

/**
 * Used to set a timezone given a string
 * s - used for meta.Timezone adding a timezone to a node
 */
func setTimezone(f *float64) {

	fmt.Print("Please enter the cities timezone: ")
	_, err := fmt.Scanf("%f\n", f)
	if err != nil { fmt.Println(err) }
}

/**
 * Used to set a population given a string
 * s - used for meta.Population adding a population to a node
 */
func setPopulation(d *int) {

	fmt.Print("Please enter the cities population: ")
	_, err := fmt.Scanf("%d\n", d)
	if err != nil { fmt.Println(err) }
}

/**
 * Used to set a region given a string
 * s - used for meta.Region adding a region to a node
 */
func setRegion(d *int) {

	fmt.Print("Please enter the cities region: ")
	_, err := fmt.Scanf("%d\n", d)
	if err != nil { fmt.Println(err) }
}

/**
 * Used to set a coordinates given a metro
 * s - used for meta.Coordinates adding a coordinates to a node
 */
func setCoordinates(newmetro *Metro){
	newmetro.Coordinates = make(map[string]int)

	s := ""
	d := 1

	setLongditude(&s, &d)
	newmetro.Coordinates[s] = d
	setLaditude(&s, &d)
	newmetro.Coordinates[s] = d
}

/**
 * Used to set a Laditude given a string
 * s - used for meta.Coordinates adding a laditude to a node
 * d - used for meta.Coordinates adding degrees to the laditude
 */
func setLaditude(s *string, d *int) {

	fmt.Print("Please enter the cities Coordinates Laditude (W/E): ")
	_, err := fmt.Scanf("%s\n", s)
	if err != nil { fmt.Println(err) }
	fmt.Print("Please enter the cities Coordinates Laditude (degrees): ")
	_, err = fmt.Scanf("%d\n", d)
	if err != nil { fmt.Println(err) }
}

/**
 * Used to set a Laditude given a string
 * s - used for meta.Coordinates adding a laditude to a node
 * d - used for meta.Coordinates adding degrees to the laditude
 */
func setLongditude(s *string, d *int) {

	fmt.Print("Please enter the cities Coordinates Longditude (N/S): ")
	_, err := fmt.Scanf("%s\n", s)
	if err != nil { fmt.Println(err) }
	fmt.Print("Please enter the cities Coordinates Longditude (degrees): ")
	_, err = fmt.Scanf("%d\n", d)
	if err != nil { fmt.Println(err) }
}

/**
 * Used to set a the outgoing edges given a metro
 * @param g - graph representing the map
 * @param m - map holding the data for CSAir
 * @param srcmetro - the source location for a metro to be added
 */
func setEdges(g *graph.Graph, m *Map, srcmetro *Metro){
	s := ""
	d := 1
	fmt.Println("Add edges to the given City: ")
	for s != "zzzz" {

		getConnectedCity(&s)

		if(s == "zzzz"){ break }

		getDistanceToCity(&d)

		route := Routes{[]string{srcmetro.Code, s}, d}
		m.Routes = append(m.Routes, route)
		AddEdge(g, route)
	}
}

/**
 * Used to get connected city given a string
 * s - used for routes.Nodes[0/1] adding a coordinates to a node
 */
func getConnectedCity(s *string) {

	fmt.Print("Please enter the code of a connected city (or zzzz to exit): ")
	_, err := fmt.Scanf("%s\n", s)
	if err != nil { fmt.Println(err) }
}

/**
 * Used to get connected city distance
 * d - distance to a city
 */
func getDistanceToCity(d *int) {

	fmt.Print("Please enter the distance to connected city: ")
	_, err := fmt.Scanf("%d\n", d)
	if err != nil { fmt.Println(err) }
}

/**
 * @Param location *Node - A particular Node
 * Used to print all graph for the given city.
 */
func printNode(node *graph.Node) {

	if node.Data == nil {
		fmt.Println("Error, city does not exist")
		return
	}

	if(node != nil){

		location := node.Data.(Metro)

		fmt.Println("Name: " + location.Name)
		fmt.Println("Code: " + location.Code)
		fmt.Println("Country: " + location.Country)
		fmt.Println("Timezone: " + strconv.FormatFloat(location.Timezone, 'f', 1, 64)) //FormatFloat(f float64, fmt byte, prec, bitSize int)
		fmt.Print("Coordinates: ")
		for key, value := range location.Coordinates { fmt.Print( key, ": ", strconv.Itoa(value), " ") }
		fmt.Println("\nPopulation:  " + strconv.Itoa(location.Population))
		fmt.Println("Region: " + strconv.Itoa(location.Region))
		fmt.Print("Possible Flights: ")
		PrintOutEdges(node)
	}
}

/**
 * @param Given a airport code, print possible Edges
 */
func PrintOutEdges(node *graph.Node){
	for i := 0; i < len(node.OutEdges); i++ {
		if node.OutEdges[i].Destination == nil { break }
		metro := node.OutEdges[i].Destination.Data.(Metro)
		fmt.Print(metro.Name)
		fmt.Print("|" + strconv.Itoa(node.OutEdges[i].Len))
		fmt.Print(", ")
	}
	fmt.Print("\n")
}

/**
 * This function returns statistical information about the graph
 * @param m - Map for the given flights
 * @param g - Graph used for representing the flights
 */
func AirlineStatistics(m *Map, g *graph.Graph){

	var shortest *Routes
	var longest  *Routes
	var smallest *Metro
	var largest  *Metro
	avg_dist := 0
	avg_pop	:= 0
	var continents []string
	flag := true

	var hubs []Metro

	for i := 0; i < len(m.Routes); i++ {

		if shortest == nil || shortest.Dist > m.Routes[i].Dist {
			shortest = &m.Routes[i]
		}
		if longest == nil || longest.Dist < m.Routes[i].Dist {
			longest = &m.Routes[i]
		}
		avg_dist += m.Routes[i].Dist
	}
	for i := 0; i < len(m.Metro); i++ {

		if smallest == nil || smallest.Population > m.Metro[i].Population {
			smallest = &m.Metro[i]
		}
		if largest == nil || largest.Population < m.Metro[i].Population {
			largest = &m.Metro[i]
		}
		avg_pop += m.Metro[i].Population
		for j := 0; j < len(continents); j++ {
			if(m.Metro[i].Continent == continents[j]){
				flag = false
				break
			}
		}
		for j := 0; j < len(continents); j++ {
			if(m.Metro[i].Continent == continents[j]){
				flag = false
				break
			}
		}
		if(flag){ continents = append(continents, m.Metro[i].Continent) }
		flag = true
	}

	for _, node := range g.Node {
		if len(node.OutEdges) > 4 {
			hubs = append(hubs, node.Data.(Metro))
		}
	}

	avg_pop /= len(m.Metro)
	avg_dist /= len(m.Routes)

	fmt.Println("Longest Flight: " + longest.Nodes[0] + " to " + longest.Nodes[1] + " at " + strconv.Itoa(longest.Dist))
	fmt.Println("Shortest Flight: " + shortest.Nodes[0] + " to " + shortest.Nodes[1] + " at " + strconv.Itoa(shortest.Dist))
	fmt.Println("Average Flight: " + strconv.Itoa(avg_dist))
	fmt.Println("Largest City: " + largest.Name + " with a population of " + strconv.Itoa(largest.Population))
	fmt.Println("Smallest City: " + smallest.Name + " with a population of " + strconv.Itoa(smallest.Population))
	fmt.Println("Average Population: " + strconv.Itoa(avg_pop))
	fmt.Print("Continents Currently Serviced: ")
	for j := 0; j < len(continents); j++ {
		fmt.Print(continents[j] + ", ")
	}
	fmt.Print("\n")
	fmt.Print("Hubs: ")
	for i := 0; i < len(hubs); i++ {
		fmt.Print(hubs[i].Name + ", ")
	}
	fmt.Print("\n")
}

/**
 * This function prints out the data for a flight, given a graph
 */
func MakeFlight(g *graph.Graph){

	var src string
	var dest string

	fmt.Println("What city would you like to fly out of?")
	setCode(&src)
	fmt.Println("What city are you flying to?")
	setCode(&dest)

	source := g.Node[src]
	destination := g.Node[dest]

	path, dist := graph.Dijkstra(g, &source, &destination)

	fmt.Print("\nPath: ")
	for j := len(path) - 1; j >= 0; j-- {
		fmt.Print(path[j].Code + "   ")
	}

	fmt.Print("\nTotal Distance: ")
	fmt.Println(dist)

	h, m, s, sum := flight.CalculateFlightStat(g, path)

	fmt.Print("Flight Time: ")
	fmt.Print(h)
	fmt.Print(" hours, ")
	fmt.Print(m)
	fmt.Print(" minutes, ")
	fmt.Print(s)
	fmt.Println(" seconds")

	fmt.Print("This Flight Costs: ")
	fmt.Printf("%6.2f", sum)

}
