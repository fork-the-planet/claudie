# Claudie `v0.8`

!!! warning "Due to updating terraform files the `v0.8.x` clusters build with claudie version `v0.7.x` will be forced to be recreated. Before updating make backups of your data"

## Deployment

To deploy Claudie `v0.8.X`, please:

1. Download claudie.yaml from [release page](https://github.com/berops/claudie/releases)

2. Verify the checksum with `sha256` (optional)

   We provide checksums in `claudie_checksum.txt` you can verify the downloaded yaml files againts the provided checksums.

3. Install claudie using `kubectl`

> We strongly recommend changing the default credentials for MongoDB, MinIO and DynamoDB before you deploy it.

```
kubectl apply -f https://github.com/berops/claudie/releases/latest/download/claudie.yaml
```

To further harden claudie, you may want to deploy our pre-defined network policies:
   ```bash
   # for clusters using cilium as their CNI
   kubectl apply -f https://github.com/berops/claudie/releases/latest/download/network-policy-cilium.yaml
   ```
   ```bash
   # other
   kubectl apply -f https://github.com/berops/claudie/releases/latest/download/network-policy.yaml
   ```


## v0.8.0

### Features

- Allow to reapply manifest after ERROR [#1337](https://github.com/berops/claudie/pull/1337)
- Allow other usernames with root access [#1335](https://github.com/berops/claudie/pull/1335)
- Fix substring match resulting in deletion of wrong nodes [#1350](https://github.com/berops/claudie/pull/1350)
- Add spec.providers validation [#1352](https://github.com/berops/claudie/pull/1352)
- Correctly change the API endpoint [#1366](https://github.com/berops/claudie/pull/1366)
- Restrict nodepool and cluster names to 14 and 28 characters respectively, and add the ability to define and use providers in a single cluster [#1348](https://github.com/berops/claudie/pull/1348)
- Prohibit changing the cloud provider in a nodepool [#1371](https://github.com/berops/claudie/pull/1371)