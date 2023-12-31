# Snippets

**Snippets** generates VSCode code snippets from snippets yam file.

> Snippets is a small tool (without dependency)

[![Open in Gitpod](https://gitpod.io/button/open-in-gitpod.svg)](https://gitpod.io/#https://github.com/bots-garden/snippets)


## Usage

```bash
snippets generate \
  --input samples/js.yml \
  --output ../.vscode/js.code-snippets 
```

You can download a remote yaml file:
```bash
snippets generate \
  --url https://raw.githubusercontent.com/bots-garden/golab-demos/main/snippets-demo/messages.yml \
  --input samples/messages.yml \
  --output ../.vscode/messages.code-snippets 
```
> if you need to use an authentication header (ex: `"PRIVATE-TOKEN: mytokenvalue"
`), use the following flags:
> ```bash
> --auth-header-name PRIVATE-TOKEN
> --auth-header-value mytokenvalue
> ``` 

### Input `yaml` file (example)

```yaml
snippet hello world:
  prefix: "js-plugin-hello-world"
  name: "hello world"
  description: "This is the hello world plugin"
  scope: "javascript"
  body: |
    function helloWorld() {
      console.log("👋 hello world 🌍")
    }

snippet good morning:
  prefix: "js-plugin-good-morning"
  name: "good morning"
  description: "this is the good morning plugin"
  scope: "javascript"
  body: |
    function goodMorning(name) {
      console.log("👋", name, "${1:message}")
      console.log("${2:message}")
    }

```

### Output `json` file (example)

```json
{
  "good morning": {
    "body": [
      "function goodMorning(name) {",
      "  console.log(\"👋\", name, \"${1:message}\")",
      "  console.log(\"${2:message}\")",
      "}",
      ""
    ],
    "description": "this is the good morning plugin",
    "prefix": "js-plugin-good-morning",
    "scope": "javascript"
  },
  "hello world": {
    "body": [
      "function helloWorld() {",
      "  console.log(\"👋 hello world 🌍\")",
      "}",
      ""
    ],
    "description": "This is the hello world plugin",
    "prefix": "js-plugin-hello-world",
    "scope": "javascript"
  }
}
```

## Install

### Linux or MacOS

```bash
SNIPPETS_VERSION="0.0.1"
SNIPPETS_OS="linux" # or darwin
SNIPPETS_ARCH="arm64" # or amd64
wget https://github.com/bots-garden/snippets/releases/download/v${SNIPPETS_VERSION}/snippets-v${SNIPPETS_VERSION}-${SNIPPETS_OS}-${SNIPPETS_ARCH}
cp snippets-v${SNIPPETS_VERSION}-${SNIPPETS_OS}-${SNIPPETS_ARCH} snippets
chmod +x snippets
rm snippets-v${SNIPPETS_VERSION}-${SNIPPETS_OS}-${SNIPPETS_ARCH}
sudo cp ./snippets /usr/bin
rm snippets
# check the version
snippets version
```

### Docker

```bash
docker run \
    -v $(pwd)/samples:/samples \
    --rm botsgarden/snippets:0.0.1  \
    ./snippets generate \
    --input samples/js.yml \
    --output samples/js.code-snippets 
```


