package ui

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/gdamore/tcell"
	"github.com/micmonay/keybd_event"
	"github.com/rivo/tview"

	"github.com/kitkat-group/katbox/pkg/kontent"
	"github.com/kitkat-group/katbox/pkg/ksettings"
)

// This package provides all of the handling of the "Main" User interface, interfaces for alternative use-cases should be in their own package

// All articles is used to hold all articles to go back to searching everything
var allArticles *kontent.Articles

// uiBugFix - This is used to fix the issue with tcell dropping a keystroke between new tcell.screens being created
// TODO (thebsdbox) remove this when tcell issue #194 is fixed
func uiBugFix() {
	kb, err := keybd_event.NewKeyBonding()
	if err != nil {
		return
	}
	//set keys
	kb.SetKeys(keybd_event.VK_SPACE)

	//launch
	kb.Launching()
}

//MainUI starts up the katbox User Interface
func MainUI(a *kontent.Articles, s *ksettings.User) error {
	// Check for a nil pointer
	if a == nil {
		return fmt.Errorf("No katbox articles were loaded")
	}

	if allArticles == nil {
		allArticles = a
	}

	// Debug output the size of the contents
	log.Debugf("%d Github repos, %d Posts, %d snippets, %d tools", a.GHRepos, a.Posts, a.Snippets, a.Tools)

	root := buildTree(a)

	tree := tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root)
	application := tview.NewApplication()

	// If a directory was selected, open it.
	tree.SetSelectedFunc(func(node *tview.TreeNode) {
		reference := node.GetReference()
		if reference == nil {
			return // Selecting the root node does nothing.
		}
		children := node.GetChildren()
		// If it has children then flip the expanded state, if it's the final child we will action it
		if len(children) != 0 {
			node.SetExpanded(!node.IsExpanded())
		} else {
			// TODO - Open the action menu on the specific article
		}
	})

	// Add additional input capturing
	tree.SetInputCapture(func(event *tcell.EventKey) *tcell.EventKey {
		switch event.Key() {
		case tcell.KeyCtrlF:
			// Stop the existing UI

			var subset *kontent.Articles
			application.Suspend(func() { subset = SearchUI(allArticles) })
			uiBugFix()
			// Get new tree
			newRoot := buildTree(subset)
			root.ClearChildren()
			root.SetChildren(newRoot.GetChildren())

		default:
			return event
		}
		return nil
	})

	if err := application.SetRoot(tree, true).Run(); err != nil {
		panic(err)
	}

	return nil
}

//SearchUI -
func SearchUI(a *kontent.Articles) *kontent.Articles {

	uiBugFix()

	title := "Search"
	label := "Search string (RegEx)"
	for {
		app := tview.NewApplication()

		form := tview.NewForm().
			AddInputField(label, "", 30, nil, nil).
			AddButton("Search", func() { app.Stop() })

		form.SetBorder(true).SetTitle(title).SetTitleAlign(tview.AlignLeft)

		if err := app.SetRoot(form, true).SetFocus(form).Run(); err != nil {
			panic(err)
		}

		searchString := form.GetFormItemByLabel(label).(*tview.InputField).GetText()

		newArticles, err := a.SearchKeywords(searchString)
		if err == nil {
			return newArticles
		}
		title = err.Error()
	}
}

// This function will take the full article set and build a tree from any search parameters
func buildTree(a *kontent.Articles) *tview.TreeNode {
	// Begin the UI Tree
	rootDir := "KatBox"
	root := tview.NewTreeNode(rootDir).
		SetColor(tcell.ColorRed)

	// Add Github articles to the tree
	ghNode := tview.NewTreeNode("GitHub Repositories").SetReference("GitHub").SetSelectable(true)
	ghNode.SetColor(tcell.ColorGreen)

	for x := range a.GHRepos {
		childNode := tview.NewTreeNode(a.GHRepos[x].Name).SetReference(x).SetSelectable(true)
		ghNode.AddChild(childNode)
	}

	// Add Post articles to the tree
	posts := tview.NewTreeNode("Posts").SetReference("Posts").SetSelectable(true)
	posts.SetColor(tcell.ColorGreen)

	for x := range a.Posts {
		childNode := tview.NewTreeNode(a.Posts[x].Name).SetReference(x).SetSelectable(true)
		posts.AddChild(childNode)
	}

	// Add Snippets  to the tree
	snippets := tview.NewTreeNode("Snippets").SetReference("Snippets").SetSelectable(true)
	snippets.SetColor(tcell.ColorGreen)

	for x := range a.Snippets {
		childNode := tview.NewTreeNode(a.Snippets[x].Name).SetReference(x).SetSelectable(true)
		snippets.AddChild(childNode)
	}

	// Add Snippets  to the tree
	tools := tview.NewTreeNode("Tools").SetReference("Tools").SetSelectable(true)
	tools.SetColor(tcell.ColorGreen)

	for x := range a.Tools {
		childNode := tview.NewTreeNode(a.Tools[x].Name).SetReference(x).SetSelectable(true)
		tools.AddChild(childNode)
	}

	// Add Team todo to the tree
	teamtodo := tview.NewTreeNode("TeamTodo").SetReference("TeamTodo").SetSelectable(true)
	teamtodo.SetColor(tcell.ColorBlue)

	for x := range a.TeamTodo {
		childNode := tview.NewTreeNode(a.TeamTodo[x].Name).SetReference(x).SetSelectable(true)
		teamtodo.AddChild(childNode)
	}

	// Add User todo to the tree
	usertodo := tview.NewTreeNode("UserTodo").SetReference("UserTodo").SetSelectable(true)
	usertodo.SetColor(tcell.ColorOrange)

	for x := range a.UserTodo {
		childNode := tview.NewTreeNode(a.UserTodo[x].Name).SetReference(x).SetSelectable(true)
		usertodo.AddChild(childNode)
	}

	// Add all of the children to the tree structure
	root.AddChild(usertodo)
	root.AddChild(teamtodo)

	root.AddChild(ghNode)
	root.AddChild(posts)
	root.AddChild(snippets)
	root.AddChild(tools)

	return root
}
