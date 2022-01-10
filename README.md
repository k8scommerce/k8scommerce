# K8sCommerce - Golang, Kubernetes & Cloud Native Ecommerce

K8sCommerce is a microservice-based ecommerce platform written in [Go](https://go.dev/)—for kubernetes deployments.

Checkout our [kubernetes operator](https://github.com/k8scommerce/cluster-operator).

## Service Overview
![K8sCommerce Overview](docs/K8sCommerceOverview.png)


## Development Progress

The following shows what has been completed and what is yet to be done:

| Service          | DB Table(s)       | .proto          | .api             | logic            |  tests           |  release        |
| ---------------- |:----------------:|:----------------:|:----------------:|:----------------:|:----------------:|:----------------:
| Cart             |:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:|                  | alpha           |
| Catalog          |:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:|                  | alpha           |
| Customer         |:white_check_mark:|:white_check_mark:|:white_check_mark:|                  |                  | pre-alpha       |
| Email            |-                 |:white_check_mark:|                  |                  |                  | pre-alpha       |
| Inventory        |:white_check_mark:|:white_check_mark:|                  |                  |                  | pre-alpha       |
| Others Bought    | -                |:white_check_mark:|                  |                  |                  | pre-alpha       |
| Payment          |:white_check_mark:|:white_check_mark:|                  |                  |                  | pre-alpha       |
| Shipping         | -                |:white_check_mark:|                  |                  |                  | pre-alpha       |
| Similar Products |:white_check_mark:|:white_check_mark:|                  |                  |                  | pre-alpha       |
| Store            |:white_check_mark:|:white_check_mark:|                  |                  |                  | pre-alpha       |
| User             |:white_check_mark:|:white_check_mark:|:white_check_mark:|                  |                  | alpha           |
| Warehouse        |-                 |:white_check_mark:|                  |                  |                  | pre-alpha       |



## Getting Started

K8sCommerce can be run in multiple ways:

- **Local single-node kubernetes kind cluster:** Recommended for evaluating K8sCommerce - [instructions]()

- **Deployment to your existing cluster:** Recommended for testing & production - [intructions]()

- **Running Clusterless:** Recommended for local development - [intructions]() 


### Quick Start
Deployment to a k8s cluster using the K8sCommerce Operator.

```sh
# Install RabbitMQ Operator for messaging
kubectl apply -f https://github.com/rabbitmq/cluster-operator/releases/latest/download/cluster-operator.yml

# Install K8sCommerce Operator
kubectl apply -f 
```

## Our Development Philosophy

We at K8sCommerce embrace the development philosophy: *"Prefer tools over conventions."*

We employ code-generation tools as often as possible to reduce repetitively writing foundational code such as CRUD operations, gRPC bindings, microservice boilerplate code. This methodology shortens the developer's learning curve while reducing human-introduced bugs.

We've incorporated the following code-generation tools in our projects to speed up development, integrations, and logic refinements, creating a desirable, streamlined, high-velocity development approach.

[Google's protocol buffers (protoc)](https://developers.google.com/protocol-buffers/): RPC client/server generation from .proto files with almost every language supported
[go-zero's goctl](https://github.com/zeromicro/go-zero): microservice framework generation from .proto files
[xo's xo](https://github.com/xo/xo): Database CRUD object generation for relational databases, with primary key, foreign key, enum, and advanced composite primary key support.

Preferring tools over conventions, combined with ever-growing documentation, empowers businesses to successfully integrate K8sCommerce's projects into their existing coding ecosystem in condensed timeframes. 


