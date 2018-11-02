def abundantNumberGen(limit): # returns list of abundant numbers up to limit
	numbers = []
	for x in range(1, limit):
		if sum(properFactors(x)) > x:
			numbers.append(x)
	return numbers

def properFactors(n): # returns list of proper factors of a number
	factors = [1] # 1 is automatically a factor
	for x in range(2, int(n**0.5)+1):
		if n % x == 0:
			factors.extend([x, n/x])
	if n**0.5 == int(n**0.5): # important! must remove the extra square root if n is a perfect square
		factors.remove(n**0.5)
	return factors

result = list(range(28123)) # list of numbers up to 28123; result[12] = 12, result[28122] = 28122
abundantNumbers = abundantNumberGen(28123)
for n in abundantNumbers:
	for m in abundantNumbers:
		if n + m < 28123:
			result[n + m] = 0 # result[100] = 100 --> result[100] = 0

print(sum(result))
