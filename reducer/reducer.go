package reducer

//Reducer get result from all mapper and reduce to final result
func Reducer(mapList chan map[string]int, result chan map[string]int) {
	final := make(map[string]int)
	for list := range mapList {
		for element, value := range list {
			final[element] += value
		}
	}
	result <- final
}
