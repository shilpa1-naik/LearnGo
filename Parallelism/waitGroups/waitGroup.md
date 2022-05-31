// From educative


A WaitGroup blocks a program and waits for a set of goroutines to finish before moving to the next steps of execution.


We can use WaitGroups through the following functions:

.Add(int): This function takes in an integer value which is essentially the number of goroutines which the waitgroup has to wait for. This function must be called before we execute a goroutine.
.Done(): This function is called within the goroutine to signal that the goroutine has successfully executed.
.Wait(): This function blocks the program until all the goroutines specified by Add() have invoked Done() from within.

```
var wg sync.WaitGroup

wg.Add(2) //called before running the goroutines

go func() {
    // Do work.
    wg.Done() //goroutine reached its completion; calling Done() to signal
}()
go func() {
    // Do work.
    wg.Done() //goroutine reached its completion; calling Done() to signal

}()

wg.Wait()
```