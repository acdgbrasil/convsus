# Task 01: Motor de Transformação de Dados

**Epic:** Core de Processamento de Dados
**Feature:** Transformação e Mapeamento de Dados

**Narrativa de Usuário:**

> **Como** um desenvolvedor,
> **Eu quero** uma função ou módulo confiável e puro para transformar um registro de dados de origem em um registro de dados de destino,
> **Para que** eu possa reutilizar essa lógica central em diferentes fluxos de trabalho (como processamento de arquivos JSON, eventos de mensageria, etc.).

---

### Critérios de Aceite Funcionais (BDD)

**Cenário 1: Mapeamento de campos diretos**
```gherkin
Dado um objeto de origem: `{"fullName": "João da Silva", "birthDate": "1990-05-15"}`
E uma regra de mapeamento: `{"nome_completo": "fullName", "data_nascimento": "birthDate"}`
Quando eu processo o objeto com a regra de mapeamento
Então o resultado deve ser: `{"nome_completo": "João da Silva", "data_nascimento": "1990-05-15"}`
```

**Cenário 2: Mapeamento de campos aninhados (Flattening)**
```gherkin
Dado um objeto de origem: `{"identificacao": {"pessoal": {"nome": "Maria Oliveira"}}, "local": "Urbano"}`
E uma regra de mapeamento: `{"nome": "identificacao.pessoal.nome", "zona": "local"}`
Quando eu processo o objeto com a regra de mapeamento
Então o resultado deve ser: `{"nome": "Maria Oliveira", "zona": "Urbano"}`
```

**Cenário 3: Mapeamento com valor padrão para campos ausentes**
```gherkin
Dado um objeto de origem: `{"nome": "José"}`
E uma regra de mapeamento que espera um sobrenome: `{"primeiro_nome": "nome", "ultimo_nome": "sobrenome"}`
E a regra define que um campo ausente deve ter o valor `null`
Quando eu processo o objeto com a regra de mapeamento
Então o resultado deve ser: `{"primeiro_nome": "José", "ultimo_nome": null}`
```

**Cenário 4: Mapeamento com transformação de valor**
```gherkin
Dado um objeto de origem: `{"genero": "F"}`
E uma regra de mapeamento que inclui uma transformação: `{"sexo": {"campo": "genero", "transformacao": {"F": "Feminino", "M": "Masculino"}}}`
Quando eu processo o objeto com a regra de mapeamento
Então o resultado deve ser: `{"sexo": "Feminino"}`
```

---

### Critérios de Aceite Não Funcionais

*   **Pureza:** A função de transformação não deve ter efeitos colaterais (ex: não deve fazer I/O de rede ou disco, nem modificar estado global). Sua saída deve depender exclusivamente de suas entradas.
*   **Desempenho:** A transformação de um único objeto deve ter uma média de execução inferior a 1 milissegundo.
*   **Clareza:** O formato das regras de mapeamento deve ser claro, bem-definido e documentado.

---

### Testes de Alto Nível

#### Testes Unitários (BDD Steps)

*   **Feature: Validação do Motor de Transformação**
    *   **Step:** `Dado um dicionário de origem e um dicionário de mapeamento`
    *   **Step:** `Quando a função de transformação é invocada`
    *   **Step:** `Então o dicionário de saída deve corresponder à estrutura e aos valores esperados`
    *   *(Nota: Cobrir todos os cenários funcionais acima, incluindo mapeamento direto, aninhado, campos ausentes, transformações de valor e tipos de dados diferentes).*

#### Testes de Integração (Para QA)

*   **Objetivo:** Validar o motor de transformação com um exemplo real e complexo antes de integrá-lo com os módulos de I/O.
*   **Procedimento:**
    1.  **Setup:** Criar um script de teste (`test_engine_integration.go`).
    2.  **Input:** O script deve carregar um único objeto do arquivo `Schemas/complete_reference_person_data.json`.
    3.  **Config:** O script deve definir um dicionário de mapeamento realista que transforme os campos do objeto de origem para um novo formato (ex: simplificando nomes e achatando a estrutura `rg`).
    4.  **Execução:** O script deve invocar o motor de transformação com o objeto e o mapeamento.
    5.  **Validação:**
        *   O script deve usar `asserts` para verificar se o dicionário resultante contém as chaves e os valores esperados.
        *   O script deve imprimir (logar) o dicionário de saída em formato JSON legível.
    6.  **Verificação Manual (QA):** O QA deve revisar o output logado para confirmar que a transformação ocorreu conforme o esperado nas regras de mapeamento.
