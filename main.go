package main

import (
	"fmt"
//	"math/rand"
	"slices"
	"os"
	"encoding/json"
	"io/ioutil"
)

type Contact struct {
	// Defining the way we store users in the json
	Name string `json:"contact_name"`
	Age int `json:"age"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
var Contacts []Contact

func saveFile(file []Contact) (int){
	marshalled, err := json.MarshalIndent(file, "", "\t") // making it prettier
	if err != nil {
		fmt.Println(err)
		return -1
	} else {
		// Write this shit
		err = os.WriteFile("save.json", marshalled, 0644)
		if err != nil {
			fmt.Println(err)
			return -1
		} else { 
			fmt.Println("Succesfully saved contact")
			reloadFile(1) 
			return 0
		}
	}
}

func listContacts(contacts []Contact) {
	for i:= 0; i < len(contacts); i++ {
		fmt.Printf("The name at %v is %v \n", i, contacts[i].Name)
	}

}
func createContact(contacts []Contact) (int) {
	var addit Contact
	fmt.Print("Contact Name: ")
	fmt.Scan(&addit.Name)
	fmt.Print("Age: ")
	fmt.Scan(&addit.Age)
	fmt.Print("Email: ")
	fmt.Scan(&addit.Email)
	// num := len(contacts.Contacts) + 1
	// var ncontacts [num]Contact
	ncontacts := append(contacts, addit)
	// fmt.Println(ncontacts)
	var ei int = saveFile(ncontacts)
	return ei
}

func findPerson(contacts []Contact) {
	var nameI string
	fmt.Println("Who's profile would you like to access?")
	fmt.Scan(&nameI)
	var nameInd int = -1
	for i := 0; i < len(contacts); i++ {
		if nameI == contacts[i].Name {
			nameInd = i
		}
	}
	if nameInd >= 0 {
		cCon := contacts[nameInd]
		fmt.Println("Found the person: ")
		fmt.Printf("\tName: %v \n\tAge: %v \n\tEmail: %v \n", cCon.Name, cCon.Age, cCon.Email) 
	} else {
		fmt.Println("No one by that name")
	}
}
func removeContact(contacts []Contact) {
	var nme string
	fmt.Println("Which contact would you like to remove")
	fmt.Scan(&nme)
	var cntI int = -1
	for i := 0; i < len(contacts); i++ {
		if nme == contacts[i].Name {
			cntI = i
		}
	}
	if cntI >= 0 {
		var newClist []Contact
		//var upd int = cntI + 1
		newClist = slices.Delete(contacts, cntI, cntI+1)
		err := saveFile(newClist)
		if err < -1 {
			fmt.Println("Error while saving the File")
		} else {fmt.Println("Contact list successfully changed and saved.") }
	} else { fmt.Println("Couldnt find contact with that name...") }
}
func menu() {
	var ans int
	fmt.Println("[1] Add New Contact")
	fmt.Println("[2] Remove Contact")
	fmt.Println("[3] List Contacts")
	fmt.Println("[4] Find Contact")
	fmt.Println("[5] Exit")
	fmt.Scan(&ans)
	if ans == 1 {
		state := createContact(Contacts)
		if state < 0 {
			fmt.Println("Error encountered, Please Report")
		} else {
			fmt.Println("Contact Added & Saved")
		}
		menu()
	} else if ans == 2{
		removeContact(Contacts)
		menu()
	} else if ans == 3{
		if len(Contacts) > 0 {
			listContacts(Contacts)
		} else {
			fmt.Println("You have no contacts currently")
		}
		menu()
	} else if ans == 4 {
		findPerson(Contacts)
		menu()
	} else if ans == 5{
		os.Exit(0)
	} else {
		fmt.Println("Invalid input")
		menu()
	}
}

func reloadFile(run int) {
	//if run < 0 { jsonFile.Close() }
	jsonFile, err := os.Open("save.json")
	if err != nil {
	 	fmt.Println(err)
	}
	defer jsonFile.Close()
	// Reading json into byteArray
	byteValue, _ := ioutil.ReadAll(jsonFile)
	// Converting into readable form and appending to contacts array
	json.Unmarshal(byteValue, &Contacts)
}

func main() {
	reloadFile(0)
	menu()
}
