// Code generated by the Pulumi Terraform Bridge (tfgen) Tool DO NOT EDIT.
// *** WARNING: Do not edit by hand unless you're certain you know what you are doing! ***

package flux

import (
	"context"
	"reflect"

	"github.com/pulumi/pulumi/sdk/v3/go/pulumi"
)

func GetFluxSync(ctx *pulumi.Context, args *GetFluxSyncArgs, opts ...pulumi.InvokeOption) (*GetFluxSyncResult, error) {
	var rv GetFluxSyncResult
	err := ctx.Invoke("flux:index/getFluxSync:getFluxSync", args, &rv, opts...)
	if err != nil {
		return nil, err
	}
	return &rv, nil
}

// A collection of arguments for invoking getFluxSync.
type GetFluxSyncArgs struct {
	Branch            *string  `pulumi:"branch"`
	Commit            *string  `pulumi:"commit"`
	GitImplementation *string  `pulumi:"gitImplementation"`
	Interval          *int     `pulumi:"interval"`
	Name              *string  `pulumi:"name"`
	Namespace         *string  `pulumi:"namespace"`
	PatchNames        []string `pulumi:"patchNames"`
	Secret            *string  `pulumi:"secret"`
	Semver            *string  `pulumi:"semver"`
	Tag               *string  `pulumi:"tag"`
	TargetPath        string   `pulumi:"targetPath"`
	Url               string   `pulumi:"url"`
}

// A collection of values returned by getFluxSync.
type GetFluxSyncResult struct {
	Branch            *string `pulumi:"branch"`
	Commit            *string `pulumi:"commit"`
	Content           string  `pulumi:"content"`
	GitImplementation *string `pulumi:"gitImplementation"`
	// The provider-assigned unique ID for this managed resource.
	Id               string            `pulumi:"id"`
	Interval         *int              `pulumi:"interval"`
	KustomizeContent string            `pulumi:"kustomizeContent"`
	KustomizePath    string            `pulumi:"kustomizePath"`
	Name             *string           `pulumi:"name"`
	Namespace        *string           `pulumi:"namespace"`
	PatchFilePaths   map[string]string `pulumi:"patchFilePaths"`
	PatchNames       []string          `pulumi:"patchNames"`
	Path             string            `pulumi:"path"`
	Secret           *string           `pulumi:"secret"`
	Semver           *string           `pulumi:"semver"`
	Tag              *string           `pulumi:"tag"`
	TargetPath       string            `pulumi:"targetPath"`
	Url              string            `pulumi:"url"`
}

func GetFluxSyncOutput(ctx *pulumi.Context, args GetFluxSyncOutputArgs, opts ...pulumi.InvokeOption) GetFluxSyncResultOutput {
	return pulumi.ToOutputWithContext(context.Background(), args).
		ApplyT(func(v interface{}) (GetFluxSyncResult, error) {
			args := v.(GetFluxSyncArgs)
			r, err := GetFluxSync(ctx, &args, opts...)
			var s GetFluxSyncResult
			if r != nil {
				s = *r
			}
			return s, err
		}).(GetFluxSyncResultOutput)
}

// A collection of arguments for invoking getFluxSync.
type GetFluxSyncOutputArgs struct {
	Branch            pulumi.StringPtrInput   `pulumi:"branch"`
	Commit            pulumi.StringPtrInput   `pulumi:"commit"`
	GitImplementation pulumi.StringPtrInput   `pulumi:"gitImplementation"`
	Interval          pulumi.IntPtrInput      `pulumi:"interval"`
	Name              pulumi.StringPtrInput   `pulumi:"name"`
	Namespace         pulumi.StringPtrInput   `pulumi:"namespace"`
	PatchNames        pulumi.StringArrayInput `pulumi:"patchNames"`
	Secret            pulumi.StringPtrInput   `pulumi:"secret"`
	Semver            pulumi.StringPtrInput   `pulumi:"semver"`
	Tag               pulumi.StringPtrInput   `pulumi:"tag"`
	TargetPath        pulumi.StringInput      `pulumi:"targetPath"`
	Url               pulumi.StringInput      `pulumi:"url"`
}

func (GetFluxSyncOutputArgs) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFluxSyncArgs)(nil)).Elem()
}

// A collection of values returned by getFluxSync.
type GetFluxSyncResultOutput struct{ *pulumi.OutputState }

func (GetFluxSyncResultOutput) ElementType() reflect.Type {
	return reflect.TypeOf((*GetFluxSyncResult)(nil)).Elem()
}

func (o GetFluxSyncResultOutput) ToGetFluxSyncResultOutput() GetFluxSyncResultOutput {
	return o
}

func (o GetFluxSyncResultOutput) ToGetFluxSyncResultOutputWithContext(ctx context.Context) GetFluxSyncResultOutput {
	return o
}

func (o GetFluxSyncResultOutput) Branch() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFluxSyncResult) *string { return v.Branch }).(pulumi.StringPtrOutput)
}

func (o GetFluxSyncResultOutput) Commit() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFluxSyncResult) *string { return v.Commit }).(pulumi.StringPtrOutput)
}

func (o GetFluxSyncResultOutput) Content() pulumi.StringOutput {
	return o.ApplyT(func(v GetFluxSyncResult) string { return v.Content }).(pulumi.StringOutput)
}

func (o GetFluxSyncResultOutput) GitImplementation() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFluxSyncResult) *string { return v.GitImplementation }).(pulumi.StringPtrOutput)
}

// The provider-assigned unique ID for this managed resource.
func (o GetFluxSyncResultOutput) Id() pulumi.StringOutput {
	return o.ApplyT(func(v GetFluxSyncResult) string { return v.Id }).(pulumi.StringOutput)
}

func (o GetFluxSyncResultOutput) Interval() pulumi.IntPtrOutput {
	return o.ApplyT(func(v GetFluxSyncResult) *int { return v.Interval }).(pulumi.IntPtrOutput)
}

func (o GetFluxSyncResultOutput) KustomizeContent() pulumi.StringOutput {
	return o.ApplyT(func(v GetFluxSyncResult) string { return v.KustomizeContent }).(pulumi.StringOutput)
}

func (o GetFluxSyncResultOutput) KustomizePath() pulumi.StringOutput {
	return o.ApplyT(func(v GetFluxSyncResult) string { return v.KustomizePath }).(pulumi.StringOutput)
}

func (o GetFluxSyncResultOutput) Name() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFluxSyncResult) *string { return v.Name }).(pulumi.StringPtrOutput)
}

func (o GetFluxSyncResultOutput) Namespace() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFluxSyncResult) *string { return v.Namespace }).(pulumi.StringPtrOutput)
}

func (o GetFluxSyncResultOutput) PatchFilePaths() pulumi.StringMapOutput {
	return o.ApplyT(func(v GetFluxSyncResult) map[string]string { return v.PatchFilePaths }).(pulumi.StringMapOutput)
}

func (o GetFluxSyncResultOutput) PatchNames() pulumi.StringArrayOutput {
	return o.ApplyT(func(v GetFluxSyncResult) []string { return v.PatchNames }).(pulumi.StringArrayOutput)
}

func (o GetFluxSyncResultOutput) Path() pulumi.StringOutput {
	return o.ApplyT(func(v GetFluxSyncResult) string { return v.Path }).(pulumi.StringOutput)
}

func (o GetFluxSyncResultOutput) Secret() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFluxSyncResult) *string { return v.Secret }).(pulumi.StringPtrOutput)
}

func (o GetFluxSyncResultOutput) Semver() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFluxSyncResult) *string { return v.Semver }).(pulumi.StringPtrOutput)
}

func (o GetFluxSyncResultOutput) Tag() pulumi.StringPtrOutput {
	return o.ApplyT(func(v GetFluxSyncResult) *string { return v.Tag }).(pulumi.StringPtrOutput)
}

func (o GetFluxSyncResultOutput) TargetPath() pulumi.StringOutput {
	return o.ApplyT(func(v GetFluxSyncResult) string { return v.TargetPath }).(pulumi.StringOutput)
}

func (o GetFluxSyncResultOutput) Url() pulumi.StringOutput {
	return o.ApplyT(func(v GetFluxSyncResult) string { return v.Url }).(pulumi.StringOutput)
}

func init() {
	pulumi.RegisterOutputType(GetFluxSyncResultOutput{})
}
