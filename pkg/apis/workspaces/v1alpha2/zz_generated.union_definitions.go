package v1alpha2

import (
	"reflect"
)

var commandUnion reflect.Type = reflect.TypeOf(CommandUnionVisitor{})

func (union CommandUnion) Visit(visitor CommandUnionVisitor) error {
	return visitUnion(union, visitor)
}
func (union *CommandUnion) discriminator() *string {
	return (*string)(&union.CommandType)
}
func (union *CommandUnion) Normalize() error {
	return normalizeUnion(union, commandUnion)
}
func (union *CommandUnion) Simplify() {
	simplifyUnion(union, commandUnion)
}

// +k8s:deepcopy-gen=false
type CommandUnionVisitor struct {
	Exec         func(*ExecCommand) error
	Apply        func(*ApplyCommand) error
	VscodeTask   func(*VscodeConfigurationCommand) error
	VscodeLaunch func(*VscodeConfigurationCommand) error
	Composite    func(*CompositeCommand) error
	Custom       func(*CustomCommand) error
}

var vscodeConfigurationCommandLocation reflect.Type = reflect.TypeOf(VscodeConfigurationCommandLocationVisitor{})

func (union VscodeConfigurationCommandLocation) Visit(visitor VscodeConfigurationCommandLocationVisitor) error {
	return visitUnion(union, visitor)
}
func (union *VscodeConfigurationCommandLocation) discriminator() *string {
	return (*string)(&union.LocationType)
}
func (union *VscodeConfigurationCommandLocation) Normalize() error {
	return normalizeUnion(union, vscodeConfigurationCommandLocation)
}
func (union *VscodeConfigurationCommandLocation) Simplify() {
	simplifyUnion(union, vscodeConfigurationCommandLocation)
}

// +k8s:deepcopy-gen=false
type VscodeConfigurationCommandLocationVisitor struct {
	Uri     func(string) error
	Inlined func(string) error
}

var componentUnion reflect.Type = reflect.TypeOf(ComponentUnionVisitor{})

func (union ComponentUnion) Visit(visitor ComponentUnionVisitor) error {
	return visitUnion(union, visitor)
}
func (union *ComponentUnion) discriminator() *string {
	return (*string)(&union.ComponentType)
}
func (union *ComponentUnion) Normalize() error {
	return normalizeUnion(union, componentUnion)
}
func (union *ComponentUnion) Simplify() {
	simplifyUnion(union, componentUnion)
}

// +k8s:deepcopy-gen=false
type ComponentUnionVisitor struct {
	Container  func(*ContainerComponent) error
	Kubernetes func(*KubernetesComponent) error
	Openshift  func(*OpenshiftComponent) error
	Plugin     func(*PluginComponent) error
	Volume     func(*VolumeComponent) error
	Custom     func(*CustomComponent) error
}

var pluginComponentsOverrideUnion reflect.Type = reflect.TypeOf(PluginComponentsOverrideUnionVisitor{})

func (union PluginComponentsOverrideUnion) Visit(visitor PluginComponentsOverrideUnionVisitor) error {
	return visitUnion(union, visitor)
}
func (union *PluginComponentsOverrideUnion) discriminator() *string {
	return (*string)(&union.ComponentType)
}
func (union *PluginComponentsOverrideUnion) Normalize() error {
	return normalizeUnion(union, pluginComponentsOverrideUnion)
}
func (union *PluginComponentsOverrideUnion) Simplify() {
	simplifyUnion(union, pluginComponentsOverrideUnion)
}

// +k8s:deepcopy-gen=false
type PluginComponentsOverrideUnionVisitor struct {
	Container  func(*ContainerComponent) error
	Volume     func(*VolumeComponent) error
	Kubernetes func(*KubernetesComponent) error
	Openshift  func(*OpenshiftComponent) error
}

var importReferenceUnion reflect.Type = reflect.TypeOf(ImportReferenceUnionVisitor{})

func (union ImportReferenceUnion) Visit(visitor ImportReferenceUnionVisitor) error {
	return visitUnion(union, visitor)
}
func (union *ImportReferenceUnion) discriminator() *string {
	return (*string)(&union.ImportReferenceType)
}
func (union *ImportReferenceUnion) Normalize() error {
	return normalizeUnion(union, importReferenceUnion)
}
func (union *ImportReferenceUnion) Simplify() {
	simplifyUnion(union, importReferenceUnion)
}

// +k8s:deepcopy-gen=false
type ImportReferenceUnionVisitor struct {
	Uri        func(string) error
	Id         func(string) error
	Kubernetes func(*KubernetesCustomResourceImportReference) error
}

var k8sLikeComponentLocation reflect.Type = reflect.TypeOf(K8sLikeComponentLocationVisitor{})

func (union K8sLikeComponentLocation) Visit(visitor K8sLikeComponentLocationVisitor) error {
	return visitUnion(union, visitor)
}
func (union *K8sLikeComponentLocation) discriminator() *string {
	return (*string)(&union.LocationType)
}
func (union *K8sLikeComponentLocation) Normalize() error {
	return normalizeUnion(union, k8sLikeComponentLocation)
}
func (union *K8sLikeComponentLocation) Simplify() {
	simplifyUnion(union, k8sLikeComponentLocation)
}

// +k8s:deepcopy-gen=false
type K8sLikeComponentLocationVisitor struct {
	Uri     func(string) error
	Inlined func(string) error
}

var projectSource reflect.Type = reflect.TypeOf(ProjectSourceVisitor{})

func (union ProjectSource) Visit(visitor ProjectSourceVisitor) error {
	return visitUnion(union, visitor)
}
func (union *ProjectSource) discriminator() *string {
	return (*string)(&union.SourceType)
}
func (union *ProjectSource) Normalize() error {
	return normalizeUnion(union, projectSource)
}
func (union *ProjectSource) Simplify() {
	simplifyUnion(union, projectSource)
}

// +k8s:deepcopy-gen=false
type ProjectSourceVisitor struct {
	Git    func(*GitProjectSource) error
	Github func(*GithubProjectSource) error
	Zip    func(*ZipProjectSource) error
	Custom func(*CustomProjectSource) error
}

var componentUnionParentOverride reflect.Type = reflect.TypeOf(ComponentUnionParentOverrideVisitor{})

func (union ComponentUnionParentOverride) Visit(visitor ComponentUnionParentOverrideVisitor) error {
	return visitUnion(union, visitor)
}
func (union *ComponentUnionParentOverride) discriminator() *string {
	return (*string)(&union.ComponentType)
}
func (union *ComponentUnionParentOverride) Normalize() error {
	return normalizeUnion(union, componentUnionParentOverride)
}
func (union *ComponentUnionParentOverride) Simplify() {
	simplifyUnion(union, componentUnionParentOverride)
}

// +k8s:deepcopy-gen=false
type ComponentUnionParentOverrideVisitor struct {
	Container  func(*ContainerComponentParentOverride) error
	Kubernetes func(*KubernetesComponentParentOverride) error
	Openshift  func(*OpenshiftComponentParentOverride) error
	Plugin     func(*PluginComponentParentOverride) error
	Volume     func(*VolumeComponentParentOverride) error
}

var projectSourceParentOverride reflect.Type = reflect.TypeOf(ProjectSourceParentOverrideVisitor{})

func (union ProjectSourceParentOverride) Visit(visitor ProjectSourceParentOverrideVisitor) error {
	return visitUnion(union, visitor)
}
func (union *ProjectSourceParentOverride) discriminator() *string {
	return (*string)(&union.SourceType)
}
func (union *ProjectSourceParentOverride) Normalize() error {
	return normalizeUnion(union, projectSourceParentOverride)
}
func (union *ProjectSourceParentOverride) Simplify() {
	simplifyUnion(union, projectSourceParentOverride)
}

// +k8s:deepcopy-gen=false
type ProjectSourceParentOverrideVisitor struct {
	Git    func(*GitProjectSourceParentOverride) error
	Github func(*GithubProjectSourceParentOverride) error
	Zip    func(*ZipProjectSourceParentOverride) error
}

var commandUnionParentOverride reflect.Type = reflect.TypeOf(CommandUnionParentOverrideVisitor{})

func (union CommandUnionParentOverride) Visit(visitor CommandUnionParentOverrideVisitor) error {
	return visitUnion(union, visitor)
}
func (union *CommandUnionParentOverride) discriminator() *string {
	return (*string)(&union.CommandType)
}
func (union *CommandUnionParentOverride) Normalize() error {
	return normalizeUnion(union, commandUnionParentOverride)
}
func (union *CommandUnionParentOverride) Simplify() {
	simplifyUnion(union, commandUnionParentOverride)
}

// +k8s:deepcopy-gen=false
type CommandUnionParentOverrideVisitor struct {
	Exec         func(*ExecCommandParentOverride) error
	Apply        func(*ApplyCommandParentOverride) error
	VscodeTask   func(*VscodeConfigurationCommandParentOverride) error
	VscodeLaunch func(*VscodeConfigurationCommandParentOverride) error
	Composite    func(*CompositeCommandParentOverride) error
}

var vscodeConfigurationCommandLocationParentOverride reflect.Type = reflect.TypeOf(VscodeConfigurationCommandLocationParentOverrideVisitor{})

func (union VscodeConfigurationCommandLocationParentOverride) Visit(visitor VscodeConfigurationCommandLocationParentOverrideVisitor) error {
	return visitUnion(union, visitor)
}
func (union *VscodeConfigurationCommandLocationParentOverride) discriminator() *string {
	return (*string)(&union.LocationType)
}
func (union *VscodeConfigurationCommandLocationParentOverride) Normalize() error {
	return normalizeUnion(union, vscodeConfigurationCommandLocationParentOverride)
}
func (union *VscodeConfigurationCommandLocationParentOverride) Simplify() {
	simplifyUnion(union, vscodeConfigurationCommandLocationParentOverride)
}

// +k8s:deepcopy-gen=false
type VscodeConfigurationCommandLocationParentOverrideVisitor struct {
	Uri     func(string) error
	Inlined func(string) error
}

var k8sLikeComponentLocationParentOverride reflect.Type = reflect.TypeOf(K8sLikeComponentLocationParentOverrideVisitor{})

func (union K8sLikeComponentLocationParentOverride) Visit(visitor K8sLikeComponentLocationParentOverrideVisitor) error {
	return visitUnion(union, visitor)
}
func (union *K8sLikeComponentLocationParentOverride) discriminator() *string {
	return (*string)(&union.LocationType)
}
func (union *K8sLikeComponentLocationParentOverride) Normalize() error {
	return normalizeUnion(union, k8sLikeComponentLocationParentOverride)
}
func (union *K8sLikeComponentLocationParentOverride) Simplify() {
	simplifyUnion(union, k8sLikeComponentLocationParentOverride)
}

// +k8s:deepcopy-gen=false
type K8sLikeComponentLocationParentOverrideVisitor struct {
	Uri     func(string) error
	Inlined func(string) error
}

var importReferenceUnionParentOverride reflect.Type = reflect.TypeOf(ImportReferenceUnionParentOverrideVisitor{})

func (union ImportReferenceUnionParentOverride) Visit(visitor ImportReferenceUnionParentOverrideVisitor) error {
	return visitUnion(union, visitor)
}
func (union *ImportReferenceUnionParentOverride) discriminator() *string {
	return (*string)(&union.ImportReferenceType)
}
func (union *ImportReferenceUnionParentOverride) Normalize() error {
	return normalizeUnion(union, importReferenceUnionParentOverride)
}
func (union *ImportReferenceUnionParentOverride) Simplify() {
	simplifyUnion(union, importReferenceUnionParentOverride)
}

// +k8s:deepcopy-gen=false
type ImportReferenceUnionParentOverrideVisitor struct {
	Uri        func(string) error
	Id         func(string) error
	Kubernetes func(*KubernetesCustomResourceImportReferenceParentOverride) error
}

var componentUnionPluginOverrideParentOverride reflect.Type = reflect.TypeOf(ComponentUnionPluginOverrideParentOverrideVisitor{})

func (union ComponentUnionPluginOverrideParentOverride) Visit(visitor ComponentUnionPluginOverrideParentOverrideVisitor) error {
	return visitUnion(union, visitor)
}
func (union *ComponentUnionPluginOverrideParentOverride) discriminator() *string {
	return (*string)(&union.ComponentType)
}
func (union *ComponentUnionPluginOverrideParentOverride) Normalize() error {
	return normalizeUnion(union, componentUnionPluginOverrideParentOverride)
}
func (union *ComponentUnionPluginOverrideParentOverride) Simplify() {
	simplifyUnion(union, componentUnionPluginOverrideParentOverride)
}

// +k8s:deepcopy-gen=false
type ComponentUnionPluginOverrideParentOverrideVisitor struct {
	Container  func(*ContainerComponentPluginOverrideParentOverride) error
	Kubernetes func(*KubernetesComponentPluginOverrideParentOverride) error
	Openshift  func(*OpenshiftComponentPluginOverrideParentOverride) error
	Volume     func(*VolumeComponentPluginOverrideParentOverride) error
}

var commandUnionPluginOverrideParentOverride reflect.Type = reflect.TypeOf(CommandUnionPluginOverrideParentOverrideVisitor{})

func (union CommandUnionPluginOverrideParentOverride) Visit(visitor CommandUnionPluginOverrideParentOverrideVisitor) error {
	return visitUnion(union, visitor)
}
func (union *CommandUnionPluginOverrideParentOverride) discriminator() *string {
	return (*string)(&union.CommandType)
}
func (union *CommandUnionPluginOverrideParentOverride) Normalize() error {
	return normalizeUnion(union, commandUnionPluginOverrideParentOverride)
}
func (union *CommandUnionPluginOverrideParentOverride) Simplify() {
	simplifyUnion(union, commandUnionPluginOverrideParentOverride)
}

// +k8s:deepcopy-gen=false
type CommandUnionPluginOverrideParentOverrideVisitor struct {
	Exec         func(*ExecCommandPluginOverrideParentOverride) error
	Apply        func(*ApplyCommandPluginOverrideParentOverride) error
	VscodeTask   func(*VscodeConfigurationCommandPluginOverrideParentOverride) error
	VscodeLaunch func(*VscodeConfigurationCommandPluginOverrideParentOverride) error
	Composite    func(*CompositeCommandPluginOverrideParentOverride) error
}

var vscodeConfigurationCommandLocationPluginOverrideParentOverride reflect.Type = reflect.TypeOf(VscodeConfigurationCommandLocationPluginOverrideParentOverrideVisitor{})

func (union VscodeConfigurationCommandLocationPluginOverrideParentOverride) Visit(visitor VscodeConfigurationCommandLocationPluginOverrideParentOverrideVisitor) error {
	return visitUnion(union, visitor)
}
func (union *VscodeConfigurationCommandLocationPluginOverrideParentOverride) discriminator() *string {
	return (*string)(&union.LocationType)
}
func (union *VscodeConfigurationCommandLocationPluginOverrideParentOverride) Normalize() error {
	return normalizeUnion(union, vscodeConfigurationCommandLocationPluginOverrideParentOverride)
}
func (union *VscodeConfigurationCommandLocationPluginOverrideParentOverride) Simplify() {
	simplifyUnion(union, vscodeConfigurationCommandLocationPluginOverrideParentOverride)
}

// +k8s:deepcopy-gen=false
type VscodeConfigurationCommandLocationPluginOverrideParentOverrideVisitor struct {
	Uri     func(string) error
	Inlined func(string) error
}

var k8sLikeComponentLocationPluginOverrideParentOverride reflect.Type = reflect.TypeOf(K8sLikeComponentLocationPluginOverrideParentOverrideVisitor{})

func (union K8sLikeComponentLocationPluginOverrideParentOverride) Visit(visitor K8sLikeComponentLocationPluginOverrideParentOverrideVisitor) error {
	return visitUnion(union, visitor)
}
func (union *K8sLikeComponentLocationPluginOverrideParentOverride) discriminator() *string {
	return (*string)(&union.LocationType)
}
func (union *K8sLikeComponentLocationPluginOverrideParentOverride) Normalize() error {
	return normalizeUnion(union, k8sLikeComponentLocationPluginOverrideParentOverride)
}
func (union *K8sLikeComponentLocationPluginOverrideParentOverride) Simplify() {
	simplifyUnion(union, k8sLikeComponentLocationPluginOverrideParentOverride)
}

// +k8s:deepcopy-gen=false
type K8sLikeComponentLocationPluginOverrideParentOverrideVisitor struct {
	Uri     func(string) error
	Inlined func(string) error
}

var componentUnionPluginOverride reflect.Type = reflect.TypeOf(ComponentUnionPluginOverrideVisitor{})

func (union ComponentUnionPluginOverride) Visit(visitor ComponentUnionPluginOverrideVisitor) error {
	return visitUnion(union, visitor)
}
func (union *ComponentUnionPluginOverride) discriminator() *string {
	return (*string)(&union.ComponentType)
}
func (union *ComponentUnionPluginOverride) Normalize() error {
	return normalizeUnion(union, componentUnionPluginOverride)
}
func (union *ComponentUnionPluginOverride) Simplify() {
	simplifyUnion(union, componentUnionPluginOverride)
}

// +k8s:deepcopy-gen=false
type ComponentUnionPluginOverrideVisitor struct {
	Container  func(*ContainerComponentPluginOverride) error
	Kubernetes func(*KubernetesComponentPluginOverride) error
	Openshift  func(*OpenshiftComponentPluginOverride) error
	Volume     func(*VolumeComponentPluginOverride) error
}

var commandUnionPluginOverride reflect.Type = reflect.TypeOf(CommandUnionPluginOverrideVisitor{})

func (union CommandUnionPluginOverride) Visit(visitor CommandUnionPluginOverrideVisitor) error {
	return visitUnion(union, visitor)
}
func (union *CommandUnionPluginOverride) discriminator() *string {
	return (*string)(&union.CommandType)
}
func (union *CommandUnionPluginOverride) Normalize() error {
	return normalizeUnion(union, commandUnionPluginOverride)
}
func (union *CommandUnionPluginOverride) Simplify() {
	simplifyUnion(union, commandUnionPluginOverride)
}

// +k8s:deepcopy-gen=false
type CommandUnionPluginOverrideVisitor struct {
	Exec         func(*ExecCommandPluginOverride) error
	Apply        func(*ApplyCommandPluginOverride) error
	VscodeTask   func(*VscodeConfigurationCommandPluginOverride) error
	VscodeLaunch func(*VscodeConfigurationCommandPluginOverride) error
	Composite    func(*CompositeCommandPluginOverride) error
}

var vscodeConfigurationCommandLocationPluginOverride reflect.Type = reflect.TypeOf(VscodeConfigurationCommandLocationPluginOverrideVisitor{})

func (union VscodeConfigurationCommandLocationPluginOverride) Visit(visitor VscodeConfigurationCommandLocationPluginOverrideVisitor) error {
	return visitUnion(union, visitor)
}
func (union *VscodeConfigurationCommandLocationPluginOverride) discriminator() *string {
	return (*string)(&union.LocationType)
}
func (union *VscodeConfigurationCommandLocationPluginOverride) Normalize() error {
	return normalizeUnion(union, vscodeConfigurationCommandLocationPluginOverride)
}
func (union *VscodeConfigurationCommandLocationPluginOverride) Simplify() {
	simplifyUnion(union, vscodeConfigurationCommandLocationPluginOverride)
}

// +k8s:deepcopy-gen=false
type VscodeConfigurationCommandLocationPluginOverrideVisitor struct {
	Uri     func(string) error
	Inlined func(string) error
}

var k8sLikeComponentLocationPluginOverride reflect.Type = reflect.TypeOf(K8sLikeComponentLocationPluginOverrideVisitor{})

func (union K8sLikeComponentLocationPluginOverride) Visit(visitor K8sLikeComponentLocationPluginOverrideVisitor) error {
	return visitUnion(union, visitor)
}
func (union *K8sLikeComponentLocationPluginOverride) discriminator() *string {
	return (*string)(&union.LocationType)
}
func (union *K8sLikeComponentLocationPluginOverride) Normalize() error {
	return normalizeUnion(union, k8sLikeComponentLocationPluginOverride)
}
func (union *K8sLikeComponentLocationPluginOverride) Simplify() {
	simplifyUnion(union, k8sLikeComponentLocationPluginOverride)
}

// +k8s:deepcopy-gen=false
type K8sLikeComponentLocationPluginOverrideVisitor struct {
	Uri     func(string) error
	Inlined func(string) error
}
