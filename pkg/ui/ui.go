package ui

import (
	"fmt"

	log "github.com/Sirupsen/logrus"
	"github.com/gdamore/tcell"
	"github.com/rivo/tview"

	"github.com/kitkat-group/katbox/pkg/kontent"
	"github.com/kitkat-group/katbox/pkg/ksettings"
)

// This package provides all of the handling of the "Main" User interface, interfaces for alternative use-cases should be in their own package

//MainUI starts up the katbox User Interface
func MainUI(a *kontent.Articles, s *ksettings.User) error {
	// Check for a nil pointer
	if a == nil {
		return fmt.Errorf("No katbox articles were loaded")
	}

	// Debug output the size of the contents
	log.Debugf("%d Github repos, %d Posts, %d snippets, %d tools", a.GHRepos, a.Posts, a.Snippets, a.Tools)

	// Begin the UI Tree
	rootDir := "KatBox"
	root := tview.NewTreeNode(rootDir).
		SetColor(tcell.ColorRed)
	tree := tview.NewTreeView().
		SetRoot(root).
		SetCurrentNode(root)

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

	// Add all of the children to the tree structure
	root.AddChild(ghNode)
	root.AddChild(posts)
	root.AddChild(snippets)
	root.AddChild(tools)

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

	if err := tview.NewApplication().SetRoot(tree, true).Run(); err != nil {
		panic(err)
	}

	return nil
}
