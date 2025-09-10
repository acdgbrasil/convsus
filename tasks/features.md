# Epic: Integração de Dados com o Ecossistema DATASUS

**Objetivo:** Permitir que os dados do sistema Convsus sejam diretamente utilizáveis pelas ferramentas de análise do DATASUS (TabWin e TabNet), eliminando a necessidade de conversão manual e habilitando a análise integrada de dados sociais e de saúde.

---

## Feature: Exportador de Dados JSON para o Formato TabWin/TabNet

**Narrativa de Usuário:**

> **Como** um analista de dados de saúde pública,
> **Eu quero** converter os dados do Convsus (em formato JSON) para o padrão de arquivos do DATASUS (`.DBF`, `.DEF`, `.CNV`),
> **Para que** eu possa importar e analisar esses dados diretamente no TabWin e TabNet.

### Caso de Uso 1: Exportar Dados de Pessoas de Referência

**Descrição:** Desenvolver um processo para transformar os dados contidos nos arquivos JSON da pasta `Schemas/` em um conjunto de arquivos compatível com o TabWin. O processo deve ler os dados, aplicar um mapeamento para um formato tabular e gerar os três arquivos necessários: um `.DBF` com os dados brutos, um `.DEF` com a definição da estrutura e um `.CNV` com os dicionários de conversão.

---

### Critérios de Aceite Funcionais (BDD)

**Cenário 1: Geração bem-sucedida dos arquivos para um tipo de dado**
```gherkin
Dado que eu tenho um ou mais arquivos JSON de "Pessoa de Referência" (ex: `reference_person_data.json`)
E existe uma configuração de mapeamento que define a correspondência entre os campos do JSON e as colunas do TabWin
Quando eu executo o processo de exportação para "Pessoa de Referência"
Então três arquivos devem ser gerados na saída: "PESSOA_REF.DBF", "PESSOA_REF.DEF" e "PESSOA_REF.CNV"
E o arquivo "PESSOA_REF.DBF" deve conter os dados dos arquivos JSON de origem, de forma consolidada e achatada ("flattened")
E o arquivo "PESSOA_REF.DEF" deve descrever corretamente as colunas, tipos e tamanhos dos campos presentes no arquivo `.DBF`
E o arquivo "PESSOA_REF.CNV" deve conter os dicionários para traduzir códigos em descrições (ex: mapear a sigla "SP" para "São Paulo")
```

**Cenário 2: Mapeamento e achatamento de dados aninhados**
```gherkin
Dado que um arquivo JSON de entrada contém dados aninhados, como: `{"rg": {"number": "1234567", "uf": "SP"}}`
E a configuração de mapeamento define que `rg.number` corresponde à coluna `NU_RG` e `rg.uf` à coluna `SG_UF_RG`
Quando o processo de exportação é executado
Então o arquivo `.DBF` gerado deve conter as colunas `NU_RG` e `SG_UF_RG`
E os valores para um registro de exemplo devem ser "1234567" e "SP", respectivamente
```

**Cenário 3: Geração de arquivos de conversão (`.CNV`) a partir dos dados**
```gherkin
Dado que o campo `biologicalGender` no JSON possui os valores "Masculino" e "Feminino"
E a coluna de destino no `.DBF`, chamada `TP_SEXO`, deve armazenar esses valores como códigos, por exemplo, '1' e '2'
Quando o processo de exportação é executado
Então o arquivo "PESSOA_REF.CNV" deve ser gerado contendo uma seção para `TP_SEXO` que mapeia '1' para "Masculino" e '2' para "Feminino"
E o arquivo `.DBF` deve armazenar apenas os códigos '1' e '2' na coluna `TP_SEXO`
```

---

### Critérios de Aceite Não Funcionais

*   **Compatibilidade:** Os arquivos gerados devem ser 100% compatíveis com o software TabWin (versão 3.6 ou superior) e com a plataforma TabNet.
*   **Desempenho:** A conversão de 100.000 registros JSON (aproximadamente 50MB) deve ser concluída em menos de 2 minutos.
*   **Configurabilidade:** O mapeamento de campos JSON para as colunas do `.DBF`, incluindo a definição de tipos, tamanhos e dicionários, deve ser gerenciado através de um arquivo de configuração externo e de fácil edição.
*   **Tratamento de Erros:** O sistema deve registrar logs detalhados e mensagens de erro claras em caso de falhas, como JSON malformado, inconsistências de tipo de dado ou problemas no arquivo de mapeamento.

---

### Testes de Alto Nível

1.  **Testes Unitários:**
    *   Validar a função que "achata" a estrutura JSON aninhada para um dicionário de primeiro nível.
    *   Testar a conversão de tipos de dados (ex: `string` de data para `date`, `string` de número para `numeric`).
    *   Testar o módulo de escrita de arquivos `.DBF` para garantir a conformidade com o formato dBase III.
    *   Testar a geração dos arquivos `.DEF` e `.CNV` com base em um esquema de mapeamento de exemplo.

2.  **Testes de Integração:**
    *   Executar um teste de ponta a ponta com um conjunto pequeno e controlado de arquivos JSON, validando o conteúdo e a estrutura dos três arquivos de saída (`.DBF`, `.DEF`, `.CNV`).
    *   Testar a capacidade do sistema de consolidar múltiplos arquivos JSON do mesmo tipo em um único conjunto de arquivos de saída.

3.  **Testes de Aceitação (UAT - User Acceptance Testing):**
    *   **Teste de Importação:** Um analista deve conseguir pegar os arquivos gerados e carregá-los em uma instância real do TabWin sem erros.
    *   **Teste de Tabulação:** Após a importação no TabWin, realizar uma tabulação simples (ex: contar o número de pessoas por estado) e verificar se o resultado corresponde aos dados dos arquivos JSON de origem.

4.  **Testes Negativos:**
    *   Tentar exportar usando um arquivo JSON com sintaxe inválida.
    *   Executar o processo com um arquivo de mapeamento ausente ou corrompido.
    *   Fornecer dados que violem o esquema de destino (ex: uma string com 50 caracteres para um campo definido com tamanho 40) e verificar se um erro apropriado é lançado.
