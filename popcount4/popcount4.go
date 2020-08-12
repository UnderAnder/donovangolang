package popcount

// pc[i] - количество единичный битов в i
var pc [256]byte

// PopCount возвращает степень заполнения
// (количество установленных битов) значения х.
func PopCount(x uint64) int {
	var m int

	for i := uint(0); i < 64; i++ {
		if x&(x-1) == 0 {
			m++
		}
		//m = pc[byte(x>>(i*64))]
	}

	return int(m)
}

/*	Упражнение 2.5. Выражение х&(х-1) сбрасывает крайний справа ненулевой
	бит х. Напишите версию PopCount, которая подсчитывает биты с использованием
	этого факта, и оцените ее производительность.*/
