package main

import "fmt"

type celsius float64
type fahrenherit float64
type getRowFn func(row int) (string, string) //声明新函数类型，返回值为两个string类型变量

func (c celsius) fahrenherit() fahrenherit { //声明方法
	return fahrenherit((c * 9.0 / 5.0) + 32.0)
}
func (f fahrenherit) celsius() celsius {
	return celsius((f - 32.0) * 5.0 / 9.0)

}

const (
	line         = "======================="
	rowFormat    = "| %8s | %8s |\n"
	numberFormat = "%.1f"
)

func drawTable(hdr1, hdr2 string, rows int, getRow getRowFn) {
	fmt.Println(line)
	fmt.Printf(rowFormat, hdr1, hdr2)
	fmt.Println(line)
	for row := 0; row < rows; row++ {
		cell1, cell2 := getRow(row)
		fmt.Printf(rowFormat, cell1, cell2) //相当于fmt.Printf("| %8s | %8s |\n",cell1,cell2)
	}

}

func ctof(row int) (string, string) {
	c := celsius(row*5 - 40)
	f := c.fahrenherit()
	cell1 := fmt.Sprintf(numberFormat, c) //使用Sprintf函数将运算结果转换为字符串
	cell2 := fmt.Sprintf(numberFormat, f)
	return cell1, cell2 //返回转换结果
}

func ftoc(row int) (string, string) {
	f := fahrenherit(row*5 - 40)
	c := f.celsius()
	cell1 := fmt.Sprintf(numberFormat, f)
	cell2 := fmt.Sprintf(numberFormat, c)
	return cell1, cell2
}

func main() {
	drawTable("℃", "℉", 29, ctof)
	fmt.Println()
	drawTable("℉", "℃", 29, ftoc)
}
