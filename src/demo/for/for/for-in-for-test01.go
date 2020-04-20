package main

import "fmt"

func main() {
	//输入3个班的分数，5个学生，统计每个班的总分，平均分，及年级的总分平均分
	var class_sum float64
	var nianji_sum float64
	var class_max, student_max int = 3, 5

	for class := 1; class <= class_max; class++ {
		class_sum = 0
		for student := 1; student <= student_max; student++ {
			var score float64
			fmt.Printf("请输入班级%v学生%v的成绩：", class, student)
			fmt.Scanf("%v", &score)
			class_sum += score
		}
		nianji_sum += class_sum
		fmt.Printf("班级%v的总分：%v\t平均分：%v\n", class, class_sum, class_sum/float64(student_max))
	}
	fmt.Printf("全年级总分：%v,平均分：%v", nianji_sum, nianji_sum/float64(class_max*student_max))
}
