package homework1

import (
	"fmt"
	"os"
	"strconv"
)

func SortClients(id int, name string, age int) error {
	a, erra := os.OpenFile("adults.txt", os.O_APPEND|os.O_WRONLY, 0600)
	t, errt := os.OpenFile("teenagers.txt", os.O_APPEND|os.O_WRONLY, 0600)
	defer a.Close()
	defer t.Close()
	if erra != nil {
		fmt.Println("Error is ", erra)
		return erra
	}
	if errt != nil {
		fmt.Println("Error is ", errt)
		return errt
	}
	if age > 16 {
		_, err := a.WriteString(string(strconv.Itoa(id) + " " + name + " " + strconv.Itoa(age) + "\n"))
		return err
	} else {
		_, err := t.WriteString(string(strconv.Itoa(id) + " " + name + " " + strconv.Itoa(age) + "\n"))
		return err
	}
}

func GetClients() {
	adults, err := os.ReadFile("adults.txt")
	if err != nil {
		fmt.Println("Error while reading adults.txt")
	}
	teen, err := os.ReadFile("teenagers.txt")
	if err != nil {
		fmt.Println("Error while reading teenagers.txt")
	}
	fmt.Println("------ Adults ------")
	fmt.Println(string(adults))
	fmt.Println("------ Teenagers ------")
	fmt.Println(string(teen))
}