// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.Collection":       schema_pkg_apis_kabanero_v1alpha1_Collection(ref),
		"github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.CollectionSpec":   schema_pkg_apis_kabanero_v1alpha1_CollectionSpec(ref),
		"github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.CollectionStatus": schema_pkg_apis_kabanero_v1alpha1_CollectionStatus(ref),
		"github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.Kabanero":         schema_pkg_apis_kabanero_v1alpha1_Kabanero(ref),
		"github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.KabaneroSpec":     schema_pkg_apis_kabanero_v1alpha1_KabaneroSpec(ref),
		"github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.KabaneroStatus":   schema_pkg_apis_kabanero_v1alpha1_KabaneroStatus(ref),
	}
}

func schema_pkg_apis_kabanero_v1alpha1_Collection(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Collection is the Schema for the collections API",
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
							Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.CollectionSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.CollectionStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.CollectionSpec", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.CollectionStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_kabanero_v1alpha1_CollectionSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CollectionSpec defines the desired state of Collection",
				Properties: map[string]spec.Schema{
					"repositoryUrl": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"name": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"version": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"desiredState": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
				},
			},
		},
		Dependencies: []string{},
	}
}

func schema_pkg_apis_kabanero_v1alpha1_CollectionStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "CollectionStatus defines the observed state of Collection",
				Properties: map[string]spec.Schema{
					"activeVersion": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"activeLocation": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"activeDigest": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"activeAssets": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.RepositoryAssetStatus"),
									},
								},
							},
						},
					},
					"availableVersion": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"availableLocation": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"statusMessage": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"images": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.Image"),
									},
								},
							},
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.Image", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.RepositoryAssetStatus"},
	}
}

func schema_pkg_apis_kabanero_v1alpha1_Kabanero(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Kabanero is the Schema for the kabaneros API",
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
							Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.KabaneroSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.KabaneroStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.KabaneroSpec", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.KabaneroStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_kabanero_v1alpha1_KabaneroSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KabaneroSpec defines the desired state of Kabanero",
				Properties: map[string]spec.Schema{
					"version": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"string"},
							Format: "",
						},
					},
					"targetNamespaces": {
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
					"github": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.GithubConfig"),
						},
					},
					"collections": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.InstanceCollectionConfig"),
						},
					},
					"tekton": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.TektonCustomizationSpec"),
						},
					},
					"cliServices": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.KabaneroCliServicesCustomizationSpec"),
						},
					},
					"landing": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.KabaneroLandingCustomizationSpec"),
						},
					},
					"che": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.CheCustomizationSpec"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.CheCustomizationSpec", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.GithubConfig", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.InstanceCollectionConfig", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.KabaneroCliServicesCustomizationSpec", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.KabaneroLandingCustomizationSpec", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.TektonCustomizationSpec"},
	}
}

func schema_pkg_apis_kabanero_v1alpha1_KabaneroStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "KabaneroStatus defines the observed state of the Kabanero instance.",
				Properties: map[string]spec.Schema{
					"kabaneroInstance": {
						SchemaProps: spec.SchemaProps{
							Description: "Kabanero operator instance readiness status. The status is directly correlated to the availability of resources dependencies.",
							Ref:         ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.KabaneroInstanceStatus"),
						},
					},
					"knativeEventing": {
						SchemaProps: spec.SchemaProps{
							Description: "Knative eventing instance readiness status.",
							Ref:         ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.KnativeEventingStatus"),
						},
					},
					"serverless": {
						SchemaProps: spec.SchemaProps{
							Description: "OpenShift serverless operator status.",
							Ref:         ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.ServerlessStatus"),
						},
					},
					"tekton": {
						SchemaProps: spec.SchemaProps{
							Description: "Tekton instance readiness status.",
							Ref:         ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.TektonStatus"),
						},
					},
					"cli": {
						SchemaProps: spec.SchemaProps{
							Description: "CLI readiness status.",
							Ref:         ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.CliStatus"),
						},
					},
					"landing": {
						SchemaProps: spec.SchemaProps{
							Description: "Kabanero Landing page readiness status.",
							Ref:         ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.KabaneroLandingPageStatus"),
						},
					},
					"appsody": {
						SchemaProps: spec.SchemaProps{
							Description: "Appsody instance readiness status.",
							Ref:         ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.AppsodyStatus"),
						},
					},
					"kappnav": {
						SchemaProps: spec.SchemaProps{
							Description: "Kabanero Application Navigator instance readiness status.",
							Ref:         ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.KappnavStatus"),
						},
					},
					"che": {
						SchemaProps: spec.SchemaProps{
							Description: "Che instance readiness status.",
							Ref:         ref("github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.CheStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.AppsodyStatus", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.CheStatus", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.CliStatus", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.KabaneroInstanceStatus", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.KabaneroLandingPageStatus", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.KappnavStatus", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.KnativeEventingStatus", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.ServerlessStatus", "github.com/kabanero-io/kabanero-operator/pkg/apis/kabanero/v1alpha1.TektonStatus"},
	}
}
