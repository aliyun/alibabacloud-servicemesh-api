#!/bin/bash

set -x
# set -e

OUTPUT_BASE="${OUTPUT_BASE:-generated-clients}"

INPUT="alibabacloudservicemesh/v1,alibabacloudservicemesh/v1beta1,ampere/v1"

INPUTS="./alibabacloudservicemesh/v1,./alibabacloudservicemesh/v1beta1,./ampere/v1"

bin/client-gen --clientset-name versioned \
    --go-header-file ./hack/boilerplate.go.txt \
    --output-base $OUTPUT_BASE  \
    --input-base "istio.io/api" \
    --input $INPUT \
    --output-package asm/pkg/clientset \

bin/lister-gen --go-header-file ./hack/boilerplate.go.txt \
    --input-dirs $INPUTS \
    --output-base $OUTPUT_BASE \
    --output-package asm/pkg/listers

bin/informer-gen --listers-package asm/pkg/listers \
    --go-header-file ./hack/boilerplate.go.txt \
    --input-dirs $INPUTS \
    --output-base $OUTPUT_BASE \
    --output-package asm/pkg/informers \
    --versioned-clientset-package \asm/pkg/clientset/versioned || true

find $OUTPUT_BASE/asm/pkg -type f -exec sed -i 's#\./alibabacloudservicemesh/#istio.io/api/alibabacloudservicemesh/#g' {} +
find $OUTPUT_BASE/asm/pkg -type f -exec sed -i 's#\./ampere/#istio.io/api/ampere/#g' {} +
find $OUTPUT_BASE/asm/pkg -type f -exec sed -i 's#asm/pkg/#istio.io/client-go/asm/pkg/#g' {} +

set +e
set +x