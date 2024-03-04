package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"

	"github.com/tebeka/selenium"
)

func Scrape_Company(driver selenium.WebDriver, Company string) (string, error) {
	suppressErrors() //i dont know how to fix the errors theyre in the walls
	About_Page := "https://www.linkedin.com/company/" + Company + "/about"
	abpge := driver.Get(About_Page)
	if abpge != nil {
		return "cant find people page", abpge
	}

	webSpan, err := driver.FindElement(selenium.ByXPATH, "//span[@class='link-without-visited-state']")
	if err != nil {
		return "", err
	}
	// Get the text of the web span element
	webURL, _ := webSpan.Text()

	fmt.Println(webURL)
	return webURL, nil

}

func Infinite_Scroll(driver selenium.WebDriver, Company string) {
	suppressErrors() //i dont know how to fix the errors theyre in the walls
	var People_Page = "https://www.linkedin.com/company/" + Company + "/people"
	pplpge := driver.Get(People_Page)
	if pplpge != nil {
		return
	}
	fmt.Println("Performing infinite scroll")
	jsScript := "return document.body.scrollHeight"
	lastHeight, _ := driver.ExecuteScript(jsScript, nil)

	for {
		// Scroll down to bottom
		driver.ExecuteScript("window.scrollTo(0, document.body.scrollHeight);", nil)

		time.Sleep(time.Duration(rand.Intn(4)+5) * time.Second)

		newHeight, _ := driver.ExecuteScript(jsScript, nil)

		if newHeight == lastHeight {
			_, err := driver.FindElement(selenium.ByXPATH, "//button[contains(text(), 'Show more results')]")
			if err != nil {
				break
			}
		}
		lastHeight = newHeight
	}

	peopleData, _ := ScrapePeoplePage(driver)

	for _, person := range peopleData {
		fmt.Printf("Name: %s\n", person["Name"])
		fmt.Printf("Title: %s\n", person["Title"])
		fmt.Printf("OpenToWork: %s\n", person["OpenToWork"])
	}
}

func ScrapePeoplePage(driver selenium.WebDriver) ([]map[string]string, error) {
	suppressErrors() //i dont know how to fix the errors theyre in the walls
	var peopleData []map[string]string

	elements, _ := driver.FindElements(selenium.ByCSSSelector, "div[class='org-people-profile-card__profile-info']")

	// Iterate over each element and extract information
	for _, element := range elements {
		person := make(map[string]string)
		Name_element, _ := element.FindElement(selenium.ByCSSSelector, "div.ember-view.lt-line-clamp.lt-line-clamp--single-line.org-people-profile-card__profile-title.t-black")
		person["Name"], _ = Name_element.Text()

		titleElement, _ := element.FindElement(selenium.ByCSSSelector, "div.t-14.t-black--light.t-normal div.ember-view.lt-line-clamp.lt-line-clamp--multi-line")
		person["Title"], _ = titleElement.Text()

		person["OpenToWork"] = strconv.FormatBool(strings.Contains(person["ProfilePic"], "profile-framed"))

		peopleData = append(peopleData, person)
	}

	return peopleData, nil
}
