package router

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/riphidon/evo/item"
)

func getItems(w http.ResponseWriter, r *http.Request) error {
	_, err := fmt.Fprintf(w, "Items Endpoint Hit")
	if err != nil {
		return err
	}
	return nil
}

func getItem(w http.ResponseWriter, r *http.Request) error {
	_, err := fmt.Fprintf(w, "Item Endpoint Hit")
	if err != nil {
		return err
	}
	return nil
}

func createItem(w http.ResponseWriter, r *http.Request) error {
	_, err := fmt.Fprintf(w, "Item Creation Endpoint Hit")
	w.Header().Set("content-type", "application/json")
	item := item.Create(w, r)
	fmt.Println("item : ")
	if err != nil {
		return err
	}
	fmt.Printf("item: %v", item)
	json.NewEncoder(w).Encode(item)
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
