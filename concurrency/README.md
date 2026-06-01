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
  