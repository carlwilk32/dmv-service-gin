package api

import (
	"encoding/json"
	"fmt"
	dmv "github.com/carlwilk32/dmv-service-gin/client"
	"math"
	"net/http"
	"sort"
	"strconv"
)

const limit = 9

func ByDistance(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	latStr := query.Get("lat")
	lonStr := query.Get("lon")
	if len(latStr) == 0 || len(lonStr) == 0 {
		http.Error(w, "Latitude and Longitude should present", http.StatusBadRequest)
		return
	}

	offices, err := dmv.GetFieldOffices("")
	if err != nil {
		fmt.Printf("DMV client failed to execute call.")
		http.Error(w, err.Error(), 500)
		return
	}

	officeByDistance := make(map[int]dmv.FieldOffice)
	for i := range offices {
		it := offices[i]
		distance := getDistanceToOffice(latStr, lonStr, it)
		officeByDistance[distance] = it
	}

	keys := make([]int, 0, len(officeByDistance))

	for k := range officeByDistance {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	response := [limit]Response{}
	for i := range limit {
		k := keys[i]
		it := Response{k, officeByDistance[k].String()}
		response[i] = it
		//fmt.Println(it)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
}

type Response struct {
	Distance int    `json:"distance"`
	Name     string `json:"name"`
}

func (r Response) String() string {
	return fmt.Sprintf("%v Miles --> %v", r.Distance, r.Name)
}

func getDistanceToOffice(latStr, lonStr string, office dmv.FieldOffice) int {
	fLat, _ := strconv.ParseFloat(latStr, 64)
	fLon, _ := strconv.ParseFloat(lonStr, 64)
	oLat, oLon := office.LatLon()
	return int(calcDistance(degToRad(fLat), degToRad(fLon), degToRad(oLat), degToRad(oLon)))
}

func calcDistance(lat1, lon1, lat2, lon2 float64) float64 {
	dlon := lon2 - lon1
	dlat := lat2 - lat1
	a := math.Pow(math.Sin(dlat/2), 2) + math.Cos(lat1)*math.Cos(lat2)*math.Pow(math.Sin(dlon/2), 2)
	c := 2 * math.Asin(math.Sqrt(a))

	// Radius of earth in miles, 6371 for kms
	r := float64(3956)
	return c * r
}

func degToRad(val float64) float64 {
	return val * math.Pi / 180
}
