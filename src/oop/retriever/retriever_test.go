package retriever

import (
	"fmt"
	"github.com/snowdreams1006/learn-go/oop/retriever/mock"
	"github.com/snowdreams1006/learn-go/oop/retriever/real"
	"testing"
	"time"
)


func TestDownload(t *testing.T) {
	var r Retriever
	r = mock.Retriever{"This is mock imooc.com"}

	inspect(r)

	//t.Log(Download(r))

	r = &real.Retriever{
		UserAgent: "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.100 Safari/537.36",
		TimeOut:   time.Minute,
	}

	//t.Log(Download(r))

	inspect(r)

	// Type assertion
	if  mockRetriever ,ok := r.(mock.Retriever);ok{
		t.Log(mockRetriever.Contents)
	}else{
		t.Log("not mock retriever")
	}

	realRetriever := r.(*real.Retriever)

	t.Log(realRetriever.TimeOut)
}

func inspect(r Retriever) {
	fmt.Printf("%[1]T %[1]v\n", r)

	switch v := r.(type) {
	case mock.Retriever:
		fmt.Println("mock,contents=",v.Contents)
	case *real.Retriever:
		fmt.Println("real,UserAgent=",v.UserAgent)
	default:
		fmt.Println("unknown")
	}
}

