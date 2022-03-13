package array

// PrimeNum 统计n以内素数个数，暴力解法
// 素数，只能被1和本身除的叫素数, 0，1不是素数
func PrimeNum(n int) int {
	if n < 2 {
		return 0
	}
	var count int
	for i := 2; i <= n; i++ {
		if isPrime(i) {
			count++
		}
	}
	return count
}

// isPrime 判断是否是素数
func isPrime(x int) bool {
	//for i := 2; i < x; i++ {
	for i := 2; i*i <= x; i++ {
		if x%i == 0 {
			return false
		}
	}
	return true
}

// Eratosthenes 埃筛法
func Eratosthenes(n int) int {
	if n < 2 {
		return 0
	}

	count := 0
	flag := make([]bool, n)
	for i := 2; i < n; i++ {
		if !flag[i] {
			count++
			//for j := 2 * i; j < n; j += i {
			for j := i * i; j < n; j += i {
				flag[j] = true
			}
		}
	}
	return count
}
