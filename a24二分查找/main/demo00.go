package main

import "fmt"

//对一个有序数组进行二分查找{1,8,10,89,1000,1234}，输入一个数看看该数组是否存在此数，并且求出下标，如果没有就提示"没有这个数"
func BinaryFind(arr *[6]int, leftIndex int, rightIndex int, findVal int) {
	//如果左边大于右边
	if leftIndex > rightIndex {
		fmt.Println("找不到")
		return
	}
	//找到中间的下标
	middle := (leftIndex + rightIndex) / 2
	//如果在左边
	if (*arr)[middle] > findVal {
		//递归调用
		BinaryFind(arr, leftIndex, middle-1, findVal)
	} else if (*arr)[middle] < findVal {
		//如果在右边
		BinaryFind(arr, middle+1, rightIndex, findVal)
	} else {
		fmt.Printf("找到,下标为:%v\n", middle)
	}

}

func main() {
	arr := [6]int{1, 8, 10, 89, 1000, 1234}
	BinaryFind(&arr, 0, len(arr)-1, 8)
}
