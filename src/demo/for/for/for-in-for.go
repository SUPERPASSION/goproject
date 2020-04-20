package main

import "fmt"

//统计3个班的成绩，每班有5人，求出各个班的平均分和所有班级的平均分

func main() {
	var class_total int
	var student_total int
	fmt.Printf("请输入班级数：")
	fmt.Scanln(&class_total)
	fmt.Printf("请输入学生数："
	fmt.Scanln(&student_total)
	var nianji_sum float64
	for class := 1; class <= class_total; class++ {
		var class_sum float64
		for student := 1; student <= student_total; student++ {
			var score float64
			fmt.Printf("请输入%v班学生%v的分数：", class, student)
			fmt.Scanln(&score)
			//score = 100
			//fmt.Printf("%v班学生%v的分数：%v\n", class, student, score)
			class_sum += score
		}
		nianji_sum += class_sum
		fmt.Printf("班级%v的总分是：%v\t平均分是：%v\n", class, class_sum, class_sum/float64(student_total))
	}
	fmt.Printf("年级总分是：%v\t平均分是：%v\n", nianji_sum, nianji_sum/float64(class_total*student_total))
}
