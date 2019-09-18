package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/domnikl/ifttt-webhook"
	"log"
	"net/http"
	"strings"
)

func LegoScrape() {

	key := ""
	ifttt := iftttWebhook.New(key)

	// Request the HTML page.
	url := "https://www.lego.com/en-us/product/central-perk-21319"
	//url := "https://www.lego.com/en-us/product/disney-train-and-station-71044"
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
	}

	// Load the HTML document
	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
	}

	// Find the review items
	boo := doc.Find("html body div#root div main#main-content.MinHeightContent-sc-1delze2-0.cFqXKV div div.ProductDetailsPagestyles__ProductOverviewContainer-sc-1waehzg-1.kvERzT div.sharedstyles__MaxWidthContainer-sc-3k7bob-0.dMLaFt div.ProductDetailsPagestyles__ProductOverviewLayout-sc-1waehzg-2.cmPKzB div.ProductOverviewstyles__Section-sc-1a1az6h-0.bgJmVZ p.ProductOverviewstyles__AvailabilityStatus-sc-1a1az6h-4.ACco.Text__BaseText-aa2o0i-0.bHappg span.Markup__StyledMarkup-ar1l9g-0.bTYWAd")
	if boo.Contents().Length() > 0 {
		boo.Each(func(i int, s *goquery.Selection) {
			if strings.Contains(strings.ToLower(s.Text()), "out") {
				fmt.Println("boo")
				//ifttt.Emit("scraper", "Lego Central Perk 21319", url, ":(")
			}
		})
	}

	yay := doc.Find("html body div#root div main#main-content.MinHeightContent-sc-1delze2-0.cFqXKV div div.ProductDetailsPagestyles__ProductOverviewContainer-sc-1waehzg-1.kvERzT div.sharedstyles__MaxWidthContainer-sc-3k7bob-0.dMLaFt div.ProductDetailsPagestyles__ProductOverviewLayout-sc-1waehzg-2.cmPKzB div.ProductOverviewstyles__Section-sc-1a1az6h-0.bgJmVZ p.ProductOverviewstyles__AvailabilityStatus-sc-1a1az6h-4.ACco.Text__BaseText-aa2o0i-0.khMRnm span.Markup__StyledMarkup-ar1l9g-0.bTYWAd")
	if yay.Contents().Length() > 0 {
		yay.Each(func(i int, s *goquery.Selection) {
			// For each item found, get the band and title
			if strings.Contains(strings.ToLower(s.Text()), "available") {
				fmt.Println(s.Text())
				ifttt.Emit("scraper", "Lego Central Perk 21319", url, ":)")
			}
		})
	}
}

func main() {
	LegoScrape()
}
