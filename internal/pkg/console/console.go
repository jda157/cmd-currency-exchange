package console

type Worker struct {
	amount float64
	src    string
	dst    string
}

func (cw *Worker) SetAmount(a float64) {
	cw.amount = a
}

func (cw *Worker) GetAmount() float64 {
	return cw.amount
}

func (cw *Worker) SetSrc(src string) {
	cw.src = src
}

func (cw *Worker) GetSrc() string {
	return cw.src
}

func (cw *Worker) SetDst(dst string) {
	cw.dst = dst
}

func (cw *Worker) GetDst() string {
	return cw.dst
}
