package proxy

type Subject interface {
	Do() string
}

type RealSubject struct{}

func (RealSubject) Do() string {
	return "real"
}

type Proxy struct {
	real RealSubject
}

func (p Proxy) Do() string {
	var res string

	// before calling(invoking) the real object, be careful to check the cache authority
	res += "pre:"

	// invoke the true object
	res += p.real.Do()

	// operation after invoking the true object, such as caching the result and post-processing the result
	res += ":after"

	return res
}

