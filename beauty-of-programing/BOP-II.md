# 编程之美第二章节思路小结

## 索引
- [2.1 求2进制数中1的个数](#21-求2进制数中1的个数)
- [2.2 (1) 求N阶乘中末尾0的数量？ (2) 求N阶乘二进制表示中末尾1的位数](#22-1-求n阶乘中末尾0的数量-2-求n阶乘二进制表示中末尾1的位数)
- [2.3 寻找发帖"水王"(即在一组无序id中找出超过一个出现次数超过50%的id(假设一定存在))](#23-寻找发帖水王即在一组无序id中找出超过一个出现次数超过50的id假设一定存在)
- [2.4 求1 ~ N(给定正整数)，十进制下 1 出现的次数](#24-求1--n给定正整数十进制下-1-出现的次数)
- [2.5 寻找最大的K个数](#25-寻找最大的k个数)
- [2.6 精准表达浮点数](#26-如何精确表达浮点数)
- [2.7 求最大公约数](#27-求最大公约数)
- [2.8 找寻符合条件的整数](#28-找寻符合条件的整数)
- [2.9 计算斐波那契数列](#29-计算斐波那契数列)
- [2.10 寻找数组中最大最小值](#210-寻找数组中最大最小值)
- [2.11 寻找最近点对](#211-寻找最近点对)
- [2.12 快速寻找满足条件的两个数](#212-快速寻找满足条件的两个数)
- [2.13 子数组最大的乘积](#213-子数组最大的乘积)
- [2.14 求数组的子数组之和的最大值](#214-求数组的子数组之和的最大值)
- [2.15 求二维数组的子数组之和的最大值(书p189, 未细究)](#215-求二维数组的子数组之和的最大值)
- [2.16 求数组中最长递增子序列](#216-求数组中最长递增子序列)
- [2.17 数组循环移位](#217-数组循环移位)
- [2.18 数组分割](#218-数组分割)
- [2.19 区间重合判断](#219-区间重合判断)

## 2.1 求2进制数中1的个数
- 解法1: 根据短除法十进制转二进制的思路，每除2取余，若余1，则1的个数加1，时间复杂度O(logN)
- 解法2: 基本思路同上，与1作‘与’运算，若大于0，则1的个数加1，然后数字左移一位重复上述步骤，时间复杂度O(logN)
- 解法3: 当给定的数N大于0时，N & (N - 1), 即移除最左侧的一个1， 重复此运算，直到结果等于0，在此之前每做一次该运算，1的个数加1，时间复杂度只与1的个数有关
  - 扩展: 判断给定的数是否是2的整数次幂？ A: 由于2的整数次幂1的个数为1，故解法3可用于此：N >= 0 && (N & (N - 1)) == 0
- 解法4：当给定的整数限制长度的情况下，空间换时间，用一个数组按顺序列出每一个数中1的个数，以给定的数为索引，可直接得出结果，时间复杂度O(1)

## 2.2 (1) 求N阶乘中末尾0的数量？ (2) 求N阶乘二进制表示中末尾1的位数

### 核心思路
将N阶乘的结果进行质因数分解，即结果可以分解为 N! = 2^x * 3^y * 5^z * 7^a * ... * 其他质数的整数幂的积

- 求N阶乘中末尾0的数量取决于可以被10的m次幂整除，只有2 * 5 = 10 即 m = Min(x, z), 由于能被2整除的数远多于5，故其结果基本上等同于 m = z,即 N!后有多少个0取决5的指数。
  - 可用以下公式z = f(1) + f(2) + f(3) + f(4) + ... + f(N), 其中f(x)为能被5整除的次数;
  - 上述公式可以进一步简化: z = (N / 5) + (N / 5^2) + (N / 5^3) + (N / 5^4) + ...(直到 5^k > N), 其中(N / 5)表示1 ~ N中能被5整除的数量，(N / 5^2)表示1 ~ N中被5整除后能再整除一次的数量。。。
  - 举例： 当N = 27时 1 ~ 27中有(27/5=5)个数(即5，10，15，20，25)可以被5整除，其中有(27/25=1)个数可以再被5整除一次故 z = 27/5 + 27/5^2 = 5 + 1 = 6 可以分解出 6个5
- 二进制表示中末尾1的位数等于末尾0的的数量加1, 二进制中末尾每多一个0便可以被2整除一次,即结果为(x + 1)
  - x的解法参照第一个问题:  z = (N / 2) + (N / 2^2) + (N / 2^3) + (N / 2^4) + ...(直到 2^k > N)

## 2.3 寻找发帖"水王"(即在一组无序id中找出超过一个出现次数超过50%的id(假设一定存在))
- 解法1: 排序后遍历集合， 当超过某预定阈值后返回该id, 时间复杂度O(NlogN + N),若阈值为其他值该解法也适用
  - 解法1特例: 当阈值大于50%时，排序后中间的那个Id一定为要找的Id, 时间复杂度O(NlogN)

### 核心思路

在保留集合特征的情况下缩小规模,例如: 移除2个不同的Id, 剩余的集合中某一个Id依旧超过50%, 由于会满足这个特性，当集合长度为2(长度为偶数)或1(长度为奇数)时, 剩余的Id 便是要找的ID

- 解法2: 遍历该集合, 每一次循环取一个值保留，并计数为1，若下一次循环的值不等于保留的值，则清除保留的值(即从集合中移除两个不同的值)，若和保留的值相同，则保留的值得计数为1，需要在接下来的循环中有同等数量与保留值不一样的数字出现才能清除保留的值(即当计数为0的时候，视作移除了n组不同的值), 最后保留的值中的Id便是核心思路中最后剩余的Id
  - 扩展问题: 若一组集合中有3个Id出现次数超过(1/4),找出这三个ID? A：即每次移除4个不同的ID，即解法2保留3个值

## 2.4 求1 ~ N(给定正整数)，十进制下 1 出现的次数
例如: 当N等于12时,1 2 3 4 5 6 7 8 9 10 11 12, 其中1出现了 1(1) + 1(10) + 2(11) + 1(12), 总共5次

- 解法1：遍历1 ~ N, 统计每一个数中1的数量(每一个数除10取余，若余数1为1，则1的出现次数+1，直到该值被除至小于0为止)，时间复杂度O(N * logN)
- 解法2: 找规律?
  - 设 N = abcde 若要统计百位上的1的次数:
  - 若c = 0, 则 出现1的次数为ab * 100(即100 ~ 199, 1100 ~ 1199, 2100 ~ 2199, 3100 ~ 3199 ... a(b-1)100 ~ a(b-1)199)
  - 若c = 1, 则 出现1的次数为ab * 100 + (de + 1)(即上述情况多出ab100 ~ab1de)
  - 若c > 1, 则 出现1的次数为(ab + 1) * 100(即c=0情况多出ab100 ~ab199)

## 2.4.2 上一个问题求出f(N), 求满足条件的N 最大是多少

- 解法: 作假设
  - f(9) = 1
  - f(99) = 20
  - f(999) = 300
  - f(9999) = 4000
  - 5) 导出 f(10^n - 1) = n * 10^(n-1)
  - 当n = 10时，f(10^10 - 1) > 10^10 - 1
  - 假设存在特定值(设为N)的时候, f(n) 恒大于 n
  - 即最大值存在于 1 ~ N之间, 只要从大往小遍历即可
  - 由5)的公式可看出，每增加10, f(n)至少增加1， 每增加100，f(n)至少增加20, 以此类推每增加10^k f(n)增加k * 10^(k-1)
  - f(n) > n => f(0 + a * 10^k + b * 10^(k-1) + c * 10^(k - 2)...) > a * k * 10^(k-1) + b * (k - 1) * 10^(k-2) + .... // 常数a>=1的正整数
  - => 由上式导出 => 当不等式右边 减去 省略项后，不等式依旧成立
  - f(n) > n => f(0 + a * 10^k + b * 10^(k-1) + c * 10^(k - 2)...) > a * k * 10^(k-1) + b * (k - 1) * 10^(k-2)
  - => f(n) > a * k * 10^(k-1) + b * (k - 1) * 10^(k-2)
  - 用类似规则社区n的省略项， n =  a * 10^k + b * 10^(k-1) + c * 10^(k - 2).... < a * 10^k + (b + 1) \* 10^(k-1)
  - => a * 10^k + (b + 1) \* 10^(k-1) > n
  - f(n) > a * k * 10^(k-1) + b * (k - 1) * 10^(k-2) >=  a * 10^k + (b + 1) * 10^(k-1) > n
  - 解当中部分的不等式: k>= 10 + ((b + 11) / (b + 10a)), 由于常数a>=1的正整数, 故((b + 11) / (b + 10a)) < 2, 即k>=12时满足不等式
  - 即当看k = 12时, f(n) > n恒成立, 即f(1 * 10^12) > 1 * 10^12 恒成立, 即上限值为10^12 - 1，由此数往下遍历即可
  - (此处与书p137有出入)

## 2.5 寻找最大的K个数
N个无序的数，假定各不相等，选出最大的k个数

- 解法1: 快速排序后取前k个，时间复杂度O(N * logN) + O(K) = O(N * logN)
- 解法2：利用快速排序的特性，一个基准数分成两组，若较大的一组数量大于K，则在较大的一组中找最大的k个数, 若数量少于k, 则在较小的一组中找（ K-该数量 - 1(基准点)）个最大的数 ，于较大的那组数即为求的结果, 评价时间复杂度是O(N\*logK)
- 解法3：题目可以理解成寻找第k大的整数, 先遍历一次，取出最大值和最小值，利用2分查找法查找第k大的数,基本操作类似于解法1，先查找比最大值和最小值的中值，然后查找比他大的集合, 若数量大于k, 则在该分区中继续2分寻找第k大的数，若数量小于k, 则在另外一个分区中找剩下的最大值, 时间复杂度O(N * logN)
- 解法4: 先取出前k个数构建一个最小堆(最小堆的根节点为该集合的最小值), 每从剩余的集合中取出一个值将其与根节点的值作对比，若比他小，则舍弃继续下一个值，若比他大则替换根节点的值并重新维护该最小堆，此方法不需要一次加载所有的值。
时间复杂度为O(NlogK)
- 解法5: 依旧类似空间换时间，仅限于N个数中的取值范围不大，维护一个数组记录每个数出现的次数找出第k大的值

## 2.6 如何精确表达浮点数
即用分数表达浮点数，分母尽可能小

- 解法1:
  - 先拆分成整数和小数的和，只考虑小数部分(0 ~ 1)
  - 有限小数0.a1a2a3...aN => a1a2a3...aN/10^N => 约去分子分母的最大公约数
  - 无限小数0.a1a2a3...aN(b1b2b3...bM) => (a1a2a3...aN + 0.(b1b2b3..bM))/10^N =>
  - 略去有限部分最后加回，=> 设Y = 0.(b1b2b3..bM) => 10^M Y = b1b2..bM + 0.(b1b2..bM) => 10^M Y = b1b2..bM + Y
  - 接上 => (10^M - 1)Y = b1b2b3..bM => Y= b1b2b3..bM/(10^M-1)
  - 带回原式 =>((a1a2..aN)(10^M-1) + b1b2b3...bM) / (10^M - 1)10^N
  - 求最大公约数

## 2.7 求最大公约数
求两个数字的最大公约数

- 解法1: 辗转相除法，两个整数的最大公约数为两个数的余数(大除小)和较小值得余数和较小的数的最大公约数，以此为基础递归直至较小数为0，此时较大的值是最大公约数
- 解法2: 更相减损数, 两个整数的最大公约数为两个数的差和较小数的最大公约数, 直至较小数为0时，较大的数为最大公约数
- 解法3：设求最大公约数的函数为f(x, y), 设p为质数
  - 若x=p\*x1, y=p\*y2，则f(x,y)=p\*f(x1, y1)
  - 若x=p\*x1, y%p != 0，则f(x,y)=f(x1/p, y)
  - 设p=2,结合解法1和解法2
  - 若x,y 为偶数，f(x,y) = 2 * f(x>>1, y>>1)
  - 若x 为奇数，y为偶数，f(x,y) = f(x, y>>1)
  - 若x,y 为奇数，f(x,y) = (x,x-y)
  - 以此递归

## 2.8 找寻符合条件的整数
任意给定一个正整数N,求一个最小的正整数M(M>1), 使得N\*M的十进制形式中只有1和0

- 解法0: 从1开始从小到大枚举，找到符合条件的值为止
- 解法1: 问题转换为求一个最小的一个仅有1和0表示的10进制数(设为X)且能被N整除
  - X的取值有1, 10, 11, 100, 101, 110, 111....
  - 找规律, 设 N为3，由1%3 == 10%3 => 101%3 == 110 % 3
  - 推断出规律:
  - 假设已经遍历了X个数, 其中X一共有K位, 同时也计算了10^K(K+1位的第一个数字)的结果, 设10^K % N = a
  - 现在需要计算K+1位其他的数, 即 X=10^K(K+1位) + Y(K位), 遍历的话需要遍历2^K - 1个数
  - 若是对2^K - 1个数按照除N的余数进行分组，最多有N - 1组(余数为1 ~ N-1)
  - 只需要计算10^K + 每组中最小的数是否能被N整除即可, 因为(组里其他数+10^K)除N的余数肯定相同(因为同一组数里的差值一定是N的整数倍)
  - 实际操作
  - 维护一个 0 - (N-1) 的分组, 初始化情况下每个分组中都没有数，假设需要验证K + 1位的所有的数是否可被整除，只需要验证(10^K % N) + 当前已存的余数的和除N是否有产生新的余数即可，根据结果补充分组
  - 按照此规则循环，直至连续有N次不更新数组为止(因为0 ~ N - 1共N个可能的余数)
  - 提前获得结果可提前结束循环

## 2.9 计算斐波那契数列
1, 1, 2, 3, 5, 8....

- 解法0: 递归运算
- 解法1: 类似于动态规划, 找规律，转为循环运算，降低空间复杂度至O(1)
- 解法2: 求通项式，书P162

## 2.10 寻找数组中最大最小值
设数组长度为N

- 解法1: 遍历一次数组, 比较2N次(即先遍历比较一轮求出最大值)
- 解法2: 将数组分为2组, 较大的一组和较小的一组, 比较相邻的两个数字(如(0，1) ，(2，3)...,共需要比较N/2次), 将较小的数放在偶数位, 较大的数放在奇数位, 分别在奇数位组合偶数位组比出最大值和最小值即可(各需要比较N/2次), 共需要比较1.5N次
- 解法3: 解法2优化为不改变原数组, 2个一组遍历数组, 第一组将较大值设为当前最大值，较小值设为当前最小值, 接下去每组先比较出大小后，分别与当前最大最小值比较并更新, 比较次数与解法2相等，但是不用修改原先的数组
- 解法4: 依然是用分治理思想. 分两组，分别求出最大最小值，分别比较最大最小值，得出结果，用递归实现，比较次数并未减少

## 2.11 寻找最近点对
给定平面上N个点的坐标，找出距离最近的两个点

- 解法1: 两两计算点与点之间的距离，拿出最短的一对, 时间复杂度O(N^2)
- 解法2: 题目的退化版，所有的点都在横坐标上，可以先将点排序O(N^2), 在遍历集合(时间复杂度O(N)), 找出距离最近的点对，总体时间复杂度为O(NlogN + N) = O(N(logN+1)) = O(NlogN), 但是此方法在平面坐标系上不适用
- 解法3: 运用分支思想，在水平方向上将N个点分成左右两部分Left和Right, 距离最近的点要么在Left中(设最短距离minLeft)，要么在Right(minRight)中, 又或者一个点在Left, 一个点在Right，该情况下, 符合的条件两点肯定在以分割线为中心，左右水平距离小于minDist = Min(minLeft, minRight)的范围内, 将区域内的点按纵坐标排序，依次遍历这些点之间的距离。得出结果. 通过递归实现, 时间复杂度O(NlogN)

## 2.12 快速寻找满足条件的两个数
在一个数组中找到和等于给定值的两个数

- 思路: 对于元素i, 查找sum - i,在不在剩余数组中
- 解法1: 两两枚举，时间复杂度O(N^2)
- 解法2: 先用O(NlogN)的方法排序, 再二分查找O(logN), 总体时间复杂度O(NlogN)
  - O(N)遍历hash表，O(1)查找，总体时间复杂度是O(N), 属于空间换时间
- 解法3: 先用O(NlogN)的方法排序, 用两个标识分别从首尾开始计算两个和并和给定值比较，若小了，将首坐标后移一位，反之尾坐标前移一位，时间总体复杂度O(NlogN)

## 2.13 子数组最大的乘积
给定任意一个长度为N的整数数组(有正数，负数，0), 只允许用乘法，不能用除法，找到任意(N-1)个数组的组合中乘积最大的一组, 并给出时间复杂度

- 解法0, 直接遍历所有可能, 时间复杂度O(N^2)
- 解法1, 空间换时间, 把(N-1)长度子数组看成被舍弃的那一个元素左边的集合和右边的集合的并集，子数组的乘积为两个集合内元素乘积的乘积, 故只要分别从头至尾和从尾至头遍历后并记录下左右两个集合的积，再组合后遍历一次即可获得所有结果，时间复杂度O(3N) = O(N)
- 解法2, 分析正负及0值特性,先计算出原数组所有元素的乘积P
  - 当P=0时, 若集合内0的个数大于1, 则返回0，若0的个数为1，若剩余乘积为负数，则返回0, 若剩余乘积为正数，则返回剩余乘积。
  - 当P>0时, 若集合内有正数，则移除最小的正数，若集合内全是负数，则移除绝对值最大的负数(这样N-1的集合乘积小于0，但是在负数中最大)
  - 当P<0时, 移除绝对值最小的负数, 返回剩余乘积

## 2.14 求数组的子数组之和的最大值
一个长度为N的一维整数数组, 求子数组(必须连续)之和的最大值(不需要返回子数组位置)

- 解法0: 遍历所有子数组的可能并求出和, 时间复杂度为O(N^2)
- 解法1: 运用分治思想，将数组分为2个长度n/2的数组, 分别求出2个数组最大的子数组之和, 以及跨分组的最大子数组之和(即分别以分隔点为首尾子数组之和), 递归求解，时间复杂度O(N\*logN)
- 解法2: 设数组第一个元素为A[0], 数组第2个元素到N个元素的最大子数组之和Sum(N-1), 则长度为N的子数组的和的最大值为max(A[0], Sum(N-1), Sum(N-1)+A[0]), 递归实现, 时间复杂度 O(N)

## 2.15 求二维数组的子数组之和的最大值
一个长度为M, 高度为N的二维整数数组, 求子数组(必须连续, 即矩形区域)之和的最大值(不需要返回子数组位置)

- 解法0: 遍历所有子数组的可能并求出和, 时间复杂度为O(N^2)
- 解法2: 先枚举纵向结果，再横向参照一维求解

## 2.16 求数组中最长递增子序列
子序列在数组中不需要连续，仅要求长度

- 解法0: 遍历所有的可能， 时间复杂度O(N^2 + N) = O(N^2)
- 解法1: 运用动态规划的思想, 从数组长度为1的情况最长递增子序列长度为1, 并记录下各个长度递增子序列的最大值中(同长度子序列不一定只有一个)的最小值, 不过时间复杂度依旧是O(N^2)i(因为需要遍历所有长度子序列的最大值的最小值小于新值，长度仅可能长, 若存在则新值所在位置的数组长度的最长子序列的长度为该旧子序列的长度+1, 若这边使用2分查找，时间复杂度会变为O(N\*logN)), 更新各个长度子序列的最大值

## 2.17 数组循环移位
将长度为N的数组循环右移K位, 时间复杂度需要是O(N)，只允许额外2个变量

- 解法1: 设数组为[a1, a2, a3...., k1, k2...kK], 将数组看成左右两部分，右移k位即交换两部分的数组, 关于右移k位，最终结果等同于K%N, 如何交换两部分数组:
  - 先分别单独逆序两部分数组
  - 再逆序整个数组

## 2.18 数组分割
将长度为2n的两个正整数的数组，将其分为两组长度各为n的子数组，并使两个子数组的和最接近

### 背包问题
有一个长度为N的正整数的集合，集合中选出部分元素使其的和小于等于给定值J
- 背包问题递归式为: 设S(i, j)为前i个元素可放入容量为j的背包的最大值
- S(i, j) = 0 (i = 0 or j = 0), S(i, j) = max(s(i - 1, j), s(i - 1, j - 新的一个元素大小) + 新的元素大小) (0 < i <= n, 0 < j < 给定值)


- 解法1: 设总数组长度之和是SUM, 两个子数组的和可能一个大于等于SUM/2, 另一个小于等于SUM/2, 一定程度上可转化为背包问题, 在背包问题上限制元素的个数, 设S(i, k, j)为前i个元素中挑了k个元素可放入容量为j的背包的最大值
- 递归式为S(i, k, j) = 0 (i = 0 or j = 0 or k = 0), S(i, k, j) = max(s(i - 1, k, j), s(i - 1, k -1 , j - 新的一个元素大小) + 新的元素大小) (0 < i <= 2n, 0 < j < SUM/2, 0 < k < n )
- 整体时间复杂度为O(2^N), 因为每多一个数需要计算前两个的情况, 其中也许包括一些重复计算
- 解法2(未实践)，在2的基础上从底下开始列出一个多维表格找寻规律，可以优化时间复杂度至O(N^2 * Sum)

## 2.19 区间重合判断
判断源区间[x, y]是否在给定N个无序区间内[a, b], [c, d]..., [x,y]需要完全被覆盖, 例: [1,6]在[2,3], [1,2], [3,6]内

- 解法1: 逐个从源区间中移除被给定区间覆盖的部分，([1,6]) => ([1,2], [3,6]) =>([3, 6])=>(), 总体时间复杂度O(N^2)
- 解法2: 先对给定区间(左边界)排序O(NLogN)并合并O(N)成N个不相交的区间, 判断源区间是否在某一区间内，此处用二分查找O(logN), 总体时间复杂度为O(NlogN)
