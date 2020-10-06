package router

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/pkg/errors"
	database "github.com/riphidon/evo/database/models"
)

func getItems(w http.ResponseWriter, r *http.Request) error {
	items, err := database.GetItems()
	if err != nil {
		return err
	}
	json.NewEncoder(w).Encode(items)
	return nil
}

func getItem(w http.ResponseWriter, r *http.Request) error {
	params := mux.Vars(r)
	id := params["id"]
	item, err := database.GetItemByID(id)
	if err != nil {
		return errors.Wrap(err, "Error fetching item")
	}
	json.NewEncoder(w).Encode(item)
	return nil
}

func createItem(w http.ResponseWriter, r *http.Request) error {
	r.Body = http.MaxBytesReader(w, r.Body, 1000)
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return err
	}
	err = database.CreateItem(body)
	if err != nil {
		return errors.Wrap(err, "Item Creation Failure")
	}
	return nil
}

func updateItem(w http.ResponseWriter, r *http.Request) error {
	_, err := fmt.Fprintf(w, "Item Update Endpoint Hit")
	if err != nil {
		return err
	}
	return nil
}
func deleteItem(w http.ResponseWriter, r *http.Request) error {
	_, err := fmt.Fprintf(w, "Item Deletion Endpoint Hit")
	if err != nil {
		return err
	}
	return nil
}
