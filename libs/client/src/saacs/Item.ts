import { Any, PartialMessage } from '@bufbuild/protobuf'
import { key, pb } from '@saacs/saacs-pb'

export function PrimaryToItem(m: key.PrimaryItemTypeMessage): pb.Item {
  return new pb.Item(
    {
      key: key.PrimaryToKeySchema(m),
      value: Any.pack(m),
    },
  )
}
