---
title: "Wrappers: contrato via CLI estável"
scope: "pull_request"
path: ["bindings/**", "cmd/convsus/**"]
severity_min: "medium"
bu  ckets: ["wrappers", "interoperability"]
enabled: true
---

## Instructions
- Wrappers (Python/Node/etc.) devem, inicialmente, chamar a CLI `convsus` (contrato estável de I/O).
- Evite acoplar a internals; quando gRPC/SDKs estiverem prontos, migre progressivamente.
- Documente expectativas de entrada/saída e erros retornados pela CLI.

## Examples

### Bad example
```
bindings/python/__init__.py // importando internal/transform diretamente
```

### Good example
```
bindings/python/convsus.py // chama `convsus export ...` e trata stdout/stderr
README do binding documenta flags e erros
```

