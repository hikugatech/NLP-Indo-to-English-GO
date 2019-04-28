package posttagger

import (
	"fmt"
	"nlp/src/dll"
	"nlp/src/naivebayes"
	"strconv"
	"strings"
)

func (kalimat Kalimat) English(words []string) Kalimat {

	//untuk menyimpan tags sebelumnya
	// var temptags []string
	//loop untuk check error selama 3x
	ctemp := dll.Start_array_int(len(kalimat.gram))
	for i := 0; i < 2; i++ {
		var tags []string

		// start array untuk menentukan nilai kosong
		gramx := dll.Start_array_string(len(kalimat.gram))
		// if len(temptags) > 0 {
		// 	tags = temptags
		// } else {
		tags = kalimat.Tags
		// }

		// setiap kata akan di looping
		for x, _ := range words {

			// mendeklarasikan bayes dengan dataset penempatan tag
			bayes := naivebayes.New("dataset/pstkalEN.csv")
			var match []string

			// jika jenis kata ada lebih dari 1,
			// maka akan di bagi dan dicocokkan dimulai berdasarkan urutan terakhir
			// fmt.Println()
			splits := strings.Split(kalimat.gram[x], ",")
			if len(splits) > 1 && len(splits) >= ctemp[x] {
				gramx[x] = splits[ctemp[x]]
				ctemp[x]++
			} else {
				gramx[x] = kalimat.gram[x]
			}

			// ======= Proses pembagian tags. ========
			// jika kalimat tags dan tags kosong *Proses Sementara*
			// keterangan match(grammar, tags sebelumnya(sifatnya yaitu 0 atau tidak ada), tags selanjutnya(sifatnya yaitu 0 atau tidak ada) )
			if len(tags) == 0 {
				match = []string{gramx[x], strconv.Itoa(0), strconv.Itoa(0)}

				// ini adalah proses jika kalimat.Tags belum terisi (kosong) *Proses Sementara*
				// keterangan match(grammar, tags yang terisi sebelumnya, tags selanjutnya(sifatnya yaitu 0 atau tidak ada) )
			} else if len(kalimat.Tags) == 0 {
				match = []string{gramx[x], tags[x-1], strconv.Itoa(0)}

				// ini adalah proses jika kalimat.Tags udah terisi namun index (kata) harus bernilai 0 *Proses Fixing*
				// keterangan match(grammar, tags sebelumnya (sifatnya yaitu 0 atau tidak ada), tags selanjutnya yang sudah terisi )
			} else if (len(kalimat.Tags)-1) > x && x < 1 {
				match = []string{gramx[x], strconv.Itoa(0), strconv.Itoa(0)}
				// fmt.Println(loctags[iloc])
			} else if (len(kalimat.Tags) - 1) > x {
				match = []string{gramx[x], tags[x-1], strconv.Itoa(0)}
				// fmt.Println(loctags[iloc-1])
			} else {
				match = []string{gramx[x], tags[x-1], strconv.Itoa(0)}
				// fmt.Println(loctags[iloc-1])
			}

			// fmt.Println(match)

			if len(kalimat.Tags) == 0 {
				//memasukkan match kedalam rumus Bayes jika kalimat.Tags kosong
				tags = append(tags, bayes.X(match))
			} else {
				//memasukkan match kedalam rumus Bayes jika kalimat.Tags terisi
				// tags[x] = bayes.X(match)
			}

			// fmt.Println(tags)

			if tags[x] == "verb" {
				break
			}

		}
		fmt.Println(gramx)
		// fmt.Println(tags)
		// memasukkan kedalam struct
		kalimat = Kalimat{kalimat.gram, tags}

		// calculate error
		// check := kalimat.checking_posttag_error_EN()
		check := 0.00
		// fmt.Println(check)

		//jika tidak ada error maka akan lanjut ke proses selanjutnya
		if check == 0.00 {
			break
		}

	}

	// pada proses ini, data set tidak bisa menjangkau untuk pembalikan kata jika ada "di" di predikat
	// maka kata tugas disini akan dirubah jadi predikat dan dibalik untuk subjek dan objek
	// jika adanya jumlah katanya kurang dari 3 maka kata tugas akan dirubah saja jadi predikat
	tags := kalimat.Tags

	// memasukkan kedalam struct
	kalimat = Kalimat{kalimat.gram, tags}

	return kalimat
}

func (kalimat Kalimat) checking_posttag_error_EN() float32 {

	// ini adalah hasil akhir (output) dari dataset posttager bayes
	checker := []string{"subjek", "verb"}

	// semua harus bernilai 0
	count_checker := dll.Start_array_int(len(checker))

	// perhitungannya yaitu setiap checker di cocokkan dengan tags. jika cocok maka count checker +1
	// index dari count checcker sesuai dengan checker
	for i, e := range checker {
		for _, y := range kalimat.Tags {
			y = strings.Trim(y, " ")
			e = strings.Trim(y, "\t")
			if y == e {
				count_checker[i]++
			}
		}
	}

	// menghitung seluruh tags
	var counts = len(kalimat.Tags)
	var err = 0

	// ini adalah pengecekan errornya. jika tag tidak sesuai dengan rule,
	// maka akan di tambah variable err nya
	fmt.Println(count_checker)
	for y, e := range count_checker {
		if e > 1 && checker[(len(checker)-1)] == kalimat.Tags[y] {
			continue
		} else if e > 1 {
			err++
		}
	}

	// hasilnya di jadikan float dikarenakan untuk melihat angka errornya
	// keterangan hasil (0-1)
	hasil := float32(err) / float32(counts)

	return hasil

}
