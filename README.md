# K8sCommerce - Cloud Native Ecommerce
<!-- ALL-CONTRIBUTORS-BADGE:START - Do not remove or modify this section -->
[![All Contributors](https://img.shields.io/badge/all_contributors-1-orange.svg?style=flat-square)](#contributors-)
<!-- ALL-CONTRIBUTORS-BADGE:END -->

K8sCommerce is in active development! The goal is to create a cloud native microservice-based ecommerce platform written in [Go](https://go.dev/)—for kubernetes deployments.

Project status: **pre-alpha**

## Project Origin

K8sCommerce began as a personal project of [localrivet](https://github.com/localrivet).

## Service Overview
![K8sCommerce Overview](docs/K8sCommerceOverview.png)

## Development Progress

The following shows what has been completed and what is yet to be done:

| Service          |    DB Table(s)     |       .proto       |        .api        |       logic        | tests | release |
| ---------------- | :----------------: | :----------------: | :----------------: | :----------------: | :---: | :-----: |
| Cart             | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |       |  alpha  |
| Catalog          | :white_check_mark: | :white_check_mark: | :white_check_mark: | :white_check_mark: |       |  alpha  |
| Customer         | :white_check_mark: | :white_check_mark: | :white_check_mark: |                    |       |    -    |
| Email            |         -          | :white_check_mark: |                    |                    |       |    -    |
| Inventory        | :white_check_mark: | :white_check_mark: |                    |                    |       |    -    |
| Others Bought    |         -          | :white_check_mark: |                    |                    |       |    -    |
| Payment          | :white_check_mark: | :white_check_mark: |                    |                    |       |    -    |
| Shipping         |         -          | :white_check_mark: |                    |                    |       |    -    |
| Similar Products | :white_check_mark: | :white_check_mark: |                    |                    |       |    -    |
| Store            | :white_check_mark: | :white_check_mark: |                    |                    |       |    -    |
| User             | :white_check_mark: | :white_check_mark: | :white_check_mark: |                    |       |  alpha  |
| Warehouse        |         -          | :white_check_mark: |                    |                    |       |    -    |


## Kubernetes Operator

Checkout our [kubernetes operator](https://github.com/k8scommerce/cluster-operator).

## Contributors ✨

Would you like to contribute? We need help in the following areas:
- Coding
- Testing
- Documentation
- Promotion
- Administration Development
- Kubernetes
- Terraform
- GitOps
- Etc.

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