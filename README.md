### What does this program do?
It parses yaml file with `v1.List` (made of nodes) OR `---` separated `v1.Node`s into `[]*v1.Node`.

### `v1.List` made of nodes:
```yaml
apiVersion: v1
items:
- apiVersion: v1
  kind: Node
  metadata:
    ...
    name: kind-control-plane
    ...
  spec:
    ...
- apiVersion: v1
  kind: Node
  metadata:
    ...
    name: kind-worker
    ...
  spec:
    ...
kind: List
metadata:
  resourceVersion: ""

``` 
Check [node-templates-list.yaml](./node-templates-list.yaml) for an actual example.

### `---` separated `v1.Node`s
```yaml
apiVersion: v1
kind: Node
metadata:
  ...
  name: kind-control-plane
  ...
spec:
  ...
---

apiVersion: v1
kind: Node
metadata:
  ...
  name: kind-worker
  ...
spec:
  ...

```
Check [node-templates-non-list.yaml](./node-templates-non-list.yaml) for an actual example.

### Output 
With `node-templates-non-list.yaml`
```sh
$ go run main.go
node kind-control-plane
node kind-worker
```
With `node-templates-list.yaml`
```sh
$ go run main.go
node kind-control-plane
node kind-worker
node kind-worker2
```
### Why
* This is a PoC to read in template nodes for `kwok` provider in Kubernetes cluster-autoscaler. Check https://github.com/kubernetes/autoscaler/pull/5820

### References
* decoding `---` separated objects: https://github.com/kubernetes/client-go/issues/193#issuecomment-1331874345
* decoding node objects: https://github.com/kubernetes/client-go/issues/193#issuecomment-1331845479
