package main

import "fmt"

func main() {
	plantes := [...]string{
		"Mercury",
		"Venus",
		"Earth",
		"Mars",
		"Jupiter",
		"Saturn",
		"Uranus",
		"Neptune",
	}

	terrestrial := plantes[0:4] //创建切片
	gasGiants := plantes[4:6]
	iceGiants := plantes[6:8]
	fmt.Println(terrestrial, gasGiants, iceGiants)

	fmt.Println(terrestrial[0]) //通过索引访问切片中的元素

	giants := plantes[4:8]
	gas := giants[0:2] //创建切片的切片
	ice := giants[2:4]
	fmt.Println(gas, ice)

	iceGiantsMarkII := iceGiants //赋值iceGiants切片
	iceGiants[1] = "Poseidon"    //修改iceGiants切片中的值
	fmt.Println(iceGiants, iceGiantsMarkII, ice)
}
