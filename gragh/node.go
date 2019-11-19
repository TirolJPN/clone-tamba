package gragh

type Node struct {
	submission_id string
	file_name string
	file_path string
	timestamp string
	lexical_index int
	metrical_index int
	before_node *Node
	next_node *Node
}

func newNode(submissionId string, fileName string, filePath string, timestamp string, lexicalIndex int, metricalIndex int) *Node {
	node := new(Node)
	node.submission_id = submissionId
	node.file_name = fileName
	node.file_path = filePath
	node.timestamp = timestamp
	node.lexical_index = lexicalIndex
	node.metrical_index = metricalIndex
	return node
}