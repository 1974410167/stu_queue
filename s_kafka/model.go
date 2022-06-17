package s_kafka

var fakeDB string

func saveMessage(text string) {
	fakeDB = text
}

func GetMessage() string { return fakeDB }
