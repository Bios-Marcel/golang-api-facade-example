# Golang API facade example

A simple `Greeter` interface that has two implementations, one for japanese and one for english.

The benefit of this pattern is, that the code that uses the interface can use any implementation, as
long as that implementation sticks to the contract.