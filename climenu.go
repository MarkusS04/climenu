package menu

import (
	"fmt"

	"os"

	"github.com/eiannone/keyboard"
	"github.com/fatih/color"
)

// menu holds data for the climenu
type menu struct {
	headline string
	position int
	items    []menuItem
}

// menuItem holds the text of command and the specified command
type menuItem struct {
	descr   string
	command func()
}

/*
GetMenuItems creates the items for the cli-menu

	description and command must have same length
*/
func GetMenuItems(description []string, command []func()) []menuItem {

	if len(description) != len(command) {
		return nil
	}
	var items []menuItem
	for i := range description {
		items = append(items, menuItem{descr: description[i], command: command[i]})
	}
	return items
}

/*
GetMenu creates an object of type menu

	descr and cmd must have same length
	will create the menuItems based on these
*/
func GetMenu(headline string, items []menuItem) *menu {
	return &menu{headline: headline, position: 0, items: items}
}

/*
MenuInteraction starts the menu in terminal

	use Arrow-Key Up or Down for navigation
	enter to exexute specified function

	shouldExit defines wheter ctrl+c should stop complete process or only the climenu
*/
func (m *menu) MenuInteraction(shouldExit bool) *menu {

itemSelceted:
	for {
		m.printMenu()

		// read key input
		_, key, err := keyboard.GetSingleKey()
		if err != nil {
			fmt.Println(err)
		}
		switch key {
		case keyboard.KeyCtrlC:
			if shouldExit {
				os.Exit(0)
			}
			return m

		case keyboard.KeyArrowDown:
			m.position++
			// set position to 0, if position bigger then len of items
			if m.position == len(m.items) {
				m.position = 0
			}

		case keyboard.KeyArrowUp:
			m.position--
			// set position to len(items)-1, if position smaller then 0
			if m.position < 0 {
				m.position = len(m.items) - 1
			}

		// on enter execute command and leave menu-loop
		case keyboard.KeyEnter:
			m.items[m.position].command()
			break itemSelceted
		}

		// Move Pointer to start
		fmt.Printf("\033[%dA", len(m.items))
	}

	return m
}

// printMenu wrtiwa the menu items to console
func (m *menu) printMenu() {
	for i, item := range m.items {
		// set color of terminal to default
		color.Unset()

		selector := "[ ]"
		newLine := ""
		if i == m.position {
			color.Set(color.FgYellow)
			selector = "[X]"
		}
		if i < len(m.items) {
			newLine = "\n"
		}

		fmt.Printf("%s %s%s", selector, item.descr, newLine)
	}
}
