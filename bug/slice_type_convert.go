package main

import (
	"fmt"
)

// []int32 类型转换为  []*int32时，只有最后一个元素生效的坑

func genOutputWrong(inputIDs []int32) []*int32 {
	outputIDs := make([]*int32, 0, len(inputIDs))
	for _, id := range inputIDs { //用元素本身, 只有最后一个元素赋值生效
		//fmt.Println(id)
		outputIDs = append(outputIDs, &id)
	}
	return outputIDs
}

func genOutputCorrect(inputIDs []int32) []*int32 {
	outputIDs := make([]*int32, 0, len(inputIDs))
	for i, _ := range inputIDs { //用下标, 每一个元素的赋值都生效
		//fmt.Println(inputIDs[i])
		outputIDs = append(outputIDs, &inputIDs[i])
	}
	return outputIDs
}

func main() {
	inputIDs := []int32{1, 2, 3, 4, 5}

	outputIDsWrong := genOutputWrong(inputIDs)
	outputIDsWrongStr := ""
	for _, id := range outputIDsWrong {
		outputIDsWrongStr += fmt.Sprintf("%v ", *id)
	}
	fmt.Printf("[%v]\n", outputIDsWrongStr)

	outputIDsCorrect := genOutputCorrect(inputIDs)
	outputIDsCorrectStr := ""
	for _, id := range outputIDsCorrect {
		outputIDsCorrectStr += fmt.Sprintf("%v ", *id)
	}
	fmt.Printf("[%v]\n", outputIDsCorrectStr)
}
