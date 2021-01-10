package eorzeadb

import (
	"log"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Item struct {
	URL          string
	Name         string
	Category     string
	EquipmentJob string
	RecipeURL    string
}

type Recipe struct {
	URL          string
	TotalCrafted string
	Materials    [5]Material
	ItemURL      string
}

type Material struct {
	URL  string
	Name string
	Num  string
}

func NewItem(url string) (*Item, error) {
	item := new(Item)

	doc, err := doc(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	item.URL = url
	item.Name = trim(doc.Find("h2.db-view__item__text__name").Text())
	item.Category = doc.Find("p.db-view__item__text__category").Text()
	item.EquipmentJob = doc.Find("div.db-view__item_equipment__class").Text()
	recipeURL, _ := doc.Find("a.db_popup.db-table__txt--detail_link").Attr("href")
	item.RecipeURL = "https://jp.finalfantasyxiv.com" + recipeURL

	return item, nil
}

func NewRecipe(url string) (*Recipe, error) {
	recipe := new(Recipe)

	doc, err := doc(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	recipe.URL = url
	recipe.TotalCrafted = trim(doc.Find("#eorzea_db > div.clearfix > div.db_cnts > div > div.recipe_detail.item_detail_box > div.db-view__data > ul > li:nth-child(1)").Text())
	doc.Find("#eorzea_db > div.clearfix > div.db_cnts > div > div.recipe_detail.item_detail_box > div.db-view__data > div:nth-child(3) > div > div").Each(func(i int, s *goquery.Selection) {
		dataKey, _ := s.Attr("data-key")
		recipe.Materials[i].URL = "https://jp.finalfantasyxiv.com/lodestone/playguide/db/item/" + dataKey + "/"
		recipe.Materials[i].Name, _ = s.Attr("data-name")
		recipe.Materials[i].Num, _ = s.Attr("data-num")
	})
	itemURL, _ := doc.Find("#eorzea_db > div.clearfix > div.db_cnts > div > div.recipe_detail.item_detail_box > div.db-view__data > div.db-tooltip__bt_item_detail > a").Attr("href")
	recipe.ItemURL = "https://jp.finalfantasyxiv.com" + itemURL

	return recipe, nil
}

func doc(url string) (*goquery.Document, error) {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		log.Fatalf("status code error: %d %s", res.StatusCode, res.Status)
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(res.Body)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return doc, nil
}

func trim(str string) string {
	return strings.NewReplacer(
		"\n", "",
		string(rune(0xE03C)), "",
		string(rune(0x0009)), "",
		"完成個数", "",
		string(rune(0x00A0)), "",
	).Replace(str)
}
