if [ -f exec.out ]; then
  rm exec.out
fi
go build -o exec.out github.com/SamsadSajid/go-microservice-poc
./exec.out