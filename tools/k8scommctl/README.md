# k8scomm2hugo

Queries your K8sCommerce store's API and generates Hugo Front Matter markdown files for both categories & products.

## Installation

```bash
go install github.com/k8scommerce/k8scomm2hugo@latest
```

## Example

```bash
k8scomm2hugo generate -e https://yourstoreapi.url -k yoursitekey -o ./content -p products -c categories
```

## Usage

```bash
Generates Hugo markdown files from K8sCommerce using the K8sCommerce API for category & product generation.

Usage:
  k8scomm2hugo [command]

Available Commands:
  completion  Generate the autocompletion script for the specified shell
  generate    Generates Hugo markdown files from K8sCommerce categories & products
  help        Help about any command

Flags:
  -h, --help     help for k8scomm2hugo
  -t, --toggle   Help message for toggle

Use "k8scomm2hugo [command] --help" for more information about a command.
```
