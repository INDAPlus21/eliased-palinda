# bug01.go 
Explanation: the channel was unbuffered which resulted in deadlock (because there was nothing to read it immediately), and I made it buffered (can contain max 1 message before something reads it, instead of 0)   

# bug02.go  
When in the loop it sent the numbers to the channel Print() couldn't read 11 before the function exited (because it's unbuffered everything is read immediately). I fixed it by adding a waitgroup that waits for the Print() function to complete  

# many2many.go 
## What happens if you switch the order of the statements wgp.Wait() and close(ch) in the end of the main function?
You get a "panic: send on closed channel", because Produce() still tries to send stuff on it, but it's closed immediately after exiting the for loops (which is quick because goroutines are called). 

## What happens if you move the close(ch) from the main function and instead close the channel in the end of the function Produce?
Nothing is consumed/produced after one call to the Produce() function has finished — we get a "panic: send on closed channel". 

## What happens if you remove the statement close(ch) completely?
Nothing in particular happens, because the program is already about finished 
 
## What happens if you increase the number of consumers from 2 to 4?
It consumes the data faster, because RandomSleep() is called within each thread, so we have fewer blocking read operations on the channel because there are more consumers.  

## Can you be sure that all strings are printed before the program stops? 
No, becausthe all the consumers could be waiting. Say that (by chance) all the producers send their final string to the channel at the same time, but all the consumers are sleeping. wg.Done() is then called by all the producers, and the code continues to close() and exits before the consumers have a chance to read it. I also confirmed this by printing the number of packets consumed vs number produced, and the number consumed was 1 lower sometimes. 

# Notes 
(close = not able to receive or send any more values, range pulls values "automatically" until channel is closed, and select acts kinda like a switch)