package main

const (
	chocoFailedDeployJson = `{
    "apiVersion": "apps/v1",
    "kind": "Deployment",
    "metadata": {
        "annotations": {
            "deployment.kubernetes.io/revision": "1",
            "kubectl.kubernetes.io/last-applied-configuration": "{\"apiVersion\":\"apps/v1\",\"kind\":\"Deployment\",\"metadata\":{\"annotations\":{},\"name\":\"choco\",\"namespace\":\"wow\"},\"spec\":{\"selector\":{\"matchLabels\":{\"name\":\"choco\"}},\"template\":{\"metadata\":{\"labels\":{\"name\":\"choco\"}},\"spec\":{\"containers\":[{\"args\":[\"sleep\",\"infinity\"],\"image\":\"golan:1.14\",\"name\":\"choco\"}]}}}}\n"
        },
        "creationTimestamp": "2023-03-28T01:59:25Z",
        "generation": 1,
        "name": "choco",
        "namespace": "wow",
        "resourceVersion": "247617",
        "uid": "07768167-4e3c-40c9-929b-4d36d0639b0e"
    },
    "spec": {
        "progressDeadlineSeconds": 600,
        "replicas": 1,
        "revisionHistoryLimit": 10,
        "selector": {
            "matchLabels": {
                "name": "choco"
            }
        },
        "strategy": {
            "rollingUpdate": {
                "maxSurge": "25%",
                "maxUnavailable": "25%"
            },
            "type": "RollingUpdate"
        },
        "template": {
            "metadata": {
                "creationTimestamp": null,
                "labels": {
                    "name": "choco"
                }
            },
            "spec": {
                "containers": [
                    {
                        "args": [
                            "sleep",
                            "infinity"
                        ],
                        "image": "golan:1.14",
                        "imagePullPolicy": "IfNotPresent",
                        "name": "choco",
                        "resources": {},
                        "terminationMessagePath": "/dev/termination-log",
                        "terminationMessagePolicy": "File"
                    }
                ],
                "dnsPolicy": "ClusterFirst",
                "restartPolicy": "Always",
                "schedulerName": "default-scheduler",
                "securityContext": {},
                "terminationGracePeriodSeconds": 30
            }
        }
    },
    "status": {
        "conditions": [
            {
                "lastTransitionTime": "2023-03-28T01:59:25Z",
                "lastUpdateTime": "2023-03-28T01:59:25Z",
                "message": "Deployment does not have minimum availability.",
                "reason": "MinimumReplicasUnavailable",
                "status": "False",
                "type": "Available"
            },
            {
                "lastTransitionTime": "2023-03-28T02:09:26Z",
                "lastUpdateTime": "2023-03-28T02:09:26Z",
                "message": "ReplicaSet \"choco-64cc6f5678\" has timed out progressing.",
                "reason": "ProgressDeadlineExceeded",
                "status": "False",
                "type": "Progressing"
            }
        ],
        "observedGeneration": 1,
        "replicas": 1,
        "unavailableReplicas": 1,
        "updatedReplicas": 1
    }
}`
	chocoFailedReplicasetJson = `{
  "apiVersion": "apps/v1",
  "kind": "ReplicaSet",
  "metadata": {
    "labels": {
      "name": "choco",
      "pod-template-hash": "64cc6f5678"
    },
    "name": "choco-64cc6f5678",
    "namespace": "wow",
    "ownerReferences": [
      {
        "apiVersion": "apps/v1",
        "blockOwnerDeletion": true,
        "controller": true,
        "kind": "Deployment",
        "name": "choco",
        "uid": "07768167-4e3c-40c9-929b-4d36d0639b0e"
      }
    ]
  },
  "spec": {
    "replicas": 1,
    "selector": {
      "matchLabels": {
        "name": "choco",
        "pod-template-hash": "64cc6f5678"
      }
    },
    "template": {
      "metadata": {
        "labels": {
          "name": "choco",
          "pod-template-hash": "64cc6f5678"
        }
      },
      "spec": {
        "containers": [
          {
            "args": [
              "sleep",
              "infinity"
            ],
            "image": "golan:1.14",
            "name": "choco"
          }
        ]
      }
    }
  },
  "status": {
    "fullyLabeledReplicas": 1,
    "observedGeneration": 1,
    "replicas": 1
  }
}`
	chocoFailedPodManifest = `{
    "apiVersion": "v1",
    "kind": "Pod",
    "metadata": {
        "creationTimestamp": "2023-03-28T01:59:25Z",
        "generateName": "choco-64cc6f5678-",
        "labels": {
            "name": "choco",
            "pod-template-hash": "64cc6f5678"
        },
        "name": "choco-64cc6f5678-664r6",
        "namespace": "wow",
        "ownerReferences": [
            {
                "apiVersion": "apps/v1",
                "blockOwnerDeletion": true,
                "controller": true,
                "kind": "ReplicaSet",
                "name": "choco-64cc6f5678",
                "uid": "c8e2e7bf-15cb-4f1e-9656-9ca46df07a4f"
            }
        ],
        "resourceVersion": "248089",
        "uid": "52a4023a-e13b-4612-a738-62f5a87691e4"
    },
    "spec": {
        "containers": [
            {
                "args": [
                    "sleep",
                    "infinity"
                ],
                "image": "golan:1.14",
                "imagePullPolicy": "IfNotPresent",
                "name": "choco",
                "resources": {},
                "terminationMessagePath": "/dev/termination-log",
                "terminationMessagePolicy": "File",
                "volumeMounts": [
                    {
                        "mountPath": "/var/run/secrets/kubernetes.io/serviceaccount",
                        "name": "kube-api-access-gxjs2",
                        "readOnly": true
                    }
                ]
            }
        ],
        "dnsPolicy": "ClusterFirst",
        "enableServiceLinks": true,
        "nodeName": "archlinux",
        "preemptionPolicy": "PreemptLowerPriority",
        "priority": 0,
        "restartPolicy": "Always",
        "schedulerName": "default-scheduler",
        "securityContext": {},
        "serviceAccount": "default",
        "serviceAccountName": "default",
        "terminationGracePeriodSeconds": 30,
        "tolerations": [
            {
                "effect": "NoExecute",
                "key": "node.kubernetes.io/not-ready",
                "operator": "Exists",
                "tolerationSeconds": 300
            },
            {
                "effect": "NoExecute",
                "key": "node.kubernetes.io/unreachable",
                "operator": "Exists",
                "tolerationSeconds": 300
            }
        ],
        "volumes": [
            {
                "name": "kube-api-access-gxjs2",
                "projected": {
                    "defaultMode": 420,
                    "sources": [
                        {
                            "serviceAccountToken": {
                                "expirationSeconds": 3607,
                                "path": "token"
                            }
                        },
                        {
                            "configMap": {
                                "items": [
                                    {
                                        "key": "ca.crt",
                                        "path": "ca.crt"
                                    }
                                ],
                                "name": "kube-root-ca.crt"
                            }
                        },
                        {
                            "downwardAPI": {
                                "items": [
                                    {
                                        "fieldRef": {
                                            "apiVersion": "v1",
                                            "fieldPath": "metadata.namespace"
                                        },
                                        "path": "namespace"
                                    }
                                ]
                            }
                        }
                    ]
                }
            }
        ]
    },
    "status": {
        "conditions": [
            {
                "lastProbeTime": null,
                "lastTransitionTime": "2023-03-28T01:59:25Z",
                "status": "True",
                "type": "Initialized"
            },
            {
                "lastProbeTime": null,
                "lastTransitionTime": "2023-03-28T01:59:25Z",
                "message": "containers with unready status: [choco]",
                "reason": "ContainersNotReady",
                "status": "False",
                "type": "Ready"
            },
            {
                "lastProbeTime": null,
                "lastTransitionTime": "2023-03-28T01:59:25Z",
                "message": "containers with unready status: [choco]",
                "reason": "ContainersNotReady",
                "status": "False",
                "type": "ContainersReady"
            },
            {
                "lastProbeTime": null,
                "lastTransitionTime": "2023-03-28T01:59:25Z",
                "status": "True",
                "type": "PodScheduled"
            }
        ],
        "containerStatuses": [
            {
                "image": "golan:1.14",
                "imageID": "",
                "lastState": {},
                "name": "choco",
                "ready": false,
                "restartCount": 0,
                "started": false,
                "state": {
                    "waiting": {
                        "message": "Back-off pulling image \"golan:1.14\"",
                        "reason": "ImagePullBackOff"
                    }
                }
            }
        ],
        "hostIP": "192.168.2.152",
        "phase": "Pending",
        "podIP": "10.42.0.15",
        "podIPs": [
            {
                "ip": "10.42.0.15"
            }
        ],
        "qosClass": "BestEffort",
        "startTime": "2023-03-28T01:59:25Z"
    }
}`
)
