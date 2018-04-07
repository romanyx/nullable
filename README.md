

# nullable
`import "github.com/romanyx/nullable"`

* [Overview](#pkg-overview)
* [Index](#pkg-index)
* [Subdirectories](#pkg-subdirectories)

## <a name="pkg-overview">Overview</a>
Package nullable provides types what allows to determine
if a JSON key has been set to null or not provided.


	usr := struct {
		Name nullable.String `json:"name"`
	}{}
	
	data := []byte("{}")
	if err := json.Unmarshal(data, &usr); err != nil {
		log.Fatalf("unmarshaling failed: %s\n", err)
	}
	
	fmt.Println("name present", usr.Name.Present) // false
	fmt.Println("name valid", usr.Name.Valid)     // false
	fmt.Println("name valid", usr.Name.Value)     // ""
	
	data = []byte("{\"name\":null}")
	if err := json.Unmarshal(data, &usr); err != nil {
		log.Fatalf("unmarshaling failed: %s\n", err)
	}
	
	fmt.Println("name present", usr.Name.Present) // true
	fmt.Println("name valid", usr.Name.Valid)     // false
	fmt.Println("name valid", usr.Name.Value)     // ""
	
	data = []byte("{\"name\":\"John\"}")
	if err := json.Unmarshal(data, &usr); err != nil {
		log.Fatalf("unmarshaling failed: %s\n", err)
	}
	
	fmt.Println("name present", usr.Name.Present) // true
	fmt.Println("name valid", usr.Name.Valid)     // true
	fmt.Println("name valid", usr.Name.Value)     // "John"




## <a name="pkg-index">Index</a>
* [type Bool](#Bool)
  * [func (b *Bool) UnmarshalJSON(data []byte) error](#Bool.UnmarshalJSON)
* [type Float](#Float)
  * [func (f *Float) UnmarshalJSON(data []byte) error](#Float.UnmarshalJSON)
* [type FloatSlice](#FloatSlice)
  * [func (f *FloatSlice) UnmarshalJSON(data []byte) error](#FloatSlice.UnmarshalJSON)
* [type Int](#Int)
  * [func (i *Int) UnmarshalJSON(data []byte) error](#Int.UnmarshalJSON)
* [type IntSlice](#IntSlice)
  * [func (i *IntSlice) UnmarshalJSON(data []byte) error](#IntSlice.UnmarshalJSON)
* [type String](#String)
  * [func (s *String) UnmarshalJSON(data []byte) error](#String.UnmarshalJSON)
* [type StringSlice](#StringSlice)
  * [func (s *StringSlice) UnmarshalJSON(data []byte) error](#StringSlice.UnmarshalJSON)


#### <a name="pkg-files">Package files</a>
[bool.go](/src/github.com/romanyx/nullable/bool.go) [doc.go](/src/github.com/romanyx/nullable/doc.go) [float.go](/src/github.com/romanyx/nullable/float.go) [int.go](/src/github.com/romanyx/nullable/int.go) [string.go](/src/github.com/romanyx/nullable/string.go) 






## <a name="Bool">type</a> [Bool](/src/target/bool.go?s=135:295#L10)
``` go
type Bool struct {
    Present bool // Present is true if key is present in json
    Valid   bool // Valid is true if value is not null and valid bool
    Value   bool
}
```
Bool represents a bool that may be null or not
present in json at all.










### <a name="Bool.UnmarshalJSON">func</a> (\*Bool) [UnmarshalJSON](/src/target/bool.go?s=351:398#L17)
``` go
func (b *Bool) UnmarshalJSON(data []byte) error
```
UnmarshalJSON implements json.Marshaler interface.




## <a name="Float">type</a> [Float](/src/target/float.go?s=137:302#L10)
``` go
type Float struct {
    Present bool // Present is true if key is present in json
    Valid   bool // Valid is true if value is not null and valid float
    Value   float64
}
```
Float represents a float that may be null or not
present in json at all.










### <a name="Float.UnmarshalJSON">func</a> (\*Float) [UnmarshalJSON](/src/target/float.go?s=358:406#L17)
``` go
func (f *Float) UnmarshalJSON(data []byte) error
```
UnmarshalJSON implements json.Marshaler interface.




## <a name="FloatSlice">type</a> [FloatSlice](/src/target/float.go?s=675:851#L34)
``` go
type FloatSlice struct {
    Present bool // Present is true if key is present in json
    Valid   bool // Valid is true if value is not null and valid []float64
    Value   []float64
}
```
FloatSlice represents a float slice that may be null or not
present in json at all.










### <a name="FloatSlice.UnmarshalJSON">func</a> (\*FloatSlice) [UnmarshalJSON](/src/target/float.go?s=907:960#L41)
``` go
func (f *FloatSlice) UnmarshalJSON(data []byte) error
```
UnmarshalJSON implements json.Marshaler interface.




## <a name="Int">type</a> [Int](/src/target/int.go?s=134:295#L10)
``` go
type Int struct {
    Present bool // Present is true if key is present in json
    Valid   bool // Valid is true if value is not null and valid int64
    Value   int64
}
```
Int represents an int that may be null or not
present in json at all.










### <a name="Int.UnmarshalJSON">func</a> (\*Int) [UnmarshalJSON](/src/target/int.go?s=351:397#L17)
``` go
func (i *Int) UnmarshalJSON(data []byte) error
```
UnmarshalJSON implements json.Marshaler interface.




## <a name="IntSlice">type</a> [IntSlice](/src/target/int.go?s=663:833#L34)
``` go
type IntSlice struct {
    Present bool // Present is true if key is present in json
    Valid   bool // Valid is true if value is not null and valid []int64
    Value   []int64
}
```
IntSlice represents an int slice that may be null or not
present in json at all.










### <a name="IntSlice.UnmarshalJSON">func</a> (\*IntSlice) [UnmarshalJSON](/src/target/int.go?s=889:940#L41)
``` go
func (i *IntSlice) UnmarshalJSON(data []byte) error
```
UnmarshalJSON implements json.Marshaler interface.




## <a name="String">type</a> [String](/src/target/string.go?s=166:332#L12)
``` go
type String struct {
    Present bool // Present is true if key is present in json
    Valid   bool // Valid is true if value is not null and valid string
    Value   string
}
```
String represents a string that may be null or not
present in json at all.










### <a name="String.UnmarshalJSON">func</a> (\*String) [UnmarshalJSON](/src/target/string.go?s=388:437#L19)
``` go
func (s *String) UnmarshalJSON(data []byte) error
```
UnmarshalJSON implements json.Marshaler interface.




## <a name="StringSlice">type</a> [StringSlice](/src/target/string.go?s=704:860#L36)
``` go
type StringSlice struct {
    Present bool // Present is true if key is present in json
    Valid   bool // Valid is true if value is not null
    Value   []string
}
```
StringSlice represents a []string that may be null or not
present in json at all.










### <a name="StringSlice.UnmarshalJSON">func</a> (\*StringSlice) [UnmarshalJSON](/src/target/string.go?s=916:970#L43)
``` go
func (s *StringSlice) UnmarshalJSON(data []byte) error
```
UnmarshalJSON implements json.Marshaler interface.








- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
