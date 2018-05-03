// Package nullable provides types what allows to determine
// if a JSON key has been set to null or not provided.
//
//		usr := struct {
//			Name nullable.String `json:"name"`
//		}{}
//
//		data := []byte("{}")
//		if err := json.Unmarshal(data, &usr); err != nil {
//			log.Fatalf("unmarshaling failed: %s\n", err)
//		}
//
//		fmt.Println("name present", usr.Name.Present) // false
//		fmt.Println("name valid", usr.Name.Valid)     // false
//		fmt.Println("name value", usr.Name.Value)     // ""
//
//		data = []byte("{\"name\":null}")
//		if err := json.Unmarshal(data, &usr); err != nil {
//			log.Fatalf("unmarshaling failed: %s\n", err)
//		}
//
//		fmt.Println("name present", usr.Name.Present) // true
//		fmt.Println("name valid", usr.Name.Valid)     // false
//		fmt.Println("name value", usr.Name.Value)     // ""
//
//		data = []byte("{\"name\":\"John\"}")
//		if err := json.Unmarshal(data, &usr); err != nil {
//			log.Fatalf("unmarshaling failed: %s\n", err)
//		}
//
//		fmt.Println("name present", usr.Name.Present) // true
//		fmt.Println("name valid", usr.Name.Valid)     // true
//		fmt.Println("name value", usr.Name.Value)     // "John"
package nullable
