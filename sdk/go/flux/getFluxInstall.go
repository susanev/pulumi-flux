// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package flux

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

// `getFluxInstall` can be used to generate Kubernetes manifests for deploying Flux.
//
// ## Example Usage
//
// ```go
// package main
//
// import (
//
//	"github.com/oun/pulumi-flux/sdk/go/flux"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
//	"github.com/pulumi/pulumi/sdk/v3/go/pulumi/config"
//
// )
//
//	func main() {
//		pulumi.Run(func(ctx *pulumi.Context) error {
//			cfg := config.New(ctx, "")
//			targetPath := cfg.Require("targetPath")
//			_, err := flux.GetFluxInstall(ctx, &flux.GetFluxInstallArgs{
//				TargetPath: targetPath,
//			}, nil)
//			if err != nil {
//				return err
//			}
//			return nil
//		})
//	}
//
// ```
func GetFluxInstall(ctx *pulumi.Context, args *GetFluxInstallArgs, opts ...pulumi.InvokeOption) (*GetFluxInstallResult, error) {
	opts = pkgInvokeDefaultOpts(opts)
	var rv GetFluxInstallResult
	err := ctx.Invoke("flux:index/getFluxInstall:getFluxInstall", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFluxInstall.
type GetFluxInstallArgs struct {
	// Base URL to get the install manifests from. When specifying this, `version` should also be set to the corresponding version to download from that URL, or the latest version associated with upstream flux will be requested. Defaults to `https://github.com/fluxcd/flux2/releases`.
	Baseurl *string `pulumi:"baseurl"`
	// The internal cluster domain. Defaults to `cluster.local`.
	ClusterDomain *string `pulumi:"clusterDomain"`
	// Toolkit components to include in the install manifests.
	Components []string `pulumi:"components"`
	// List of extra components to include in the install manifests.
	ComponentsExtras []string `pulumi:"componentsExtras"`
	// Kubernetes secret name used for pulling the toolkit images from a private registry.
	ImagePullSecrets *string `pulumi:"imagePullSecrets"`
	// Log level for toolkit components. Defaults to `info`.
	LogLevel *string `pulumi:"logLevel"`
	// The namespace scope for install manifests. Defaults to `flux-system`.
	Namespace *string `pulumi:"namespace"`
	// Deny ingress access to the toolkit controllers from other namespaces using network policies. Defaults to `true`.
	NetworkPolicy *bool `pulumi:"networkPolicy"`
	// Container registry where the toolkit images are published. Defaults to `ghcr.io/fluxcd`.
	Registry *string `pulumi:"registry"`
	// Relative path to the Git repository root where Flux manifests are committed.
	TargetPath string `pulumi:"targetPath"`
	// List of toleration keys used to schedule the components pods onto nodes with matching taints.
	TolerationKeys []string `pulumi:"tolerationKeys"`
	// Flux version. Defaults to `v0.41.2`.
	Version *string `pulumi:"version"`
	// If true watch for custom resources in all namespaces. Defaults to `true`.
	WatchAllNamespaces *bool `pulumi:"watchAllNamespaces"`
}

// A collection of values returned by getFluxInstall.
type GetFluxInstallResult struct {
	// Base URL to get the install manifests from. When specifying this, `version` should also be set to the corresponding version to download from that URL, or the latest version associated with upstream flux will be requested. Defaults to `https://github.com/fluxcd/flux2/releases`.
	Baseurl *string `pulumi:"baseurl"`
	// The internal cluster domain. Defaults to `cluster.local`.
	ClusterDomain *string `pulumi:"clusterDomain"`
	// Toolkit components to include in the install manifests.
	Components []string `pulumi:"components"`
	// List of extra components to include in the install manifests.
	ComponentsExtras []string `pulumi:"componentsExtras"`
	// Manifests in multi-doc yaml format.
	Content string `pulumi:"content"`
	// The ID of this resource.
	Id string `pulumi:"id"`
	// Kubernetes secret name used for pulling the toolkit images from a private registry.
	ImagePullSecrets *string `pulumi:"imagePullSecrets"`
	// Log level for toolkit components. Defaults to `info`.
	LogLevel *string `pulumi:"logLevel"`
	// The namespace scope for install manifests. Defaults to `flux-system`.
	Namespace *string `pulumi:"namespace"`
	// Deny ingress access to the toolkit controllers from other namespaces using network policies. Defaults to `true`.
	NetworkPolicy *bool `pulumi:"networkPolicy"`
	// Expected path of content in git repository.
	Path string `pulumi:"path"`
	// Container registry where the toolkit images are published. Defaults to `ghcr.io/fluxcd`.
	Registry *string `pulumi:"registry"`
	// Relative path to the Git repository root where Flux manifests are committed.
	TargetPath string `pulumi:"targetPath"`
	// List of toleration keys used to schedule the components pods onto nodes with matching taints.
	TolerationKeys []string `pulumi:"tolerationKeys"`
	// Flux version. Defaults to `v0.41.2`.
	Version *string `pulumi:"version"`
	// If true watch for custom resources in all namespaces. Defaults to `true`.
	WatchAllNamespaces *bool `pulumi:"watchAllNamespaces"`
}

func GetFluxInstallOutput(ctx *pulumi.Context, args GetFluxInstallOutputArgs, opts ...pulumi.InvokeOption) GetFluxInstallResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFluxInstallResult, error) {
			args := v.(GetFluxInstallArgs)
			r, err := GetFluxInstall(ctx, &args, opts...)
			var s GetFluxInstallResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFluxInstallResultOutput)
}

// A collection of arguments for invoking getFluxInstall.
type GetFluxInstallOutputArgs struct {
	// Base URL to get the install manifests from. When specifying this, `version` should also be set to the corresponding version to download from that URL, or the latest version associated with upstream flux will be requested. Defaults to `https://github.com/fluxcd/flux2/releases`.
	Baseurl pulumi.StringPtrInput `pulumi:"baseurl"`
	// The internal cluster domain. Defaults to `cluster.local`.
	ClusterDomain pulumi.StringPtrInput `pulumi:"clusterDomain"`
	// Toolkit components to include in the install manifests.
	Components pulumi.StringArrayInput `pulumi:"components"`
	// List of extra components to include in the install manifests.
	ComponentsExtras pulumi.StringArrayInput `pulumi:"componentsExtras"`
	// Kubernetes secret name used for pulling the toolkit images from a private registry.
	ImagePullSecrets pulumi.StringPtrInput `pulumi:"imagePullSecrets"`
	// Log level for toolkit components. Defaults to `info`.
	LogLevel pulumi.StringPtrInput `pulumi:"logLevel"`
	// The namespace scope for install manifests. Defaults to `flux-system`.
	Namespace pulumi.StringPtrInput `pulumi:"namespace"`
	// Deny ingress access to the toolkit controllers from other namespaces using network policies. Defaults to `true`.
	NetworkPolicy pulumi.BoolPtrInput `pulumi:"networkPolicy"`
	// Container registry where the toolkit images are published. Defaults to `ghcr.io/fluxcd`.
	Registry pulumi.StringPtrInput `pulumi:"registry"`
	// Relative path to the Git repository root where Flux manifests are committed.
	TargetPath pulumi.StringInput `pulumi:"targetPath"`
	// List of toleration keys used to schedule the components pods onto nodes with matching taints.
	TolerationKeys pulumi.StringArrayInput `pulumi:"tolerationKeys"`
	// Flux version. Defaults to `v0.41.2`.
	Version pulumi.StringPtrInput `pulumi:"version"`
	// If true watch for custom resources in all namespaces. Defaults to `true`.
	WatchAllNamespaces pulumi.BoolPtrInput `pulumi:"watchAllNamespaces"`
}

func (GetFluxInstallOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFluxInstallArgs)(nil)).Elem()
}

// A collection of values returned by getFluxInstall.
type GetFluxInstallResultOutput struct{ *pulumi.OutputState }

func (GetFluxInstallResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFluxInstallResult)(nil)).Elem()
}

func (o GetFluxInstallResultOutput) ToGetFluxInstallResultOutput() GetFluxInstallResultOutput {
	return o
}

func (o GetFluxInstallResultOutput) ToGetFluxInstallResultOutputWithContext(ctx context.Context) GetFluxInstallResultOutput {
	return o
}

// Base URL to get the install manifests from. When specifying this, `version` should also be set to the corresponding version to download from that URL, or the latest version associated with upstream flux will be requested. Defaults to `https://github.com/fluxcd/flux2/releases`.
func (o GetFluxInstallResultOutput) Baseurl() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFluxInstallResult) *string { return v.Baseurl }).(pulumi.StringPtrOutput)
}

// The internal cluster domain. Defaults to `cluster.local`.
func (o GetFluxInstallResultOutput) ClusterDomain() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFluxInstallResult) *string { return v.ClusterDomain }).(pulumi.StringPtrOutput)
}

// Toolkit components to include in the install manifests.
func (o GetFluxInstallResultOutput) Components() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFluxInstallResult) []string { return v.Components }).(pulumi.StringArrayOutput)
}

// List of extra components to include in the install manifests.
func (o GetFluxInstallResultOutput) ComponentsExtras() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFluxInstallResult) []string { return v.ComponentsExtras }).(pulumi.StringArrayOutput)
}

// Manifests in multi-doc yaml format.
func (o GetFluxInstallResultOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v GetFluxInstallResult) string { return v.Content }).(pulumi.StringOutput)
}

// The ID of this resource.
func (o GetFluxInstallResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFluxInstallResult) string { return v.Id }).(pulumi.StringOutput)
}

// Kubernetes secret name used for pulling the toolkit images from a private registry.
func (o GetFluxInstallResultOutput) ImagePullSecrets() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFluxInstallResult) *string { return v.ImagePullSecrets }).(pulumi.StringPtrOutput)
}

// Log level for toolkit components. Defaults to `info`.
func (o GetFluxInstallResultOutput) LogLevel() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFluxInstallResult) *string { return v.LogLevel }).(pulumi.StringPtrOutput)
}

// The namespace scope for install manifests. Defaults to `flux-system`.
func (o GetFluxInstallResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFluxInstallResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

// Deny ingress access to the toolkit controllers from other namespaces using network policies. Defaults to `true`.
func (o GetFluxInstallResultOutput) NetworkPolicy() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetFluxInstallResult) *bool { return v.NetworkPolicy }).(pulumi.BoolPtrOutput)
}

// Expected path of content in git repository.
func (o GetFluxInstallResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v GetFluxInstallResult) string { return v.Path }).(pulumi.StringOutput)
}

// Container registry where the toolkit images are published. Defaults to `ghcr.io/fluxcd`.
func (o GetFluxInstallResultOutput) Registry() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFluxInstallResult) *string { return v.Registry }).(pulumi.StringPtrOutput)
}

// Relative path to the Git repository root where Flux manifests are committed.
func (o GetFluxInstallResultOutput) TargetPath() pulumi.StringOutput {
	return o.ApplyT(func(v GetFluxInstallResult) string { return v.TargetPath }).(pulumi.StringOutput)
}

// List of toleration keys used to schedule the components pods onto nodes with matching taints.
func (o GetFluxInstallResultOutput) TolerationKeys() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFluxInstallResult) []string { return v.TolerationKeys }).(pulumi.StringArrayOutput)
}

// Flux version. Defaults to `v0.41.2`.
func (o GetFluxInstallResultOutput) Version() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFluxInstallResult) *string { return v.Version }).(pulumi.StringPtrOutput)
}

// If true watch for custom resources in all namespaces. Defaults to `true`.
func (o GetFluxInstallResultOutput) WatchAllNamespaces() pulumi.BoolPtrOutput {
	return o.ApplyT(func(v GetFluxInstallResult) *bool { return v.WatchAllNamespaces }).(pulumi.BoolPtrOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFluxInstallResultOutput{})
}
