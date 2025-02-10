#!/bin/bash

set -x
set -e

OUTPUT_BASE="${OUTPUT_BASE:-generated-clients}"

bin/client-gen --clientset-name clientset -h ./hack/boilerplate.go.txt --input-base istio.io/api --input-dirs alibabacloudservicemesh/v1 --input alibabacloudservicemesh/v1 --output-base $OUTPUT_BASE --output-package asm/pkg

bin/lister-gen -h ./hack/boilerplate.go.txt --input-dirs ./alibabacloudservicemesh/v1 --output-base $OUTPUT_BASE --output-package asm/pkg/listers

bin/informer-gen --listers-package asm/pkg/listers -h ./hack/boilerplate.go.txt --input-dirs ./alibabacloudservicemesh/v1 --output-base $OUTPUT_BASE --output-package asm/pkg/informers --versioned-clientset-package  asm/pkg/clientset

find $OUTPUT_BASE/asm/pkg -type f -exec sed -i 's#\./alibabacloudservicemesh/#istio.io/api/alibabacloudservicemesh/#g' {} +
find $OUTPUT_BASE/asm/pkg -type f -exec sed -i 's#asm/pkg/#istio.io/client-go/asm/pkg/#g' {} +

set +e
set +x