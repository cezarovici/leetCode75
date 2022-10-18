package main

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

func sumOfElements(nums []int) int {
	var res int
	for _, element := range nums {
		res = res + element
	}

	return res
}

func runningSumDecreasingTheLast(nums []int) []int {
	res := make([]int, cap(nums))
	var sum int

	sum = sumOfElements(nums)

	for i := len(nums) - 1; i >= 0; i-- {
		res[i] = sum
		sum = sum - nums[i]
	}

	return res
}

func main() {

}
