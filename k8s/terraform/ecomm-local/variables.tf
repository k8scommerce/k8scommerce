variable "kind_cluster_name" {
  type        = string
  description = "The name of the cluster."
  default     = "ecomm-local"
}

variable "kind_cluster_config_path" {
  type        = string
  description = "The location where this cluster's kubeconfig will be saved to."
  default     = "~/.kube/config"
}

# FLUX
variable "github_owner" {
  type        = string
  description = "github owner"
}

variable "github_token" {
  type        = string
  description = "github token"
}

variable "repository_name" {
  type        = string
  description = "github repository name"
}

variable "repository_visibility" {
  type        = string
  description = "How visible is the github repo"
  default     = "private"
}

variable "branch" {
  type        = string
  description = "branch name"
  default     = "master"
}

variable "target_path" {
  type        = string
  description = "flux sync target path"
}

# variable "ingress_nginx_helm_version" {
#   type        = string
#   description = "The Helm version for the nginx ingress controller."
#   default     = "4.0.6"
# }

# variable "ingress_nginx_namespace" {
#   type        = string
#   description = "The nginx ingress namespace (it will be created if needed)."
#   default     = "ingress-nginx"
# }
