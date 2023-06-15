package money

type Money int

func Change(money Money, denomination []int) []int {
	var change = make([]int, 0)

	for money > 0 {
		for i := 0; i < len(denomination); i++ {
			coin := denomination[i]
			for money-Money(coin) >= 0 {
				money -= Money(coin)
				change = append(change, coin)
			}
		}
	}

	return change
}

// }
