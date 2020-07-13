package main

import "fmt"

// 在程序中定义用户的结构体。
type user struct {
	name    string
	surname string
}

func main() {
	// ----------------------
	// Declare and initialize
	// ----------------------

	// 声明并初始化一个map，该map存储字符串类型的键用来保存用户类型的值。
	users1 := make(map[string]user)

	// 添加键值对到map中
	users1["Roy"] = user{"Rob", "Roy"}
	users1["Ford"] = user{"Henry", "Ford"}
	users1["Mouse"] = user{"Mickey", "Mouse"}
	users1["Jackson"] = user{"Michael", "Jackson"}

	// ----------------
	// 遍历map
	// ----------------

	fmt.Printf("\n=> Iterate over map\n")
	for key, value := range users1 {
		fmt.Println(key, value)
	}

	// ------------
	// Map literals
	// ------------

	// 申明并初始化一个带有初始值的map
	users2 := map[string]user{
		"Roy":     {"Rob", "Roy"},
		"Ford":    {"Henry", "Ford"},
		"Mouse":   {"Mickey", "Mouse"},
		"Jackson": {"Michael", "Jackson"},
	}

	// 遍历整个map
	fmt.Printf("\n=> Map literals\n")
	for key, value := range users2 {
		fmt.Println(key, value)
	}

	// ----------
	// 删除map中的key
	// ----------

	delete(users2, "Roy")

	// --------
	// 查找key
	// --------

	// 查找名为Roy的key
	// found1为true，u1为该类型的复制值
	// found2为false,u2为user类型，只不过值为空
	u1, found1 := users2["Roy"]
	u2, found2 := users2["Ford"]

	// 显示查找的值(u)和查找结果(found)
	fmt.Printf("\n=> Find key\n")
	fmt.Println("Roy", found1, u1)
	fmt.Println("Ford", found2, u2)

	// --------------------
	// Map key restrictions
	// --------------------

	// type users []user
	// 使用此语法，我们也可以定义一组用户
	// 这是我们定义用户的第二种方式。我们可以使用现有类型并将其用作复数类型的基础。这是两种不同的类型。这里没有关系。
	// However, when we try use it as a key, like: u := make(map[users]int)
	// the compiler says we cannot use that: "invalid map key type users"
	// The reason is: whatever we use for the key, the value must be comparable. We have to use it
	// in some sort of boolean expression in order for the map to create a hash value for it.
}
