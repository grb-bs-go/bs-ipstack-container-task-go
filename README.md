# Welcome to the geoip page!
Public Repo: https://github.com/grb-bs-go/bs-ipstack-container-task-go

## Description:
The geoip utility is a simple golang command-line tool that accepts an IP Address, queries the IPStack API (api.ipstack.com) and outputs the corresponding location lattitude and longitude in JSON format. A valid access key (access-key) is required to be input as a command line argument (register for free account with limited access on www.ipstack.com).

### Usage
Simply type 'geoip' on the IDE or command-line for usage info.

$ geoip

Usage: geoIP ip-address access-key

Example (IPv4): geopIP 80.44.77.120 a4js4jd2ld3eddKd3d

Example (IPv6): geopIP 0:0:0:0:0:ffff:502c:4d78 a4js4jd2ld3eddKd3d

JSON Response: {"IP":"80.44.77.120","Latitude":"52.569950103759766","Longitude":"1.1133400201797485"}

### Container Image
The corresponding image is available from github.

Image Name: jessltd2007/geoip-bs-go  

URL: https://hub.docker.com/repository/docker/jessltd2007/geoip-bs-go

# Run geoip
## Run using Docker

$ docker run --name geoip-demo -d jessltd2007/geoip-bs-go

$ docker exec -it geoip-demo sh -c /geoip

Usage: geoIP ip-address access-key

Example (IPv4): geopIP 80.44.77.120 a4js4jd2ld3eddKd3d

Example (IPv6): geopIP 0:0:0:0:0:ffff:502c:4d78 a4js4jd2ld3eddKd3d

JSON Response: {"IP":"80.44.77.120","Latitude":"52.569950103759766","Longitude":"1.1133400201797485"}


 

$ docker exec -it geoip-demo sh -c "/geoip 80.44.77.120 access-key"

{"IP":"80.44.77.120","Latitude":"52.569950103759766","Longitude":"1.1133400201797485"}


 
$ docker exec -it geoip-demo sh -c "/geoip 1.2.3.4 access-key"

{"IP":"1.2.3.4","Latitude":"-27.467580795288086","Longitude":"153.02789306640625"}

## Run using Kubernetes
If you have access to a K8s cluster, create the following ephemeral pod.

$ kubectl run -i geoip-demo --image=jessltd2007/geoip-bs-go --restart=Never -- /geoip 1.2.3.4 access-key

If you don't see a command prompt, try pressing enter.

{"IP":"1.2.3.4","Latitude":"-27.467580795288086","Longitude":"153.02789306640625"}

# geoip Dev & Build
The following docker commands were used to build and upload the image (refer simple Dockerfile).

$ docker build -t jessltd2007/geoip-bs-go .

$ docker push jessltd2007/geoip-bs-go

### IDE (VSCode)

To install/run app within IDE or CL...

install latest go & git

gh repo clone grb-bs-go/bs-ipstack-container-task-go

go mod download

go run main.go

Functional Testing

go test

 
### Issue with IPv6 not working (api.ipstack.com)
geoip IPv4 address location resolution works, but the IPStack API consistently returns ZERO latitude/longitude values for all IPv6 addresses. 

$ geoip 0:0:0:0:0:ffff:502c:4d78 access-key

{"IP":"::ffff:502c:4d78","Latitude":"0","Longitude":"0"}

I noticed same result when inputting all common forms of IPv6 address formats (short/long) into the www.ipstack.com website form tool, so assume IPv6 not supported by IPStack (would require further investigation).

# Security
Security is an entire subject encompassing layered infrastructure-centric security practices and conventions across the Cloud/Cluster/Container/Code domain layers. The contemporary IaC Pipeline-based Build & Deploy of containerised applications can be based on any number of different Build/Test/Image/Scan/Verify cloud-native stacks and toolsets. Container security testing should encompass both static code & dynamic environment (SAST/DAST/IAST) regimes for example Trivy & Clair (currently learning as part of CNCF CKS training/exam). In this simple golang application context, note the access-key is deliberately not included in the program logging output. This secret access key should be securely stored (cloud secrets/vault etc) and must be input by each geoip utility user. Run within a K8s cluster, we would rely on RBAC and layered cluster security best-practices to secure this containerised app, probably adopting the typical microservices-based approach.



 
