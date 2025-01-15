/** 
What we have covered:
- More TDD practice
- Learned for
- Learned how to write benchmarks
*/

package iteration

func Repeat(character string, time int) string {
	var repeated string

	for i := 0; i < time; i++ {
		repeated += character
	}

	return repeated
}