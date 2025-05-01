package log

import (
	"fmt"
	"os"
	"sync"

	"github.com/sirupsen/logrus"
)

type ToFile struct {
	fn    string
	title string
	ch    chan string
	wg    sync.WaitGroup
}

func NewToFile(fn string, title string) *ToFile {
	t := &ToFile{
		fn:    fn,
		ch:    make(chan string, 100),
		title: title,
		wg:    sync.WaitGroup{},
	}

	return t
}

func (t *ToFile) Write(msg string) {
	t.ch <- msg
}

func (t *ToFile) Init() (err error) {
	fd, err := os.OpenFile(t.fn, os.O_CREATE|os.O_TRUNC|os.O_WRONLY, 0666)
	if err != nil {
		return
	}

	t.ch = make(chan string, 1)
	t.wg.Add(1)
	_, err = fd.WriteString(fmt.Sprintf("Start %s\n", t.title))
	if err != nil {
		return err
	}

	logrus.Infof("Start logging %s into %s", t.title, t.fn)
	go func(ch chan string, fd *os.File) {
		defer fd.Close()
		defer t.wg.Done()
		for c := range ch {
			p := fmt.Sprintf("[%s] %s\n", t.title, c)
			fd.WriteString(p)
		}
	}(t.ch, fd)

	return
}

func (t *ToFile) Close() {
	close(t.ch)
	t.wg.Wait()
}
