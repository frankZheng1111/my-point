package main

// 2018chengxuyuanxiaohui: 八皇后问题
// 8 * 8的棋盘上摆放八个皇后(棋子), 其中任意一个棋子的同一行，同一列，同一斜线不能出现其他皇后(棋子)

// 相关思路, 递归回溯

import (
	"errors"
	"fmt"
	"log"
)

const WIDTH = 8
const HEIGHT = 8

type CheckerBoard [WIDTH][HEIGHT]int

func (checkerBoard *CheckerBoard) Print() {
	for i := 0; i < HEIGHT; i++ {
		for j := 0; j < WIDTH; j++ {
			fmt.Printf("%d ", checkerBoard[i][j])
		}
		fmt.Printf("\n")
	}
}

// index是第一个索引即第几行
// subIndex是第一个索引即第几列
func (checkerBoard *CheckerBoard) IsPositionValid(index int, subIndex int) bool {
	// 遍历第index行是否已有皇后(ps, 按照本题的解法无需查询横向, 但是就方法名声明的行为则需要检查)
	for i := 0; i < WIDTH; i++ {
		if checkerBoard[index][i] != 0 {
			return false
		}
	}
	// 遍历第subIndex列是否已有皇后
	for i := 0; i < HEIGHT; i++ {
		if checkerBoard[i][subIndex] != 0 {
			return false
		}
	}

	// 遍历该位置上左侧的斜线是否有皇后(包括当前位置)
	for i, j := index, subIndex; i >= 0 && j >= 0; i, j = i-1, j-1 {
		if checkerBoard[i][j] != 0 {
			return false
		}
	}

	// 遍历该位置上右侧的斜线是否有皇后
	for i, j := index+1, subIndex-1; i < WIDTH && j >= 0; i, j = i+1, j-1 {
		if checkerBoard[i][j] != 0 {
			return false
		}
	}

	// 遍历该位置下左侧的斜线是否有皇后(不包括当前位置)
	for i, j := index-1, subIndex+1; i >= 0 && j < HEIGHT; i, j = i-1, j+1 {
		if checkerBoard[i][j] != 0 {
			return false
		}
	}

	// 遍历该位置下右侧的斜线是否有皇后(不包括当前位置)
	for i, j := index+1, subIndex+1; i < WIDTH && j < HEIGHT; i, j = i+1, j+1 {
		if checkerBoard[i][j] != 0 {
			return false
		}
	}

	return true
}

// 设定第一个符合条件的结果
func (checkerBoard *CheckerBoard) SetBoard() error {
	var recursiveFunc func(currentRow int) bool
	recursiveFunc = func(currentRow int) bool {
		if currentRow == HEIGHT {
			// 利用返回一个true 来表示找到结果
			return true
		}
		for i := 0; i < WIDTH; i++ {
			// 若该行该位置不合法，尝试该行下一个位置
			if !checkerBoard.IsPositionValid(currentRow, i) {
				continue
			}
			checkerBoard[currentRow][i] = 1
			if recursiveFunc(currentRow + 1) {
				// 此处的结果用于把已经找到的消息返回最上层
				return true
			} else {
				checkerBoard[currentRow][i] = 0
			}
		}
		// 循环结束未找到合适位置，返回not ok
		return false
	}
	if !recursiveFunc(0) {
		return errors.New("MissingResult")
	}
	return nil
}

// 打印所有符合条件的结果
func PrintAllBoards() {
	var checkerBoard CheckerBoard
	var recursiveFunc func(currentRow int)
	var totalCount int
	recursiveFunc = func(currentRow int) {
		if currentRow == HEIGHT {
			// 核心在于只要逐层遍历只要找到合适的点即进行下一层遍历，直到遍历完所有层后作为一个合格的棋盘输出
			fmt.Println("----------------")
			totalCount++
			checkerBoard.Print()
			return
		}
		for i := 0; i < WIDTH; i++ {
			// 若该行该位置不合法，尝试该行下一个位置
			if !checkerBoard.IsPositionValid(currentRow, i) {
				continue
			}
			checkerBoard[currentRow][i] = 1
			recursiveFunc(currentRow + 1)
			checkerBoard[currentRow][i] = 0
		}
		return
	}
	recursiveFunc(0)
	fmt.Println("总结果一共有多少种: ", totalCount)
	return
}

func main() {
	var board CheckerBoard
	err := board.SetBoard()
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("输出棋盘: ")
	board.Print()
	fmt.Println("输出所有符合条件的棋盘: ")
	PrintAllBoards()
}
