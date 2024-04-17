import { type FormKitNode } from '@formkit/core'

export default (node: FormKitNode) => {
  if (!node.context) return
  // console.log(node.context)
  node.context.classes.label = node.context?.state.dirty
    ? 'text-red-500'
    : 'text-surface-700 dark:text-surface-0/80'
  node.on('commit', (v) => {
    if (node.context?.state.dirty) {
      node.context.help = '* Modified'
    }
    // console.log('commit', v)
  })
}
