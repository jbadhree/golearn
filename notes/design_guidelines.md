[Back](notes.md)

# Design Guidelines

https://github.com/ardanlabs/gotraining

A developer can have a mental model for only 10000 lines of code
Don't use Debugger we should have a mental model to debug and logs
Mechanical Sympathy - Think about the hardware

Lack of performance is gonna come from:
Latency - Network, IO
Memory Allocations
How efficiently you access data
Algorithm efficiency


All we are doing is Reading and Writing in memory
It must be accurate, consistent and efficient 

Data integrity problem- we should understand the data

Data Oriented Design is important 

For every 20 lines we add 1 bug. Less is more.
Handle errors right then and there 

Readability is about not hiding the cost of code 
Simplicity is about hiding complexity, new semantic level 

