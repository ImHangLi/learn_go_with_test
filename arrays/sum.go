/**
We have covered
- Arrays
- Slices
- Use make() to construct a slice
- How they have a fixed capacity but you can create new slices from old ones using append
- How to slice slices!
- len() to get the length of an array or slice
- Test coverage tool -cover
- reflect.DeepEqual and why it's useful but can reduce the type-safety of your code
*/

package main

func Sum(numbers []int) (sum int) {
	for _, number := range numbers {
		sum += number
	}

	return
}

func SumAllTails(numbersToSum ...[]int) []int {
	sum := []int{}

	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sum = append(sum, 0)
		} else {
			sum = append(sum, Sum(numbers[1:]))
		}
	}

	return sum
}
