![](https://github.com/mugioka/slack-bots/actions/workflows/ci-apps.yaml/badge.svg?branch=main)
![](https://github.com/mugioka/slack-bots/actions/workflows/ci-helm-charts.yaml/badge.svg?branch=main)
![](https://github.com/mugioka/slack-bots/actions/workflows/release-helm-charts.yaml/badge.svg?branch=main)
[![codecov](https://codecov.io/gh/mugioka/slack-bots/branch/main/graph/badge.svg?token=A343SU2DXS)](https://codecov.io/gh/mugioka/slack-bots)

# slack-bots

## How to develop

### Install tools with `asdf`

- Install an asdf if you did not install it.
  - [docs](http://asdf-vm.com/guide/getting-started.html#_1-install-dependencies)
- Install tools

  ```bash
  awk '{ print $1; }' .tool-versions | while read tool; do asdf plugin add ${tool}; done; asdf install
  ```

### Install tools with `go install`

```bash
cat tools/tools.go | awk -F'"' '/_/ {print $2}' | xargs -tI {} go install {}
```

### Setup pre-commit

```bash
pre-commit install
```
