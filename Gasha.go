package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Item struct {
	Name   string
	Rarity string
}

var gachaPool = []Item{
	{"アイテムA", "ノーマル"},
	{"アイテムB", "ノーマル"},
	{"アイテムC", "ノーマル"},
	{"アイテムD", "レア"},
	{"アイテムE", "レア"},
	{"アイテムF", "レア"},
	{"アイテムG", "スーパーレア"},
}

func main() {
	rand.Seed(time.Now().UnixNano())

	fmt.Println("ガチャを回しますか？ (1 or 10) もしくは提供割合を表示するコマンド（rates）を入力してください。")
	var input string
	fmt.Scan(&input)

	if input == "1" {
		item := drawSingleGacha()
		fmt.Printf("アイテム: %s (%s)\n", item.Name, item.Rarity)

		if item.Rarity == "スーパーレア" {
			fmt.Println("おめでとうございます！スーパーレアを引きました！")
		}
	} else if input == "10" {
		items := drawTenGacha()
		fmt.Println("10連ガチャの結果:")
		for _, item := range items {
			fmt.Printf("アイテム: %s (%s)\n", item.Name, item.Rarity)
		}
	} else if input == "rates" {
		displayRates()
	} else {
		fmt.Println("無効な選択です。1または10、もしくはratesを入力してください。")
	}
}

func drawSingleGacha() Item {
	probability := rand.Float64()
	if probability < 0.01 {
		return Item{"スーパーレアアイテム", "スーパーレア"}
	} else if probability < 0.11 {
		return Item{"レアアイテム", "レア"}
	}

	return Item{"ノーマルアイテム", "ノーマル"}
}

func drawTenGacha() []Item {
	var results []Item
	for i := 0; i < 10; i++ {
		item := drawSingleGacha()
		results = append(results, item)
	}
	return results
}

func displayRates() {
	fmt.Println("各レアリティの提供割合:")
	fmt.Println("スーパーレア: 1%")
	fmt.Println("レア: 10%")
	fmt.Println("ノーマル: 89%")
}
