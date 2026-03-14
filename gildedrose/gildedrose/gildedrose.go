package gildedrose

type Item struct {
	Name            string
	SellIn, Quality int
}

const (
	sulfuras = "Sulfuras, Hand of Ragnaros"
	agedBrie = "Aged Brie"
	concert  = "Backstage passes to a TAFKAL80ETC concert"
	elixir   = "Elixir of the Mongoose"
	conjured = "Conjured Mana Cake"
	vest     = "+5 Dexterity Vest"
)

// isNormalDecreasingQualityProduct reports whether the given product
// follows the "normal" decreasing-in-quality behavior. Special products
// (Aged Brie, Backstage passes, Sulfuras) do not follow the normal
// decreasing rules.
func isNormalDecreasingQualityProduct(product string) bool {
	switch product {
	case agedBrie, concert, sulfuras:
		return false
	default:
		return true
	}
}
func applyConcertQualityRules(item *Item) {
	if item.Name == concert {
		if item.SellIn < 11 {
			item.Quality++
		}
		if item.SellIn < 6 {
			item.Quality++
		}
	}
}

func applyPassedSellInRules(item *Item) {
	if item.Name == agedBrie {
		if item.Quality < 50 {
			item.Quality++
		}
	} else if item.Name == concert {
		item.Quality = 0
	} else if item.Name != sulfuras {
		if item.Quality > 0 {
			item.Quality--
			if item.Name == conjured && item.Quality > 0 {
				item.Quality--
			}
		}
	}
}

func applyQualityRules(item *Item) {
	if isNormalDecreasingQualityProduct(item.Name) {
		if item.Quality > 0 {
			item.Quality--
			if item.Name == conjured {
				item.Quality--
			}
		}
	} else {
		if item.Quality < 50 {
			item.Quality++
			applyConcertQualityRules(item)
		}
	}
}

func UpdateQuality(items []*Item) {
	for _, item := range items {
		applyQualityRules(item)

		if item.Name != sulfuras {
			item.SellIn--
		}
		if item.SellIn < 0 {
			applyPassedSellInRules(item)
		}
	}
}
