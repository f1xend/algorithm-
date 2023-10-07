package intersectarr

import "fmt"

func Intersect(nums1 []int, nums2 []int) []int {
	// Создаем словарь для хранения частоты элементов в nums1
	frequency := make(map[int]int)
	for _, num := range nums1 {
		frequency[num]++
	}
	fmt.Println(frequency)

	// Создаем результирующий слайс
	var res = make([]int, 0, min(len(nums1), len(nums2)))

	// Проверяем каждый элемент в nums2
	for _, num := range nums2 {
		// Если элемент присутствует в словаре и его частота больше 0, добавляем его в результирующий слайс
		if frequency[num] > 0 {
			res = append(res, num)
			frequency[num]--
		}
	}

	return res
}
