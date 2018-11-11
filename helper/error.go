package helper

type Error struct {
	Message string
}

func (er Error) Error() string {
	return er.Message
}
