package util

import (
	"fmt"
	"sort"
	"strings"

	"github.com/luhanm/bexs-backend-exam/routes/scale"
)

type Route struct {
	Steps       []string `json:"-"`
	CompleteWay string   `json:"bestRoute"`
	Cost        int      `json:"cost"`
	Invalid     bool     `json:"-"`
	Completed   bool     `json:"-"`
}

//SplitRoute recebe a rota no formato DE-PARA e devolve particionada. Caso não seja possível será retornado erro
func SplitRoute(route string) (origin, destination string, err error) {
	dePara := strings.Split(strings.ToUpper(route), "-")
	if len(dePara) != 2 {
		return "", "", fmt.Errorf("invalid route")
	}
	return dePara[0], dePara[1], nil
}

func (route *Route) addStep(steps []string, step string) {

	route.Steps = append(steps, step)
	if len(route.Steps) > 1 {
		route.CompleteWay = route.CompleteWay + " - "
	}
	route.CompleteWay = route.CompleteWay + step
}

//ContainsStep Retorna true caso o local pesquisado já esta na rota
func (route *Route) ContainsStep(place string) bool {
	for _, step := range route.Steps {
		if step == place {
			return true
		}
	}
	return false
}

func getNextPoint(routes []Route, destination string) []Route {
	for key, route := range routes {
		if route.Completed {
			continue
		}
		for _, scale := range scale.Scales {
			isOneStep := scale.Origin == route.Steps[len(route.Steps)-1]
			repited := route.ContainsStep(scale.Destination)
			if isOneStep && !repited {
				var newRoute Route
				newRoute.CompleteWay = route.CompleteWay
				newRoute.addStep(route.Steps, scale.Destination)
				newRoute.Cost = route.Cost + scale.Cost
				newRoute.Completed = scale.Destination == destination
				routes = append(routes, newRoute)
			}
		}
		routes[key].Invalid = true
	}

	//Remove
	for n := 0; n < len(routes); n++ {
		if routes[n].Invalid && !routes[n].Completed {
			routes = append(routes[:n], routes[n+1:]...)
			n--
		}
	}

	for _, route := range routes {
		if !route.Completed {
			routes = getNextPoint(routes, destination)
			break
		}
	}
	return routes
}

//FindCheapestWay encontra a rota mais barata. Caso não encontre será retornado erro
func FindCheapestWay(origin string, destination string) (bestRote Route, err error) {

	var routes []Route
	var route Route
	route.addStep(route.Steps, origin)
	route.Cost = 0
	route.Completed = false
	route.Invalid = false
	routes = append(routes, route)

	routes = getNextPoint(routes, destination)
	sort.Slice(routes, func(i, j int) bool {
		return routes[i].Cost < routes[j].Cost
	})

	if len(routes) == 0 {
		return route, fmt.Errorf("non-existent route")
	}

	if len(routes[0].Steps) < 2 {
		return route, fmt.Errorf("non-existent route")
	}

	return routes[0], nil
}
