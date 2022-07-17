package homework

func sortMapValues(input map[int]string) (result []string) {
	map_keys := make([]int, 0, len(input))	
	for key := range input {
    	map_keys = append(map_keys, key)		
	}
	for i := 0; i < len(map_keys); i++ {
		for k := len(map_keys) - 1; k > i; k-- {
			if map_keys[k-1] > map_keys[k] {
				map_keys[k-1], map_keys[k] = map_keys[k], map_keys[k-1]
			}
		}
	}
	for _, k := range map_keys {
		result = append(result, input[k])
	}
	return
}
