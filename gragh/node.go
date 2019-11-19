package gragh

type Node struct {
	SubmissionId  string
	FileName      string
	FilePath       string
	Timestamp      string
	LexicalIndex  int
	MetricalIndex int
	BeforeNode    *Node
	NextNode      *Node
}

type Nodes []*Node


func NewNode(submissionId string, fileName string, filePath string, timestamp string, lexicalIndex int, metricalIndex int) (node *Node) {
	node = new(Node)
	node.SubmissionId = submissionId
	node.FileName = fileName
	node.FilePath = filePath
	node.Timestamp = timestamp
	node.LexicalIndex = lexicalIndex
	node.MetricalIndex = metricalIndex
	return node
}