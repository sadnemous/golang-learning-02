package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"simple/types"
)

func handleRequest(payload []byte, router any) {
	fmt.Println("=== Request Handler ===")
	fmt.Printf("Payload: %s\n", string(payload))

	if r, ok := router.(types.HttpRouter); ok {
		fmt.Printf("Method: %s\n", r.Method)
		fmt.Printf("URL: %s\n", r.URL)
		fmt.Printf("Accept: %s\n", r.Accept)
		fmt.Printf("Content-Type: %s\n", r.ContentType)

		req, err := http.NewRequest(r.Method, r.URL, bytes.NewBuffer(payload))
		if err != nil {
			fmt.Println("Error creating request:", err)
			return
		}

		req.Header.Set("Accept", r.Accept)
		req.Header.Set("Content-Type", r.ContentType)

		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("Error sending request:", err)
			return
		}
		defer resp.Body.Close()

		body, _ := io.ReadAll(resp.Body)
		fmt.Printf("\nResponse Status: %s\n", resp.Status)
		fmt.Printf("Response Body: %s\n", string(body))
	}

}

func printUser(user types.User) {
	fmt.Println("  ID      :", user.ID)
	fmt.Println("  Name    :", user.Name)
	fmt.Println("  Username:", user.Username)
	fmt.Println("Name:", user.Name, "ID:", user.ID)
}
func printUsers(users map[string]types.User) {
	for k, v := range users {
		fmt.Println("Key:", k, "Value:", v)
	}
}
func printAnything(data any) {
	//fmt.Printf("Value: %v, Type: %T\n", data, data)
	// check the type of data

	data_type := reflect.TypeOf(data).String()

	// check if data_type has pointer or not
	fmt.Println(data_type)
	// Type switch
	switch v := data.(type) {
	case map[string]types.User:
		// for k, v := range v {
		// 	tmpUser := v
		// 	fmt.Println("Key:", k, "Value:", v)
		// 	fmt.Println("  ID      :", tmpUser.ID)
		// 	fmt.Println("  Name    :", tmpUser.Name)
		// 	fmt.Println("  Username:", tmpUser.Username)
		// }
		printUsers(v)
	case map[string]*types.User:
		for k, v := range v {
			tmpUser := v
			fmt.Println("Key:", k, "Value:", v)
			fmt.Println("  ID      :", tmpUser.ID)
			fmt.Println("  Name    :", tmpUser.Name)
			fmt.Println("  Username:", tmpUser.Username)
		}
	case types.User:
		fmt.Println("Name:", v.Name, "ID:", v.ID)
	case int:
		fmt.Println("Integer:", v)
	case string:
		fmt.Println("String:", v)
	default:
		fmt.Println("Unknown type")
	}

	// Or type assertion
	if user, ok := data.(types.User); ok {
		fmt.Println("User name:", user.Name)
	}

}

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

func (d Dog) Speak() string {
	return "Woof! I'm " + d.Name
}

type Cat struct {
	Name string
}

func (c Cat) Speak() string {
	return "Meow! I'm " + c.Name
}

func makeItSpeak(s Speaker) {
	fmt.Println(s.Speak())
}

func main() {

	jsonRequest := `{"name":"Bob","email":"bob@example.com"}`
	payload := []byte(jsonRequest)
	/*
			curl -v -X POST http://127.0.0.1:5000/users \
		  -H "Content-Type: application/json" \
		  -d '{"name":"Bob","email":"bob@example.com"}'
	*/
	router := types.HttpRouter{
		URL:         "http://127.0.0.1:5000/users",
		Method:      "POST",
		Accept:      "application/json",
		ContentType: "application/json",
	}

	handleRequest(payload, router)
	return

	var user_1 types.User = types.User{ID: 1, Name: "John Doe", Username: "johndoe"}
	var user_2 types.User = types.User{ID: 1, Name: "John Doe", Username: "johndoe"}
	var user_3 types.User = types.User{ID: 1, Name: "John Doe", Username: "johndoe"}
	var users = make(map[string]types.User)
	users["user_1"] = user_1
	users["user_2"] = user_2
	users["user_3"] = user_3

	fmt.Println("\nMap example:")

	// printAnything(42)
	// printAnything("Hello")
	// printAnything(3.14)
	// printAnything([]int{1, 2, 3})
	// printAnything(users)

	fmt.Println("\n=== Value vs Pointer Map ===")

	valueMap := make(map[string]types.User)
	pointerMap := make(map[string]*types.User)

	valueMap["alice"] = types.User{ID: 1, Name: "Alice Cooper", Username: "username-alice"}
	pointerMap["bob"] = &types.User{ID: 2, Name: "Bob Dylan", Username: "username-bob"}

	fmt.Println("\nTrying to modify through value map:")
	u := valueMap["alice"]
	u.Name = "Alice Modified"
	fmt.Printf("After modification - Map: %v, Variable: %v\n", valueMap["alice"].Name, u.Name)

	fmt.Println("\nTrying to modify through pointer map:")
	p := pointerMap["bob"]
	p.Name = "Bob Modified"
	fmt.Printf("After modification - Map: %v, Variable: %v\n", pointerMap["bob"].Name, p.Name)

	fmt.Printf("\n")
	fmt.Printf("=== valueMap===\n")
	printAnything(valueMap)

	fmt.Printf("\n")
	fmt.Printf("=== pointerMap===\n")
	printAnything(pointerMap)

	fmt.Printf("=====\n")
	fmt.Printf("\n")

	fmt.Println("\nInterface example:")
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}

	makeItSpeak(dog)
	makeItSpeak(cat)
}
