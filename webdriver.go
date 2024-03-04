package main

import (
	"fmt"
	"os/exec"
	"runtime"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func startdriver(CompanyPage string) {
	Login_Page := "https://www.linkedin.com/login/"
	About_Page := CompanyPage
	var chromedriverPath string
	switch runtime.GOOS {
	case "windows":
		chromedriverPath = "./chromedriver.exe"
	case "linux":
		chromedriverPath = "./chromedriver"
	case "darwin":
		chromedriverPath = "./chromedriver_mac"
	}

	service, err := selenium.NewChromeDriverService(chromedriverPath, 4444)
	if err != nil {
		panic(err)
	}
	defer service.Stop()

	caps := selenium.Capabilities{}
	caps.AddChrome(chrome.Capabilities{Args: []string{
		"window-size=1920x1080",
		"--no-sandbox",
		"--disable-dev-shm-usage",
		"disable-gpu",
		"--ignore-certificate-errors",
		"--disable-3d-apis",
		"--log-level=3",
	}})

	driver, _ := selenium.NewRemote(caps, "")

	driver.Get(Login_Page)
	Search_Login_Page(driver)
	webURL, _ := Scrape_Company(driver, About_Page)
	Infinite_Scroll(driver, CompanyPage)
	ScrapePeoplePage(driver)
	peopleData, _ := ScrapePeoplePage(driver)
	fmt.Println("writing to CSV")
	WriteToCSV(peopleData, CompanyPage+".csv", webURL, CompanyPage)
}
