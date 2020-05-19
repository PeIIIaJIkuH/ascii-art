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

// func (a *Art) Apply(c byte, b Banner) {
// 	big := b.ToBig(c)
// 	a.arr[a.i] = append(a.arr[a.i], big)
// }

func (a *Art) Apply(str string, b Banner) {
	for i := 0; i < len(str); i++ {
		if str[i] == '\\' && str[i+1] == 'n' {
			a.Update()
			i += 2
		}
		big := b.ToBig(str[i])
		a.arr[a.i] = append(a.arr[a.i], big)
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
	for ; size+len(a.arr[index][i+1][0]) < width; i++ {
		size += len(a.arr[index][i][0])
	}
	temp := Art{0, [][][]string{a.arr[index][:i]}, [][]string{a.color[index][:i]}}
	a.arr[index] = a.arr[index][i:]
	a.color[index] = a.color[index][i:]
	return temp
}

// temp := Art{}
// temp.Init()
// temp.i = 0
// for j := 0; j < i; j++ {
// 	temp.arr[0] = append(temp.arr[0], a.arr[index][j])
// 	temp.color[0] = append(temp.color[0], a.color[index][j])
// }

func (a Art) Print(align string) {
	width := terminalWidth()
	for i := 0; i <= a.i; i++ {
		for a.Size(i) > width {
			c := a.copy(i)
			c.Print("")
		}
		left := 0
		if align == "right" {
			left = width - a.Size(i)
		} else if align == "center" {
			left = (width - a.Size(i)) / 2
		}
		for j := 0; j < 8; j++ {
			printStr(" ", left)
			for k := range a.arr[i] {
				fmt.Printf(a.color[i][k], a.arr[i][k][j])
			}
			fmt.Println()
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

func (a *Art) InitColors(str string, colors []string, slices *[][]int) {
	j, index := 0, 0
	for i := 0; i < len(str); i++ {
		if str[i] == '\\' {
			if str[i+1] == 'n' {
				j++
				i += 2
				if len(*slices) > 0 {
					(*slices)[index][1] += 2
					for j := index + 1; j < len(*slices); j++ {
						(*slices)[j][0] += 2
						(*slices)[j][1] += 2
					}
				}
			}
		}
		color := "white"
		if len(colors) == 0 {
			color = "white"
		} else if len(colors) == 1 {
			color = colors[0]
		} else if i >= (*slices)[index][0] && i < (*slices)[index][1] {
			color = colors[index]
			if i == (*slices)[index][1]-1 && index < len(colors)-1 {
				index++
			}
			if i >= (*slices)[len((*slices))-1][1] {
				color = "white"
			}
		}

		rgb := generateRgb(color)
		if len(rgb) == 0 {
			fmt.Println("Wrong color")
			os.Exit(3)
		}

		a.color[j] = append(a.color[j], "\033[38;2;"+rgb+"%s\033[0m")
	}
}

// func (a *Art) ToArt(filename string) {
// 	file, err := ioutil.ReadFile(filename)
// 	if err != nil {
// 		log.Fatal(err)
// 		return
// 	}
// 	str := string(file)

// }
