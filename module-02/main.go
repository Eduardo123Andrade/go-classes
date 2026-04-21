package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name string `json:"name"`
	Id   uint64 `json:"id"`
}

func (u User) UpdateName(name string) {
	u.Name = name
}

func main() {
	user := User{Name: "Teste", Id: 0}
	user.UpdateName("Pedro")
	
	res, err := json.Marshal(user)
	
	if(err != nil){
		panic(err)
	}

	fmt.Println(string(res))
}
