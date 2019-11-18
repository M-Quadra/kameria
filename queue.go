package kameria

type queue struct {
	ary []interface{}
}

func (slf queue) Size() int {
	return len(slf.ary)
}

func (slf queue) Empty() bool {
	return slf.Size() <= 0
}

func (slf *queue) push(v interface{}) {
	slf.ary = append(slf.ary, v)
}

func (slf *queue) pop() interface{} {
	if slf.Empty() {
		return nil
	}

	v := slf.ary[0]
	slf.ary = slf.ary[1:]
	return v
}

//Queue4String Queue<string>
type Queue4String struct {
	queue
	UsingDefaultValue bool
	DefaultValue      string
}

//Push add a new value
func (slf *Queue4String) Push(v string) {
	slf.queue.push(v)
	// fmt.Println(slf.ary)
}

//Pop remove and return first value
func (slf *Queue4String) Pop() (string, bool) {
	v := slf.queue.pop()
	if v == nil {
		if slf.UsingDefaultValue {
			return slf.DefaultValue, true
		}

		return "", false
	}

	opt, ok := v.(string)
	return opt, ok
}
