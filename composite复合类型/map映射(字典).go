package main

import "fmt"

/* map(映射,字典)是一个 无序的 key - value(键值对)集合
每个map中的 key(键)必须具有唯一性,必须是支持==或!=的类型(切片,函数等不支持)
map中的 value(值)可以是 任意类型
map中的 属性为 len(长度) , cap (容量)
cap不使用make指定可忽略,cap()在map中无法使用,容量自动扩容
map 返回的 结果是 无序的
map作为函数参数时为 指针传递
*/
func main() {
	// * map的创建和初始化 :
	//	创建 :
	//	1) 标准语法 :
	var m map[int]string // var mapName map[keyType]valueType
	//	2) 自动推导 :
	m1 := map[int]string{}
	m2 := make(map[int]string, 5) //指定 m3容量为5 ,不需要可忽略
	//	初始化 :
	// 1)
	m = map[int]string{1: "比较麻烦", 2: "的"}
	// 2) 常用 :
	m3 := map[int]string{555: "简化", 666: "创建", 888: "map"}
	fmt.Println(m, m1, m2, m3)
	// * map的遍历 :
	for k, v := range m3 { // k ==key, v ==value
		fmt.Printf("m3的键为%v,值为%v\n", k, v)
	}
	// * map的增删改查
	// 1) 增,改 :
	m1[11] = "单次"
	m1[22] = "赋值"   //追加
	m1[11] = "还是单次" // 对已存在的 键进行 赋值,会修改对应的 值
	// 2) 删
	delete(m3, 888) //删除 M3中的 888 键
	fmt.Println(m3)
	// 3) 查 (检查一个 key值是否存在)
	v, b := m3[555] // b 返回true 或 false,如果 b为 true则 v返回 该 值(b为false则忽略v)
	fmt.Println(b, v)
}
