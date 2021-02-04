package calculation

import (
	"go/token"
	"go/types"
	"strconv"
	"strings"
	"sync"
)

func Calculate(value string) string {
	vArr := strings.Split(value, " ")
	var wg sync.WaitGroup
	for i := 0; i < len(vArr); i++ {
		toComputeInt, err := strconv.Atoi(strings.TrimSpace(strings.TrimRight(vArr[i], "!")))
		if err == nil {
			wg.Add(1)
			go calculateFactorial(toComputeInt, vArr, i, &wg)
		}
	}
	wg.Wait()

	return evaluateExp(strings.Join(vArr, ""))
}

func evaluateExp(s string) string {
	fs := token.NewFileSet()
	tv, err := types.Eval(fs, nil, token.NoPos, s)
	if err != nil {
		return s
	}
	return tv.Value.String()
}

func calculateFactorial(n int, vArr []string, index int, wg *sync.WaitGroup) {
	defer wg.Done()
	var res int64 = 1
	for i := 1; i<=n; i++ {
		res *= int64(i)
	}
	vArr[index] = strconv.FormatInt(res, 10)
}
