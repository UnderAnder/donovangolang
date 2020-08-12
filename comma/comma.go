package comma

import "bytes"

func Comma(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	return Comma(s[:n-3]) + "," + s[n-3:]
}

func Comma2(s string) string {
	n := len(s)
	if n <= 3 {
		return s
	}
	var buf bytes.Buffer
	buf.WriteByte(s[:n-3])
	buf.WriteByte(",")
	buf.WriteByte(s[n-3:])

	return buf.String()
}

/* Упражнение 3.10. Напишите нерекурсивную версию функции comma, использую­
щую b y te s . B uffer вместо конкатенации строк.
Упражнение 3.11. Усовершенствуйте функцию comma так, чтобы она корректно
работала с числами с плавающей точкой и необязательным знаком.
Упражнение 3.12. Напишите функцию, которая сообщает, являются ли две строки
анаграммами одна другой, т.е. состоят ли они из одних и тех же букв в другом порядке.
*/
