package v1alpha1

import "github.com/parryjacob/k8s"

func init() {
	k8s.Register("storage.k8s.io", "v1alpha1", "volumeattachments", false, &VolumeAttachment{})

	k8s.RegisterList("storage.k8s.io", "v1alpha1", "volumeattachments", false, &VolumeAttachmentList{})
}
