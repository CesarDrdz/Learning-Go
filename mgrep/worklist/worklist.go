package worklist

type Entery struct {
	Path string
}

type Worklist struct {
	jobs chan Entery
}

func (w *Worklist) Add(work Entery) {
	w.jobs <- work
}

func (w *Worklist) Next() Entery {
	j := <-w.jobs
	return j
}

func New(bufSize int) Worklist {
	return Worklist{make(chan Entery, bufSize)}
}

func NewJob(path string) Entery {
	return Entery{path}
}

func (w *Worklist) Finalize(numWorkers int) {
	for i := 0; i < numWorkers; i++ {
		w.Add(Entery{""})
	}
}
