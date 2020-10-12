package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
	"github.com/luhanm/bexs-backend-exam/routes/scale"
	sca "github.com/luhanm/bexs-backend-exam/routes/scale"
	"github.com/luhanm/bexs-backend-exam/routes/util"
)

func HttpInsertRoute(w http.ResponseWriter, r *http.Request) {
	var scale sca.TypeScale
	var err error
	err = json.NewDecoder(r.Body).Decode(&scale)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	scale.Origin = strings.ToUpper(strings.TrimSpace(scale.Origin))
	scale.Destination = strings.ToUpper(strings.TrimSpace(scale.Destination))

	if scale.Origin == "" || scale.Destination == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("origin or destination invalid"))
		return
	}

	_, err = sca.GetScale(scale.Origin, scale.Destination)
	if err != nil {
		err := sca.AddScale(scale)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
		}
	}

	w.WriteHeader(http.StatusNoContent)
}

func HttpGetRoute(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	paramRoute := params["route"]
	cheapestWay, _ := strconv.ParseBool(r.URL.Query().Get("cheapest"))
	origin, destination, err := util.SplitRoute(paramRoute)
	if err != nil {
		fmt.Println(err.Error())
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	var bodyResponse []byte

	if cheapestWay {
		bestRoute, err := util.FindCheapestWay(origin, destination)
		if err != nil {
			fmt.Println(err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		fmt.Printf("best route: %s > $%d", bestRoute.CompleteWay, bestRoute.Cost)
		fmt.Println("")

		bodyResponse, _ = json.Marshal(bestRoute)
	} else {
		scale, err := scale.GetScale(origin, destination)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte(err.Error()))
			return
		}
		fmt.Printf("route: %s - %s > $%d", scale.Origin, scale.Destination, scale.Cost)
		fmt.Println("")
		bodyResponse, _ = json.Marshal(scale)
	}

	w.Write(bodyResponse)
}
