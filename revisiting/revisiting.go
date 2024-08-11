package revisiting

type Account struct {
	Name    string
	Balance float64
}

type Person struct {
	Name string
}

type Transaction struct {
	From string
	To   string
	Sum  float64
}

func NewTransaction(from, to Account, sum float64) Transaction {
	return Transaction{From: from.Name, To: to.Name, Sum: sum}
}

func NewBalanceFor(account Account, transactions []Transaction) Account {
	return Reduce(
		transactions,
		applyTransaction,
		account,
	)
}

func Find[A any](items []A, predicate func(A) bool) (value A, found bool) {
	for _, v := range items {
		if predicate(v) {
			return v, true
		}
	}

	return
}

func applyTransaction(a Account, transaction Transaction) Account {
	if transaction.From == a.Name {
		a.Balance -= transaction.Sum
	}

	if transaction.To == a.Name {
		a.Balance += transaction.Sum
	}

	return a
}

func Sum(numbers []int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum
}

func SumAll(numbersToSum ...[]int) []int {
	var sums []int
	for _, number := range numbersToSum {
		sums = append(sums, Sum(number))
	}

	return sums
}

func SumAllTails(numbersToSum ...[]int) []int {
	var sums []int
	for _, number := range numbersToSum {
		if len(number) == 0 {
			sums = append(sums, 0)
		} else {
			sums = append(sums, Sum(number[1:]))
		}
	}

	return sums
}

func Reduce[A, B any](collection []A, f func(B, A) B, initialValue B) B {
	result := initialValue

	for _, c := range collection {
		result = f(result, c)
	}

	return result
}
