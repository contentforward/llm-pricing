import { Models } from '../index'

describe('Models', () => {
    test('OpenAIModels is an array of Models', () => {
        expect(Array.isArray(Models)).toBe(true)
        Models.forEach(model => {
            expect(typeof model.provider).toBe('string')
            expect(typeof model.model).toBe('string')
            expect(typeof model.priceQuery).toBe('number')
            expect(typeof model.priceOutput).toBe('number')
        })
    })
})
