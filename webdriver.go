package main

import (
	"fmt"
	"runtime"

	"github.com/tebeka/selenium"
	"github.com/tebeka/selenium/chrome"
)

func startdriver(CompanyPage string, chromedriverPath string) {
	Login_Page := "https://www.linkedin.com/login/"
	About_Page := CompanyPage
	if chromedriverPath == "" {
		switch runtime.GOOS {
		case "windows":
			chromedriverPath = "chromedriver.exe"
		case "linux":
			chromedriverPath = "chromedriver"
		case "darwin":
			chromedriverPath = "chromedriver_mac"
		}
	}

	if chromedriverPath == "" {
		fmt.Println("Chromedriver path not provided. Using default path.")
	} else {
		fmt.Println("Using Chromedriver path:", chromedriverPath)
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

	// Create a new WebDriver using the ChromeDriver service
	driver, err := selenium.NewRemote(caps, "")
	if err != nil {
		panic(err)
	}
	defer driver.Quit()

	driver.Get(Login_Page)
	Search_Login_Page(driver)
	webURL, _ := Scrape_Company(driver, About_Page)
	Infinite_Scroll(driver, CompanyPage)
	ScrapePeoplePage(driver)
	peopleData, _ := ScrapePeoplePage(driver)

	fmt.Println("writing to CSV")
	WriteToCSV(peopleData, CompanyPage+".csv", webURL, CompanyPage)
}
