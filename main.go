package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const monitorings = 1
const monitoringDelay = 1

func main() {

	showIntro()

	for {

		showMenu()

		option := readOption()

		switch option {

		case 1:

			startMonitoring()

		case 2:

			printLogs()

		case 0:

			fmt.Println("Exiting...")
			os.Exit(0)

		default:

			fmt.Println("Unknown Option")
			os.Exit(-1)

		}

	}

}

func showIntro() {

	fmt.Println("")
	fmt.Println("Welcome!")

}

func showMenu() {

	fmt.Println("")
	fmt.Println("1 - Start Monitoring")
	fmt.Println("2 - Show Logs")
	fmt.Println("0 - Exit")

}

func readOption() int {

	var option int

	fmt.Print("Enter your option: ")
	fmt.Scan(&option)
	fmt.Println("Selected option: ", option)

	return option

}

func startMonitoring() {

	fmt.Println("Monitoring...")

	sites := readSitesFromFile()

	for i := 0; i < monitorings; i++ {

		for _, site := range sites {

			testSite(site)

		}

		fmt.Println("")

		time.Sleep(monitoringDelay * time.Second)

	}

}

func testSite(site string) {

	res, err := http.Get(site)

	if err != nil {

		fmt.Println("An error has occured:", err)

	}

	if res.StatusCode == 200 {

		fmt.Println("Site:", site, "loaded successfully!")
		registerLog(site, true)

	} else {

		fmt.Println("Site:", site, "has errors. Status Code:", res.StatusCode)
		registerLog(site, false)

	}

}

func readSitesFromFile() []string {

	sites := []string{}

	file, err := os.Open("sites.txt")

	if err != nil {

		fmt.Println("An error has occured:", err)

	}

	reader := bufio.NewReader(file)

	for {

		line, err := reader.ReadString('\n')

		line = strings.TrimSpace(line)

		sites = append(sites, line)

		if err == io.EOF {

			break

		}

	}

	file.Close()

	return sites

}

func registerLog(site string, status bool) {

	file, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)

	if err != nil {

		fmt.Println("An error has occured:", err)

	}

	file.WriteString(time.Now().Format("Monday 02/01/2006 15:04:05") + " - " + site + " - online: " + strconv.FormatBool(status) + "\n")

	file.Close()

}

func printLogs() {

	fmt.Println("Showing Logs...")

	file, err := ioutil.ReadFile("log.txt")

	if err != nil {

		fmt.Println("An error has occured:", err)

	}

	fmt.Println(string(file))

}
