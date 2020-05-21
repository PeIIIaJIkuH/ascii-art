package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	art ".."
)

func contains(str string) bool {
	for _, i := range os.Args[1:] {
		if i == str {
			return true
		}
	}
	return false
}

func flags(find string) string {
	for _, s := range os.Args[1:] {
		if strings.Index(s, find) != -1 {
			return s[len(find):]
		}
	}
	return ""
}

func checkColor(color, slice *[]string) bool {
	if len(*slice) == 0 {
		*color = (*color)[:1]
	} else if len(*slice) != len(*color) {
		min := len(*color)
		if min > len(*slice) {
			min = len(*slice)
		}
		*color = (*color)[:min]
		*slice = (*slice)[:min]
	}
	prevEnd := 0
	for i := range *slice {
		arr := strings.Split((*slice)[i], ":")
		if len(arr) != 2 {
			fmt.Println("Slice must be in 'a:b' format, where a and b are non-negative integers, a < b, each slice must be different")
			return false
		}
		if arr[0] == "" || arr[1] == "" {
			fmt.Println("Slice must be in 'a:b' format, where a and b are non-negative integers, a < b, each slice must be different")
			return false
		}
		start, e1 := strconv.Atoi(arr[0])
		end, e2 := strconv.Atoi(arr[1])
		if start >= end || start < 0 || end < 0 || e1 != nil || e2 != nil || start < prevEnd {
			fmt.Println("Slice must be in 'a:b' format, where a and b are non-negative integers, a < b, each slice must be different")
			return false
		}
		prevEnd = end
	}
	return true
}

func parseColors() ([]string, [][]int) {
	str := os.Args[1]
	colorstr := flags("--color=")
	colors := []string{}
	if len(colorstr) != 0 {
		colors = strings.Split(colorstr, ",")
	}
	slicestr := flags("--slice=")
	slice := []string{}
	if len(slicestr) != 0 {
		slice = strings.Split(slicestr, ",")
	}
	if len(colors) != 0 {
		if !checkColor(&colors, &slice) {
			os.Exit(3)
		}
	}

	if len(slice) > 0 {
		last := strings.Split(slice[len(slice)-1], ":")
		lastStart, _ := strconv.Atoi(last[0])
		lastEnd, _ := strconv.Atoi(last[1])
		if lastStart >= art.LenWithoutNewline(str) || lastEnd > art.LenWithoutNewline(str) {
			fmt.Println("Slices must be in range of the word")
			os.Exit(3)
		}
	}

	slices := [][]int{}
	for _, i := range slice {
		arr := strings.Split(i, ":")
		start, _ := strconv.Atoi(arr[0])
		end, _ := strconv.Atoi(arr[1])
		slices = append(slices, []int{start, end})
	}

	return colors, slices
}

func checkAlign(align *string) {
	if len(*align) == 0 {
		*align = "left"
	}
	if *align != "left" && *align != "right" && *align != "center" && *align != "justify" {
		fmt.Println("Wrong align!")
		os.Exit(3)
	}
}

func main() {
	args := os.Args[1:]
	if len(args) >= 1 {
		str := args[0]
		if len(str) == 0 {
			fmt.Println("String must contain at least 1 character")
			return
		}

		a := art.Art{}
		a.Init()
		b := art.Banner{}
		b.Init("../standard.txt")

		count := 0
		if contains("standard") {
			count++
		}
		if contains("shadow") {
			count++
			b.Init("../shadow.txt")
		}
		if contains("thinkertoy") {
			count++
			b.Init("../thinkertoy.txt")
		}
		if count > 1 {
			fmt.Println("Please choose only 1 font style")
			return
		}

		// reverse := flags("--reverse=")
		// if len(reverse) != 0 {

		// }

		a.Apply(str, b)

		colors, slices := parseColors()
		a.InitColors(str, colors, &slices)
		fmt.Println("Colors:")
		fmt.Println(a.GetColor())
		fmt.Println("_______")

		align := flags("--align=")
		checkAlign(&align)

		filename := flags("--output=")
		if len(filename) == 0 {
			a.Print(align)
		} else {
			if align != "" && align != "left" || len(colors) != 0 {
				fmt.Println("Can not write to file with these flags")
				return
			}
			a.Fprint(filename)
			fmt.Println("Done")
		}
	}
}
