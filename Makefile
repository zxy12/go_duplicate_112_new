
ZDEBUG_A = '/Users/zhouxinyu/www/localhost/pkg/darwin_amd64/github.com/zxy12/go_duplicate_112_new/src/zdebug.a'

all:
	rm -f $(ZDEBUG_A)
	export GOROOT_BOOTSTRAP=/Users/zhouxinyu/www/localhost/src/github.com/golang/go1.4 && cd ./src && ./make.bash

clean:

.PHONY:
	all test
