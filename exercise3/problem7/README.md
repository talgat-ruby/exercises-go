# Problem 7

Create methods for `BankAccount`, `FedexAccount` and `KazPostAccount`. From `BankAccount` and `KazPostAccount` can be
withdrawn money. From `FedexAccount` and `KazPostAccount` can be sent packages.

```go
normanOsborne := BankAccount{
    name:    "Norman Osborne",
    balance: 1_000_000,
}
peterParker := FedexAccount{
    name: "Peter Parker",
	packages: make(string[])
}
auntMay := KazPostAccount{
    name:    "Aunt May",
    balance: 200,
	packages: make(string[])
}

withdrawMoney(10, normanOsborne, auntMay) // can be withdrawn by multiple people
sendPackagesTo("Mary Jane", peterParker, auntMay) // can be sent by multiple people
```
