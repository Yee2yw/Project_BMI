package main

import "Project_BMI/BMI"

func main() {
	h := BMI.BMI(&BMI.Human{Height: 1.8, Weight: 70}) // 注意高度单位应为米
	BMI.Compare(h)
}
