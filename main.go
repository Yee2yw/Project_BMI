package main

import (
	"Project_BMI/BMItest"
	"bufio"
	. "fmt"
	"os"
	"strconv"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	Print("Please input your height(in meters): ")
	heightInput, _ := reader.ReadString('\n')
	height, err := strconv.ParseFloat(heightInput[:len(heightInput)-2], 64)
	//回车键(Enter)在命令行中是\r\n，所以要去掉最后两个字符
	if err != nil {
		Print("Invalid input")
		Print(err)
		return
	}

	Print("Please input your weight(in kilograms): ")
	weightInput, _ := reader.ReadString('\n')
	weight, err := strconv.ParseFloat(weightInput[:len(weightInput)-2], 64)
	if err != nil {
		Print("Invalid input")
		return
	}

	h := BMItest.BMI(&BMItest.Human{Height: height, Weight: weight}) // 注意高度单位应为米
	BMItest.Compare(h)
}
