# 编程之美第二章节思路小结

## 索引
- [3.1 字符串移位包含问题](#31-字符串移位包含问题)

## 3.1 字符串移位包含问题
给定字符串s1和s2，要求判定s2是否能够被s1作循环移位(如abcd 左移1位变成bcda)的字符串所包含

- 解法1: 遍历所有循环移位的结果并判定是否包含，时间O(N^2)
- 解法2: 找寻规律，若保留所有循环位移的结果，则循环位移的字符串位abcdabcd，判定字符串长度不超过原长度且是否包含在新字符串内即可