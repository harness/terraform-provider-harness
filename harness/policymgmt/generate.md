This is how you generate the go files for the policy management module.

```bash
cp ~/workspace/policy-mgmt/design/gen/http/openapi3.json swagger/
docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli-v3 generate \                                           
    -i /local/swagger/openapi3.json --additional-properties packageName=policymgmt \
    -l go \
    -o /local/
```
