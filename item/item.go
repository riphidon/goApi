package item

import (
	"encoding/json"
	"math/rand"
	"net/http"
	"strconv"
)

type Item struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Version string `json:"version"`
	Maker   *Maker `json:"maker"`
}

type Maker struct {
	Brand    string `json:"brand"`
	Country  string `json:"country"`
	AreaCode string `json:"areaCode"`
}

var items []Item

func Create(w http.ResponseWriter, r *http.Request) Item {
	var item Item
	_ = json.NewDecoder(r.Body).Decode(&item)
	item.ID = strconv.Itoa(rand.Intn(10000))
	items = append(items, item)
	return item
}
