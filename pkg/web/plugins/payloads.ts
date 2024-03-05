export default definePayloadPlugin(() => {
  definePayloadReducer('Specimen', (data) => data === '<original-blink>' && '_')
  definePayloadReviver('BlinkingText', () => '<revivified-blink>')
})
