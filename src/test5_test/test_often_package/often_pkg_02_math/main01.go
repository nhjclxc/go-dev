package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

func main() {

	fmt.Printf("1")
	math.Abs(2)

	rand.Seed(time.Now().UnixNano()) // 用当前时间设置种子
	fmt.Println(rand.Int())
	fmt.Println(rand.Intn(10)) // 返回一个 [0, n) 范围内的随机数
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))
	fmt.Println(rand.Intn(10))

	/*
	   Pi    = 3.14159265358979323846  // 圆周率
	   E     = 2.7182818284590452354   // 自然常数
	   Phi   = 1.6180339887498948482   // 黄金分割比
	   Sqrt2 = 1.4142135623730950488   // √2


	✅ 二、基础数值运算函数
	| 函数                     | 说明          | 示例                                 |
	| ---------------------- | ----------- | ---------------------------------- |
	| `math.Abs(x)`          | 绝对值         | `math.Abs(-5.2) // 5.2`            |
	| `math.Max(x, y)`       | 返回最大值       | `math.Max(2, 5) // 5`              |
	| `math.Min(x, y)`       | 返回最小值       | `math.Min(2, 5) // 2`              |
	| `math.Mod(x, y)`       | 浮点数取模 x % y | `math.Mod(7, 2.5) // 2.0`          |
	| `math.Remainder(x, y)` | IEEE 余数     | `math.Remainder(5.5, 2.0) // -0.5` |


	✅ 三、指数与对数函数
	| 函数                 | 说明                  | 示例                         |
	| ------------------ | ------------------- | -------------------------- |
	| `math.Exp(x)`      | e 的 x 次幂            | `math.Exp(1) // ≈ 2.71828` |
	| `math.Log(x)`      | 自然对数 ln(x)          | `math.Log(math.E) // 1`    |
	| `math.Log10(x)`    | 以10为底的对数            | `math.Log10(100) // 2`     |
	| `math.Pow(x, y)`   | 幂运算 x^y             | `math.Pow(2, 3) // 8`      |
	| `math.Sqrt(x)`     | 平方根 √x              | `math.Sqrt(16) // 4`       |
	| `math.Cbrt(x)`     | 立方根                 | `math.Cbrt(27) // 3`       |
	| `math.Hypot(x, y)` | 返回 √(x² + y²)，即斜边长度 | `math.Hypot(3, 4) // 5`    |


	✅ 四、三角函数与角度转换
	| 函数                | 说明       | 示例                           |
	| ----------------- | -------- | ---------------------------- |
	| `math.Sin(x)`     | 正弦函数     | `math.Sin(math.Pi/2) // ≈ 1` |
	| `math.Cos(x)`     | 余弦函数     | `math.Cos(0) // 1`           |
	| `math.Tan(x)`     | 正切函数     | `math.Tan(math.Pi/4) // ≈ 1` |
	| `math.Asin(x)`    | 反正弦      | `math.Asin(1) // ≈ π/2`      |
	| `math.Acos(x)`    | 反余弦      | `math.Acos(0) // ≈ π/2`      |
	| `math.Atan(x)`    | 反正切      | `math.Atan(1) // ≈ π/4`      |
	| `math.Atan2(y,x)` | y/x 的反正切 | `math.Atan2(1,1) // π/4`     |


	✅ 五、取整与舍入
	| 函数              | 说明            | 示例                       |
	| --------------- | ------------- | ------------------------ |
	| `math.Floor(x)` | 向下取整          | `math.Floor(2.9) // 2`   |
	| `math.Ceil(x)`  | 向上取整          | `math.Ceil(2.1) // 3`    |
	| `math.Trunc(x)` | 截断小数部分（取整数部分） | `math.Trunc(-2.9) // -2` |
	| `math.Round(x)` | 四舍五入          | `math.Round(2.5) // 3`   |


	✅ 六、特殊数值处理
	| 函数                    | 说明                      | 示例                           |
	| --------------------- | ----------------------- | ---------------------------- |
	| `math.IsNaN(x)`       | 判断是否为 NaN               | `math.IsNaN(0/0)`            |
	| `math.IsInf(x, sign)` | 是否为正无穷/负无穷              | `math.IsInf(math.Inf(1), 1)` |
	| `math.Copysign(x, y)` | 返回 `x` 的绝对值，并附加 `y` 的符号 | `math.Copysign(3, -1) // -3` |


	✅ 七、浮点比较技巧
	由于浮点数精度误差，直接用 == 比较常常不可靠。推荐
	func floatEquals(a, b float64) bool {
	    return math.Abs(a - b) < 1e-9
	}




	*/
	math.Hypot(3, 4)




}
