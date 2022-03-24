# Instructions
✓ Implementation for the WordCount function 
✓ Reading a text file into a string in the main function  
✓ Check that the unittest passes
✓ Log the runtime performance in the table below

✓ Update the WordCount function with the Map and Reduce tasks, using goroutines to parallelise and a channel to gather partial results
✓ Check that the unittest passes
✓ Find the optimal amount of gorountines before you encounter diminishing returns in performance improvements
✓ Log the runtime performance in the table below


| Variant      | Runtime (ms)       |
| ------------ | ------------:      |
| singleworker |         3233       | 
| mapreduce    | 2501 (13 threads)  |