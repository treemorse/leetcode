package golang

func fib(n int, cache []int) int {
    if cache[n] != -1 {
        return cache[n]
    }

    cache[n] = fib(n-1, cache) + fib(n-2, cache)
    return cache[n]
}

