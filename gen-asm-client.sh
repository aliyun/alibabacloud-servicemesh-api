client-gen --clientset-name clientset -h ./hack/boilerplate.go.txt --input-base istio.io/api --input-dirs alibabacloudservicemesh/v1 --input alibabacloudservicemesh/v1 --output-base ../alibabacloud-servicemesh-go-client --output-package asm/pkg

lister-gen -h ./hack/boilerplate.go.txt --input-dirs ./alibabacloudservicemesh/v1 --output-base ../alibabacloud-servicemesh-go-client --output-package asm/pkg/listers

informer-gen --listers-package asm/pkg/listers -h ./hack/boilerplate.go.txt --input-dirs ./alibabacloudservicemesh/v1 --output-base ../alibabacloud-servicemesh-go-client --output-package asm/pkg/informers --versioned-clientset-package  asm/pkg/clientset

find ../alibabacloudservicemesh-go-client/asm/pkg -type f -exec sed -i '' 's#\./alibabacloudservicemesh/#istio.io/api/alibabacloudservicemesh/#g' {} +
find ../alibabacloud-servicemesh-go-client/asm/pkg -type f -exec sed -i '' 's#asm/pkg/#istio.io/client-go/asm/pkg/#g' {} +