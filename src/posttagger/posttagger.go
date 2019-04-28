package posttagger

import (
	"fmt"
	"nlp/src/dll"
	"regexp"
	"strings"
)

type Kata struct {
	Nama    string
	Grammar []string
	Susunan string
}

//rule
var subjek = []string{"KG", "KB", "FN", "SNK"}                                                   // pengganti nomina, nomina, Frasa Nominal
var predikat = []string{"KB", "KK", "KKI", "KKT", "KS", "BLP", "OD", "FPre", "KT", "SRP", "TLK"} //Kata Tanya, nomina, verba, adjektiva, Numeralia, Frasa Preposisional
var objek = []string{"KB", "KG", "SRP"}                                                          //nomina, pengganti nomina
var pel = []string{"KB", "KK", "KS", "KBil", "BLP", "OD", "FPre", "SMB"}                         //nomina, verba, adjektiva, numeralia, frasa preposisional
var keterangan = []string{"KET", "SFT", "NGR", "ARH", "HR", "BN"}                                //Keterangan, kata sifat, ARAH
var KTugas = []string{"KD", "KP", "KS", "Ksan", "SYM"}                                           //Kata Depan, Kata Penghubung, Kata Seru, Kata Sandang

func News(words []string, dataset []string, grammar []string) []Kata {
	var wordhasil []Kata
	wordhasil = grammar_in(words, dataset, grammar)
	wordhasil = cek_kata_dasar(wordhasil)
	// fmt.Println(wordhasil)
	wordhasil = susunan_kalimat(wordhasil)
	return wordhasil
}

func cek_kata_dasar(words []Kata) []Kata {
	for i := 0; i < len(words); i++ {
		// fmt.Println(len(words))
		// fmt.Println(words[i].Nama)
		if len(words[i].Grammar) == 1 {
			// return words
			continue
		} else if len(words) == 2 && dll.Match_array(words[i].Grammar, "KKI") {
			for j := 0; j < len(words[i].Grammar); j++ {
				words[i].Grammar[j] = ""
			}
			words[i].Grammar = words[i].Grammar[:len(words[i].Grammar)-1]
			words[i].Grammar[0] = "KKI"
			continue
		} else if dll.Match_array(words[i].Grammar, "KK") {
			words[i].Grammar = words[i].Grammar[:len(words[i].Grammar)-1]
			words[i].Grammar[0] = "KK"
			continue
		}

		if i == 0 && dll.Match_array(words[i].Grammar, "KG") {
			for k := 0; k < len(words[i].Grammar); k++ {
				words[i].Grammar = words[i].Grammar[:len(words[i].Grammar)-1]
			}
			words[i].Grammar[0] = "KG"
			continue
		}

		if search_grammar(words[i], "KET") && len(words) > 2 && (i+1) <= len(words) {
			if search_grammar(words[(i+1)], "HR") {

				for k := 1; k < len(words[i].Grammar); k++ {
					words[i].Grammar = words[i].Grammar[:len(words[i].Grammar)-1]
				}
				// fmt.Print((i + 1))
				words[i].Grammar[0] = "KET"
				continue
			}
		}
	}
	return words
}

func susunan_kalimat(words []Kata) []Kata {
	var temp []string
	var tempsub string = ""
	stats := dll.Start_array_bool(5)
	for i := 0; i < len(words); i++ {
		if len(words[i].Grammar) < 1 {
			for y := 0; y < len(words); y++ {
				words[y].Susunan = ""
			}
			return words
		} else {

			if dll.Match_array(subjek, words[i].Grammar[0]) && stats[0] == false {
				if i == 0 {
					words[i].Susunan = "Subjek"
					stats[0] = true
					continue
				} else if words[i-1].Grammar[0] == "KD" {
					words[i-1].Susunan = "Keterangan"
					words[i].Susunan = "Keterangan"
					stats[3] = true
					continue
				} else if words[i-1].Grammar[0] == "KT" {
					words[i-1].Susunan = "Predikat"
					words[i].Susunan = "Subjek"
					stats[0] = true
					stats[1] = true
					continue
				} else if len(words) > 2 && (i+1) < 2 {
					fmt.Println((i + 1))
					if words[(i + 1)].Grammar[0] == "KG" {
						words[(i + 1)].Susunan = "Subjek"
						words[i].Susunan = "Subjek"
						stats[0] = true
						continue
					}
				} else {
					words[i].Susunan = "Subjek"
					stats[0] = true
					temp = append(temp, words[i].Grammar[0])
				}
			}

			if dll.Match_array(predikat, words[i].Grammar[0]) && stats[0] && stats[1] == false {
				regex, _ := regexp.MatchString("di.*", words[i].Nama)
				if dll.Match_array(temp, words[i].Grammar[0]) {
					words[i].Susunan = "Predikat"
					stats[1] = true
					continue
				} else if regex {
					words[i].Susunan = "Predikat"
					if len(words) > 2 {
						search := search_words(words, "Subjek", "susunan")
						searcha := search_words(words, "Objek", "susunan")
						if len(search) != 0 {
							for y := 0; y < len(search); y++ {
								// fmt.Println(search[y])
								words[search[y]].Susunan = "Objek"
							}
							tempsub = "subjek"
							stats[1] = true
							continue
						} else if len(searcha) != 0 {
							for y := 0; y < len(search); y++ {
								words[search[y]].Susunan = "Subjek"
							}
							stats[1] = true
							continue
						}
					}
				} else if i != 0 {
					words[i].Susunan = "Predikat"
					stats[1] = true
					continue
				}
			}

			if dll.Match_array(objek, words[i].Grammar[0]) && len(words) > 2 && stats[2] == false && stats[1] && stats[0] {
				// fmt.Println(tempsub)
				if tempsub == "subjek" {
					words[i].Susunan = "Subjek"
					stats[2] = true
					continue
				} else if dll.Match_array(temp, words[i].Grammar[0]) {
					words[i].Susunan = "Objek"
					stats[2] = true
					continue
				} else if i != 0 {
					words[i].Susunan = "Objek"
					stats[2] = true
				}
			}

			if dll.Match_array(keterangan, words[i].Grammar[0]) && len(words) > 2 && stats[3] == false && stats[1] && stats[2] && stats[0] {
				if len(words) > 2 && (i+1) < len(words) {
					if search_grammar(words[(i+1)], "HR") {
						words[i].Susunan = "Keterangan"
						words[(i + 1)].Susunan = "Keterangan"
						stats[3] = true
						continue
					}
				} else {
					words[i].Susunan = "Keterangan"
					stats[3] = true
					continue
				}
			}
		}
	}
	return words
}

func grammar_in(words []string, dataset []string, grammar []string) []Kata {
	grammarhasil := []Kata{}
	for i := 0; i < len(words); i++ {
		var array []string
		for k := 0; k < len(dataset); k++ {
			if dataset[k] == words[i] { //jika sama kata dgn dataset database
				if len(array) == 0 {
					array = append(array, grammar[k])
				} else {
					for j := 0; j < len(array); j++ {
						if array[j] != grammar[k] {
							array = append(array, grammar[k])
						}
					}
				}

			}
		}
		grammarhasil = append(grammarhasil, Kata{Nama: words[i], Grammar: array})
	}

	return grammarhasil
}

func search_words(words []Kata, matcher string, pattern string) []int {
	var index []int
	for i := 0; i < len(words); i++ {
		if pattern == "susunan" {
			if strings.ToLower(words[i].Susunan) == strings.ToLower(matcher) {
				index = append(index, i)
			}
		}
	}
	// fmt.Println(words)
	return index
}

func search_grammar(words Kata, matcher string) bool {
	var index bool = false
	for i := 0; i < len(words.Grammar); i++ {
		if strings.ToLower(words.Grammar[i]) == strings.ToLower(matcher) {
			index = true
		}
	}
	// fmt.Println(words)
	return index
}
