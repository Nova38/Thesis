import { auth, pb } from '@saacs/saacs-pb'

export const ActionsAlias = {
  all: {
    type: {
      primary: [
        pb.Action.VIEW,
        pb.Action.CREATE,
        pb.Action.UPDATE,
        pb.Action.DELETE,
      ],
      subitem: {
        history: [
          pb.Action.VIEW_HISTORY,
          pb.Action.VIEW_HIDDEN_TXS,
          pb.Action.HIDE_TX,
          pb.Action.UNHIDE_TX,
        ],
        suggest: [
          pb.Action.SUGGEST_VIEW,
          pb.Action.SUGGEST_CREATE,
          pb.Action.SUGGEST_APPROVE,
          pb.Action.SUGGEST_DELETE,
        ],
      },
    },
    level: {
      view: [
        pb.Action.VIEW,
        pb.Action.SUGGEST_VIEW,
        pb.Action.VIEW_HISTORY,
      ],
      create: [
        pb.Action.CREATE,
        pb.Action.SUGGEST_CREATE,
        pb.Action.SUGGEST_APPROVE,
      ],
      update: [
        pb.Action.UPDATE,
        pb.Action.SUGGEST_APPROVE,
      ],

    },
  },

} as const
