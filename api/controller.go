package api

import (
	"fmt"
	dmv "github.com/carlwilk32/dmv-service-gin/client"
	"math"
	"net/http"
	"sort"
	"strconv"
)

// todo just a sample implementqtion for now
func ByDistance(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	latStr := query.Get("lat")
	lonStr := query.Get("lon")
	if len(latStr) == 0 || len(lonStr) == 0 {
		panic("empty lat, lon")
	}

	offices := dmv.GetFieldOffices("")
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

	for i := range 10 {
		k := keys[i]
		fmt.Println(i+1, "->", k, "Miles from", officeByDistance[k])
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	//err := json.NewEncoder(w).Encode(offices[0]) //write one
	//if err != nil {
	//	panic(err)
	//}
}

type OfficeResponse struct {
	id      string
	IdShort string `json:"id"`
	Name    string `json:"name"`
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
