## To Build

### Warning!!!

in file "config/ReadConfig.go" do this:

```Golang
	// uncomment this code when you build
	// ex, err := os.Executable()
	// if err != nil {
	// 	panic(err)
	// }
	// exePath := filepath.Dir(ex)
	// jsonFile, err := os.Open(exePath + "/config.json")

	// and comment this code when you build
	jsonFile, err := os.Open("./config.json")
```

isso ocorre para que o arquivo compilado consiga reconhecer o arquivo de configurações ao seu lado, e não deixo esse valor como padrão pois gera um bug em desenvolvimento, estou pensando em uma solução para tal.

### Build

run:
`n`

## Config Json Options

The configuration file follows these orders:

- The file needs to be next to the executable

- It has an object called "questions" with a question list and "templateCommit" of type string.

  - The type "question" has an "id", we will talk about it later, a type which can be "select" or "text",a "label", which is the question to be asked (e.g. "What is the scope of this change?") and the optional value "options", if the question has only some valid options (e.g. "front-end", "back-end", "mobile"...). Still have tow optional values, they are used when not declared "options", they define min and max of the string response (e.g. `{ "min": 1, "max": 66}`)

    - the type option need a "name" of type string (e.g. name: "front-end") and "desc" string (e.g. "desc": "Change in front-end scope"). - the "id" value is how to you call the response value in "templateCommit" (e.g "id": "scope").

  - the unique rule for TemplateCommit is the value of each question being called with its id between <> (e.g. `<scope>`)

scheme of what was quoted above:

```Typescript
type config = {
  "questions": []question,
  "templateCommit": string,
}

type question = {
  "id": string,
  "type": string,
  "label": string,
  "options"?: []option,
  "min": int,
  "max": int,
}

type option = {
  "name": string,
  "desc": string
}
```

a full file of configuration for example:
(he works, you can copy and use if you are not in order to create your)

```JSON
{
  "questions": [
    {
      "id": "scope",
      "type": "select",
      "label": "What scope of this change? (e.g. backend or frontend)",
      "options": [
        { "name": "front-end", "desc": "Change in front-end scope" },
        { "name": "back-end", "desc": "Change in back-end scope" }
      ]
    },
    {
      "id": "type",
      "type": "select",
      "label": "Select the type of change you're committing",
      "options": [
        { "name": "feat", "desc": "feat: A new feature" },
        { "name": "fix", "desc": "fix: A bug fix" },
        { "name": "docs", "desc": "docs: Documentation only changes" },
        {
          "name": "style",
          "desc": "style: Changes that do not affect the meaning of the code\n       (white-space, formatting, missing semi-colons, etc)"
        },
        {
          "name": "refactor",
          "desc": "refactor: A code change that neither fixes a bug nor adds a feature"
        },
        {
          "name": "perf",
          "desc": "perf: A code change that improves performance"
        },
        { "name": "test", "desc": "test: Adding missing tests" },
        {
          "name": "chore",
          "desc": "chore: Changes to the build process or auxiliary tools\n       and libraries such as documentation generation"
        },
        { "name": "revert", "desc": "revert: Revert to a commit" },
        { "name": "WIP", "desc": "WIP: Work in progress" }
      ]
    },

    {
      "id": "subject",
      "type": "text",
      "label": "Write a short, imperative tense description of change (max 66 chars)",
      "errorMsg": "Write a minimal 1 and max 66 chars",
      "min": 1,
      "max": 66
    },
    {
      "id": "desc",
      "type": "text",
      "label": "Provide a large description of the changes: (press enter for skip)",
      "errorMsg": "write a valid text",
      "min": 0,
      "max": 1000
    }
  ],
  "templateCommit": "<type>(<scope>): <subject>\n\n<desc>\n"
}
```
