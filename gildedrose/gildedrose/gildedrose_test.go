package gildedrose

import (
	"testing"
)

func Test_Foo(t *testing.T) {
	var items = []*Item{
		{"+5 Dexterity Vest", 10, 20},
		{"Aged Brie", 2, 0},
		{"Elixir of the Mongoose", 5, 7},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
		{"Backstage passes to a TAFKAL80ETC concert", 10, 49},
		{"Backstage passes to a TAFKAL80ETC concert", 5, 49},
		{"Conjured Mana Cake", 3, 45}, // <-- :O
	}
	var want = []*Item{
		{"+5 Dexterity Vest", -5, 0},
		{"Aged Brie", -13, 28},
		{"Elixir of the Mongoose", -10, 0},
		{"Sulfuras, Hand of Ragnaros", 0, 80},
		{"Backstage passes to a TAFKAL80ETC concert", 0, 50},
		{"Backstage passes to a TAFKAL80ETC concert", -5, 0},
		{"Backstage passes to a TAFKAL80ETC concert", -10, 0},
		{"Conjured Mana Cake", -12, 0},
	}

	for i := 0; i < 15; i++ {
		UpdateQuality(items)
	}

	for i, item := range items {
		wantItem := want[i]
		if item.Name != wantItem.Name {
			t.Errorf("Name: Expected name failed, got '%s', want '%s'", item.Name, wantItem.Name)
		}
		if item.Quality != wantItem.Quality {
			t.Errorf("Quality: Expected quality for '%s' failed, got '%d', want '%d'", item.Name, item.Quality, wantItem.Quality)
		}
		if item.SellIn != wantItem.SellIn {
			t.Errorf("Sellin: Expected sellin for '%s' failed, got %d, want %d", item.Name, item.SellIn, wantItem.SellIn)
		}
	}
}
