package main

//判断整数大于0小于0还是等于0
//func main() {
//	var num int
//	fmt.Printf("请输入一个整数：")
//	fmt.Scanf("%d", &num)
//	if num > 0 {
//		fmt.Printf("%d大于0", num)
//	} else if num < 0 {
//		fmt.Printf("%d小于0", num)
//	} else {
//		fmt.Printf("%d等于0", num)
//	}
//}

//判断年份是否为闰年
//func main() {
//	var year int
//	fmt.Printf("请输入一个年份：")
//	fmt.Scanf("%d", &year)
//	if year%4 == 0 && year%100 != 0 {
//		fmt.Printf("%d年是闰年", year)
//	} else {
//		fmt.Printf("%d年不是闰年", year)
//	}
//}

//判断整数是否为水仙花数，水仙花数是指一个三位数，其各个位数上的数字的立方和等于其本身，例如153=1*1*1 + 5*5*5 + 3*3*3
//func main() {
//	var baiwei, shiwei, gewei, num int
//	fmt.Printf("请输入一个数字：")
//	fmt.Scanf("%d", &num)
//	if num >= 100 && num <= 999 {
//		baiwei = num / 100
//		shiwei = num % 100 / 10
//		gewei = num % 100 % 10
//		if baiwei*baiwei*baiwei+shiwei*shiwei*shiwei+gewei*gewei*gewei == num {
//			fmt.Printf("%d是一个水仙花数", num)
//		} else {
//			fmt.Printf("%d不是一个水仙花数", num)
//		}
//	} else {
//		fmt.Printf("请输入一个三位数")
//	}
//}

//写出输出结果
//func main() {
//	m, n := 0, 3
//	if m > 0 {
//		if n > 2 {
//			fmt.Printf("A")
//		} else {
//			fmt.Printf("B")
//		}
//	}
//}

//保存用户名和密码判断是不是张三和1234，输出登录成功
//func main() {
//	user:="张三"
//	password:="1234"
//	if user == "张三" && password =="1234"{
//		fmt.Printf("登录成功")
//	}
//
//}

//编写程序，根据输入的年份和月份，求出该月的天数（1-12），用switch
//func main() {
//	var year, month int
//	fmt.Printf("请输入年份：")
//	fmt.Scanf("%d", &year)
//	fmt.Printf("请输入月份：")
//	fmt.Scanf("%d", &month)
//
//	switch month {
//	case 1, 3, 5, 7, 8, 10, 12:
//		fmt.Printf("%d年%d月当月有31天", year, month)
//	case 2:
//		if year%4 == 0 && year%100 != 0 {
//			fmt.Printf("%d年%d月当月有29天", year, month)
//		} else {
//			fmt.Printf("%d年%d月当月有28天", year, month)
//		}
//	default:
//		fmt.Printf("%d年%d月当月有20天", year, month)
//	}
//}

//开发一款软件，根据公式，（身高-108）*2=体重，可以有十斤左右的浮动，来观察测试者体重是否合适
//func main() {
//	var weight, height float64
//	fmt.Printf("请输入身高（cm）：")
//	fmt.Scanf("%f", &height)
//	fmt.Printf("请输入体重（g）：")
//	fmt.Scanf("%f", &weight)
//	if (height-108)*2 > weight-5 && (height-108)*2 < weight+5 {
//		fmt.Printf("您的体重合适")
//	}
//}

//判断此次考试，为什么等级（用switch和if elseif 分支）
//90-100优秀，80-89优良，70-79良好，60-69合格，60以下不及格
//func main() {
//	var score float64
//	fmt.Printf("请输入你的考试成绩：")
//	fmt.Scanf("%f", &score)
//
//	switch {
//	case score <= 100 && score >= 90:
//		fmt.Printf("您的等级为优秀")
//	case score < 90 && score >= 80:
//		fmt.Printf("您的等级为优良")
//	case score < 80 && score >= 70:
//		fmt.Printf("您的等级为良好")
//	case score < 70 && score >= 60:
//		fmt.Printf("您的等级为合格")
//	case score < 60:
//		fmt.Printf("您的等级为不合格")
//	default:
//		fmt.Printf("你他妈的考的是个j8	")
//	}
//}

//有两个数a和b，如果a能被b整除，或者a加b大于1000，则输出a，否则输出b
//func main() {
//	var a, b int = 100, 9
//	if a%b == 0 || a+b > 1000 {
//		fmt.Printf("%v", a)
//	} else {
//		fmt.Printf("%v", b)
//	}
//}

//对三个数，进行从小到大的输出
//func main() {
//	var a, b, c float64
//	fmt.Printf("请输入三个数字abc，用空格分隔：")
//	fmt.Scanf("%v %v %v", &a, &b, &c)
//	if a > b {
//		a, b = b, a
//	}
//	if c < a {
//		a, b, c = c, a, b
//	} else if c < b && c > a {
//		a, b, c = a, c, b
//	}
//	fmt.Printf("排序为：\n%v\n%v\n%v", a, b, c)
//}
