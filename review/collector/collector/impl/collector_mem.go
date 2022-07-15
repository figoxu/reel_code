package impl

import "anyway/collector"

type Collector struct {
	ch          chan interface{}
	chain       collector.ProcessorChain
	workerCount int
}

func NewCollector(capacity int, chain collector.ProcessorChain, workerCount int) collector.Collector {
	c := &Collector{
		ch:    make(chan interface{}, capacity),
		chain: chain,
	}
	for i := 0; i < workerCount; i++ {
		go c.working()
	}
	return c
}

func (p *Collector) working() {
	defer func() {
		// do some catch and notification
	}()
	for v := range p.ch {
		var err error
		processData := v
		for _, f := range p.chain {
			processData, err = f.Hdl(processData)
			if err != nil {
				// hdl error
				break
			}
		}
	}
}

func (p *Collector) Write(data interface{}) error {
	p.ch <- data
	return nil
}
