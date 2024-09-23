# Conditional Statements in Go

This repository contains sample code demonstrating the usage of conditionals, if statements with initial statement, and switch statements in Go.

## if Statements

The if statements in Go do not use parentheses around the condition:

```
go
if height > 4 {
    fmt.Println("You are tall enough!")
} else if height > 6 {
    fmt.Println("You are super tall!")
} else {
    fmt.Println("You are not tall enough!")
}
```

## Initial Statement in if Block

- An if statement in Go can have an "initial" statement. The variable(s) created in the initial statement are only defined within the scope of the if body. For example:

```go

if length := getLength(email); length < 1 {
fmt.Println("Email is invalid")
}
```

## Switch Statements

- Switch statements in Go are a way to compare a value against multiple options. They are similar to if-else statements but are more concise and readable when the number of options is more than 2:

```go

func getCreator(os string) string {
var creator string
switch os {
case "linux":
creator = "Linus Torvalds"
case "windows":
creator = "Bill Gates"
case "mac":
creator = "A Steve"
default:
creator = "Unknown"
}
return creator
}
```

- Notice that in Go, the break statement is not required at the end of a case to stop it from falling through to the next case. The break statement is implicit in Go.

### If you do want a case to fall through to the next case, you can use the fallthrough keyword.

```go
func getCreator(os string) string {
    var creator string
    switch os {
    case "linux":
        creator = "Linus Torvalds"
    case "windows":
        creator = "Bill Gates"

    // all three of these cases will set creator to "A Steve"
    case "Mac OS":
        fallthrough
    case "Mac OS X":
        fallthrough
    case "mac":
        creator = "A Steve"

    default:
        creator = "Unknown"
    }
    return creator
}
```

- The default case does what you'd expect: it's the case that runs if none of the other cases match. For more details and examples, please refer to the code files in this repository.
