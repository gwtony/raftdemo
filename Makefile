all: 
	CGO_CFLAGS="-I/home/vagrant/job/rocksdb/include" \
	CGO_LDFLAGS="-lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd" \
		go build

#CGO_LDFLAGS="-L/home/vagrant/job/rocksdb -L/lib64 -lrocksdb -lstdc++ -lm -lz -lbz2 -lsnappy -llz4 -lzstd" \
