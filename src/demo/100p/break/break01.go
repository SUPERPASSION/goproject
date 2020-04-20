package main

import "fmt"

func main() {
	////为了生成一个随机数，还需要rand设置一个种子
	//rand.Seed(time.Now().Unix())
	//for {
	//	//随机生成1~100的整数
	//	random_num := rand.Intn(101)
	//	fmt.Printf("%v\n", random_num)
	//	if random_num == 99 {
	//		break
	//	}
	//}

	for i:=1;i<=100;i++{
		fmt.Println(i)
		for j:=1;j<=100;j++{
			fmt.Println(j)
			if j==50{
				break
			}
		}
	}
}
