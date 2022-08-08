# Okteto Golang Starter
The purpose of this repo is to provide required files to directly develop Golang Web API on Kubernetes with Okteto tool.

## Requirements
- Latest [Okteto](https://github.com/okteto/okteto/releases)  binary installed on your machine
- An operational Kubernetes or k3s cluster
- An Ingress Controller in your cluster (nginx in my case)
- A StorageClass to automaticaly create pv with pvc object. Your StorageClass should be set as default with this command:
```
kubectl patch storageclass MY-STORAGECLASS -p '{"metadata": {"annotations":{"storageclass.kubernetes.io/is-default-class":"true"}}}'
```

I personnaly use cert-manager and HashiCorp Vault to get a valid TLS certificate for my subdomain. If you don't use these tools you need to rewrite **backend-ing.yml** in manifests directory to work with letsencrypt or http (no tls).<br>
These documentations can help you:
- For cert-manager letsencrypt issuer: https://cert-manager.io/docs/configuration/acme
- For HTTP ingress: https://kubernetes.io/docs/concepts/services-networking/ingress/#the-ingress-resource

## First Usage
If it's the first time you use okteto and this repo you just need to do these commands:
```
git clone https://github.com/aamoyel/okteto-go-starter.git
kubectl apply -f manifests
okteto up
```
When you call 'okteto up' for the first time you need to select a kubernetes context.

## Air Binary
You need to have Air binary to automatically build and run your go project. Your can find [Air](https://github.com/cosmtrek/air) installation instructions down below.<br>
After 'okteto up' (in dev container), run this command :
```
curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin
```

After this, simply type 'air' in your terminal

## Results
Go to http://localhost:8080 to see local development container process and at your ingress url, https://dev.amoyel.loc in my case.<br>
You can change domain and subdomain in **backend-cert.yml** file in manifests directory.

## End Dev Session:
Simply enter CTRL+C to stop air binary.<br>
Then do CTRL+D or type 'exit'.<br>
And finaly type 'okteto down' to avoid resources consumption.

## Common usage
```
okteto up
air
```
Develop your Golang API ...<br>
Do [this step](https://git.amoyel.fr/opsfinity/okteto-go-starter#end-dev-session) when you want to stop your dev session.

## Change Golang Version
If you want to change Go version, simply change:
- Image tag in **okteto.yml**, eg: "_image: golang:1.18.5-bullseye_"
- Image tag in **manifests/backend-deploy.yml**, eg: "_- image: golang:1.18.5-bullseye_"

You can find other tags directly on Docker Hub.

## More Info
To find more information on okteto.yml and okteto cli go to https://www.okteto.com/docs/reference/cli/
