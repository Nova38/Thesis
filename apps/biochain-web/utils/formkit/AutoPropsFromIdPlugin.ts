import { titleCase } from 'scule'
import { type FormKitNode } from '@formkit/core'

export default (node: FormKitNode) => {
  if (!node.props.id) return
  console.log('AutoPropsFromIdPlugin', node.props.id)
  node.name = node.props.id // auto set name to id
  node.props.label = titleCase(node.props.id) // auto set label
  if (!['button', 'submit'].includes(node.props.type)) {
    // automatically set help text, but exclude buttons
  }
}
