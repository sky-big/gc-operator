# CustomResourceDefinition资源详细说明

## 基本字段定义说明
```
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  # CRD名称必须是: <plural>.<group>格式
  name: crd.example.dev
spec:
  # 决定了REST API路径：/apis/<group>/<version>
  group: example.dev
  # 此CRD支持的版本
  versions:
    - name: v1
      # 每个版本都可以被禁用或启用
      served: true
      # 只有一个版本可以被标记为true，表示以此版本来存储
      storage: true
  # 可选Namespaced或Cluster，表示此资源是命名空间限定的，还是全局的
  scope: Namespaced
  names:
    # 决定了REST API路径： /apis/<group>/<version>/<plural>
    plural: crd
    # 在CLI中的别名
    singular: crontab
    # 驼峰式大小写的，正式的资源类型
    kind: CRD1
    # 在CLI中可以使用的短名称
    shortNames:
    - cr
```

## 声明kubernetes在kubectl获取该CRD资源显示额外的字段

```
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: crd.example.dev
spec:
  # 在资源名称后面，额外显示以下字段
  additionalPrinterColumns:
    - name: Status
      type: string
      JSONPath: .status.phase
```

## 子资源

自定义资源可以支持/status和/scale子资源。此特性在1.11版本中处于Beta状态且默认启用。你需要在CRD中进行定义才能启用这些子资源。

scale子资源支持让其他K8S组件（例如HorizontalPodAutoscaler和PodDisruptionBudget控制器）与你的CR进行交互。kubectl scale也可以利用该子资源对CR进行扩容。

status子资源可以让你把资源的规格和状态分开。

/status
启用状态子资源后，自定义资源的/status URL可用：

(1). 数据对应资源的.status字段
(2). PUT /status仅仅会修改.status字段，也仅仅对该字段进行验证
(3). 对资源本身进行PUT/POST/PATCH操作，会忽视.status字段
(4). 每次修改.spec字段，都导致.metadata.generation ++ 

/scale
启用扩容子资源后，自定义资源的/scale URL可用，RESTful载荷类型为autoscaling/v1.Scale

要启用扩容子资源，CRD需要指定：

(1). SpecReplicasPath，指定自定义资源中对应Scale.Spec.Replicas的JSON路径。必须值
(2). StatusReplicasPath，指定自定义资源中对应Scale.Status.Replicas的JSON路径。必须值
(3). LabelSelectorPath，指定自定义资源中对应Scale.Status.Selector的JSON路径。可选值，和HPA联用则必须设置

例如:
```
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
spec:
  subresources:
    # 启用状态子资源
    status: {}
    # 启用扩容子资源
    scale:
      specReplicasPath: .spec.replicas
      statusReplicasPath: .status.replicas
      labelSelectorPath: .status.labelSelector
```

## Categories用于指定资源所属的类别

```
通过kubectl get all可以访问到上述CRD的自定义资源

apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
spec:
  names:
    categories:
    - all
```

## 声明Kubernetes规定自定义资源的字段验证规则

```
apiVersion: apiextensions.k8s.io/v1beta1
kind: CustomResourceDefinition
metadata:
  name: crd.example.dev
spec:
  validation:
    openAPIV3Schema:
      # 根对象包含的属性
      properties:
        # spec属性规则
        spec:
          # 必须具有以下子属性
          required:
            - targetRef
            - service
            - canaryAnalysis
          # spec的子属性列表
          properties:
            # 简单属性，指定类型
            progressDeadlineSeconds:
              type: number
            # 对象属性
            targetRef:
              type: object
              required: ['apiVersion', 'kind', 'name']
              properties:
                apiVersion:
                  type: string
                kind:
                  type: string
                name:
                  type: string
            # 复杂属性：可以是字符串，也可以是对象
            autoscalerRef:
              anyOf:
                - type: string
                - type: object
              required: ['apiVersion', 'kind', 'name']
              properties:
                apiVersion:
                  type: string
                kind:
                  type: string
                name:
                  type: string
            service:
              type: object
              required: ['port']
              properties:
                port:
                  type: number
            skipAnalysis:
              type: boolean
            canaryAnalysis:
              properties:
                interval:
                  type: string
                  # 基于正则式验证
                  pattern: "^[0-9]+(m|s)"
                threshold:
                  type: number
                maxWeight:
                  type: number
                stepWeight:
                  type: number
                metrics:
                  type: array
                  properties:
                    items:
                      type: object
                      required: ['name', 'threshold']
                      properties:
                        name:
                          type: string
                        interval:
                          type: string
                          pattern: "^[0-9]+(m|s)"
                        threshold:
                          type: number
                        query:
                          type: string
                webhooks:
                  type: array
                  properties:
                    items:
                      type: object
                      required: ['name', 'url', 'timeout']
                      properties:
                        name:
                          type: string
                        url:
                          type: string
                          format: url
                        timeout:
                          type: string
                          pattern: "^[0-9]+(m|s)"
```