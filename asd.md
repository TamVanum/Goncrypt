## 🧱 1. **Constructor Conventions (`NewX`)**

We touched on this already, but here's the full pattern:

```go
type MyService struct {
    repo MyRepo
}

func NewMyService(repo MyRepo) *MyService {
    return &MyService{repo: repo}
}
```

### Why `New`?
- `NewX` is an idiomatic name for factory/constructor functions in Go.
- If it's the only public function in the package, you can even name it just `New()`:

```go
// In package `myservice`
func New(repo MyRepo) *MyService
```

💡 `New` is not a keyword, it's just a **convention**.

---

## 📁 2. **Package Naming**

- **All lowercase**, no underscores or camelCase.
- Prefer **short, meaningful names**: `auth`, `repo`, `service`, `http`, `db`, `mail`, `user`.

```go
// Not this
package UserService

// This
package userservice
```

But better:

```go
package service
```

Because consumers will use: `service.NewUserService()`.

---

## 🧑‍💻 3. **Interfaces: Behavior > Identity**

Go favors **small, behavior-based interfaces**.  
Conventionally:

- Interface names use `-er` suffix: `Reader`, `Writer`, `Fetcher`, `Storer`, `Authenticator`.
- Never prefix with `I` like in Java or C#: ❌ `IUserService`.

Example:

```go
type Storer interface {
    Store(data []byte) error
}
```

Good practice: Define interfaces **where they're used**, not where they're implemented.

---

## ⚠️ 4. **Error Handling**

Go **embraces explicit error handling**, but you can follow these conventions to keep things clean:

```go
data, err := repo.Get(id)
if err != nil {
    return nil, fmt.Errorf("getting user: %w", err)
}
```

### Custom error variables:

```go
var ErrNotFound = errors.New("not found")
```

### Use `%w` to wrap errors:
```go
return fmt.Errorf("saving user: %w", err)
```

---

## 🧪 5. **Testing Convention**

- Test files end in `_test.go`.
- Use `TestXxx(t *testing.T)` for unit tests.

```go
func TestNewUserService(t *testing.T) {
    repo := NewMockRepo()
    svc := NewUserService(repo)
    if svc == nil {
        t.Fatal("expected service to be non-nil")
    }
}
```

- Place tests in **the same package** for black-box testing.
- Or use `package mypkg_test` to test like an external consumer.

---

## ⚙️ 6. **Receiver Naming Conventions**

Use **short, meaningful variable names** for method receivers, like:

```go
type UserService struct {}

func (s *UserService) CreateUser(...) { ... }
```

| Type | Receiver |
|------|----------|
| `UserService` | `s` |
| `UserRepository` | `r` |
| `Handler` | `h` |
| `Config` | `c` |
| `Validator` | `v` |

---

## ✍️ 7. **Struct Tags Convention**

Use backticks with `json`, `db`, `validate`, etc.:

```go
type User struct {
    ID   int    `json:"id" db:"id"`
    Name string `json:"name" validate:"required"`
}
```

---

## 📦 8. **Exported vs Unexported**

In Go:

- Uppercase identifiers are **exported** (public).
- Lowercase identifiers are **unexported** (private to package).

```go
type UserService struct {}     // exported
type userRepository struct {}  // internal

func NewUserService() *UserService // exported
func newUserRepo() *userRepository // internal
```

---

## 📐 9. **Project Layout Tips (Minimal)**

```
/cmd/api         → entrypoint (main.go)
/internal/user   → user-specific logic
/internal/repo   → db layer
/pkg/crypto      → shared packages
```

But you can go simpler:

```
/main.go
/user/
    handler.go
    service.go
    repo.go
```

---

## 🧹 10. **Other Subtle Conventions**

- Avoid unnecessary getters/setters.
    ```go
    // Instead of
    func (u *User) GetName() string

    // Just use
    func (u *User) Name() string
    ```

- If your type implements `String() string`, it's printed nicely:

    ```go
    func (u User) String() string {
        return fmt.Sprintf("User<%d:%s>", u.ID, u.Name)
    }
    ```

- Use `context.Context` as the **first argument** in methods that need it:

    ```go
    func (s *UserService) CreateUser(ctx context.Context, user User) error
    ```

---

## 👉 TL;DR Summary

| Concept            | Convention                           |
|--------------------|--------------------------------------|
| Constructor        | `NewX()`                             |
| Package names      | lowercase, no underscores            |
| Interfaces         | Small, behavior-based (`-er`)        |
| Struct tags        | Backticks with `json`, `db`, etc.    |
| Receivers          | Short and meaningful (`s`, `r`, `h`) |
| Exported symbols   | Capitalized                          |
| Error variables    | `var ErrX = errors.New("...")`       |
| Method context     | First arg: `ctx context.Context`     |

---

I can go deeper into any of these or show you **real-world examples** (e.g., mimicking how popular Go projects like Kubernetes, Gin, or Go kit are structured).

Where do you want to dive next — architecture, testing, Go generics, interfaces in practice, or idiomatic patterns like middleware or decorators?