package topic

type TopicSession interface {
}

type Unit struct {
	register chan TopicSession
	Name     string
}
