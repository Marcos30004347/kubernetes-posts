deps:
	go get -d . 

compile: deps
	CGO_ENABLED=0 GOOS=linux go build .

codegen:
	./hack/scripts/codegen.sh

run:
	go run . --etcd-servers localhost:2379 --authentication-kubeconfig ~/.kube/config --authorization-kubeconfig ~/.kube/config --kubeconfig ~/.kube/config
