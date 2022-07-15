package collector

type Collector interface {
	Write(data interface{}) error
}

type Processor interface {

	//Write(data interface{}) error
	Hdl(data interface{}) (interface{}, error)
}

type ProcessorChain []Processor
