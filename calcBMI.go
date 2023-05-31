package calcBMI

func calcBMI(heigh, weight int) double {
	return double(weight) / pow(heigh)
}

func pow(num int) int {
	return num * num
}