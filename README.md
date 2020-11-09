# Go-Quay

Generated go client for quay.io

## regenerate client

```shell
curl -L https://quay.io/api/v1/discovery |\
  jq '.definitions.NewRepo.properties.repo_kind.enum=["image","application","null"]'|\
  jq '.definitions.AddLabel.properties.media_type.enum=["text/plain","application/json","null"]'|\
  jq . > swagger.json

# patch swagger.json: replace `null` by `"null"`

swagger generate client -f swagger.json -A client -c quay -a quay
```
