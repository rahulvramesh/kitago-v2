## Building Package 

1, Install Build Tool

```sh
    go get github.com/laher/goxc
```

2, Install Dependency 

```sh
    go mod download
```

3, Build 

```sh
    make build
```

4, Run

if the host machine is mac
```sh
    ./build/0.0.2/darwin_amd64/kitago sum 1 2
```

please not as per the configuration goxc will generate the binary for all major os, since expect little time to compile and build