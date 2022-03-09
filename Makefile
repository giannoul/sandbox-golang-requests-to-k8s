KIND_CLUSTER	=$(or $(shell printenv KIND_CLUSTER_NAME), sandbox)
PWD		= $(shell pwd)

.PHONY: --apply-ingress-nginx
--apply-ingress-nginx: 
	kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/main/deploy/static/provider/kind/deploy.yaml

.PHONY: --wait-for-ingress-to-get-ready
--wait-for-ingress-to-get-ready:
	kubectl rollout status deployment ingress-nginx-controller -n ingress-nginx

.PHONY: --deploy-application
--deploy-application: 
	kubectl apply -f manifests

.PHONY: --kind-create
--kind-create: 
	kind create cluster --name ${KIND_CLUSTER} --config .kind/cluster-config.yml

.PHONY: infra-create
infra-create: --kind-create --apply-ingress-nginx --wait-for-ingress-to-get-ready --deploy-application


.PHONY: infra-delete
infra-delete:
	kind delete cluster --name ${KIND_CLUSTER}

.PHONY: kind-list
kind-list:
	kind get clusters



.PHONY: expose-ingress-controller
expose-ingress-controller: 
	kubectl port-forward svc/ingress-nginx-controller -n ingress-nginx 8080:80
	
.PHONY: golang-container-start
golang-container-start:
	docker run --rm -ti --network host -v $(PWD)/connector:/app --workdir /app golang:1.16-buster bash
