# List
go mod init client-example
go mod tidy

# Dynamic List

## operator-sdk
operator-sdk create api --group cache --version v1 --kind Redis --resource --controller
operator-sdk create webhook --group cache --version v1 --kind Redis --defaulting --programmatic-validation

## Kubebuilder
kubebuilder init --domain ha-ving.store --repo ha-ving.store/Redis
```
├── Dockerfile
├── Makefile
├── PROJECT
├── README.md
├── cmd
│   └── main.go
├── config
│   ├── default
│   │   ├── kustomization.yaml
│   │   ├── manager_metrics_patch.yaml
│   │   └── metrics_service.yaml
│   ├── manager
│   │   ├── kustomization.yaml
│   │   └── manager.yaml
│   ├── prometheus
│   │   ├── kustomization.yaml
│   │   └── monitor.yaml
│   └── rbac
│       ├── kustomization.yaml
│       ├── leader_election_role.yaml
│       ├── leader_election_role_binding.yaml
│       ├── role.yaml
│       ├── role_binding.yaml
│       └── service_account.yaml
├── go.mod
├── go.sum
├── hack
│   └── boilerplate.go.txt
└── test
    ├── e2e
    │   ├── e2e_suite_test.go
    │   └── e2e_test.go
    └── utils
        └── utils.go

```
kubebuilder create api --group cafe24 --version v1 --kind Redis

`--group` : 자원이 속한 그룹
`--kind` : 자원 이름
`--version` : 자원 버전 (stable: v1)

```
INFO Create Resource [y/n]                        
y
INFO Create Controller [y/n]                      
y
INFO Writing kustomize manifests for you to edit... 
INFO Writing scaffold for you to edit...          
INFO api/v1/redis_types.go                        
INFO api/v1/groupversion_info.go                  
INFO internal/controller/suite_test.go            
INFO internal/controller/redis_controller.go      
INFO internal/controller/redis_controller_test.go 
INFO Update dependencies:
$ go mod tidy           
INFO Running make:
$ make generate                
mkdir -p /home/workspace/example/bin
Downloading sigs.k8s.io/controller-tools/cmd/controller-gen@v0.15.0
/home/workspace/example/bin/controller-gen-v0.15.0 object:headerFile="hack/boilerplate.go.txt" paths="./..."
```

```
├── Dockerfile
├── Makefile
├── PROJECT
├── README.md
├── api
│   └── v1
│       ├── groupversion_info.go
│       ├── redis_types.go
│       └── zz_generated.deepcopy.go
├── bin
│   └── controller-gen-v0.15.0
├── cmd
│   └── main.go
├── config
│   ├── crd
│   │   ├── kustomization.yaml
│   │   └── kustomizeconfig.yaml
│   ├── default
│   │   ├── kustomization.yaml
│   │   ├── manager_metrics_patch.yaml
│   │   └── metrics_service.yaml
│   ├── manager
│   │   ├── kustomization.yaml
│   │   └── manager.yaml
│   ├── prometheus
│   │   ├── kustomization.yaml
│   │   └── monitor.yaml
│   ├── rbac
│   │   ├── kustomization.yaml
│   │   ├── leader_election_role.yaml
│   │   ├── leader_election_role_binding.yaml
│   │   ├── redis_editor_role.yaml
│   │   ├── redis_viewer_role.yaml
│   │   ├── role.yaml
│   │   ├── role_binding.yaml
│   │   └── service_account.yaml
│   └── samples
│       ├── cafe24_v1_redis.yaml
│       └── kustomization.yaml
├── go.mod
├── go.sum
├── hack
│   └── boilerplate.go.txt
├── internal
│   └── controller
│       ├── redis_controller.go
│       ├── redis_controller_test.go
│       └── suite_test.go
└── test
    ├── e2e
    │   ├── e2e_suite_test.go
    │   └── e2e_test.go
    └── utils
        └── utils.go
```
# 디렉토리
- `crd` : CRD 메니페스트 파일

make manifests
```/home/workspace/example/bin/controller-gen-v0.15.0 rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases```
```
├── Dockerfile
├── Makefile
├── PROJECT
├── README.md
├── api
│   └── v1
│       ├── groupversion_info.go
│       ├── redis_types.go
│       └── zz_generated.deepcopy.go
├── bin
│   └── controller-gen-v0.15.0
├── cmd
│   └── main.go
├── config
│   ├── crd
│   │   ├── bases
│   │   │   └── cafe24.ha-ving.store_redis.yaml     // 새로 생김.
│   │   ├── kustomization.yaml
│   │   └── kustomizeconfig.yaml
│   ├── default
│   │   ├── kustomization.yaml
│   │   ├── manager_metrics_patch.yaml
│   │   └── metrics_service.yaml
│   ├── manager
│   │   ├── kustomization.yaml
│   │   └── manager.yaml
│   ├── prometheus
│   │   ├── kustomization.yaml
│   │   └── monitor.yaml
│   ├── rbac
│   │   ├── kustomization.yaml
│   │   ├── leader_election_role.yaml
│   │   ├── leader_election_role_binding.yaml
│   │   ├── redis_editor_role.yaml
│   │   ├── redis_viewer_role.yaml
│   │   ├── role.yaml
│   │   ├── role_binding.yaml
│   │   └── service_account.yaml
│   └── samples
│       ├── cafe24_v1_redis.yaml                // 새로 생김.
│       └── kustomization.yaml
├── go.mod
├── go.sum
├── hack
│   └── boilerplate.go.txt
├── internal
│   └── controller
│       ├── redis_controller.go
│       ├── redis_controller_test.go
│       └── suite_test.go
└── test
    ├── e2e
    │   ├── e2e_suite_test.go
    │   └── e2e_test.go
    └── utils
        └── utils.go

```

kubebuilder create webhook --group cafe24 --version v1 --kind Redis --defaulting --programmatic-validation
```
INFO Writing kustomize manifests for you to edit...
INFO Writing scaffold for you to edit...          
INFO api/v1/redis_webhook.go                      
INFO api/v1/redis_webhook_test.go                 
INFO api/v1/webhook_suite_test.go                 
INFO Update dependencies:
$ go mod tidy           
INFO Running make:
$ make generate                
/home/workspace/example/bin/controller-gen-v0.15.0 object:headerFile="hack/boilerplate.go.txt" paths="./..."
Next: implement your new Webhook and generate the manifests with:
$ make manifests
```
make manifests
```/home/workspace/example/bin/controller-gen-v0.15.0 rbac:roleName=manager-role crd webhook paths="./..." output:crd:artifacts:config=config/crd/bases```

```
├── Dockerfile
├── Makefile
├── PROJECT
├── README.md
├── api
│   └── v1
│       ├── groupversion_info.go
│       ├── redis_types.go
│       ├── redis_webhook.go
│       ├── redis_webhook_test.go
│       ├── webhook_suite_test.go
│       └── zz_generated.deepcopy.go
├── bin
│   └── controller-gen-v0.15.0
├── cmd
│   └── main.go
├── config
│   ├── certmanager
│   │   ├── certificate.yaml
│   │   ├── kustomization.yaml
│   │   └── kustomizeconfig.yaml
│   ├── crd
│   │   ├── bases
│   │   │   └── cafe24.ha-ving.store_redis.yaml
│   │   ├── kustomization.yaml
│   │   ├── kustomizeconfig.yaml
│   │   └── patches
│   │       ├── cainjection_in_redis.yaml
│   │       └── webhook_in_redis.yaml
│   ├── default
│   │   ├── kustomization.yaml
│   │   ├── manager_metrics_patch.yaml
│   │   ├── manager_webhook_patch.yaml
│   │   ├── metrics_service.yaml
│   │   └── webhookcainjection_patch.yaml
│   ├── manager
│   │   ├── kustomization.yaml
│   │   └── manager.yaml
│   ├── prometheus
│   │   ├── kustomization.yaml
│   │   └── monitor.yaml
│   ├── rbac
│   │   ├── kustomization.yaml
│   │   ├── leader_election_role.yaml
│   │   ├── leader_election_role_binding.yaml
│   │   ├── redis_editor_role.yaml
│   │   ├── redis_viewer_role.yaml
│   │   ├── role.yaml
│   │   ├── role_binding.yaml
│   │   └── service_account.yaml
│   ├── samples
│   │   ├── cafe24_v1_redis.yaml
│   │   └── kustomization.yaml
│   └── webhook
│       ├── kustomization.yaml
│       ├── kustomizeconfig.yaml
│       ├── manifests.yaml
│       └── service.yaml
├── go.mod
├── go.sum
├── hack
│   └── boilerplate.go.txt
├── internal
│   └── controller
│       ├── redis_controller.go
│       ├── redis_controller_test.go
│       └── suite_test.go
└── test
    ├── e2e
    │   ├── e2e_suite_test.go
    │   └── e2e_test.go
    └── utils
        └── utils.go

```

### Makefile
```make help```

```
Usage:
  make <target>

General
  help             Display this help.

Development
  manifests        Generate WebhookConfiguration, ClusterRole and CustomResourceDefinition objects.
  generate         Generate code containing DeepCopy, DeepCopyInto, and DeepCopyObject method implementations.
  fmt              Run go fmt against code.
  vet              Run go vet against code.
  test             Run tests.
  lint             Run golangci-lint linter
  lint-fix         Run golangci-lint linter and perform fixes

Build
  build            Build manager binary.
  run              Run a controller from your host.
  docker-build     Build docker image with the manager.
  docker-push      Push docker image with the manager.
  docker-buildx    Build and push docker image for the manager for cross-platform support
  build-installer  Generate a consolidated YAML with CRDs and deployment.

Deployment
  install          Install CRDs into the K8s cluster specified in ~/.kube/config.
  uninstall        Uninstall CRDs from the K8s cluster specified in ~/.kube/config. Call with ignore-not-found=true to ignore resource not found errors during deletion.
  deploy           Deploy controller to the K8s cluster specified in ~/.kube/config.
  undeploy         Undeploy controller from the K8s cluster specified in ~/.kube/config. Call with ignore-not-found=true to ignore resource not found errors during deletion.

Dependencies
  kustomize        Download kustomize locally if necessary.
  controller-gen   Download controller-gen locally if necessary.
  envtest          Download setup-envtest locally if necessary.
  golangci-lint    Download golangci-lint locally if necessary.
```

### 주석
```
//+kubebuilder:scaffold:imports
//+kubebuilder:scaffold:scheme
//+kubebuilder:scaffold:builder
```
### Prometheus 
`Prometheus Operator` 사용하는 경우, 자동으로 Controller Metric 수집한다.

### RBAC
`kube-auth-proxy` 사용하면, metric 엔드포인트를 RBAC로 제한.
`leader_election_role.yaml`, `leader_election_role_binding.yaml` 선거기능에 필요한 권한 설정
`role.yaml`, `role_binding.yaml` 컨트롤러가 리소스에 액세스 할 수 있는 권한 설정
ps. 필요 없으면, 파일을 삭제하면 된다. `kustomization.yaml`에서 제외 해야한다.

### Sample
- Sample Manifest

### Webhook ( Admission Webhook )
특정 리소스를 만들고, 업데이트 할 때 webhook API 호출 -> 유효성 검사 -> 리소스 업데이트

`--programmatic-validation` : 리소스 검증
`--defaulting` : 리소스 필드에 기본값 설정
`--conversion` : 사용자 정의리소스 버전을 업그레이드 할 때 리소스를 변환 하기 위한 webhook?

`api/v1/redis_webhook.go` : 이 파일에 webhook 구현

### `kustomization.yaml` 편집
make manifest 이후
`config/webhook/manifest.yaml` 생성된다.
`config/default/kustomization.yaml` 을 편집해야한다.


make docker-build

make install 
make deploy

## webhook
- `path` : 
- `mutating` : true (유효성 검사) | false
- `failurePolicy` : 


## AdmissionReview

Admission Controller : 허락하는 컨트롤러
Mutate(변조), Validate(검증) 허가 여부



Ref. 
- https://malwareanalysis.tistory.com/704
- https://medium.com/cloud-native-daily/working-with-kubernetes-using-golang-a3069d51dfd6 
- https://github.com/kubernetes/client-go/blob/master/examples/dynamic-create-update-delete-deployment/main.go
- https://medium.com/@souravchoudhary0306/designing-a-dynamic-rate-limiter-in-go-188d57e3526f
