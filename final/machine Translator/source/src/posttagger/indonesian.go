package posttagger

import (
	"nlp/src/dll"
	"nlp/src/naivebayes"
	"strconv"
	"strings"
)

/*
==============================================================================================

	Pada Proses ini. untuk menentukan jenis kalimat SPOK (Subjek, Predikan, Objek, Keterangan)
	maka akan dilakukan beberapa proses yaitu

	1. mengisi kalimat.Tags secara sementara
	2. pengecheckan untuk mengetahui tingkat error penempatan jenis kalimat yang salah
	   (Aku cuma meniru gaya JST (Neural Network) untuk pengecheckan)
	3. Fixing.

	Dasarannya ini berasal dari rule base bahasa indonesia yang di tambah Naive bayes untuk
	mempermudah saja ;)

==============================================================================================
*/

func (kalimat Kalimat) Indonesian(words []string) Kalimat {

	//untuk menyimpan tags sebelumnya
	// var temptags []string
	var loctags []int
	//loop untuk check error selama 3x
	for i := 0; i < 2; i++ {
		var tags, gram []string

		// start array untuk menentukan nilai kosong
		ctemp := dll.Start_array_int(len(kalimat.gram))
		// if len(temptags) > 0 {
		// 	tags = temptags
		// } else {
		tags = kalimat.Tags
		// }

		gram = kalimat.gram
		iloc := 0

		// setiap kata akan di looping
		for x, _ := range words {
			if len(loctags) > 0 {
				if x != loctags[iloc] {
					continue
				}
			}

			// fmt.Println(loctags[iloc])

			// mendeklarasikan bayes dengan dataset penempatan tag
			bayes := naivebayes.New("dataset/pstkalID.csv")
			var match []string

			// jika jenis kata ada lebih dari 1,
			// maka akan di bagi dan dicocokkan dimulai berdasarkan urutan terakhir
			splits := strings.Split(gram[x], ",")
			if len(splits) > 1 {
				ctemp[x] = (len(splits) - 1) - i
				gram[x] = splits[ctemp[x]]
			}

			// ======= Proses pembagian tags. ========
			// jika kalimat tags dan tags kosong *Proses Sementara*
			// keterangan match(grammar, tags sebelumnya(sifatnya yaitu 0 atau tidak ada), tags selanjutnya(sifatnya yaitu 0 atau tidak ada) )
			if len(tags) == 0 {
				match = []string{gram[x], strconv.Itoa(0), strconv.Itoa(0)}

				// ini adalah proses jika kalimat.Tags belum terisi (kosong) *Proses Sementara*
				// keterangan match(grammar, tags yang terisi sebelumnya, tags selanjutnya(sifatnya yaitu 0 atau tidak ada) )
			} else if len(kalimat.Tags) == 0 {
				match = []string{gram[x], tags[x-1], strconv.Itoa(0)}

				// ini adalah proses jika kalimat.Tags udah terisi namun index (kata) harus bernilai 0 *Proses Fixing*
				// keterangan match(grammar, tags sebelumnya (sifatnya yaitu 0 atau tidak ada), tags selanjutnya yang sudah terisi )
			} else if (len(kalimat.Tags)-1) > x && x < 1 {
				match = []string{gram[x], strconv.Itoa(0), tags[loctags[iloc+1]]}
				// fmt.Println(loctags[iloc])
			} else if (len(tags) - 1) > x {
				match = []string{gram[x], tags[loctags[iloc-1]], tags[loctags[iloc+1]]}
				// fmt.Println(loctags[iloc-1])
			} else {
				match = []string{gram[x], tags[loctags[iloc-1]], strconv.Itoa(0)}
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
			iloc++
			// fmt.Println(tags)
		}

		// memasukkan kedalam struct
		kalimat = Kalimat{gram, tags}

		// calculate error
		check := kalimat.checking_posttag_error()
		// fmt.Println(check)

		//jika tidak ada error maka akan lanjut ke proses selanjutnya
		if check == 0 || check == 0.00 {
			break
		}

		matcher := []string{"imbuhan", "Kdepan", "partikel"}

		for x, e := range tags {
			if !dll.Match_array(matcher, e) {
				loctags = append(loctags, x)
				// temptags = append(temptags, e)
			}
		}
	}

	// pada proses ini, data set tidak bisa menjangkau untuk pembalikan kata jika ada "di" di predikat
	// maka kata tugas disini akan dirubah jadi predikat dan dibalik untuk subjek dan objek
	// jika adanya jumlah katanya kurang dari 3 maka kata tugas akan dirubah saja jadi predikat
	tags := kalimat.Tags
	for x, e := range tags {
		if e == "ktugas" && len(words) > 3 {
			tags[x-1] = "objek"
			tags[x+2] = "subjek"
			tags[x] = "predikat"
		} else if e == "ktugas" {
			tags[x] = "predikat"
		}
	}

	// memasukkan kedalam struct
	kalimat = Kalimat{kalimat.gram, tags}

	return kalimat
}

/*
==============================================================================================

	Pada Proses ini. untuk mengecek suatu kalimat salah atau benar.
	di perlukan perhitungan yaitu dengan menyiapkan tags hasil akhir (SPOK) dan dicocokkan
	di setiap tags yang berdasarkan Rule. jika tidak cocok maka akan di kalkulasi jadi error

	Rule :
	subjek 		= 1 tag
	predikat 	= 1 tag
	Objek		= 1 tag
	Keterangan 	= Dinamis (tidak terbatas Tag)

==============================================================================================
*/

func (kalimat Kalimat) checking_posttag_error() float32 {

	// ini adalah hasil akhir (output) dari dataset posttager bayes
	checker := []string{"subjek", "predikat", "objek", "pelengkap", "keterangan"}

	// semua harus bernilai 0
	count_checker := dll.Start_array_int(len(checker))

	// perhitungannya yaitu setiap checker di cocokkan dengan tags. jika cocok maka count checker +1
	// index dari count checcker sesuai dengan checker
	for i, e := range checker {
		for _, y := range kalimat.Tags {
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
	// fmt.Println(count_checker)
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
