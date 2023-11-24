export interface Model {
    provider: string
    model: string
    priceQuery: number
    priceOutput: number
}

export const OpenAIModels: Model[] = [
    {
        provider: 'openai',
        model: 'gpt-4-1106-preview',
        priceQuery: 0.01,
        priceOutput: 0.03,
    },
    {
        provider: 'openai',
        model: 'gpt-4-1106-vision-preview',
        priceQuery: 0.01,
        priceOutput: 0.03,
    },
    {
        provider: 'openai',
        model: 'gpt-4',
        priceQuery: 0.03,
        priceOutput: 0.06,
    },
    {
        provider: 'openai',
        model: 'gpt-4-32k',
        priceQuery: 0.06,
        priceOutput: 0.12,
    },
    {
        provider: 'openai',
        model: 'gpt-3.5-turbo-1106',
        priceQuery: 0.001,
        priceOutput: 0.002,
    },
    {
        provider: 'openai',
        model: 'gpt-3.5-turbo-instruct',
        priceQuery: 0.0015,
        priceOutput: 0.002,
    },
]

export const Models: Model[] = [...OpenAIModels]
