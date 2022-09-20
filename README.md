# Stock Ticker SRE Exercise

## Local Development and Deployment
Project was developed and tested locally with the following software installed:

```
docker >v4.2.0
grpcurl > v1.8.7
minikube >v1.27.0
kubernetes >1.25.1
kubectl >v1.14 (req for native kustomise support)
```

To run the project locally within docker, run the `docker-local.sh` script in the `scripts` directory.

To deploy the project to a vanilla Kubernetes environment, run the `start.sh` file at the project root directory. 

To then test the project ingress, run `grpcurl -plaintext localhost:9000 Ticker/getAvgPrice`

_N.B. My interpretation of the guidance was the last N days worth of data meant that if there was no data for one of the days, i.e. a weekend, the day should still to be counted. Both approaches were considered, but this was the most logical interpretation to me, and the code required makes the project much more extensible in the future. I wanted to highlight this now as given the `NDAYS=7` environment variable, it may not be immediately obvious why, for example, only five days may be returned._

## Initial exercise thoughts
Overall this was a fun exercise, and one of the most enjoyable coding challenges I've been provided.

While I intend to discuss this more in-depth during the next interview stage, the time series format of the API response proved surprisingly challenging to handle compared to most APIs, and proved the bulk of the challenge to overcome here.

I chose to implement this using the gRPC framework for a multitude of reasons despite no experience with it prior to the exercise. It appeared rather versatile for the scope of this task and particularly Part 3: Resilience. I look forward to discussing the approach with you all soon.

## Known Issues
While attempting to leverage the `grpc_health_v1` package for Part 2: Kubernetes in order to polish my deployment, the implemtation was not successful within the 5 hours I prescribed myself for this coding exercise and was a question of diminishing returns. This was nevertheless left in to demonstrate intent.