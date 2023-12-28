package workers

import (
	"bytes"
	"hugoproxy/binaryTree"
	"math/rand"
	"os"
	"time"
)

var (
	binaryTreePath = `/app/static/tasks/binary.md`
	head           = `{{< mermaid [class="text-center"]>}}\n`
	tail           = `{{< /mermaid >}}`
)

func BinaryTreeWorker() {
	buff := &bytes.Buffer{}
	for {
		tree := binaryTree.GenerateTree(5)
		data := makeMermaidToFileData(buff, tree)
		err := os.WriteFile(binaryTreePath, data, 0644)
		checkError(err)
		time.Sleep(time.Second * 5)
		for i := 0; i < 95; i++ {
			tree.Insert(rand.Intn(1000))
			data = makeMermaidToFileData(buff, tree)
			err = os.WriteFile(binaryTreePath, data, 0644)
			checkError(err)
			time.Sleep(time.Second * 5)
		}
	}
}

func makeMermaidToFileData(buff *bytes.Buffer, tree *binaryTree.AVLTree) []byte {
	buff.Reset()
	_, err := buff.Write(stringToByte(head))
	checkError(err)
	_, err = buff.Write(stringToByte(tree.ToMermaid()))
	checkError(err)
	_, err = buff.Write(stringToByte(tail))
	checkError(err)
	return buff.Bytes()
}

func stringToByte(str string) []byte {
	return []byte(str)
}
