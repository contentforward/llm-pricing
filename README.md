# LLM Pricing as a package

## Supported languages
- Golang
- Javascript / Typescript

## Available models

| Provider | Model                      | Price Query | Price Output |
|----------|----------------------------|-------------|--------------|
| openai   | gpt-4-1106-preview         | 0.01        | 0.03         |
| openai   | gpt-4-1106-vision-preview  | 0.01        | 0.03         |
| openai   | gpt-4                      | 0.03        | 0.06         |
| openai   | gpt-4-32k                  | 0.06        | 0.12         |
| openai   | gpt-3.5-turbo-1106         | 0.0010      | 0.0020       |
| openai   | gpt-3.5-turbo-instruct     | 0.0015      | 0.0020       |

### Golang

```go

package main

import (
    "fmt"
    "github.com/ignastech/llm-pricing"
)

func main() {
	fmt.Printf("%+v", llm_pricing.Models) // all models
	fmt.Printf("%+v", llm_pricing.OpenAIModels) // just openai models
}
```

### Javascript / Typescript

```js

import { Models, OpenAIModels } from 'llm-pricing';

console.log(Models); // all models
console.log(OpenAIModels); // just openai models
```