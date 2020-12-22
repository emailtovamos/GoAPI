docker build --no-cache -t emailtovamos/goapi:v1 .

docker push emailtovamos/goapi:v1

kubectl delete deployment goapi

kubectl apply -f deployment.yaml

kubectl apply -f service.yaml