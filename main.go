package main

import (
	"flag"
	"fmt"
)

func main() {
	var Company string
	flag.StringVar(&Company, "c", "", "The linkedin company you want to scrape. EG: https://linkedin.com/company/COMPANYNAME/")
	flag.StringVar(&Company, "CompanyPage", "", "The linkedin company you want to scrape. EG: https://linkedin.com/company/COMPANYNAME/")
	flag.Parse()
	var chromedriverPath string
	flag.StringVar(&chromedriverPath, "d", "", "Path to the Your chromedriver eg: /bin/epic/chromedriver")
	flag.StringVar(&chromedriverPath, "driver", "", "Path to the Your chromedriver eg: /bin/epic/chromedriver")
	flag.Parse()

	fmt.Println("	PISTOLFISH\n	Just a silly little scraper.")
	if Company == "" {
		fmt.Println("Wheres your Company??")
		return
	}

	fmt.Println("CompanyPage:", Company)

	startdriver(Company, chromedriverPath)
}
