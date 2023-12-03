package main

import (
	"fmt"
	"log/slog"

	"github.com/nova38/thesis/packages/tree"
)

// func printNode[T A]() func(node *tree.Node[T]) {
//	return func(node *tree.Node[T]) {
//		fmt.Println(node.Path, ":", node.Value)
//	}
//}

type A struct {
	name string
}

func main() {
	fmt.Println("Hello, world!")

	t := tree.NewTree[A]()

	c := &A{name: "struct"}
	// a := &A{name: "a"}
	b := new(A)
	b.name = "b"

	err := t.AddPath("a.b", b)
	if err != nil {
		return
	}
	// t.Print()

	err = t.AddPath("a.b.c", &A{name: "cme"})
	if err != nil {
		return
	}
	fmt.Printf("a.b.c: %s \n", t.GetPath("a.b.c").Value)

	// // fmt.Printf("a.b.c: %v\n", c)
	err = t.AddPath("a.b", b)
	if err != nil {
		return
	}
	err = t.AddPath("1.2.3", c)
	if err != nil {
		return
	}

	// t.AddPath("z.z", &A{name: "zap"})
	fmt.Printf("a.b: %s \n", t.GetPath("a.b").Value)
	b.name = "bnew"
	fmt.Printf("a.b: %s \n", t.GetPath("a.b").Value)

	err = t.AddPath("a.b", b)
	if err != nil {
		return
	}
	fmt.Printf("a.b: %s \n", t.GetPath("a.b").Value)

	t.Visit(tree.VisitorOptions{},
		func(node *tree.Node[A]) string {
			if node.Value == nil {
				node.Value = new(A)
			}
			node.Value.name = "new"
			slog.Info("Final", "node", slog.AnyValue(node), slog.Any("value", node.Value))
			return node.Path
		},
	)

	t.Print()
	slog.Info("Final", " %+v", t.GetPath("a.b").Value)
	// t.Print()
}
