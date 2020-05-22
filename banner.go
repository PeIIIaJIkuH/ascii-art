package art

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Banner struct {
	arr [][]string
}

func (b Banner) GetArr() [][]string {
	return b.arr
}

func (b *Banner) Clear() {
	b.arr = [][]string{}
}

func (b *Banner) Init(filename string) {
	b.Clear()
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	str := string(file)
	for i := 0; i < 96; i++ {
		b.arr = append(b.arr, []string{})
		str = str[strings.Index(str, "\n")+1:]
		b.arr[i] = append(b.arr[i], []string{"", "", "", "", "", "", "", ""}...)
		for j := 0; j < 8; j++ {
			b.arr[i][j] += str[:strings.Index(str, "\n")]
			str = str[strings.Index(str, "\n")+1:]
		}
	}
}

func (b Banner) Print() {
	for i := 0; i < len(b.arr); i++ {
		fmt.Println(i)
		for j := 0; j < len(b.arr[i]); j++ {
			fmt.Println(b.arr[i][j])
		}
	}
}

func isEqual(a1, a2 []string) bool {
	if len(a1) != len(a2) {
		return false
	}
	for i := range a1 {
		if a1[i] != a2[i] {
			return false
		}
	}
	return true
}

func (b Banner) Index(symbol []string) int {
	for i, j := range b.arr {
		if isEqual(symbol, j) {
			return i
		}
	}
	return -1
}

func (b Banner) ToBig(symbol byte) []string {
	alpha := Alphabet()
	index := strings.Index(alpha, string(symbol))
	return b.arr[index]
}

func (b Banner) Find(big []string) int {
	for i, j := range b.arr {
		if isEqual(big, j) {
			return i
		}
	}
	return -1
}
