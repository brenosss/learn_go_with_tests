package main

func Sum(numbers []int) int {
    result := 0
    for _, number := range numbers {
        result += number
    }
    return result
}

func SumAll(numbersToSum ...[]int) (sums []int) {
    var sum []int
    for _, numbers := range numbersToSum {
        sum = append(sum, Sum(numbers))
    }
    return sum
}
