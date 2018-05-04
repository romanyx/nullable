package main

import (
	"encoding/json"
	"fmt"
	"log"

	"gopkg.in/romanyx/nullable.v1"
)

const (
	updateRequestJSON = `{
		"title": "Nullable package",
		"status": null
	}`
)

type updateRequest struct {
	Title       nullable.String `json:"title"`
	Description nullable.String `json:"description"`
	Status      nullable.Int    `json:"status"`
}

func main() {
	ur := updateRequest{}
	if err := json.Unmarshal([]byte(updateRequestJSON), &ur); err != nil {
		log.Fatal(err)
	}

	updateDB(&ur)

	usr := struct {
		Name nullable.String `json:"name"`
	}{}

	data := []byte("{}")
	if err := json.Unmarshal(data, &usr); err != nil {
		log.Fatalf("unmarshaling failed: %s\n", err)
	}

	fmt.Println("name present", usr.Name.Present) // false
	fmt.Println("name valid", usr.Name.Valid)     // false
	fmt.Println("name value", usr.Name.Value)     // ""

	data = []byte("{\"name\":null}")
	if err := json.Unmarshal(data, &usr); err != nil {
		log.Fatalf("unmarshaling failed: %s\n", err)
	}

	fmt.Println("name present", usr.Name.Present) // true
	fmt.Println("name valid", usr.Name.Valid)     // false
	fmt.Println("name value", usr.Name.Value)     // ""

	data = []byte("{\"name\":\"John\"}")
	if err := json.Unmarshal(data, &usr); err != nil {
		log.Fatalf("unmarshaling failed: %s\n", err)
	}

	fmt.Println("name present", usr.Name.Present) // true
	fmt.Println("name valid", usr.Name.Valid)     // true
	fmt.Println("name value", usr.Name.Value)     // "John"
}

func updateDB(ur *updateRequest) {
	updateColumn("title", ur.Title.Present, ur.Title.Valid, ur.Title.Value)
	updateColumn("description", ur.Description.Present, ur.Description.Valid, ur.Description.Value)
	updateColumn("status", ur.Status.Present, ur.Status.Valid, ur.Status.Value)
}

func updateColumn(field string, present, valid bool, value interface{}) {
	if present {
		if valid {
			fmt.Printf("set %s value to: %v\n", field, value)
			return
		}

		fmt.Printf("set %s value to null\n", field)
		return
	}

	fmt.Printf("%s value did not modified\n", field)
}
