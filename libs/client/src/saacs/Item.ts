import { Any } from '@bufbuild/protobuf'
import { key, pb } from '@saacs/saacs-pb'

export function ToItem(m: key.PrimaryItemTypeMessage) {
  return new pb.Item(
    {
      key: key.PrimaryToKeySchema(m),
      value: Any.pack(m),
    },
  )
}
