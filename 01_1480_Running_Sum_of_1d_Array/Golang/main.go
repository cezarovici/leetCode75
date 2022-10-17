package main

///tmp/GoLand/___gobench_test.test -test.v -test.paniconexit0 -test.bench . -test.run ^$
//goos: linux
//goarch: amd64
//pkg: test
//cpu: Intel(R) Core(TM) i7-4600M CPU @ 2.90GHz

func runningSumDoubleFor(nums []int) []int {
	var res []int
	var sum int

	for i := 0; i < len(nums); i++ {
		sum = 0
		for j := 0; j < i+1; j++ {
			sum = sum + nums[j]
		}
		res = append(res, sum)
	}

	return res
}

//Runtime: 2 ms, faster than 77.41% of Go online submissions for Running Sum of 1d Array.
//Memory Usage: 2.8 MB, less than 14.97% of Go online submissions for Running Sum of 1d Array.

//BenchmarkRunningSumDoubleFor
//BenchmarkRunningSumDoubleFor/example1
//BenchmarkRunningSumDoubleFor/example1-4         	 9224781	       143.3 ns/op
//BenchmarkRunningSumDoubleFor/example2
//BenchmarkRunningSumDoubleFor/example2-4         	 6368623	       175.4 ns/op
//BenchmarkRunningSumDoubleFor/example3
//BenchmarkRunningSumDoubleFor/example3-4         	 6589051	       170.0 ns/op

func sumOfArray(nums []int) int {
	var res int
	for _, element := range nums {
		res = res + element
	}

	return res
}

func runningSumDecreasingTheLast(nums []int) []int {
	res := make([]int, len(nums))
	var sum int

	sum = sumOfArray(nums)

	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = sum
		sum = sum - nums[i]
	}

	return res
}

//Runtime: 2 ms, faster than 77.41% of Go online submissions for Running Sum of 1d Array.
//Memory Usage: 2.8 MB, less than 47.44% of Go online submissions for Running Sum of 1d Array.

//BenchmarkDecreasingTheLatest
//BenchmarkDecreasingTheLatest/example1
//BenchmarkDecreasingTheLatest/example1-4         	35549641	        39.08 ns/op
//BenchmarkDecreasingTheLatest/example2
//BenchmarkDecreasingTheLatest/example2-4         	26626332	        45.36 ns/op
//BenchmarkDecreasingTheLatest/example3
//BenchmarkDecreasingTheLatest/example3-4         	23636972	        42.80 ns/op

func main() {
	//PASS
	//
	//Process finished with the exit code 0
}
