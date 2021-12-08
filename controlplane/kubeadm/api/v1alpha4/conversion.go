/*
Copyright 2020 The Kubernetes Authors.

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

package v1alpha4

import (
	"sigs.k8s.io/controller-runtime/pkg/conversion"

	"sigs.k8s.io/cluster-api/controlplane/kubeadm/api/v1beta1"
	utilconversion "sigs.k8s.io/cluster-api/util/conversion"
)

func (src *KubeadmControlPlane) ConvertTo(destRaw conversion.Hub) error {
	dest := destRaw.(*v1beta1.KubeadmControlPlane)

	if err := Convert_v1alpha4_KubeadmControlPlane_To_v1beta1_KubeadmControlPlane(src, dest, nil); err != nil {
		return err
	}

	// Manually restore data.
	restored := &v1beta1.KubeadmControlPlane{}
	if ok, err := utilconversion.UnmarshalData(src, restored); err != nil || !ok {
		return err
	}

	dest.Spec.KubeadmConfigSpec.Ignition = restored.Spec.KubeadmConfigSpec.Ignition

	return nil
}

func (dest *KubeadmControlPlane) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.KubeadmControlPlane)

	if err := Convert_v1beta1_KubeadmControlPlane_To_v1alpha4_KubeadmControlPlane(src, dest, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion except for metadata
	return utilconversion.MarshalData(src, dest)
}

func (src *KubeadmControlPlaneList) ConvertTo(destRaw conversion.Hub) error {
	dest := destRaw.(*v1beta1.KubeadmControlPlaneList)

	return Convert_v1alpha4_KubeadmControlPlaneList_To_v1beta1_KubeadmControlPlaneList(src, dest, nil)
}

func (dest *KubeadmControlPlaneList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.KubeadmControlPlaneList)

	return Convert_v1beta1_KubeadmControlPlaneList_To_v1alpha4_KubeadmControlPlaneList(src, dest, nil)
}

func (src *KubeadmControlPlaneTemplate) ConvertTo(destRaw conversion.Hub) error {
	dest := destRaw.(*v1beta1.KubeadmControlPlaneTemplate)

	if err := Convert_v1alpha4_KubeadmControlPlaneTemplate_To_v1beta1_KubeadmControlPlaneTemplate(src, dest, nil); err != nil {
		return err
	}

	// Manually restore data.
	restored := &v1beta1.KubeadmControlPlaneTemplate{}
	if ok, err := utilconversion.UnmarshalData(src, restored); err != nil || !ok {
		return err
	}

	dest.Spec.Template.Spec.KubeadmConfigSpec.Ignition = restored.Spec.Template.Spec.KubeadmConfigSpec.Ignition

	return nil
}

func (dest *KubeadmControlPlaneTemplate) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.KubeadmControlPlaneTemplate)

	if err := Convert_v1beta1_KubeadmControlPlaneTemplate_To_v1alpha4_KubeadmControlPlaneTemplate(src, dest, nil); err != nil {
		return err
	}

	// Preserve Hub data on down-conversion except for metadata
	return utilconversion.MarshalData(src, dest)
}

func (src *KubeadmControlPlaneTemplateList) ConvertTo(destRaw conversion.Hub) error {
	dest := destRaw.(*v1beta1.KubeadmControlPlaneTemplateList)

	return Convert_v1alpha4_KubeadmControlPlaneTemplateList_To_v1beta1_KubeadmControlPlaneTemplateList(src, dest, nil)
}

func (dest *KubeadmControlPlaneTemplateList) ConvertFrom(srcRaw conversion.Hub) error {
	src := srcRaw.(*v1beta1.KubeadmControlPlaneTemplateList)

	return Convert_v1beta1_KubeadmControlPlaneTemplateList_To_v1alpha4_KubeadmControlPlaneTemplateList(src, dest, nil)
}
