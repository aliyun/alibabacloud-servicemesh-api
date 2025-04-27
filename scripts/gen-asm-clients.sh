#!/bin/bash

set -x
set -e

OUTPUT_BASE="${OUTPUT_BASE:-generated-clients}"

INPUT="alibabacloudservicemesh/v1,alibabacloudservicemesh/v1beta1"

INPUTS="./alibabacloudservicemesh/v1,./alibabacloudservicemesh/v1beta1"

bin/client-gen --clientset-name versioned -h ./hack/boilerplate.go.txt --input-base "istio.io/api" --input-dirs $INPUTS --input $INPUT --output-base $OUTPUT_BASE --output-package asm/pkg/clientset

bin/lister-gen -h ./hack/boilerplate.go.txt --input-dirs $INPUTS --output-base $OUTPUT_BASE --output-package asm/pkg/listers

bin/informer-gen --listers-package asm/pkg/listers -h ./hack/boilerplate.go.txt --input-dirs $INPUTS --output-base $OUTPUT_BASE --output-package asm/pkg/informers --versioned-clientset-package asm/pkg/clientset/versioned

find $OUTPUT_BASE/asm/pkg -type f -exec sed -i 's#\./alibabacloudservicemesh/#istio.io/api/alibabacloudservicemesh/#g' {} +
find $OUTPUT_BASE/asm/pkg -type f -exec sed -i 's#asm/pkg/#istio.io/client-go/asm/pkg/#g' {} +

set +e
set +x