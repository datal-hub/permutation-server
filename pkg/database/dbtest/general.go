package dbtest

func (db *DbTest) IsEmpty() bool {
	return true
}

func (db *DbTest) Clear() {
}

func (db *DbTest) Init(force bool) error {
	return nil
}
