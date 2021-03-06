## organ - Go Library

<p align="center">
  <img src="https://user-images.githubusercontent.com/55533920/130828283-fb8242fa-545e-4ae3-ad56-4ccc4ffe68a8.png" />
</p>

a lib for organizing and managing files in go.

#### striving for a simple to use approach.

Begin by creating an instance of root by either making such root the directory of your source or by assigning a certain path.

``` go
package main

import(
	"organLib/paths"
)

r := paths.NewWithPath("c:/path/work")
fmt.Printf("root size : %d", r.Size)

```

This basic lib with utility functionalities is divided in *4* main fields, a decision made to simplify it's usage.

They are

- paths
- files
- generators
- organizers

**paths** possesses as it's main type, root, which provides the name, size and list of files in the given root. You have the following functions to fiddle with this type.

```go
func New() *RootDir
```

New(), enables you to create a root in the main's source directory.

````go
func NewWithPath(path string) *RootDir
````

NewWithPath(), enables you to create a root in the specified path.



### ***Disclaimer,***

###### This project starts as a challenge in order to improve my knowledge in the beautiful language of go. It's by no means a serious or dense in usability library, it's more of an experiment and exploration piece.