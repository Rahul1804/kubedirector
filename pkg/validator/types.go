// Copyright 2019 Hewlett Packard Enterprise Development LP

// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at

//     http://www.apache.org/licenses/LICENSE-2.0

// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package validator

import (
	"github.com/bluek8s/kubedirector/pkg/shared"
	av1beta1 "k8s.io/api/admission/v1beta1"
)

// admitFunc is used as the type for all the callback validators
type admitFunc func(*av1beta1.AdmissionReview) *av1beta1.AdmissionResponse

type checkFunc func() error

const (
	validatorServiceName                  = "kubedirector-validator"
	validatorWebhook                      = "kubedirector-webhook"
	validatorSecret                       = "kubedirector-validator-secret"
	webhookHandlerName                    = "validate-cr.kubedirector.hpe.com"
	validationPort                        = 8443
	validationPath                        = "/validate"
	healthPath                            = "/healthz"
	defaultNativeSystemd                  = false
	defaultBackupClusterStatus            = false
	defaultAllowRestoreWithoutConnections = false

	appCrt  = "app.crt"
	appKey  = "app.pem"
	rootCrt = "ca.crt"

	allowDeleteLabel = shared.KdDomainBase + "/allow-delete-while-restoring"

	multipleSpecChange = "Change to spec not allowed before previous spec change has been processed."
	pendingNotifies    = "Change to spec not allowed because some members have not processed notifications of previous change."

	appInUse           = "KubeDirectorApp resource cannot be deleted or modified while referenced by the following KubeDirectorCluster resources: %s"
	invalidAppMessage  = "Invalid app(%s). This app resource ID has not been registered."
	invalidCardinality = "Invalid member count for role(%s). Specified member count:%d Role cardinality:%s"
	invalidRole        = "Invalid role(%s) in app(%s) specified. Valid roles: \"%s\""
	unconfiguredRole   = "Active role(%s) in app(%s) must have its configuration included in the roles array."

	modifiedProperty = "The %s property is read-only."
	modifiedRole     = "Role(%s) properties other than the members count cannot be modified while role members exist."

	invalidNodeRoleID     = "Invalid roleID(%s) in roleServices array in config section. Valid roles: \"%s\""
	invalidSelectedRoleID = "Invalid element(%s) in selectedRoles array in config section. Valid roles: \"%s\""
	invalidServiceID      = "Invalid service_id(%s) in roleServices array in config section. Valid services: \"%s\""

	nonUniqueRoleID       = "Each id in the roles array must be unique."
	nonUniqueServiceID    = "Each id in the services array must be unique."
	nonUniqueSelectedRole = "Each element of selectedRoles array in config section must be unique."
	nonUniqueServiceRole  = "Each roleID in roleServices array in config section must be unique."

	invalidDefaultSecretPrefix = "defaultSecret(%s) does not have the required name prefix(%s)."
	invalidDefaultSecret       = "Unable to find defaultSecret(%s) in namespace(%s)."
	invalidSecretPrefix        = "Secret(%s) for role(%s) does not have the required name prefix(%s)."
	invalidSecret              = "Unable to find secret(%s) for role(%s) in namespace(%s)."

	noDefaultImage  = "Role(%s) has no specified image, and no top-level default image is specified."
	ttyWithoutStdin = "Role(%s) requested TTY without STDIN."

	noURLScheme = "The endpoint for service(%s) must include a urlScheme value because isDashboard is true."

	failedToPatch = "Internal error: failed to populate default values for unspecified properties."

	failedToPatchPVC = "Internal error: failed to apply ownerReference to PVC for kdcluster."

	invalidStorageDef   = "Storage size for role (%s) is incorrectly defined."
	invalidStorageSize  = "Storage size for role (%s) should be greater than zero."
	invalidStorageClass = "Unable to fetch storageClass object with the provided name(%s)."

	invalidMinStorageDef = "Minimum storage size for role (%s) is incorrectly defined."

	invalidRoleStorageClass = "Unable to fetch storageClassName(%s) for role(%s)."
	noDefaultStorageClass   = "storageClassName is not specified for one or more roles, and no default storage class is available."
	badDefaultStorageClass  = "storageClassName is not specified for one or more roles, and default storage class (%s) is not available on the system."

	invalidShmemDef  = "Shared memory size for role (%s) is incorrectly defined."
	invalidShmemSize = "Shared memory size for role (%s) should be greater than zero."

	invalidResource = "Specified resource(\"%s\") value(\"%s\") for role(\"%s\") is invalid. Minimum value must be \"%s\"."
	invalidStorage  = "Specified persistent storage size(\"%s\") for role(\"%s\") is invalid. Minimum size must be \"%s\"."
	invalidSrcURL   = "Unable to access the specified URL(\"%s\") in file injection spec for the role (%s). error: %s."

	maxMemberLimit = "Maximum number of total members per KD cluster supported is %d."

	failedSecretKeyEncryption                       = "cannot encrypt secret key %s"
	forbiddenManualSecretKeyEncryptedValuePlacement = "manually setting secret key (%s) encrypted value is forbidden"
	invalidMasterEncryptionKey                      = "masterEncryptionKey is invalid. error: %s."
	masterEncryptionKeyChange                       = "masterEncryptionKey value cannot be changed while kdclusters exist"

	invalidConfigDelete = "kd-global-config cannot be deleted while kdclusters exist"

	invalidPVC        = "Unable to find persistentvolumeclaim(%s) in namespace(%s) as specified for role(%s)."
	invalidVolumeMode = "Specified persistentvolumeclaim(%s) for role (%s) is invalid. VolumeMode(%s) for the underlying volume must be configured as Filesystem."
	invalidAccessMode = "Specified persistentvolumeclaim(%s) is invalid. AccessModes for this volume must contain either ReadWriteMany or ReadOnlyMany, since its consumed by more than 1 member of the cluster."

	duplicateMountPath     = "Specified mountPath(%s) for role(%s) is invalid. It must be unique within the role."
	systemMountPathClash   = "Specified mountPath(%s) for role(%s) is invalid. It clashes with system generated mountPath."
	failedVolumeMountCheck = "Unexpected error while validating for unique volume mount paths for role(%s)."
)

type dictValue map[string]string
