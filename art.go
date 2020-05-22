package art

import (
	"fmt"
	"os"
)

type Art struct {
	i     int
	arr   [][][]string
	color [][]string
}

func (a Art) GetIndex() int {
	return a.i
}

func (a Art) GetArr() [][][]string {
	return a.arr
}

func (a Art) GetColor() [][]string {
	return a.color
}

func (a *Art) Init() {
	a.arr = append(a.arr, [][]string{})
	a.color = append(a.color, []string{})
}

func (a *Art) Update() {
	a.i++
	a.Init()
}

func (a *Art) Apply(str string, b Banner) {
	for i := 0; i < len(str); i++ {
		if str[i] == '\\' && i+1 < len(str) {
			if str[i+1] == 'n' {
				a.Update()
				i++
				if i == len(str)-1 {
					break
				}
				continue
			}
		}
		big := b.ToBig(str[i])
		a.arr[a.i] = append(a.arr[a.i], big)
	}
	if len(a.arr[a.i]) == 0 {
		a.arr[a.i] = append(a.arr[a.i], []string{"", "", "", "", "", "", "", ""})
	}
}

func (a Art) Size(index int) int {
	count := 0
	for _, i := range a.arr[index] {
		count += len(i[0])
	}
	return count
}

func (a *Art) copy(index int) Art {
	width := terminalWidth()
	i, size := 0, 0
	for ; size < width; i++ {
		size += len(a.arr[index][i][0])
	}
	i--
	size -= len(a.arr[index][i][0])
	temp := Art{0, [][][]string{a.arr[index][:i]}, [][]string{a.color[index][:i]}}
	a.arr[index] = a.arr[index][i:]
	a.color[index] = a.color[index][i:]
	return temp
}

func (a Art) simplePrint(align string, index int) {
	width := terminalWidth()
	left := 0
	if align == "right" {
		left = width - a.Size(index)
	} else if align == "center" {
		left = (width - a.Size(index)) / 2
	}
	for j := 0; j < 8; j++ {
		printStr(" ", left)
		for k := range a.arr[index] {
			fmt.Printf(a.color[index][k], a.arr[index][k][j])
		}
		fmt.Println()
	}
}

func (a *Art) TrimLeadSpaces(index int, b Banner) {
	for _, i := range a.arr[index] {
		if b.Find(i) != 0 {
			return
		}
		a.arr[index] = a.arr[index][1:]
		temp := a.color[index][0]
		fmt.Println(index)
		for j := 1; j < len(a.color[index]); j++ {
			temp, a.color[index][j] = a.color[index][j], temp
		}
		a.color[index] = a.color[index][1:]
	}
}

func (a *Art) TrimTailSpaces(index int, b Banner) {
	for i := len(a.arr[index]) - 1; i >= 0; i-- {
		if b.Find(a.arr[index][i]) != 0 {
			return
		}
		a.arr[index] = a.arr[index][:i]
		a.color[index] = a.color[index][:i]
	}
}

func (a *Art) TrimSpaces(index int, b Banner) {
	a.TrimLeadSpaces(index, b)
	a.TrimTailSpaces(index, b)
}

func (a Art) printJustify(b Banner) {
	for i := range a.arr {
		a.TrimSpaces(i, b)
		a.simplePrint("", i)
	}
}

func (a Art) Print(align string, b Banner) {
	width := terminalWidth()
	for i := 0; i <= a.i; i++ {
		for a.Size(i) > width {
			c := a.copy(i)
			if align == "justify" {
				c.printJustify(b)
			} else {
				c.simplePrint(align, 0)
			}
		}
		if align == "justify" {
			a.printJustify(b)
		} else {
			a.simplePrint(align, i)
		}
	}
}

func (a Art) Fprint(filename string) {
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	for i := 0; i <= a.i; i++ {
		for j := 0; j < 8; j++ {
			for k := range a.arr[i] {
				file.WriteString(a.arr[i][k][j])
			}
			file.WriteString("\n")
		}
	}
	err = file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (a *Art) InitColors(colors []string, slices *[][]int, b Banner) {
	fmt.Println(colors)
	fmt.Println(*slices)
	index, colorIndex := 0, 0
	for i := range a.arr {
		for j := range a.arr[i] {
			if b.Find(a.arr[i][j]) != 0 {
				if colorIndex >= len(colors) {
					a.color[i] = append(a.color[i], "\033[38;2;255;255;255m%s\033[0m")
					continue
				}
				if index >= (*slices)[colorIndex][0] && index < (*slices)[colorIndex][1] {
					rgb := generateRgb(colors[colorIndex])
					if len(rgb) == 0 {
						fmt.Println("Wrong color")
						os.Exit(3)
					}
					a.color[i] = append(a.color[i], "\033[38;2;"+rgb+"%s\033[0m")
					if index == (*slices)[colorIndex][1]-1 {
						colorIndex++
					}
				}
				index++
			} else if b.Find(a.arr[i][j]) == 0 {
				a.color[i] = append(a.color[i], "\033[38;2;255;255;255m%s\033[0m")
			}
		}
	}
}

// func (a *Art) InitColors(str string, colors []string, slices *[][]int, align string) {
// 	fmt.Println("Printing colors:", a.color)
// 	fmt.Println(slices)
// 	j, index := 0, 0
// 	for i := 0; i < len(str); i++ {
// 		if str[i] == '\\' && i+1 < len(str) {
// 			if str[i+1] == 'n' {
// 				if len(a.color[j]) == 0 {
// 					a.color[j] = append(a.color[j], "\033[38;2;255;255;255m%s\033[0m")
// 				}
// 				j++
// 				i++
// 				if len(*slices) > 0 {
// 					for k := index; k < len(*slices); k++ {
// 						(*slices)[k][0] += 2
// 						(*slices)[k][1] += 2
// 					}
// 				}
// 				continue
// 			}
// 		}
// 		if align == "justify" && str[i] == ' ' {
// 			if len(*slices) > 0 {
// 				for k := index; k < len(*slices); k++ {
// 					(*slices)[k][0]++
// 					(*slices)[k][1]++
// 				}
// 			}
// 			continue
// 		}

// 		color := "white"
// 		if len(colors) == 0 {
// 			color = "white"
// 		} else if len(*slices) == 0 {
// 			color = colors[index]
// 		} else if i >= (*slices)[index][0] && i < (*slices)[index][1] {
// 			color = colors[index]
// 			if i == (*slices)[index][1]-1 && index < len(colors)-1 {
// 				index++
// 			}
// 		}
// 		if len((*slices)) >= 1 && i >= (*slices)[len((*slices))-1][1] {
// 			color = "white"
// 		}

// 		rgb := generateRgb(color)
// 		if len(rgb) == 0 {
// 			fmt.Println("Wrong color")
// 			os.Exit(3)
// 		}
// 		a.color[j] = append(a.color[j], "\033[38;2;"+rgb+"%s\033[0m")
// 	}
// 	if len(a.color[j]) == 0 {
// 		a.color[j] = append(a.color[j], "\033[38;2;255;255;255m%s\033[0m")
// 	}
// 	fmt.Println(slices)
// }
