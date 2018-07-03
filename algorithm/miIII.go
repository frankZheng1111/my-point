package main

// 小米三面Q2：
// 一副从1到n的牌，每次从牌堆顶取一张放桌子上，再取一张牌放在牌堆底, 直到手里没牌，最后桌子上的牌是从1到n有序, 设计程序, 输入n, 输出牌堆的顺序数组

// 题目解析(说人话):
// 一副从1到n的牌，(现在有一个牌堆，牌堆内的牌面1到n, 牌堆顺序未知),
// 每次从牌堆顶取一张放桌子上，再取一张牌放在牌堆底, (按照此顺序出牌, 桌面上的牌横向排列), 直到手里没牌
// (若要求)最后桌子上的牌是从1到n有序，(则)设计程序, 输入n, 输出(满足该条件的初始)牌堆的顺序数组

// 解法:
// 构建0 ~ n - 1 的数组, 0为牌堆顶
// 按照题述的方式出牌, 变动数组顺序, 例: 设n=5, [0,1,2,3,4] = 数组A0, 出牌顺序后为 [0, 2, 4, 3, 1] = 数组A1
// 替换变动后的数组的值和下标(括弧内为下标)[0(0), 2(1), 4(2), 3(3), 1(4)] = 交换值与下标 => [0(0), 1(2), 2(4), 3(3), 4(1)] = 根据变化后的下标调整顺序 => [0, 4, 1, 3, 2]
// 最后调整完的数组即满足条件的初始牌堆顺序

// 解题思路
//
// 数组A1 值表示的是符合递增顺序的原数组A0的下标, A1的下标表示的是出牌后的新下标(位置),

// ***** 核心思路 *****
// 比如A1[1] = 2, A1[2] = 4, 分别表示原先A0下标为2的值出牌后位置在1, 原先A0下标为4的值出牌后位置在2, 构成一组映射关系
// 为了满足出牌后的顺序, 我们只需将正确位置的值，替换安置到变换前的数组A0的对应位置(即A1的值代表的是A0的对应位置)上
// ***** 核心思路 *****

// 就此题而言，我们的解题目标是出牌后的顺序是0 ~ n - 1有序，数组A1的下标刚好是0 ~ n - 1, 而数组A1的值又代表的是A0的对应位置, 故只要交换A1的值与下标即可(个人认为交换下标和值只是这一题的个例, 不容易扩展和理解)

import "fmt"

// 按照特定顺序出牌
func BuildNewCards(initialCards *[]int) []int {
	var cardNum int = len(*initialCards)
	var cards []int = make([]int, 0, cardNum)
	for len(*initialCards) > 0 { // ps 这边只为了直观的描述出牌操作顺序，性能上存在优化的余地
		// 牌顶出一张牌, 设第一个元素是牌顶
		cards = append(cards, (*initialCards)[0])
		if len(*initialCards) == 1 {
			*initialCards = (*initialCards)[0:0] // 移除最后一张牌
			break
		} else {
			*initialCards = (*initialCards)[1:]
		}
		// 置一张牌到牌底
		if len(*initialCards) >= 1 {
			*initialCards = append(*initialCards, (*initialCards)[0])
			*initialCards = (*initialCards)[1:]
		}
	}
	return cards
}

func InitialCards(cardNum int) []int {
	var initialCards []int = make([]int, cardNum, cardNum)
	var targetCards []int = make([]int, cardNum, cardNum) // 最后要求的目标数组, 就该题而言完全不需要，但是若要题目扩展为其他顺序的结果时可用到此数组
	var cardsAfterChange []int
	// 将卡组的下标显示设置在数组里
	// 构建目标数组
	for i := 0; i < cardNum; i++ {
		initialCards[i] = i
		targetCards[i] = i
	}

	// 获取变换后，原先数组各个位置的数的新位置, 这边出牌函数可替换任意方式
	cardsAfterChange = BuildNewCards(&initialCards)
	fmt.Println(cardsAfterChange)

	// 将目标数组根据变化后的数组记录的原数组位置信息将对应位置的值插入原数组中
	// 若仅该题而言如此循环即可
	// for i := 0; i < len(cardsAfterChange); i++ {
	// 	initialCards[cardsAfterChange[i]] = i
	// }
	//
	initialCards = make([]int, cardNum, cardNum)
	for i := 0; i < len(targetCards); i++ {
		initialCards[cardsAfterChange[i]] = targetCards[i]
	}
	return initialCards
}

func main() {
	fmt.Println("第一个元素为牌顶")
	var n int = 5
	fmt.Printf("n=%d时, 初始牌堆为%v\n", n, InitialCards(n))
	n = 10
	fmt.Printf("n=%d时, 初始牌堆为%v\n", n, InitialCards(n))
}
