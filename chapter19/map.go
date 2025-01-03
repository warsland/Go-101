package main

import "fmt"

func main() {
	temperature := map[string]int{
		"Earth": 15, //通过复合字面量为映射提供键值对
		"Mars":  -65,
	}

	temp := temperature["Earth"] //按键查值
	fmt.Printf("On average the Earth is %v℃\n", temp)
	temperature["Earth"] = 16  //修改映射的键值对的值
	temperature["Venus"] = 464 //添加新的键值对到映射
	fmt.Println(temperature)

	/*moon := temperature["Moon"]//查找映射中不存在的键值对。
	fmt.Println(moon)*/
	if moon, ok := temperature["Moon"]; ok { //当ok值为true时，表示键"Moon"存在且键的值为零值
		fmt.Printf("On average the Moon is %v℃\n", moon)
	} else { //当ok值为false时，表示键"Moon"不存在
		fmt.Println("Where is the moon?")
	}
}
