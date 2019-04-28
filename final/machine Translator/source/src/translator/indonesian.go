package translator

import (
	"fmt"
	"nlp/src/dll"
	"regexp"
	"strings"
)

func (tl WordsTL) Translate_ID(words []string) []string {

	var temp []string

	tenses := tl.cek_tenses_EN(words)
	fmt.Println(tenses)

	for _, e := range words {
		temp = append(temp, tl.search_translate_EN(e))
	}

	// co := 0
	for x, y := range temp {
		// matcher := []string{"imbuhan", "Kdepan", "partikel", "penjelas"}

		if tenses == "presentcontinues" {
			if tl.tags[x] == "tobe" {
				temp[x] = strings.Replace(y, y, "sedang", -1)
			}
			if tl.tags[x] == "verb" {
				split := []string{"es", "s", "ing", "ed"}
				for n, m := range split {
					regex, _ := regexp.MatchString("!?"+m, y)
					// fmt.Print(regex)
					if regex {
						y = strings.Replace(y, split[n], "", -1)
						y = strings.Replace(y, "\"", "", -1)
						y = tl.search_translate_EN(y)
						temp[x] = strings.Replace(y, y, "me"+y, -1)
						// fmt.Println(y)
					}

				}

			}
		}

		if tl.tags[x] == "verb" {
			break
		}

	}

	return temp

}

func (tl WordsTL) cek_tenses_EN(words []string) string {

	// matcher := []string{"am", "is", "are"}
	var tenses string
	var vaa string
	for x, e := range tl.tags {

		split := []string{"es", "s", "ing", "ed"}
		if e != "verb" && len(words)-1 > x {
			for _, m := range split {
				regex, _ := regexp.MatchString("!?"+m+"", words[x+1])
				// fmt.Print(regex)
				if regex && tl.tags[x+1] == "verb" {
					vaa = m
				}

			}
		}

		// fmt.Println(vaa)
		// fmt.Println(e)

		if e == "verb" {
			break
		}
	}

	if vaa == "ing" {
		tenses = "presentcontinues"
	} else {
		tenses = "simplepresent"
	}

	return tenses
}

func (tl WordsTL) search_translate_EN(kata string) string {
	index := dll.Search_array(tl.id, kata)
	// fmt.Println(tl.id)
	if index == -1 {
		return "\"" + kata + "\""
	}
	// fmt.Println(tl)
	return tl.en[index]
}
