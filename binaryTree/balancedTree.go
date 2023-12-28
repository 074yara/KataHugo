package binaryTree

import (
	"fmt"
	"math/rand"
)

// Node представляет узел бинарного дерева.
type Node struct {
	Key    int   // Ключ узла
	Height int   // Высота узла (для сбалансированности)
	Left   *Node // Левое поддерево
	Right  *Node // Правое поддерево
}

// AVLTree представляет сбалансированное бинарное дерево.
type AVLTree struct {
	Root *Node // Корень дерева
}

// NewNode создает новый узел с заданным ключом.
func NewNode(key int) *Node {
	return &Node{Key: key, Height: 1}
}

// Insert вставляет новый ключ в дерево и поддерживает его сбалансированность.
func (t *AVLTree) Insert(key int) {
	t.Root = insert(t.Root, key)
}

// ToMermaid возвращает строковое представление дерева для использования в Mermaid.
func (t *AVLTree) ToMermaid() string {
	return toMermaid(t.Root)
}

func toMermaid(node *Node) string {
	if node == nil {
		return ""
	}

	mermaidCode := fmt.Sprintf("graph TD;\n  %d", node.Key)
	if node.Left != nil {
		mermaidCode += fmt.Sprintf(" --> %d", node.Left.Key)
	}
	if node.Right != nil {
		mermaidCode += fmt.Sprintf(" --> %d", node.Right.Key)
	}

	mermaidCode += ";\n"
	mermaidCode += toMermaid(node.Left)
	mermaidCode += toMermaid(node.Right)

	return mermaidCode
}

// height возвращает высоту узла (0 для nil).
func height(node *Node) int {
	if node == nil {
		return 0
	}
	return node.Height
}

// max возвращает максимум из двух чисел.
func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// updateHeight обновляет высоту узла на основе высот его поддеревьев.
func updateHeight(node *Node) {
	node.Height = 1 + max(height(node.Left), height(node.Right))
}

// getBalance возвращает разницу высот между левым и правым поддеревьями узла.
func getBalance(node *Node) int {
	if node == nil {
		return 0
	}
	return height(node.Left) - height(node.Right)
}

// leftRotate выполняет левый поворот узла для балансировки.
func leftRotate(x *Node) *Node {
	y := x.Right
	x.Right = y.Left
	y.Left = x
	updateHeight(x)
	updateHeight(y)
	return y
}

// rightRotate выполняет правый поворот узла для балансировки.
func rightRotate(y *Node) *Node {
	x := y.Left
	y.Left = x.Right
	x.Right = y
	updateHeight(y)
	updateHeight(x)
	return x
}

// insert вставляет новый ключ в дерево и выполняет балансировку.
func insert(node *Node, key int) *Node {
	if node == nil {
		return NewNode(key)
	}

	if key < node.Key {
		node.Left = insert(node.Left, key)
	} else if key > node.Key {
		node.Right = insert(node.Right, key)
	} else {
		// Дубликаты не разрешены, можно обработать иначе, если нужно
		return node
	}

	updateHeight(node)

	balance := getBalance(node)

	// Левая сторона тяжелее
	if balance > 1 {
		if key < node.Left.Key {
			return rightRotate(node)
		} else if key > node.Left.Key {
			node.Left = leftRotate(node.Left)
			return rightRotate(node)
		}
	}

	// Правая сторона тяжелее
	if balance < -1 {
		if key > node.Right.Key {
			return leftRotate(node)
		} else if key < node.Right.Key {
			node.Right = rightRotate(node.Right)
			return leftRotate(node)
		}
	}

	return node
}

// GenerateTree генерирует сбалансированное дерево с указанным количеством элементов.
func GenerateTree(count int) *AVLTree {
	tree := AVLTree{}
	for i := 0; i < count; i++ {
		key := rand.Intn(1000)
		tree.Insert(key)
	}
	return &tree
}
