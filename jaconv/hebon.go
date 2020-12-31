package jaconv

import (
	"unicode/utf8"

	"golang.org/x/exp/utf8string"
)

func ToHebon(kana string) string {
	isOmitted := map[string]bool{
		"aa": false, "ee": false, "ii": false, // i は連続しても省略しない
		"oo": false, "ou": false, "uu": false,
	}

	var hebon string
	var lastHebon string

	i := 0
	for {
		ch := charHebonByIndex(kana, i)
		if ch.Char == "っ" {
			// "っち"
			nextCh := charHebonByIndex(kana, i+1)
			if nextCh.Hebon != "" {
				ch.Hebon = "t"
			}
		} else if ch.Char == "ん" {
			nextCh := charHebonByIndex(kana, i+1)
			if nextCh.Hebon != "" {
				ch.Hebon = "n"
			}
		} else if ch.Char == "ー" {
			// 長音は無視
			ch.Hebon = ""
		}

		if ch.Hebon != "" {
			// 変換できる文字の場合
			if lastHebon != "" {
				// 連続する母音の除去
				joinedHebon := lastHebon + ch.Hebon
				if len(joinedHebon) > 2 {
					joinedHebon = joinedHebon[len(joinedHebon)-2:]
				}
				if isOmitted[joinedHebon] {
					ch.Hebon = ""
				}
			}
			hebon += ch.Hebon
		} else {
			if ch.Char != "ー" {
				// 変換できない文字の場合
				hebon += ch.Char
			}
		}

		lastHebon = ch.Hebon
		i += utf8.RuneCountInString(ch.Char)
		if i >= utf8.RuneCountInString(kana) {
			break
		}
	}

	return hebon
}

type CharHebon struct {
	Char  string
	Hebon string
}

func charHebonByIndex(kana string, index int) CharHebon {
	hebonMap := map[string]string{
		"あ": "a", "い": "i", "う": "u", "え": "e", "お": "o",
		"か": "a", "き": "i", "く": "u", "け": "e", "こ": "o",
		"さ": "a", "し": "i", "す": "u", "せ": "e", "そ": "o",
		"た": "a", "ち": "i", "つ": "u", "て": "e", "と": "o",
		"な": "a", "に": "i", "ぬ": "u", "ね": "e", "の": "o",
		"は": "a", "ひ": "i", "ふ": "u", "へ": "e", "ほ": "o",
		"ま": "a", "み": "i", "む": "u", "め": "e", "も": "o",
		"や": "a", "ゆ": "u", "よ": "o",
		"ら": "a", "り": "i", "る": "u", "れ": "e", "ろ": "o",
		"わ": "a", "ゐ": "i", "ゑ": "e", "を": "o",
		"ぁ": "a", "ぃ": "i", "ぅ": "u", "ぇ": "e", "ぉ": "o",
		"が": "a", "ぎ": "i", "ぐ": "u", "げ": "e", "ご": "o",
		"ざ": "a", "じ": "i", "ず": "u", "ぜ": "e", "ぞ": "o",
		"だ": "a", "ぢ": "i", "づ": "u", "で": "e", "ど": "o",
		"ば": "a", "び": "i", "ぶ": "u", "べ": "e", "ぼ": "o",
		"ぱ": "a", "ぴ": "i", "ぷ": "u", "ぺ": "e", "ぽ": "o",
		"きゃ": "ya", "きゅ": "yu", "きょ": "yo",
		"しゃ": "ya", "しゅ": "yu", "しょ": "yo",
		"ちゃ": "ya", "ちゅ": "yu", "ちょ": "yo", "ちぇ": "ye",
		"にゃ": "ya", "にゅ": "yu", "にょ": "yo",
		"ひゃ": "ya", "ひゅ": "yu", "ひょ": "yo",
		"みゃ": "ya", "みゅ": "yu", "みょ": "yo",
		"りゃ": "ya", "りゅ": "yu", "りょ": "yo",
		"ぎゃ": "ya", "ぎゅ": "yu", "ぎょ": "yo",
		"じゃ": "ja", "じゅ": "yu", "じょ": "yo",
		"びゃ": "ya", "びゅ": "yu", "びょ": "yo",
		"ぴゃ": "ya", "ぴゅ": "yu", "ぴょ": "yo",
	}

	var hebon string
	var char string
	utfstr := utf8string.NewString(kana)
	// 2文字ヒットするとき
	if index+1 < utf8.RuneCountInString(kana) {
		char = utfstr.Slice(index, index+2)
		hebon = hebonMap[char]
	}
	// 2文字はヒットしないが1文字はヒットするとき
	if hebon == "" && index < utfstr.RuneCount() {
		char = utfstr.Slice(index, index+1)
		hebon = hebonMap[char]
	}
	return CharHebon{Char: char, Hebon: hebon}
}
