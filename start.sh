minikube start
kubectl delete deploy stock-ticker
kubectl delete configmap stock-ticker-config
kubectl delete -k internal/templates

kubectl create -f internal/templates/configmap.yaml
kubectl apply -k internal/templates
kubectl apply -f deployment.yaml
echo "Sleeping for 15 seconds to allow the pod to start"
sleep 15
kubectl port-forward deployment/stock-ticker 9000:9000