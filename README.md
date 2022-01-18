# K8sCommerce - Cloud Native Ecommerce
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-1-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

K8sCommerce is a cloud native microservice-based ecommerce platform written in [Go](https://go.dev/)—for kubernetes deployments.

Checkout our [kubernetes operator](https://github.com/k8scommerce/cluster-operator).

## Service Overview
![K8sCommerce Overview](docs/K8sCommerceOverview.png)


## Development Progress

The following shows what has been completed and what is yet to be done:

| Service          | DB Table(s)       | .proto          | .api             | logic            |  tests           |  release        |
| ---------------- |:----------------:|:----------------:|:----------------:|:----------------:|:----------------:|:----------------:
| Cart             |:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:|                  | alpha           |
| Catalog          |:white_check_mark:|:white_check_mark:|:white_check_mark:|:white_check_mark:|                  | alpha           |
| Customer         |:white_check_mark:|:white_check_mark:|:white_check_mark:|                  |                  | -               |
| Email            |-                 |:white_check_mark:|                  |                  |                  | -               |
| Inventory        |:white_check_mark:|:white_check_mark:|                  |                  |                  | -               |
| Others Bought    | -                |:white_check_mark:|                  |                  |                  | -               |
| Payment          |:white_check_mark:|:white_check_mark:|                  |                  |                  | -               |
| Shipping         | -                |:white_check_mark:|                  |                  |                  | -               |
| Similar Products |:white_check_mark:|:white_check_mark:|                  |                  |                  | -               |
| Store            |:white_check_mark:|:white_check_mark:|                  |                  |                  | -               |
| User             |:white_check_mark:|:white_check_mark:|:white_check_mark:|                  |                  | alpha           |
| Warehouse        |-                 |:white_check_mark:|                  |                  |                  | -               |



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

## Project Development Philosophy

The K8sCommerce project embraces the development philosophy: *"Prefer tools over conventions."*

Employing code-generation tools as often as possible reduces repetitively writing foundational code such as CRUD operations, gRPC bindings, microservice boilerplate code. etc. This methodology shortens the developer's learning curve while reducing, and oftentimes eliminating, human-introduced bugs.

Incorporating the following code-generation tools in our project has helped speed up development, integrations, and logic refinements, creating a desirable, streamlined, high-velocity development approach.

- [Google's protocol buffers (protoc)](https://developers.google.com/protocol-buffers/): RPC client/server generation from .proto files with almost every language supported
- [go-zero's goctl](https://github.com/zeromicro/go-zero): microservice framework generation from .proto and .api files
- [xo's xo](https://github.com/xo/xo): Database CRUD object generation for relational databases, with primary key, foreign key, enum, and advanced composite primary key support.

Preferring tools over conventions, combined with ever-growing documentation, empowers businesses to successfully integrate K8sCommerce's projects into their existing coding ecosystem in condensed timeframes. 

## Contributors ✨

Thanks goes to these wonderful people ([emoji key](https://allcontributors.org/docs/en/emoji-key)):

<!-- ALL-CONTRIBUTORS-LIST:START - Do not remove or modify this section -->
<!-- prettier-ignore-start -->
<!-- markdownlint-disable -->
<table>
  <tr>
    <td align="center"><a href="https://github.com/localrivet"><img src="https://avatars.githubusercontent.com/u/833950?v=4?s=100" width="100px;" alt=""/><br /><sub><b>localrivet</b></sub></a><br /><a href="#infra-localrivet" title="Infrastructure (Hosting, Build-Tools, etc)">🚇</a> <a href="https://github.com/k8scommerce/k8scommerce/commits?author=localrivet" title="Code">💻</a> <a href="https://github.com/k8scommerce/k8scommerce/commits?author=localrivet" title="Documentation">📖</a> <a href="https://github.com/k8scommerce/k8scommerce/commits?author=localrivet" title="Tests">⚠️</a> <a href="#maintenance-localrivet" title="Maintenance">🚧</a> <a href="#content-localrivet" title="Content">🖋</a></td>
  </tr>
</table>

<!-- markdownlint-restore -->
<!-- prettier-ignore-end -->

<!-- ALL-CONTRIBUTORS-LIST:END -->

This project follows the [all-contributors](https://github.com/all-contributors/all-contributors) specification. Contributions of any kind welcome!