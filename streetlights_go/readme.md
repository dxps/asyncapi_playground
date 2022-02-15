## Streetlights - Using Go

This is based on the official [Streetlights](https://www.asyncapi.com/docs/tutorials/streetlights) tutorial, but it uses Go instead of Node.js.



<br/>

### Prereqs

- [AsyncAPI Generator](https://github.com/asyncapi/generator/) (`ag`)
  - Used `pnpm install -g @asyncapi/generator` that installed ver. 1.9.0 (available at the time of this writing).



<br/>

### Steps

The following steps were performed:

1. Created `asyncapi.yaml` spec file.
2. Generated the Go code using `ag asyncapi.yaml @asyncapi/go-watermill-template -o ./goapp -p moduleName=github.com/dxps/asyncapi_playground/streetlights_go/goapp`