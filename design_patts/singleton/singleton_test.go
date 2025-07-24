package singleton

import (
	"fmt"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSingletonGorut(t *testing.T) {
	t.Parallel()
	// s1 := NewSingleton("PETR")
	// fmt.Println("===> ", s1.GetName())
	// s2 := NewSingleton("OLIA")
	// fmt.Println("===> ", s2.GetName())
	// s3 := NewSingleton("SANIYA")
	// fmt.Println("===> ", s3.GetName())

	wg := sync.WaitGroup{}
	// chT := make(chan struct{})

	for i := range 1000 {
		i := i
		// t.Run(strconv.Itoa(i), func(t *testing.T) {
		// 	t.Parallel()
		// 	s := NewSingleton(fmt.Sprintf("x%d", i))
		// 	fmt.Println("===> ", s.GetName())
		// })

		wg.Add(1)

		go func() {

			defer wg.Done()
			s := NewSingleton(fmt.Sprintf("x%d", i))
			fmt.Println("===> ", s.GetName())
		}()
	}
	wg.Wait()
}

func TestSingletonParall(t *testing.T) {

	tests := [100]struct {
		name string
		id   int
	}{}

	for i := range 100 {
		tests[i].name = fmt.Sprintf("case: %d", i)
		tests[i].id = i
	}
	var lincPrivSingl *singleton

	for _, test := range tests {

		t.Run(test.name, func(t *testing.T) {
			t.Parallel()
			fmt.Println(test.id, " ", test.name)

			s := NewSingleton(fmt.Sprintf("%v: %d", test.name, test.id))

			if lincPrivSingl == nil {
				lincPrivSingl = s
			}

			require.Equal(t, lincPrivSingl, s)

			fmt.Println("===> ", s.GetName())
		})

	}

}
