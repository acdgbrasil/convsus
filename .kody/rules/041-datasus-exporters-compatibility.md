---
title: "Exportadores DATASUS: compatibilidade TabWin/TabNet garantida"
scope: "pull_request"
path: ["internal/format/**", "cmd/**", "testdata/**"]
severity_min: "high"
bu  ckets: ["interoperability", "quality-gates"]
enabled: true
---

## Instructions
- `.DBF` em dBase III (tipos válidos, tamanhos e padding corretos; codepage consistente; datas YYYYMMDD).
- `.DEF` descreve fielmente colunas/tipos/tamanhos do `.DBF` gerado.
- `.CNV` cobre mapeamentos de códigos ↔ descrições usados; sem chaves órfãs.
- Inclua testes/golden e evidência de importação no TabWin/TabNet quando aplicável.

## Examples

### Bad example
```
Added: internal/format/dbf/write.go // sem testes nem DEF/CNV correspondente
```

### Good example
```
Added: internal/format/dbf/write.go
Added: internal/format/def/write.go
Added: internal/format/cnv/write.go
Added: testdata/pessoa_ref.{dbf,def,cnv}
Docs: instruções de validação no TabWin
```

