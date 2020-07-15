package main

import "fmt"

//定义二维数组，用于保存三个班，每个班五名同学成绩，并求出每个班级平均分、以及所有班级平均分
func main() {
	var scores [3][5]float64
	for i := 0; i < len(scores); i++ {
		for j := 0; j < len(scores[i]); j++ {
			fmt.Printf("请输入第%d班,第%d号学生成绩:\n", i+1, j+1)
			fmt.Scanln(&scores[i][j])
		}
	}

	totlaSum := 0.0
	for i := 0; i < len(scores); i++ {
		sum := 0.0
		for j := 0; j < len(scores[i]); j++ {
			sum += scores[i][j]
		}
		totlaSum += sum
		fmt.Printf("第%v班总分:%v,平均分:%v\n", i+1, sum, sum/float64(len(scores[i])))
	}
	fmt.Printf("所有班级总分:%v,平均分:%v\n", totlaSum, totlaSum/15)
}
