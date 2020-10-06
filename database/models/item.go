package database

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	"github.com/pkg/errors"
)

type Item struct {
	ID   string `json:"id"`
	Name string `json:"name"`
	Vers string `json:"version"`
}

var items []Item

// GetItems return a list of items from a json file
func GetItems() ([]Item, error) {
	var items []Item
	jsonItems, err := ioutil.ReadFile("database/itemCatalog.json")
	if err != nil {
		return items, err
	}
	if err := json.Unmarshal(jsonItems, &items); err != nil {
		return items, err
	}
	return items, nil

}

// GetItemByID returns an item given a specific id
func GetItemByID(id string) (Item, error) {
	var i Item
	var items []Item
	items, err := GetItems()
	if err != nil {
		return Item{}, err
	}
	for _, item := range items {
		if item.ID == id {
			i = item
			return i, nil
		}
	}
	return i, errors.Wrap(err, "No Item Found")
}

// CreateItem creates a new item in the database
func CreateItem(i []byte) error {
	var item Item
	err := json.Unmarshal(i, &item)
	if err != nil {
		return errors.Wrap(err, "couldn't marshal item")
	}
	fmt.Printf("Item Unmarshaled: %v", item)
	err = appendData(item)
	if err != nil {
		return errors.Wrap(err, "Item Creation Fail")
	}
	return nil
}

func appendData(i Item) error {
	filename := "database/itemCatalog.json"
	err := checkFile(filename)
	if err != nil {
		return errors.Wrap(err, "an error occured while checking file name")
	}

	// Read File
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		return errors.Wrap(err, "Couldn't read file")

	}

	data := []Item{}
	json.Unmarshal(file, &data)
	data = append(data, i)
	dataBytes, err := json.Marshal(data)
	if err != nil {
		return errors.Wrap(err, "Couldn't perform marshal")
	}

	// Write To File
	err = ioutil.WriteFile(filename, dataBytes, 0644)
	if err != nil {
		return errors.Wrap(err, "Couldn't write to file")
	}
	fmt.Printf("New Item Created: %v", i.Name)
	return nil
}

func checkFile(filename string) error {
	_, err := os.Stat(filename)
	if os.IsNotExist(err) {
		_, err := os.Create(filename)
		if err != nil {
			return err
		}
	}
	return nil
}
