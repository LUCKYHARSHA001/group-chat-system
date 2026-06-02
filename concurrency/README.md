# oops concepts

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
