package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name         string   `json:"name"`
	PhoneNumbers []string `json:"phoneNumbers"`
}

func (o *Person) UnmarshalJSON(b []byte) error {

	//parse object that only has a name
	var name string
	if err := json.Unmarshal(b, &name); err == nil {
		*o = Person{}
		o.Name = name
		return nil
	}

	//name parsing failed, so let's parse the whole thing
	type person Person
	var p person
	if err := json.Unmarshal(b, &p); err != nil {
		return err
	}
	
	*o = Person(p)
	return nil
}

const payload = `[
    {"name": "Peter Johnson", "phoneNumbers": ["555-333-3333", "555-555-5555"]},
    "Vojtek Dobry",
    {"name": "Adeline Smith"}
]
`

func main() {
	var persons []*Person
	if err := json.Unmarshal([]byte(payload), &persons); err != nil {
		panic(err)
	}
	for _, obj := range persons {
		fmt.Printf("Name:%v, Numbers:%v\n", obj.Name, obj.PhoneNumbers)
	}
}
