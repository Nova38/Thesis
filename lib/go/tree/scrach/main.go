package main

import (
	"fmt"
	"log/slog"

	"github.com/nova38/thesis/lib/go/tree"
)

func printNode[T A]() func(node *tree.Node[T]) {
	return func(node *tree.Node[T]) {
		fmt.Println(node.Path, ":", node.Value)
	}
}

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

	t.AddPath("a.b", b)
	// t.Print()

	t.AddPath("a.b.c", &A{name: "cme"})
	fmt.Printf("a.b.c: %s \n", t.GetPath("a.b.c").Value)

	// // fmt.Printf("a.b.c: %v\n", c)
	t.AddPath("a.b", b)
	t.AddPath("1.2.3", c)

	// t.AddPath("z.z", &A{name: "zap"})
	fmt.Printf("a.b: %s \n", t.GetPath("a.b").Value)
	b.name = "bnew"
	fmt.Printf("a.b: %s \n", t.GetPath("a.b").Value)

	t.AddPath("a.b", b)
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
