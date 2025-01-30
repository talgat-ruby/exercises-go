package main
3
 
​
4
 
import (
5
 
  "fmt"
6
 
  "sync"
7
 
)
8
 
​
9
 
// bankAccount represents a simple bank account with a balance and mutex for concurrent safety
10
 
type bankAccount struct {
11
 
  balance int
12
 
  mu      sync.Mutex
13
 
}
14
 
​
15
 
// Deposit adds an amount to the account balance, ensuring concurrent safety
16
 
func (a *bankAccount) Deposit(amount int) {
17
 
  a.mu.Lock()
18
 
  defer a.mu.Unlock()
19
 
  a.balance += amount
20
 
}
21
 
​
22
 
// Withdraw subtracts an amount from the account balance if funds are available, ensuring concurrent safety
23
 
func (a *bankAccount) Withdraw(amount int) bool {
24
 
  a.mu.Lock()
25
 
  defer a.mu.Unlock()
26
 
  if a.balance >= amount {
27
 
    a.balance -= amount
28
 
    return true
29
 
  }
30
 
  return false // insufficient funds
31
 
}
32
 
​
33
 
// Balance safely returns the current balance of the account
34
 
func (a *bankAccount) Balance() int {
35
 
  a.mu.Lock()
36
 
  defer a.mu.Unlock()
37
 
  return a.balance
38
 
}
39
 
​
40
 
func main() {
41
 
  account := &bankAccount{}
42
 
  account.Deposit(1000)
43
 
  fmt.Println("Current Balance:", account.Balance()) // Should print 1000
44
 
​
45
 
  account.Withdraw(200)
46
 
  fmt.Println("Current Balance after withdrawal:", account.Balance()) // Should print 800
47
 

48
 
package main
49
 
​
50
 
import (
51
 
  "time"
52
 
)
53
 
​
54
 
var readDelay = 10 * time.Millisecond
55
 
​
56
 
type bankAccount struct {
57
 
  blnc int
58
 
}
59
 
​
60
 
func newAccount(blnc int) *bankAccount {
61
 
  return &bankAccount{blnc}
62
 
}
Footer
© 2024 GitHub, Inc.
Footer navigation
Terms
Privacy
Security
Status
Docs
Contact
Manage cookies
Do not share my personal information
HEAD
Resolve Conflicts · Pull Request #210 · talgat-ruby/exercises-go

Resolve Conflicts · Pull Request #210 · talgat-ruby/exercises-go

package main

type bankAccount struct {
	blnc int
}

func newAccount(blnc int) *bankAccount {
	return &bankAccount{blnc}
}
main
6786b69556adaf3ccffc518c7655a26812d55e64
