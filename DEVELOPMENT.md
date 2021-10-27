
## Environment Setup

Docker for Desktop on MacOS is highly recommended for KubeRay development. You can also use `minikube` or `Kind` to bring up local kubernetes cluster.


## Development

1. Deploy Ray Operator

```
kubectl apply -k "github.com/ray-project/kuberay/ray-operator/config/default"
```

> Note: If you need to test ray operator in dev branch, please check `ray-operator/README.md` for more details.

2. Build and run backend service 


```
cd backend
go build -a -o service cmd/main.go
./service
```

> Note: by default, it uses the `~/.kube/config` to connect to your cluster. Please make the corresponding if you need a custom path

```
# pkg/client/cluster.go
defaultKubeConfigPath := filepath.Join(home, ".kube", "config")
```


3. Build Cli

```
cd cli
go build -a -o kuberay main.go
./kuberay help
```

> Note: cli uses default local port `localhost:8887` as conn string. 


```
# cli/pkg/cmdutil/client.go
func GetGrpcConn() (*grpc.ClientConn, error) {
	address := "localhost:8887"
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithTimeout(5 * time.Second))
	if err != nil {
		log.Fatalf("can not connect: %v", err)
	}

	return conn, err
}
```


