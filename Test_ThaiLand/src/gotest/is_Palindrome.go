package gotest

//- Hồ sơ CPU: Xác định các hàm sử dụng nhiều CPU nhất bằng cách mỗi giây hệ thống lấy thông tin tầm 100 lần, mỗi lần là một mẫu hồ sơ. Cách gọi: go test -cpuprofile=<tên file>
//- Hồ sơ bộ nhớ heap: Xác định các hàm chiếm dụng nhiều bộ nhớ nhất. Cách gọi: go test -memprofile=<tên file>
//- Hồ sơ khóa goroutines: Xác định các hàm khóa các goroutines lâu nhất. Cách gọi: go test -blockprofile=<tên file>
//- go test -bench=. -cpuprofile=cpu.log
//- go tool pprof gotest.test cpu.log
func IsPalindrome(s string) bool {
	for i := range s {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}

//hight performance
func IsPalindrome1(s string) bool {
	n := len(s) / 2
	for i := 0; i <= n; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
