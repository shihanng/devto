```
docker run --rm -v ${PWD}:/local openapitools/openapi-generator-cli generate \
    -i https://raw.githubusercontent.com/thepracticaldev/dev.to/7d0aeeefe5cf6250c5b58ae6b631bfe41fe5bf4a/docs/api_v0.yml \
    -g go \
    -o /local/out/go
```
