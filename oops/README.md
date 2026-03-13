# oops concepts

---

**Aurthor:** A.K.Hasha Vardhan

---

## 🚀Overview:
   This folder is created to understand the concepts of OOPs(Object-Oriented-Programming) in GO language. where OOPs in go language is different compared to other languages like c++/java. as even though it **can do encapsulation,abstraction and Polymorphism** it does not uses classes or inheritance but **works on Structs , methods and interfaces** 

---

# Struct:
   normally in GO language classees are not supported so we will use something called **Struct** and we will define struct using some keywords like type and struct
   -> while giving thee name for the struct we need to remember that if the name was **capital** that means it is **exported** which means it can be accessed by other files in that project and if the name is **small** that means it is **unexported** which means it is private and accessed by that file only.

---
# Encapsulation:
   It is used to hide the sensitive code from the users, in GO encapsulation is implemented by capitalizing the feilds so that the functions will be make them public if the functions are public then are said to be on package level which means remaining functions can use them.

   #### Pointer Reciever:
   in GO language Pointer receiver is used to make the **permenant changes** making it changes in the original memory unlike value receiver which makes the changes the data in copy of the data making the original data remains the same

---

# Inheritance:
   In GO language there is no allow for inheritance directly so instead we implement using a concept called **Composition** through **Struct Embedding** and a struct embed another struct aalowing it to inherit the feilds and metholds 

---

# Abstraction:
   Abstraction is a process of hiding the complexity to the user and just showing the required interface to the user
   But unlike other languages it doesnt use abstarct classes but we achive it through **Interface** and this Interface shows what it can do but not how it will be done 
   ex:Think of a TV remote. You know that pressing the "Power" button turns the TV on. You don't need to know the electrical engineering inside the TV to use it. The button is the abstraction.

---

# polymorphism:
   it allows you to treat different objects as if they were the same type, as long as they share a specific behavior.
   unlike other languages which achives polymorphism using inheritance in go language we can get polymorphism using the **interfaces**\n
   ->in this it uses **duck typing**: if it implements all the methods in the interface then it woeks as polymorphism      
