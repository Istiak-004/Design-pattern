# What is Strategy
The Strategy Design Pattern is a behavioral design pattern that enables selecting an algorithm’s behavior at runtime. It defines a family of algorithms, encapsulates each one, and makes them interchangeable. The strategy pattern lets the algorithm vary independently from clients that use it.


## Concepts Behind Strategy Pattern
1. Context: Uses a Strategy object. It maintains a reference to the Strategy object and interacts with it only via the Strategy interface.

2. Strategy Interface: Common interface for all supported algorithms.

3. Concrete Strategies: Implementations of the Strategy interface.


## Example of Implementation

    payment-strategy/
    │
    ├── strategy/
        |-- implementation
          ├── credit_card.go      // Concrete Strategy
           ├── paypal.go           // Concrete Strategy
           └── crypto.go           // Concrete Strategy
        ├── strategy.go         // Strategy interface

    │
    ├── context/
    │   └── payment_context.go  // Context using the strategy
    │
    └── main.go



## Benefits of Strategy Pattern
1. Open/Closed Principle: Easily add new strategies without modifying existing code.

2. Single Responsibility: Separate the algorithm from usage.

3. Runtime flexibility: Switch behavior dynamically.

## ❗ When to Use It
1. You have many related classes that differ only in behavior.

2. You want to avoid using large conditional blocks (if-else or switch) to select behaviors.

3. You want to encapsulate related algorithms.