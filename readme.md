# climenu
CLI-Menu is a tool for creating small and simple menus in cli.
Works on linux and windows 11

# Examples
Create menu and start it
```go
package main

import (
	"climenu"
	"fmt"
)

func main() {
	climenu.GetMenu("Test", climenu.GetMenuItems([]string{"Option 1", "Option 2", "Option 3"}, []func(){
		func() {
			fmt.Println("Option 1 ausgewählt.")
		},
		func() {
			fmt.Println("Option 2 ausgewählt.")
		},
		func() {
			fmt.Println("Option 3 ausgewählt.")
		},
	})).MenuInteraction(true)
}

```

This will look like this:
<pre><code class="sh"><span style="color:yellow">[X] Option 1</span>
[ ] Option 2
[ ] Option 3</code></pre>