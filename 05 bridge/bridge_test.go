package bridge

func ExampleCommonSMS() {
	m := NewCommonMessage(ViaSMS())
	m.SendMessage("have a drink?", "bob")
}

func ExampleCommonEmail() {
	m := NewCommonMessage(ViaEmail())
	m.SendMessage("have a drink?", "bob")
}

func ExampleUrgencySMS() {
	m := NewUrgencyMessage(ViaSMS())
	m.SendMessage("have a drink?", "bob")
}

func ExampleUrgencyEmail() {
	m := NewUrgencyMessage(ViaEmail())
	m.SendMessage("have a drink?", "bob")
}

