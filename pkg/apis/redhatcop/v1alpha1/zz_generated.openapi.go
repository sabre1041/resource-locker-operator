// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/redhat-cop/resource-locker-operator/pkg/apis/redhatcop/v1alpha1.LockingStatus":        schema_pkg_apis_redhatcop_v1alpha1_LockingStatus(ref),
		"github.com/redhat-cop/resource-locker-operator/pkg/apis/redhatcop/v1alpha1.Patch":                schema_pkg_apis_redhatcop_v1alpha1_Patch(ref),
		"github.com/redhat-cop/resource-locker-operator/pkg/apis/redhatcop/v1alpha1.Resource":             schema_pkg_apis_redhatcop_v1alpha1_Resource(ref),
		"github.com/redhat-cop/resource-locker-operator/pkg/apis/redhatcop/v1alpha1.ResourceLocker":       schema_pkg_apis_redhatcop_v1alpha1_ResourceLocker(ref),
		"github.com/redhat-cop/resource-locker-operator/pkg/apis/redhatcop/v1alpha1.ResourceLockerSpec":   schema_pkg_apis_redhatcop_v1alpha1_ResourceLockerSpec(ref),
		"github.com/redhat-cop/resource-locker-operator/pkg/apis/redhatcop/v1alpha1.ResourceLockerStatus": schema_pkg_apis_redhatcop_v1alpha1_ResourceLockerStatus(ref),
	}
}

func schema_pkg_apis_redhatcop_v1alpha1_LockingStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"objectReference": {
						SchemaProps: spec.SchemaProps{
							Description: "the name of the locked configuration",
							Ref:         ref("k8s.io/api/core/v1.ObjectReference"),
						},
					},
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Type of deployment condition.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status of the condition, one of True, False, Unknown.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"lastUpdateTime": {
						SchemaProps: spec.SchemaProps{
							Description: "The last time this condition was updated.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "A human readable message indicating details about the transition.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"objectReference", "status"},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.ObjectReference", "k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}

func schema_pkg_apis_redhatcop_v1alpha1_Patch(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Patch describe a pacth to be enforced at runtim",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"sourceObjectRefs": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "SourceObject refs is an arrays of refereces to source objects that will be used as input for the template processing",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("k8s.io/api/core/v1.ObjectReference"),
									},
								},
							},
						},
					},
					"targetObjectRef": {
						SchemaProps: spec.SchemaProps{
							Description: "TargetObjectRef is a reference to the object to which the pacth should be applied.",
							Ref:         ref("k8s.io/api/core/v1.ObjectReference"),
						},
					},
					"patchType": {
						SchemaProps: spec.SchemaProps{
							Description: "PatchType is the type of patch to be applied, one of \"application/json-patch+json\"'\"application/merge-patch+json\",\"application/strategic-merge-patch+json\",\"application/apply-patch+yaml\"",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"patchTemplate": {
						SchemaProps: spec.SchemaProps{
							Description: "PatchTemplate is a go template that will be resolved using the SourceObjectRefs as parameters. The result must be a valid patch based on the pacth type and the target object.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"targetObjectRef", "patchTemplate"},
			},
		},
		Dependencies: []string{
			"k8s.io/api/core/v1.ObjectReference"},
	}
}

func schema_pkg_apis_redhatcop_v1alpha1_Resource(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Type: []string{"object"},
				Properties: map[string]spec.Schema{
					"object": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/runtime.RawExtension"),
						},
					},
					"excludedPaths": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Type:   []string{"string"},
										Format: "",
									},
								},
							},
						},
					},
				},
				Required: []string{"object"},
			},
		},
		Dependencies: []string{
			"k8s.io/apimachinery/pkg/runtime.RawExtension"},
	}
}

func schema_pkg_apis_redhatcop_v1alpha1_ResourceLocker(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ResourceLocker is the Schema for the resourcelockers API",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/redhat-cop/resource-locker-operator/pkg/apis/redhatcop/v1alpha1.ResourceLockerSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/redhat-cop/resource-locker-operator/pkg/apis/redhatcop/v1alpha1.ResourceLockerStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/redhat-cop/resource-locker-operator/pkg/apis/redhatcop/v1alpha1.ResourceLockerSpec", "github.com/redhat-cop/resource-locker-operator/pkg/apis/redhatcop/v1alpha1.ResourceLockerStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_redhatcop_v1alpha1_ResourceLockerSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ResourceLockerSpec defines the desired state of ResourceLocker",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"resources": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Resources is a list of resource manifests that should be locked into the specified configuration",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/redhat-cop/resource-locker-operator/pkg/apis/redhatcop/v1alpha1.Resource"),
									},
								},
							},
						},
					},
					"patches": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Patches is a list of pacthes that should be encforced at runtime.",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/redhat-cop/resource-locker-operator/pkg/apis/redhatcop/v1alpha1.Patch"),
									},
								},
							},
						},
					},
					"serviceAccountRef": {
						SchemaProps: spec.SchemaProps{
							Description: "ServiceAccountRef is the service account to be used to run the controllers associated with this configuration",
							Ref:         ref("k8s.io/api/core/v1.LocalObjectReference"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/redhat-cop/resource-locker-operator/pkg/apis/redhatcop/v1alpha1.Patch", "github.com/redhat-cop/resource-locker-operator/pkg/apis/redhatcop/v1alpha1.Resource", "k8s.io/api/core/v1.LocalObjectReference"},
	}
}

func schema_pkg_apis_redhatcop_v1alpha1_ResourceLockerStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "ResourceLockerStatus defines the observed state of ResourceLocker",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"type": {
						SchemaProps: spec.SchemaProps{
							Description: "Type of deployment condition.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Description: "Status of the condition, one of True, False, Unknown.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"lastUpdateTime": {
						SchemaProps: spec.SchemaProps{
							Description: "The last time this condition was updated.",
							Ref:         ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"message": {
						SchemaProps: spec.SchemaProps{
							Description: "A human readable message indicating details about the transition.",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"resourceStatuses": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/redhat-cop/resource-locker-operator/pkg/apis/redhatcop/v1alpha1.LockingStatus"),
									},
								},
							},
						},
					},
					"patchStatuses": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "atomic",
							},
						},
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/redhat-cop/resource-locker-operator/pkg/apis/redhatcop/v1alpha1.LockingStatus"),
									},
								},
							},
						},
					},
				},
				Required: []string{"type", "status"},
			},
		},
		Dependencies: []string{
			"github.com/redhat-cop/resource-locker-operator/pkg/apis/redhatcop/v1alpha1.LockingStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}
