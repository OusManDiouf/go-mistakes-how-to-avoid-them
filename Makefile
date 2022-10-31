# dont support module yet so dont forget to make go install github.com/kisielk/godepgraph
# and refer to the local GOBIN to look for the executable
depgraph:
	 godepgraph github.com/OusManDiouf/go-mistakes-how-avoid-them | dot -Tpng -o depgraph.png

test:
	go test -v .

itest:
	INTEGRATION=true go test --tags=integration -v .

itest-using-env:
	INTEGRATION=true go test -v .

short-test:
	go test -short -v .
