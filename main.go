package main

import (
	"bufio"
	"fmt"
	"os"
	"startrader/client"
	"startrader/globals"
	"strings"
	"time"
)

func intro() {
	fmt.Println("Welcome to...")
	time.Sleep(1 * time.Second)
	fmt.Println("")
	fmt.Print("|  hello world|") // See if the space before "hello" shows

// 	fmt.Println(`  _________ __                 ___________                  .___            
//  /   _____//  |______ _______  \__    ___/___________     __| _/___________ 
//  \_____  \\   __\__  \\_  __ \   |    |  \_  __ \__  \   / __ |/ __ \_  __ \
//    /        \|  |  / __ \|  | \/   |    |   |  | \// __ \_/ /_/ \  ___/|  | \/
// /_______  /|__| (____  /__|      |____|   |__|  (____  /\____ |\___  >__|   
//         \/           \/                              \/      \/    \/        `)
	time.Sleep(3 * time.Second)

	reader := bufio.NewReader(os.Stdin)

	fmt.Println("")
	fmt.Println("What is your name? :")
	player_name, _ := reader.ReadString('\n')
	player_name = strings.TrimSpace(player_name)

	time.Sleep(1 * time.Second)

	fmt.Println("")
	fmt.Println("What is your company name? :")
	company_name, _ := reader.ReadString('\n')
	company_name = strings.TrimSpace(company_name)

	time.Sleep(1 * time.Second)

	fmt.Println("")
	fmt.Printf("In the year 2350, %s began their conquest of the Sol System with the founding of %s", player_name, company_name)
}

func main() {
	cfg := &globals.Config{
		SaveFile: "test",
	}

	client.StartRepl(cfg)
}
