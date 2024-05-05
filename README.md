# Напишите программу, подсчитывающую сколько раз буква встречается в предложении, а также частоту встречаемости в процентах.

То есть в предложении hello world результатом будет.
h - 1 0.1
e - 1 0.1
l - 3 0.33
o - 2 0.2
w - 1 0.1
r - 1 0.1
d - 1 0.1

Мы уже решали похожую задачу на предыдущем семинаре.
Помните, что нужно подсчитывать только буквы, не учитывая пробельные символы или знаки препинания?
Для решения задачи используйте составной тип map.
## Подсказка 1: 
H и h это разные символы, приведите их к общему виду с помощью функции unicode.ToLower().
## Подсказка 2: 
для вывода с нужной точностью используйте специальные глаголы %f, например для вывода числа с плавающей точкой с точностью 2 можно использовать функцию fmt.Printf("%0.2f\n", 10.33333).
## Подсказка 3: 
для вывода руны, нужно использовать глагол %c, то есть fmt.Printf("%c\n", 'Н').