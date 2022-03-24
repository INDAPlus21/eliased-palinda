# How many CPUs does you program use? How much faster is your parallel version?
My current go version is 1.14, so I'm NOT using as many OS threads as possible. So I need to tell the program how many threads it should use with GOMAXPROCS, but the other people have it for free (but they still need to do the main bit of the task). It took 11.04 seconds to complete 7 images, and 35 seconds all the 8 images without the use of any optimizations. 

One possible thing to do is to devide the work units into the pngs (the most naive method). Let's try. 

Ok. the thing where everything is blocking, but it takes a bit less time makes some sense. 

Yes. And when not capturing the loop variable fn I get (while getting the same images) much faster execution times (7 seconds before second to last is completed, 24 seconds last). And if I try to use all the threads... I get the same identical result. And if I try to use goroutines within the julia function as well... It should be slightly different execution times I assume (???), why do everyone take roughly the same time... seems suspicious? yes... and it's only picture 5 and 7 that are created... No, this was because I also captured n (picture number). Ok. As I suspected. The second goroutine WITHIN julia didn't work (it does something suspicious). But I'm capturing I again!!!.. normal function (not anonymous incline functions) are probably easier not to make errors with... These times make a LOT more sense. Yes! and this only took 5.67 seconds for the second to last, and 5.79 seconds for the last 

Even MORE goroutines??? (within iterate!!!)... yeah yeah... i'm bored with this now. let's continue to the next task 