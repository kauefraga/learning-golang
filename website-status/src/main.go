package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func getStatus(url string) {
	fmt.Println("> Scanning", url)

	response, error := http.Head(url)

	if error != nil {
		fmt.Println("Something went wrong.", error)
		os.Exit(-1)
	}

	if response.StatusCode == 200 {
		fmt.Println("[monitor] The website is up!")
	} else {
		fmt.Println("[monitor] The website is down!", response.StatusCode)
	}

	fmt.Println("")
}

func main() {
	const version float32 = 1.1

	fmt.Println("Version:", version)

	for {
		fmt.Println("[1] Start monitoration")
		fmt.Println("[2] Show logs")
		fmt.Println("[0] Exit")

		var option int
		fmt.Scan(&option)

		if option == 1 {
			fmt.Println("\033[1m> Monitoring...\033[0m")

			urls := [3]string{
				"https://httpstat.us/Random/200,203,404,500",
				"https://httpstat.us/Random/200",
				"https://httpstat.us/Random/200,201,500",
			}

			for i := 0; i < 3; i++ {
				for _, url := range urls {
					getStatus(url)
				}

				if i < 2 {
					fmt.Println("We are going to scan again. Wait 10 seconds.")
					time.Sleep(10 * time.Second)
				}
			}
		}

		if option == 2 {
			fmt.Println("\033[1m> Fetching logs...\033[0m")
			fmt.Println("[logger] uibgaiiguurbgueigbre")
			fmt.Println("[logger] failed fonsugbariugd")
		}

		if option == 0 {
			fmt.Println("\033[90m> Exiting...\033[0m")
			fmt.Println("\033[90m---------------\033[0m")
			os.Exit(0)
		}
	}
}
