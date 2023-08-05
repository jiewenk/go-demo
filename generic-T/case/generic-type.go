package _case

import "fmt"

// getMaxIntNum 普通的int类型的比较
func getMaxIntNum(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// getMaxLongNum 普通的long类型的比较
func getMaxLongNum(a, b int64) int64 {
	if a >= b {
		return a
	}
	return b
}

// getMaxFloatNum 普通的float类型的比较
func getMaxFloatNum(a, b float32) float32 {
	if a >= b {
		return a
	}
	return b
}

// getMaxDoubleNum 普通的double类型的比较
func getMaxDoubleNum(a, b float64) float64 {
	if a >= b {
		return a
	}
	return b
}

func PlainTypeCase() {
	intNum := getMaxIntNum(1, 5)
	longNum := getMaxLongNum(1999999999, 57897464649879)
	floatNum := getMaxFloatNum(1789.5, 5789798.7)
	doubleNum := getMaxDoubleNum(19879454649674987946.89898989, 546546498798.4654674987)
	fmt.Printf("普通类型：提供的两个int数字中最大的是[%d]\n", intNum)
	fmt.Printf("普通类型：提供的两个long数字中最大的是[%d]\n", longNum)
	fmt.Printf("普通类型：提供的两个float数字中最大的是[%f]\n", floatNum)
	fmt.Printf("普通类型：提供的两个double数字中最大的是[%.10f]\n", doubleNum)
}

func genericFunction[T int | int64 | float32 | float64](a, b T) T {
	if a >= b {
		return a
	}
	return b
}

func GenericFunctionCase() {
	intNum := genericFunction(1, 5)
	longNum := genericFunction(1999999999, 57897464649879)
	floatNum := genericFunction(1789.5, 5789798.7)
	doubleNum := genericFunction(19879454649674987946.89898989, 546546498798.4654674987)
	fmt.Printf("泛型类型函数：提供的两个int数字中最大的是[%d]\n", intNum)
	fmt.Printf("泛型类型函数：提供的两个long数字中最大的是[%d]\n", longNum)
	fmt.Printf("泛型类型函数：提供的两个float数字中最大的是[%f]\n", floatNum)
	fmt.Printf("泛型类型函数：提供的两个double数字中最大的是[%.10f]\n", doubleNum)
}

// CustomGenericTypeOfNumber 自定义泛型类型接口
type CustomGenericTypeOfNumber interface {
	uint8 | uint16 | int32 | uint64 | ~uint32 | float32
	uint32 | float32 | float64 | int32
}

func customGenericFunction[T CustomGenericTypeOfNumber](a, b T) T {
	if a >= b {
		return a
	}
	return b
}

func CustomGenericCase() {
	intNum := customGenericFunction(int32(1), int32(5))
	//longNum := customGenericFunction(1999999999, 57897464649879)
	floatNum := customGenericFunction(float32(1789.5), float32(5789798.7))
	//doubleNum := customGenericFunction(19879454649674987946.89898989, 546546498798.4654674987)
	fmt.Printf("自定义泛型类型：提供的两个int数字中最大的是[%d]\n", intNum)
	//fmt.Printf("泛型类型函数：提供的两个long数字中最大的是[%d]\n", longNum)
	fmt.Printf("自定义泛型类型：提供的两个float数字中最大的是[%f]\n", floatNum)
	//fmt.Printf("泛型类型函数：提供的两个double数字中最大的是[%.10f]\n", doubleNum)
}
