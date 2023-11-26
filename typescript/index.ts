export interface Model {
    provider: string
    model: string
    price_query: number
    price_output: number
}

export const OpenAIModels: Model[] = [
    {
        provider: 'openai',
        model: 'gpt-4-1106-preview',
        price_query: 0.01,
        price_output: 0.03,
    },
    {
        provider: 'openai',
        model: 'gpt-4-1106-vision-preview',
        price_query: 0.01,
        price_output: 0.03,
    },
    {
        provider: 'openai',
        model: 'gpt-4',
        price_query: 0.03,
        price_output: 0.06,
    },
    {
        provider: 'openai',
        model: 'gpt-4-32k',
        price_query: 0.06,
        price_output: 0.12,
    },
    {
        provider: 'openai',
        model: 'gpt-3.5-turbo-1106',
        price_query: 0.001,
        price_output: 0.002,
    },
    {
        provider: 'openai',
        model: 'gpt-3.5-turbo-instruct',
        price_query: 0.0015,
        price_output: 0.002,
    },
]

export const Models: Model[] = [...OpenAIModels]
