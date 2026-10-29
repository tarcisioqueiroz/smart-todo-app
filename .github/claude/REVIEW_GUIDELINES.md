# Guia de revisão de código

Este arquivo define como o Claude revisa os PRs deste repositório.
Edite à vontade: as mudanças valem a partir do próximo PR revisado.

## Idioma e tom

- Escreva em **português do Brasil**.
- Seja direto e construtivo. Explique o *porquê* de cada apontamento.
- Não elogie trechos sem motivo e não repita o que o diff já mostra.

## Formato do comentário geral (um por revisão)

```markdown
## 🤖 Revisão do Claude

### Resumo
<2 a 4 frases: o que o PR faz e qual a impressão geral>

### Walkthrough
| Arquivo | O que mudou |
|---|---|
| `caminho/arquivo` | <uma linha> |

### Diagrama de sequência
<somente se o PR altera um fluxo entre componentes (ex.: app → API → banco).
Use um bloco ```mermaid sequenceDiagram. Omita a seção se não se aplicar.>

### Checklist
- 🔴 **Bloqueia o merge:** <lista ou "nada">
- 🟠 **Importante:** <lista ou "nada">
- 🟡 **Sugestões / 🔵 nitpicks:** <quantidade, os detalhes estão nos comentários inline>
```

## Formato dos comentários inline

Comece cada comentário com a severidade:

- 🔴 **Crítico**: bug, perda de dados, falha de segurança, quebra de build.
- 🟠 **Importante**: erro não tratado, caso de borda, problema de performance real.
- 🟡 **Sugestão**: legibilidade, design, testes que faltam.
- 🔵 **Nitpick**: detalhe pequeno, opcional.

Quando a correção for objetiva, inclua um bloco de sugestão do GitHub
para aplicar com um clique:

````markdown
```suggestion
<código corrigido>
```
````

Agrupe problemas repetidos num comentário só, citando os outros locais.
Limite: no máximo ~15 comentários inline por revisão; priorize os mais graves.

## O que verificar

### Backend (Go — `backend/`)
- Todo `error` retornado é tratado ou propagado com contexto (`fmt.Errorf("...: %w", err)`).
- Consultas SQL usam parâmetros (`$1`, `$2`), nunca concatenação de strings.
- `context.Context` é repassado do handler até o banco; timeouts quando faz sentido.
- Recursos fechados com `defer` (`rows.Close()`, `resp.Body.Close()`).
- Handlers HTTP validam entrada e retornam status corretos (400, 404, 500).
- Sem goroutines que vazam nem acesso concorrente sem proteção.
- Nada de segredo ou credencial no código; configuração vem de variáveis de ambiente.

### Mobile (React Native / Expo — `mobile/`)
- Siga `mobile/AGENTS.md`: Expo Router, rotas em `src/app/`, código não-rota fora dela.
- Dependências de `useEffect`/`useMemo`/`useCallback` corretas; sem loops de render.
- Estados de carregamento e erro tratados nas chamadas à API.
- Nenhuma chave de API ou segredo embutido no app (tudo no bundle é público).
- Tipagem TypeScript sem `any` desnecessário.

### Integração com IA
- Entrada do usuário em linguagem natural é tratada como não confiável:
  a resposta do modelo é validada antes de virar operação no banco.
- Chaves de provedores de IA ficam só no backend.

### Geral
- Testes para lógica nova ou corrigida, quando fizer sentido.
- Migrações de banco reversíveis e compatíveis com os dados existentes.
- `.env.example` atualizado quando surgir variável nova.

## O que ignorar

- Arquivos gerados e lockfiles (`package-lock.json`, `go.sum`).
- Imagens e fontes em `mobile/assets/`.
- Formatação que `gofmt` ou o lint do Expo já corrigem.
