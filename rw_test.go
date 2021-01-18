package mutex

import (
	"fmt"
	"math"
	"os"
	"sync"
	"testing"
	"text/tabwriter"
)

func TestMutex(t *testing.T) {
	tw := tabwriter.NewWriter(os.Stdout, 0, 1, 2, ' ', 0)
	defer tw.Flush()

	var m sync.RWMutex
	fmt.Fprintf(tw, "Readers\tRWMutex\tMutex\n")
	for i := 0; i < 20; i++ {
		count := int(math.Pow(2, float64(i)))

		fmt.Fprintf(
			tw,
			"%d\t%v\t%v\n",
			count,
			society(count, &m, m.RLocker()),
			society(count, &m, &m),
		)
	}
}
