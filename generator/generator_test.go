package generator

import (
	"testing"

	"github.com/ikawaha/kagome/v2/filter"
)

func TestKatakanaKatsuyou1(t *testing.T) {
	expected := "なんちゃッテ"
	actual := katakanaKatsuyou("なんちゃって", 2)
	t.Log(actual)
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestKatakanaKatsuyou2(t *testing.T) {
	expected := "どうしちゃったノカナ{EMOJI_POS}"
	actual := katakanaKatsuyou("どうしちゃったのかな{EMOJI_POS}", 3)
	t.Log(actual)
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestKatakanaKatsuyou3(t *testing.T) {
	expected := "どうしちゃったのかな{EMOJI_POS}"
	actual := katakanaKatsuyou("どうしちゃったのかな{EMOJI_POS}", 0)
	t.Log(actual)
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestKatakanaKatsuyou4(t *testing.T) {
	expected := "東西南北{EMOJI_POS}"
	actual := katakanaKatsuyou("東西南北{EMOJI_POS}", 2)
	t.Log(actual)
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestInsertPunctuations0(t *testing.T) {
	expected := "どうしちゃったのかな"
	actual := insertPunctuations("どうしちゃったのかな", PunctuationConfig{TargetHinshis: filter.NewPOSFilter(filter.POS{"助動詞"}, filter.POS{"名詞"})})
	t.Log(actual)
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func TestInsertPunctuations2(t *testing.T) {
	expected := "どうしちゃった、のかな"
	actual := insertPunctuations("どうしちゃったのかな", PunctuationConfig{TargetHinshis: filter.NewPOSFilter(filter.POS{"助動詞"}), Rate: 100})
	t.Log(actual)
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}
func TestInsertPunctuations3(t *testing.T) {
	expected := "どうしちゃった、のか、な、"
	actual := insertPunctuations("どうしちゃったのかな", PunctuationConfig{TargetHinshis: filter.NewPOSFilter(filter.POS{"助動詞"}, filter.POS{"助詞"}), Rate: 100})
	t.Log(actual)
	if actual != expected {
		t.Errorf("handler returned unexpected body: got %v want %v", actual, expected)
	}
}

func BenchmarkPunctuation(b *testing.B) {
	for i := 0; i < b.N; i++ {
		insertPunctuations("どうしちゃったのかな", PunctuationConfig{TargetHinshis: filter.NewPOSFilter(filter.POS{"助動詞"}, filter.POS{"助詞"}), Rate: 100})
	}
}
