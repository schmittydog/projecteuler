#!/usr/bin/python3

def parts(n):
    arr = [1] + [0]*n
    for i in range(1, n+1):
        for j in range(n-i+1):
            if arr[j]:
                arr[i+j] += arr[j]
    return arr[n]


for x in range(1, 100_000):
    print(x, parts(x))
