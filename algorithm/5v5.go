package main

import "fmt"

// 现有10个明星各有一定的粉丝数, 现在需要分成2组，各组粉丝数量最好接近;
// 设状态转移函数
// 转换为背包问题: 10个物件中挑5件物品放入容积为(总粉丝数/2)的背包中
// F(i, k, n) F为前i个人中挑了k个人放入了剩余容积

func GroupForVersus(starFans []int) (map[int]int, map[int]int) {
	if len(starFans)%2 != 0 {
		panic("count versus")
	}
	gp1 := make(map[int]int, len(starFans)/2)
	gp2 := make(map[int]int, len(starFans)/2)
	// 在fans的明星中选出starlimit使他们的粉丝数的和接近fansLimit
	var tmpFunc func(fans []int, starLimit, fansLimit int) (int, map[int]int)
	tmpFunc = func(fans []int, starLimit, fansLimit int) (int, map[int]int) {
		lastIndex := len(fans) - 1
		if starLimit < 1 || fansLimit < 1 || lastIndex < 0 {
			return 0, map[int]int{}
		}
		// 未选择最后一位明星时分组粉丝和
		notChoose, gpNotChoose := tmpFunc(fans[:lastIndex], starLimit, fansLimit)
		if fansLimit < fans[lastIndex] {
			return notChoose, gpNotChoose
		}
		// 选择最后一位明星时分组粉丝和
		choose, gpChoose := tmpFunc(fans[:lastIndex], starLimit-1, fansLimit-fans[lastIndex])
		choose += fans[lastIndex]
		gpChoose[lastIndex] = fans[lastIndex]

		if choose > notChoose {
			return choose, gpChoose
		} else {
			return notChoose, gpNotChoose
		}
	}
	var totalFans int
	for _, v := range starFans {
		totalFans += v
	}
	_, gp1 = tmpFunc(starFans, len(starFans)/2, totalFans/2)
	for index, value := range starFans {
		if _, ok := gp1[index]; !ok {
			gp2[index] = value
		}
	}
	return gp1, gp2
}

func main() {
	starFans := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	gp1, gp2 := GroupForVersus(starFans)
	fmt.Println("明星分组问题")
	for star, fans := range starFans {
		fmt.Println("明星", star, "粉丝数: ", fans)
	}
	fmt.Println("分组1:")
	for star, fans := range gp1 {
		fmt.Println("明星", star, "粉丝数: ", fans)
	}
	fmt.Println("分组2:")
	for star, fans := range gp2 {
		fmt.Println("明星", star, "粉丝数: ", fans)
	}
}
