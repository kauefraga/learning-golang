# `Learning Golang` (:brazil:)

> 🐿 Aprendendo a linguagem Go. Escolhi essa linguagem porque queria testar o [Fyne](https://fyne.io) e também porque já experimentei-a há um tempo atrás e adorei a modularização, gerenciador de pacotes, performance e o ecossistema no geral. Decidi escrever este learning-{algumacoisa} em português pois eu costumo escrever meus projetos em inglês e vi aqui a oportunidade de agregar a comunidade brasileira.

## ⬇ Começando

### Pré-requisitos

Abaixo eu vou citar os pré-requisitos para usar a linguagem Go, como instalá-los e recomendar algumas ferramentas e extensões.

- Ter a linguagem [Go](https://go.dev) instalada. Veja na seção [#instalando a linguagem go](#instalando-a-linguagem-go)

- Ter um editor de texto (estou utilizando [Visual Studio Code](https://code.visualstudio.com)). Veja [aqui](#instalando-e-configurando-o-visual-studio-code).

### Instalando a [linguagem Go](https://go.dev)

#### Windows

- Você pode instalar com um gerenciador de pacote como o [Scoop](https://scoop.sh), [Chocolatey](https://chocolatey.org) ou outro.
  - Exemplo: `scoop install main/go`

- Entre na [página de download da linguagem](https://go.dev/dl), procure por "Microsoft Windows" e baixe o "go{x.xx.x}.windows-amd64.msi".

#### Linux

Não estou usando sistemas Linux atualmente. Leia:

- [How to install Go in Linux](https://golangdocs.com/install-go-linux)
- [Como instalar Go no Linux](https://dev.to/alexandreliberato/como-instalar-golang-no-linux-3pl9)

### Instalando e configurando o [Visual Studio Code](https://code.visualstudio.com)

Assim como na instalação do Go, você pode usar `scoop install extras/vscode`.
Se não estiver com vontade de baixar um gerenciador de pacote, então simplesmente entre no [site do Visual Studio Code](https://code.visualstudio.com) e clique no botão azul "Download for xxxxxx". Você vai ser redirecionado para [outra página](https://code.visualstudio.com/docs), tente ler um pouco.

Para instalar no Linux, veja [a documentação do Visual Studio Code sobre Linux](https://code.visualstudio.com/docs/setup/linux).

Vamos para as extensões. Se tem uma coisa que o Visual Studio Code é bom, então isso é a quantidade de customizações disponíveis! Aqui irei citar algumas recomendações de extensões para você baixar:

- [Go](https://marketplace.visualstudio.com/items?itemName=golang.Go);
- [Code Spell Checker](https://marketplace.visualstudio.com/items?itemName=streetsidesoftware.code-spell-checker);
  - Inglês já vem por padrão.
  - [Português Brasileiro](https://marketplace.visualstudio.com/items?itemName=streetsidesoftware.code-spell-checker-portuguese-brazilian)
- [Tabnine AI](https://marketplace.visualstudio.com/items?itemName=TabNine.tabnine-vscode) ou [Codeium](https://marketplace.visualstudio.com/items?itemName=Codeium.codeium);
- [Editor Config](https://marketplace.visualstudio.com/items?itemName=EditorConfig.EditorConfig);
- [Git Lens](https://marketplace.visualstudio.com/items?itemName=eamodio.gitlens);

- [Markdown Lint](https://marketplace.visualstudio.com/items?itemName=DavidAnson.vscode-markdownlint);
- [Markdown Checkboxes](https://marketplace.visualstudio.com/items?itemName=bierner.markdown-checkbox);
- [Markdown Emoji](https://marketplace.visualstudio.com/items?itemName=bierner.markdown-emoji);
- [Markdown Preview Github](https://marketplace.visualstudio.com/items?itemName=bierner.markdown-preview-github-styles);
- [Markdown Preview Mermaid](https://marketplace.visualstudio.com/items?itemName=bierner.markdown-mermaid);

- [Fluent Icons](https://marketplace.visualstudio.com/items?itemName=miguelsolorio.fluent-icons);
- [Material Theme Icons](https://marketplace.visualstudio.com/items?itemName=Equinusocio.vsc-material-theme-icons) ou [Material Icon Theme](https://marketplace.visualstudio.com/items?itemName=PKief.material-icon-theme);
- [Dracula](https://marketplace.visualstudio.com/items?itemName=dracula-theme.theme-dracula), [One Dark Pro](https://marketplace.visualstudio.com/items?itemName=zhuangtongfa.Material-theme) e [Atom One Dark Theme](https://marketplace.visualstudio.com/items?itemName=akamud.vscode-theme-onedark).

### Instalando este projeto

1. Clonar o repositório
2. Compilar e rodar os códigos

```bash
# (1)
git clone https://github.com/kauefraga/learning-golang.git

# (2)
go run website-status/src/main.go
# ou
go build website-status/src/main.go
./main.exe
```

## 📚 Conhecimento adquirido

Lista de aprendizados

- Entrada padrão

```go
// Pacote principal
package main

// Função de entrada
func main() {

}
```

- Tipos primitivos
  - Inteiros (`int`)
  - Ponto-flutuante (`float32` e `float64`)
  - String
  - Boolianos
- Variáveis
  - Declaração normal: `var nome_da_variável tipo_da_variável`
  - Declaração encurtada: `nome_da_variável := 0` (precisa ser inicializada)
- Inferência de tipos
- Funções
  - Declaração: `func nome_da_função(...parâmetros) tipo_de_retorno {}`
  - Múltiplos retornos
- Arrays (fixos) e slices (dinâmicos)
- Controle de fluxo
  - `if, elseif e else`
  - `switch`
- Repetições
  - for loop: `for ... {}`
  - while loop não existe (estranho, porém acontece) mas pode ser reproduzido usando um loop infinito `for {}` e uma simples lógica
- Encerrar programas
  - Use `os.Exit(código int)`
- Fazer requisições HTTP
  - Use `http.Get(url string)`

## 🧻 Recursos

Lista de conteúdos que consumi para escrever este projeto:

- [Como baixar a linguagem Go](https://go.dev/doc/install)
- [Documentação da linguagem Go](https://go.dev/doc/effective_go)
- [Formação da Alura](https://cursos.alura.com.br/formacao-go)
  - [Variáveis em Go](https://www.alura.com.br/artigos/variaves-com-go-lang)
  - [Estruturas de controle](https://www.alura.com.br/artigos/estruturas-basicas-de-controle-com-go)
  - [Datas](https://www.alura.com.br/artigos/golang-trabalhando-com-datas)
  - [Conversão de tipos](https://www.alura.com.br/artigos/conversao-de-tipos-com-go)

## 📝 Licença

Este projeto está sob licença do MIT - Veja a [LICENÇA](https://github.com/kauefraga/learning-golang/blob/main/LICENSE) para mais informações.
