# The Repository Pattern

The Repository Pattern is a design pattern that mediates between the domain (business logic) and data mapping layers of your application. It acts like an in-memory collection of domain objects while abstracting away the details of data persistence.

## Key Concept

    Abstraction: Hides the details of how data is persisted

    Separation of Concerns: Business logic doesn't need to know about database operations

    Testability: Makes it easier to unit test business logic by mocking the repository

    Flexibility: Allows you to change data sources without affecting business logic


## Structure

    Business Layer
        ↓
    Repository Interface (Abstract)
        ↓
    Repository Implementation (Concrete)
        ↓
    Data Source (Database, API, etc.)


# In-Depth Analysis of Repository Pattern Layers

Lets dive into it. 

## Domain Layer

### Purpose 

1. Defines the core business entities
2. Contains pure business logic structures
3. Free from infrastructure concerns

```go
type User struct {
    ID        int
    FirstName string
    LastName  string
    Email     string
}
```

### Why it's important:

    Single Source of Truth: All other layers reference this definition

    Business Focus: Contains only what's relevant to business operations

    Decoupling: Changes to storage don't require changes to domain model


## Repository Interface Layer

### purpose

1. Defines the contract for data access operations
2. Abstracts persistence details from business logic
3. Enables dependency inversion

```go
type UserRepository interface {
    GetAll() ([]domain.User, error)
    GetByID(id int) (*domain.User, error)
    Create(user *domain.User) error
    Update(user *domain.User) error
    Delete(id int) error
}
```

### Why it's important:

    Abstraction: Business logic depends on abstraction, not implementation
    Testability: Enables easy mocking for unit tests
    Flexibility: Allows switching data sources without changing business logic

## Repository Implementation Layer


### Purpose

1. Concrete implementation of repository interface
2. Handles all data persistence details
3. Translates between domain objects and storage format

```go
type userRepositoryMemory struct {
    users map[int]domain.User
    mu    sync.RWMutex
}

func NewUserRepositoryMemory() UserRepository {
    return &userRepositoryMemory{
        users: make(map[int]domain.User),
    }
}
```

### Why it's important:

    Encapsulation: Hides all persistence details

    Thread Safety: The mutex ensures safe concurrent access

    Lifecycle Management: Constructor provides controlled initialization

## Service Layer

### Purpose:

1. Contains business logic
2. Coordinates domain objects and repositories
3. Application-specific operations

```go
type UserService struct {
    repo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
    return &UserService{repo: repo}
}

func (s *UserService) RegisterUser(firstName, lastName, email string) (*domain.User, error) {
    user := &domain.User{
        ID:        len(firstName) + len(lastName) + len(email),
        FirstName: firstName,
        LastName:  lastName,
        Email:     email,
    }
    
    if err := s.repo.Create(user); err != nil {
        return nil, err
    }
    
    return user, nil
}
```

### Why it's important:

    Business Logic Home: Where domain rules are enforced
    Transaction Boundary: Often manages transaction scope
    Use Case Implementation: Represents specific application workflows

## Presentation/API Layer 

### Purpose

1. Application entry point
2. Dependency wiring
3. Top-level composition
4. Presentation/API handling

```go 
func main() {
    // Initialize repository
    userRepo := repository.NewUserRepositoryMemory()
    
    // Initialize service with repository
    userService := service.NewUserService(userRepo)
    
    // Use the service
    user, err := userService.RegisterUser("John", "Doe", "john@example.com")
    if err != nil {
        panic(err)
    }
    
    fmt.Printf("Registered user: %+v\n", user)
}
```

## Why it's important:

    Composition Root: Where dependencies are wired together

    Configuration: Sets up the application environment

    Lifecycle Management: Controls startup/shutdown