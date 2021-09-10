
compile: deps
	CGO_ENABLED=0 GOOS=linux go build .

deps:
	go get -d . 

codegen:
	./hack/scripts/codegen.sh

build:
	docker build -t marcos30004347/kubernetes-posts .

run:
	go run . --etcd-servers localhost:2379 --authentication-kubeconfig ~/.kube/config --authorization-kubeconfig ~/.kube/config --kubeconfig ~/.kube/config

deploy:
	kubectl apply -f ./artifacts/deploy/namespace.yaml
	kubectl apply -f ./artifacts/deploy/service-account.yaml
	kubectl apply -f ./artifacts/deploy/

undeploy:
	kubectl delete -f ./artifacts/deploy/

