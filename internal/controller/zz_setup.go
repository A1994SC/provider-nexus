// SPDX-FileCopyrightText: 2023 The Crossplane Authors <https://crossplane.io>
//
// SPDX-License-Identifier: Apache-2.0

package controller

import (
	ctrl "sigs.k8s.io/controller-runtime"

	"github.com/crossplane/upjet/pkg/controller"

	azure "github.com/a1994sc/provider-nexus/internal/controller/blobstore/azure"
	file "github.com/a1994sc/provider-nexus/internal/controller/blobstore/file"
	group "github.com/a1994sc/provider-nexus/internal/controller/blobstore/group"
	s3 "github.com/a1994sc/provider-nexus/internal/controller/blobstore/s3"
	providerconfig "github.com/a1994sc/provider-nexus/internal/controller/providerconfig"
	yumproxy "github.com/a1994sc/provider-nexus/internal/controller/repository/yumproxy"
)

// Setup creates all controllers with the supplied logger and adds them to
// the supplied manager.
func Setup(mgr ctrl.Manager, o controller.Options) error {
	for _, setup := range []func(ctrl.Manager, controller.Options) error{
		azure.Setup,
		file.Setup,
		group.Setup,
		s3.Setup,
		providerconfig.Setup,
		yumproxy.Setup,
	} {
		if err := setup(mgr, o); err != nil {
			return err
		}
	}
	return nil
}
