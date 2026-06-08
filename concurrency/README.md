# Concurrency

---

**Aurthor:** A.K.Hasha Vardhan

---

## 🚀Overview:
  This is created to understand the concepts of Concurrency in GO language.In Go Concurrency means structuring the program so it can handle multiple tasks at once.It has a slogan saying **"Do not Communicate by sharing memory;instead,share memory by communicating"** So in this we will learn concepts like:  
  1. Goroutines: this shows how to start concurrent work using **go keyword**.
  2. Channels  : this shows how goroutines communicate safely
  3. sync.WaitGroup : this tells how to wait for many goroutines to finish. 
  4. select  :this shows how handle multiple channel operations.
  5. sync.Mutex and sync.RWMutex  : this shows how to protect the shared data
  6. Race detector : this shows how to find the concurrency bugs early

---

## Goroutine:  
  Goroutines are go's lightweight way to run functions concurrently. we can start one by keeping **"go"keyword** before the function call and Go runtime handles scheduling it for execution.  
  ex: go functioncall()  
  ->In Go the main function itselfs runs as a goroutine too. and when main finishes the program exits even if other goroutines are still running.  

  ### use Cases:  
   1. Webserver:  
      imagine a 1000 users all click "Buy now" at exact moment  
      -> without goroutine the server should process one by one user.  
      -> with goroutine there is no need of waiting for every new one gets their own goroutine to handle  
  ### mistakes we may do:  
  1. starting a goroutine and forgetting to wait for it to finish.  
  2. Reading or writing shared variables without synchronization.
  3. using goroutines for work that is simple and sequential, where the extra complexity is unnecessary.  

---

## Channels:  
  Channels are Go's built-in way for goroutine to communicate and synchronize.  
  ->Think of them like a pipes: where one goroutine sends value and another recives them and channel helps to coordinate the exchange safely.  
  ->A channel carries values of a specific type so a **chan int** can only send and receive int values.  
  ->In Go lang **<-** is the both send and receive operator which is depending on direction.  
  ex: 
    ch:=make(chan int)  // creating channel  
    ch<-10   // sending  
    x:=<-ch  // receiving  
  
  ->these channels are especially used when one goroutine produces work and another consumes it or when multiple goroutines need to coordinate progress.  

  ### use cases:  
  1. when you need to pass data between goroutines.    
  2. when you need to synchronize through message passing.  
  3. when you are building producer-consumer flows,pipelines or background workers.  
  4. when you want one goroutine to wait for another without explicit mutexes.

---

## sync.waitGroup:  
  It is Go's standard way to wait for a group od goroutines to finish before your program continues.  
  -> It is especially useful when you start several background tasks and need to block untill all of them are done.  
  ### How it works:  
  -> A waitGroup works like a counter.You increase the counter when you start a goroutine, decrease it when that goroutine finishes, and call Wait() to pause until the counter reaches zero.  
  1. Add(n)  //to register work  
  2. Done()  //inside goroutine when it completes  
  3. wait()  //in main goroutine to block until everything is finished  

  ->without a waitGroup your main function may exit before other goroutines complete.  

  ### use cases:  
  1. launching multiple API calls and waiting for all of them to finish.  
  2. Running several workers that process jobs in parallel.  
  3. Starting background initialization tasks.  
  4. Waiting for cleanup tasks, file writes, or batch jobs to complete.   

  ### code:  
  var wg sync.waitGroup  
  for i:=0;i<5;i++{  
    wg.Add(1)  
    go func(){  
      defer wg.Done()
    }()  
  }  
  wg.Wait()  

---  

## Select:  
  In go **select** is the way you wait on multiple channel operations at the same time.  
  ->It is mainly used for writing concurrent code because it lets a goroutine respond to whichever channel becomes ready first  
  ex:  
  select{  
    case msg:=<-ch1:  
      fmt.println(msg)  
    case ch2 <- 10:  
      fmt.println("sent")
    default:  
      fmt.println("no channel ready")  
  }  
  -> we can think "select" as similar to a switch, but instead of checking value , it checks channel operations. Each case is a send or receive on a channel and Go runs the first case that is ready.  
  -> If one of the channel operations can proceed, select chooses it.  
  -> if more than one are ready Go picks one at random.  
  -> if none are ready then immediately the default block will run.  
   
   ### use cases:  
   1. A server waiting for work,shutdown or timeout signals.  
   2. A request that should fail if it takes tht very long.  
   3. long running goroutines that need to stop when a cancel signal arrives.  
  
---

## sync.Mutex:
  Sync.Mutex is Go's basic Mutual exclusion lock.It protects shared data that only one goroutine can access a critical section at a time which helps to prevent race conditions.  
  -> A mutex is a lock with two operations
  1. lock()  
  2. unlock()  
  ->when one goroutine locks it ,other goroutines trying to lock the same mutex must wait until its unlocked.  
  ->Go's mutex has a usefull zero value: an unlocked mutex works immediately so usually there is no need extra initiallization.  
  ->Shared memory becomes dangerous when multiple goroutines reaad and write the same data at same time. A mutex fixes that by allowing only one goroutine into a protected code section which prevents inconsistance updates and data races.  
  ex:  
  mu.lock()  
  defer mu.unlock()

  ### use cases:  
  1. multiple goroutines needss to access the same data  
  2. when you need fast ,diirect protection around shared memory.  
  3. the data structure is easy to guard than to redesign with channels.
