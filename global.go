package art

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func terminalWidth() int {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, _ := cmd.Output()
	output := string(out)[:len(out)-1]
	width, _ := strconv.Atoi(output[strings.Index(output, " ")+1:])
	return width
}

func Alphabet() string {
	str := ""
	for i := 32; i <= 126; i++ {
		str += string(byte(i))
	}
	return str
}

func LenWithoutNewline(str string) int {
	size := len(str)
	for i := 0; i < len(str); i++ {
		if str[i] == '\\' {
			if str[i+1] == 'n' {
				i += 2
				size -= 2
			}
		}
	}
	return size
}

func generateRgb(color string) string {
	switch color {
	case "":
		return "255;255;255m"
	case "white":
		return "255;255;255m"
	case "black":
		return "0;0;0m"
	case "red":
		return "255;0;0m"
	case "green":
		return "0;128;0m"
	case "yellow":
		return "255;0255;0m"
	case "blue":
		return "0;0;255m"
	case "magenta":
		return "255;0;255m"
	case "cyan":
		return "0;255;255m"
	case "lime":
		return "0;255;0m"
	case "silver":
		return "192;192;192m"
	case "gray":
		return "128;128;128m"
	case "maroon":
		return "128;0;0m"
	case "olive":
		return "128;128;0m"
	case "purple":
		return "128;0;128m"
	case "teal":
		return "0;128;128m"
	case "mint":
		return "170;255;195m"
	case "lavender":
		return "230;190;255m"
	case "pink":
		return "250;190;190m"
	case "brown":
		return "170;110;40m"
	case "orange":
		return "245;130;48m"
	case "apricot":
		return "255;215;180m"
	case "beige":
		return "255;250;200m"
	case "tomato":
		return "255;99;71m"
	case "gold":
		return "255;215;0m"
	case "salmon":
		return "250;128;114m"
	default:
		arr := strings.Split(color, ".")
		if len(arr) == 3 {
			r, e1 := strconv.Atoi(arr[0])
			g, e2 := strconv.Atoi(arr[1])
			b, e3 := strconv.Atoi(arr[2])
			if e1 != nil || e2 != nil || e3 != nil || r < 0 || g < 0 || b < 0 || r > 255 || g > 255 || b > 255 {
				return ""
			}
			return strconv.Itoa(r) + ";" + strconv.Itoa(g) + ";" + strconv.Itoa(b) + "m"
		}
		return ""
	}
}

func printStr(str string, count int) {
	for i := 0; i < count; i++ {
		fmt.Print(str)
	}
}

func toArr(str string) [][]string {
	index, newlines := 0, 0
	arr := [][]string{{"", "", "", "", "", "", "", ""}}
	for len(str) > 0 {
		arr[index][newlines%8] += str[:strings.Index(str, "\n")]
		str = str[strings.Index(str, "\n")+1:]
		if newlines%8 == 7 && len(str) > 1 {
			index++
			arr = append(arr, []string{"", "", "", "", "", "", "", ""})
		}
		newlines++
	}
	return arr
}

func areEqual(arr []string) bool {
	for i := 0; i < 8; i++ {
		for j := i + 1; j < 8; j++ {
			if len(arr[i]) != len(arr[j]) {
				return false
			}
		}
	}
	return true
}

func checkReverse(arr [][]string, str string) bool {
	if strings.Count(str, "\n")%8 != 0 {
		return false
	}
	for _, i := range arr {
		if len(i) != 8 {
			return false
		}
	}
	for _, i := range arr {
		if !areEqual(i) {
			return false
		}
	}
	return true
}

func printArr(arr [][]string) {
	for _, i := range arr {
		for _, j := range i {
			fmt.Println("|" + j + "|")
		}
	}
}

func toBig(arr [][]string, index, start, end int) []string {
	big := []string{}
	for i := 0; i < 8; i++ {
		big = append(big, arr[index][i][start:end])
	}
	return big
}

var fonts []string = []string{"../standard.txt", "../shadow.txt", "../thinkertoy.txt"}
var fontIndex int = 0

func Reverse(filename string, b Banner) {
	file, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
		os.Exit(3)
	}
	str := string(file)
	arr := toArr(str)
	if !checkReverse(arr, str) {
		fmt.Println("Incorrect reverse file!")
		os.Exit(3)
	}
	alpha := Alphabet()
	res := ""
	for i := 0; i < len(arr); i++ {
		start, add := 0, 1
		for start < len(arr[i][0]) {
			for b.Find(toBig(arr, i, start, start+add)) == -1 {
				if start+add >= len(arr[i][0]) {
					fontIndex++
					if fontIndex > 2 {
						fmt.Println("Incorrect reverse file!")
						os.Exit(3)
					}
					b.Init(fonts[fontIndex])
					Reverse(filename, b)
					return
				}
				add++
			}
			find := b.Find(toBig(arr, i, start, start+add))
			res += string(alpha[find])
			start += add
			add = 1
		}
		if i < len(arr)-1 {
			res += "\n"
		}
	}
	fmt.Println(res)
}
