package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

func Search_Login_Page(driver selenium.WebDriver) {
	suppressErrors()
	fmt.Println("Waiting for feed to load")
	i := 0
	for {
		currentURL, _ := driver.CurrentURL()

		fmt.Printf("%s\r", strings.Repeat(".", i))
		i++
		time.Sleep(1 * time.Second) //its my first day
		if i == 10000 {
			fmt.Println("Checking for Captcha or incorrect credentials")

		}
		if currentURL == "https://www.linkedin.com/feed/" {
			fmt.Println("Feed loaded!")
			break
		}
	}
	fmt.Println("Feed loaded!")

}
