package problems

func canPlaceFlowers(flowerbed []int, n int) bool {
	if n == 0 {
		return true
	}

	if len(flowerbed) < 3 {
		if n >= 1 {
			return flowerbed[0] == 0 && flowerbed[1] == 0
		}
		return false
	}

	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 0 && leftEmpty(flowerbed, i) && rightEmpty(flowerbed, i) {
			n--
			flowerbed[i] = 1
		}
		if n == 0 {
			break
		}
	}

	return n <= 0
}

func leftEmpty(list []int, i int) bool {
	if i == 0 {
		return true
	}
	return list[i-1] == 0
}

func rightEmpty(list []int, i int) bool {
	if i >= len(list)-1 {
		return true
	}
	return list[i+1] == 0
}
