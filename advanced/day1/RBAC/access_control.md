# RBAC and ABAC

RBAC = Role Based Access Control
ABAC = Atribute Based Access Control

### Examples
**RBAC**
```
apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  namespace: projectCaribou
  name: developer
rules:
- apiGroups: [""] # "" indicates the core API group
  resources: ["pods", "deployments"]
  verbs: ["create", "update", "delete"]
```
```
apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: developer-edit
  namespace: projectCaribou
subjects:
# You can specify more than one "subject"
- kind: User
  name: jane # "name" is case sensitive
  apiGroup: rbac.authorization.k8s.io
roleRef:
  kind: Role #this must be Role or ClusterRole
  name: developer # this must match the name of the Role/CR to bind to
  apiGroup: rbac.authorization.k8s.io
```

**ABAC**
```
{
    "apiVersion": "abac.authorization.kubernetes.io/v1beta1",
    "kind": "Policy",
    "spec": {
        "user":"jane",
        "namespace": "projectCaribou",
        "resource": "pods/deployments",
        "apiGroup": "*",
        "readonly": false,
        "verbs": "create/update/delete"
    }
}
```
