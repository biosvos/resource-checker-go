package unstructure

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
)
