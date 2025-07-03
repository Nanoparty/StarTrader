package utils

import (
	"os"
	// "fmt"
	"log"
	"time"
	"math/rand"
	"bufio"
)

func Random_Name_From_File(inputFile string) string {
	wd, _ := os.Getwd()
	
	file, err := os.Open(wd + "/utils/" + inputFile)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}

	var names []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		names = append(names, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatalf("Error reading file: %v", err)
	}

	rand.Seed(time.Now().UnixNano())
	if len(names) == 0 {
		log.Fatal("No names in name file.")
	}
	name := names[rand.Intn(len(names))]
	return name
}

func Generate_Employee_Name() string {
	return Random_Name_From_File("names/names_employees.txt")
}

func Generate_Transport_Ship_Name() string {
	return Random_Name_From_File("names/names_transport_ship.txt")
}

func Generate_Mining_Ship_Name() string {
	return Random_Name_From_File("names/names_mining_ship.txt")
}

func Generate_Combat_Ship_Name() string {
	return Random_Name_From_File("names/names_combat_ship.txt")
}