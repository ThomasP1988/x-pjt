list-port:
	kubectl get svc --all-namespaces -o go-template='{{range .items}}{{range.spec.ports}}{{if .nodePort}}{{.nodePort}}{{"\n"}}{{end}}{{end}}{{end}}'
k8s-local:
	minikube start
	minikube addons enable ingress
	minikube tunnel
expose-influx:
	minikube service --url influxdb-influxdb2
expose-orderbook:
	minikube service --url crypto-dai
install:
	helm repo add nginx-stable https://helm.nginx.com/stable
	helm repo add influxdata https://helm.influxdata.com/
dev:
	skaffold dev 