[![GoDoc](https://godoc.org/github.com/romanyx/nullable?status.svg)](https://godoc.org/github.com/romanyx/nullable)
[![Build Status](https://travis-ci.org/romanyx/nullable.png)](https://travis-ci.org/romanyx/nullable)
[![Go Report Card](https://goreportcard.com/badge/github.com/romanyx/nullable)](https://goreportcard.com/report/github.com/romanyx/nullable)

# nullable

Provides ability to determine if a json key has been set to null or not provided.
Inspired by [How to determine if a JSON key has been set to null or not provided](https://www.calhoun.io/how-to-determine-if-a-json-key-has-been-set-to-null-or-not-provided/) article by **Jon Calhoun**

## Example

``` go
	usr := struct {
		Name nullable.String `json:"name"`
	}{}

	data := []byte("{}")
	if err := json.Unmarshal(data, &usr); err != nil {
		log.Fatalf("unmarshaling failed: %s\n", err)
	}

	fmt.Println(usr.Name.Present) // false
	fmt.Println(usr.Name.Valid)   // false
	fmt.Println(usr.Name.Value)   // ""

	data = []byte("{\"name\":null}")
	if err := json.Unmarshal(data, &usr); err != nil {
		log.Fatalf("unmarshaling failed: %s\n", err)
	}

	fmt.Println(usr.Name.Present) // true
	fmt.Println(usr.Name.Valid)   // false
	fmt.Println(usr.Name.Value)   // ""

	data = []byte("{\"name\":\"John\"}")
	if err := json.Unmarshal(data, &usr); err != nil {
		log.Fatalf("unmarshaling failed: %s\n", err)
	}

	fmt.Println(usr.Name.Present) // true
	fmt.Println(usr.Name.Valid)   // true
	fmt.Println(usr.Name.Value)   // "John"
```