# TwoSidePrime

## Description:

This API accept a number and return response as 'true' if the given number is a Two-sided Prime 
(https://prime-numbers.info/article/two-sided-primes), otherwise 'false'.

## How to test:
Once you run the program, you can either test the program in browser as `http://localhost:8082/twoSidePrime/<num>` 
or can make curl request as
`curl http://localhost:8082/twoSidePrime/<num>` i.e 
`curl http://localhost:8082/twoSidePrime/37` . It will return `true` or `false` based on it's Two-sided prime or not. 
### To Run test cases:
1. Go to directory `TwoSidePrime/src/TwoSidedPrime`.
2. Execute command `go test -v`.
