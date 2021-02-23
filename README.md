# Statefulset PoC

A simple microservice deployed as a statefulset which is scalable using
Horizontal Pod Autoscaler. 

To deploy this example, you need `ko`. run the below command from project's root
directory

```shell
ko apply -n <namespace> config/
```

after deploying every pod of a staefulset gets a usnique cluster local FQDN and can
be reached from any of the cluster workloads. because all the pods would have their
own persistence, session context can be isolated.

to reach one of the pod from a workload, use the below command:
```shell
curl -v http://sfs-sample-0.sfs-service.test.svc.cluster.local:8080/session

curl -v http://sfs-sample-1.sfs-service.test.svc.cluster.local:8080/session

curl -v http://sfs-sample-2.sfs-service.test.svc.cluster.local:8080/session

curl -v http://sfs-sample-3.sfs-service.test.svc.cluster.local:8080/session
```