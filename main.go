/*
jadwal-go LitFill <email>
*/
package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"slices"
	"strings"
	"time"
)

var namaHari = []string{"minggu", "senin", "selasa", "rabu", "kamis", "jum'at", "sabtu"}

type Jadwal struct {
	kelas  string
	jadwal []any
}

type jadwalMap map[string]any

func main() {
	hariJadwal := time.Now().Local().Weekday()
	if !(len(os.Args) < 2) {
		ind := slices.Index(namaHari, strings.ToLower(os.Args[1]))
		if ind < 0 {
			log.Fatal("input salah, gunakan nama hari")
		}
		hariJadwal = time.Weekday(ind)
	}
	fmt.Println(namaHari[hariJadwal])

	jmap := make(jadwalMap)
	fileContent, err := os.ReadFile("./small.json")
	if err != nil {
		log.Fatal("nope!")
	}
	if err := json.Unmarshal(fileContent, &jmap); err != nil {
		log.Fatal("cant Unmarshal")
	}
	fmt.Printf("%+v\n", jmap)

	for key, val := range jmap {
		fmt.Println(key + ":")

		val, ok := val.([]any)
		if !ok {
			log.Fatal("not ok!")
		}

		for i, v := range val {
			v, ok := v.([]any)
			if !ok {
				log.Fatal("vvvvvvvv")
			}
			hari := namaHari[(i/3)+1]
			fmt.Printf("%s jam ke-%d: ", hari, i%3+1)

			for j, w := range v {
				w, ok := w.(map[string]any)
				if !ok {
					log.Fatal("weird!")
				}
				if j == 0 {
					fmt.Printf("Ust %s ", w["nama"])
				}
				if j == 1 {
					fmt.Printf("mengajar fan %s.", w["nama"])
				}

			}
			fmt.Println()
		}
	}
}
