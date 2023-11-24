import { OpenAIModels } from '../index'

describe('OpenAI Models', () => {
    test('OpenAIModels is an array of Models', () => {
        expect(Array.isArray(OpenAIModels)).toBe(true)
        OpenAIModels.forEach(model => {
            expect(typeof model.provider).toBe('string')
            expect(typeof model.model).toBe('string')
            expect(typeof model.priceQuery).toBe('number')
            expect(typeof model.priceOutput).toBe('number')
        })
    })
    test('each Model has the correct provider', () => {
        OpenAIModels.forEach(model => {
            expect(model.provider).toMatch(/openai/)
        })
    })
    test('Model prices are positive numbers', () => {
        OpenAIModels.forEach(model => {
            expect(model.priceQuery).toBeGreaterThan(0)
            expect(model.priceOutput).toBeGreaterThan(0)
        })
    })
    test('the array contains specific models', () => {
        const modelNames = OpenAIModels.map(model => model.model)
        expect(modelNames).toEqual(
            expect.arrayContaining([
                'gpt-4-1106-preview',
                'gpt-4-1106-vision-preview',
                'gpt-4',
                'gpt-4-32k',
                'gpt-3.5-turbo-1106',
                'gpt-3.5-turbo-instruct',
            ])
        )
    })
})
