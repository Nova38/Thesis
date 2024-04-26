import { auth, pb } from '@saacs/saacs-pb'

export const ActionsAlias: Record<string, pb.Action[]> = {
  all: [pb.Action.CREATE, pb.Action.VIEW, pb.Action.],
} as const
