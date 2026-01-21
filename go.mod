module github.com/Hyodar/tdxs

go 1.24.6

// BEGIN: Constellation replace directives -----------------------------------
// TODO(daniel-weisse): revert after merging https://github.com/martinjungblut/go-cryptsetup/pull/16.
replace github.com/martinjungblut/go-cryptsetup => github.com/daniel-weisse/go-cryptsetup v0.0.0-20230705150314-d8c07bd1723c

// Kubernetes replace directives are required because we depend on k8s.io/kubernetes/cmd/kubeadm
// k8s discourages usage of k8s.io/kubernetes as a dependency, but no external staging repositories for kubeadm exist.
// Our code does not actually depend on these packages, but `go mod download` breaks without this replace directive.
// See this issue: https://github.com/kubernetes/kubernetes/issues/79384
// And this README: https://github.com/kubernetes/kubernetes/blob/master/staging/README.md
replace (
	k8s.io/cloud-provider => k8s.io/cloud-provider v0.32.3
	k8s.io/controller-manager => k8s.io/controller-manager v0.32.3
	k8s.io/cri-client => k8s.io/cri-client v0.32.3
	k8s.io/csi-translation-lib => k8s.io/csi-translation-lib v0.32.3
	k8s.io/dynamic-resource-allocation => k8s.io/dynamic-resource-allocation v0.32.3
	k8s.io/endpointslice => k8s.io/endpointslice v0.32.3
	k8s.io/externaljwt => k8s.io/externaljwt v0.32.3
	k8s.io/kube-aggregator => k8s.io/kube-aggregator v0.32.3
	k8s.io/kube-controller-manager => k8s.io/kube-controller-manager v0.32.3
	k8s.io/kube-proxy => k8s.io/kube-proxy v0.32.3
	k8s.io/kube-scheduler => k8s.io/kube-scheduler v0.32.3
	k8s.io/legacy-cloud-providers => k8s.io/legacy-cloud-providers v0.30.11
	k8s.io/pod-security-admission => k8s.io/pod-security-admission v0.32.3
	k8s.io/sample-apiserver => k8s.io/sample-apiserver v0.32.3
)

// END: Constellation replace directives -------------------------------------

require (
	github.com/coreos/go-systemd/v22 v22.5.0
	github.com/google/go-tdx-guest v0.3.2-0.20250505161510-9efd53b4a100
	github.com/google/go-tpm-tools v0.4.4
	github.com/spf13/cobra v1.9.1
	gopkg.in/yaml.v3 v3.0.1
)

require (
	github.com/google/go-cmp v0.7.0 // indirect
	github.com/google/go-sev-guest v0.13.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/kr/pretty v0.3.1 // indirect
	github.com/rogpeppe/go-internal v1.14.1 // indirect
	github.com/spf13/pflag v1.0.7 // indirect
	golang.org/x/crypto v0.40.0 // indirect
	google.golang.org/protobuf v1.36.6 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
)
