/*
Copyright 2016 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-generated-swagger-docs.sh

// AUTO-GENERATED FUNCTIONS START HERE
var map_StorageClass = map[string]string{
	"":                     "StorageClass describes the parameters for a class of storage for which PersistentVolumes can be dynamically provisioned.\n\nStorageClasses are non-namespaced; the name of the storage class according to etcd is in ObjectMeta.Name.",
	"metadata":             "Standard object's metadata. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
	"provisioner":          "Provisioner indicates the type of the provisioner.",
	"parameters":           "Parameters holds the parameters for the provisioner that should create volumes of this storage class.",
	"reclaimPolicy":        "Dynamically provisioned PersistentVolumes of this storage class are created with this reclaimPolicy. Defaults to Delete.",
	"mountOptions":         "Dynamically provisioned PersistentVolumes of this storage class are created with these mountOptions, e.g. [\"ro\", \"soft\"]. Not validated - mount of the PVs will simply fail if one is invalid.",
	"allowVolumeExpansion": "AllowVolumeExpansion shows whether the storage class allow volume expand",
<<<<<<< HEAD
	"volumeBindingMode":    "VolumeBindingMode indicates how PersistentVolumeClaims should be provisioned and bound.  When unset, VolumeBindingImmediate is used. This field is alpha-level and is only honored by servers that enable the VolumeScheduling feature.",
=======
>>>>>>> Initial dep workover
}

func (StorageClass) SwaggerDoc() map[string]string {
	return map_StorageClass
}

var map_StorageClassList = map[string]string{
	"":         "StorageClassList is a collection of storage classes.",
	"metadata": "Standard list metadata More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#metadata",
	"items":    "Items is the list of StorageClasses",
}

func (StorageClassList) SwaggerDoc() map[string]string {
	return map_StorageClassList
}

// AUTO-GENERATED FUNCTIONS END HERE