# k8s-oidc-discovery-server

## Motivation

I wanted to host [the OIDC discovery server that amazon-eks-pod-identity-webhook uses for self-hosting](https://github.com/aws/amazon-eks-pod-identity-webhook/blob/d4269f48c6d6427f83e31ca52957b8c67d3c2fcf/SELF_HOSTED_SETUP.md) in Cloud Run instead of S3.

## How to use

This is intended to run on [Google Cloud Run](https://cloud.google.com/run).

PKCS #8 public key must be entered in the `PKCS_KEY` environment variable.

## License

Apache 2.0 - Copyright 2019 Amazon.com, Inc. or its affiliates. All Rights Reserved. See [LICENSE](LICENSE)

[pkg/key/key.go](pkg/key/key.go) is originally copied from https://github.com/aws/amazon-eks-pod-identity-webhook/blob/d4269f48c6d6427f83e31ca52957b8c67d3c2fcf/hack/self-hosted/main.go .