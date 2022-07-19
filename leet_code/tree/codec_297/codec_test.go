package codec_297

import (
	"strconv"
	"strings"
	"testing"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

type Codec struct {
}

func Constructor() Codec {
	return Codec{}
}

var seq []string
var index int

// Serializes a tree to a single string.
func (codec *Codec) serialize(root *TreeNode) string {
	seq = []string{}
	dfs(root)
	return strings.Join(seq, ",")
}

// Deserializes your encoded data to tree.
func (codec *Codec) deserialize(data string) *TreeNode {
	if data == "" {
		return nil
	}
	index = 0
	seq = strings.Split(data, ",")

	return buildTree()
}

func dfs(root *TreeNode) {
	if root == nil {
		seq = append(seq, "null")
		return
	}

	seq = append(seq, strconv.Itoa(root.Val))
	dfs(root.Left)
	dfs(root.Right)
}

func buildTree() *TreeNode {
	if index >= len(seq) || seq[index] == "null" {
		index++
		return nil
	}

	value, _ := strconv.Atoi(seq[index])
	root := &TreeNode{Val: value}

	index++
	root.Left = buildTree()
	root.Right = buildTree()

	return root
}

/**
 * Your Codec object will be instantiated and called as such:
 * ser := Constructor();
 * deser := Constructor();
 * data := ser.serialize(root);
 * ans := deser.deserialize(data);
 */

func TestCodec(t *testing.T) {
	data := []string{"1", "2", "3", "null", "null", "4", "5"}
	codec := Constructor()

	str := strings.Join(data, ",")
	root := codec.deserialize(str)

	t.Log(root)

	ans := codec.serialize(root)
	t.Log(ans)
}
