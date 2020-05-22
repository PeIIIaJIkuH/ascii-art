package art

import (
	"fmt"
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
