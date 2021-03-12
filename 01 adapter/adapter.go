package adapter

type Target interface {
	Request() string
}

type Adaptee interface {
	SpecificRequest() string
}

// do not forget to write the factory method
func NewAdaptee() Adaptee {
	return &adapteeImpl{}
}

// adapteeImpl is the implemented object that is adapted
type adapteeImpl struct{}

func (*adapteeImpl) SpecificRequest() string {
	return "adaptee method"
}

func NewAdapter(adaptee Adaptee) Target {
	return &adapter{
		Adaptee: adaptee,
	}
}

type adapter struct {
	Adaptee
}

func (a *adapter) Request() string {
	return a.SpecificRequest()
}
