package homework

func average(input [15]float32) (result float32) {
	var sum float32
	//var average float32
	for _, elem := range input {
		sum += elem
	}
	result = sum / float32(len(input))
	return
}
