package impl_test

import (
	"anyway/collector"
	"anyway/collector/impl"
	"errors"
	"fmt"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
	"testing"
)

func TestImpl(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecsWithDefaultAndCustomReporters(t, "RepoImpl Suite", []Reporter{reporters.NewJUnitReporter("junit.xml")})
}

type MockData struct {
	Content string
}

type Notify interface {
	SavePoint(tilte string)
	Done(maxDuration int) // finish
}

var _ = Describe("Collector", func() {
	var n Notify
	defer n.Done(10000)
	n.SavePoint("a")
	fooHdl := buildMockOperator("_foo")
	barHdl := buildMockOperator("_bar")
	n.SavePoint("b")
	c := impl.NewCollector(10, collector.ProcessorChain{
		fooHdl,
		barHdl,
	}, 1)
	n.SavePoint("c")
	It("Write", func() {
		c.Write(&MockData{
			Content: "helloworld",
		})
	})
})

type MockDataRepo struct {
}

func (p *MockDataRepo) Upsert(in *MockData) error {
	return nil
}

type MockSink struct {
	Repo *MockDataRepo
}

func (p *MockSink) Hdl(in interface{}) (interface{}, error) {
	md, ok := in.(*MockData)
	if !ok {
		return nil, errors.New(`DataType Not Match`)
	}
	if err := p.Repo.Upsert(md); err != nil {
		return nil, err
	}
	return in, nil
}

type MockHdl struct {
	FakerAppend string
}

func (p *MockHdl) Hdl(in interface{}) (interface{}, error) {
	md, ok := in.(*MockData)
	if !ok {
		return nil, errors.New(`DataType Not Match`)
	}
	fmt.Println(`debug @f`, p.FakerAppend, " @v", md.Content)
	return &MockData{
		Content: md.Content + p.FakerAppend,
	}, nil
}

func buildMockOperator(fakerAppend string) collector.Processor {
	return &MockHdl{
		FakerAppend: fakerAppend,
	}
}
