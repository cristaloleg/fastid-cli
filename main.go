package main

import (
	"flag"

	"github.com/cristalhq/fastid"
)

func main() {
	epoch := flag.Int64("epoch", fastid.DefaultEpoch, "epoch of the generator")
	workerID := flag.Int("worker_id", 1, "id of the worker")
	flag.Parse()

	gen, err := fastid.NewGenerator(*epoch, *workerID)
	if err != nil {
		panic(err)
	}

	id := gen.Next()
	println(id)
}
