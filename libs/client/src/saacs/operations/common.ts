import { auth, pb } from '@saacs/saacs-pb'

export type Alias = pb.Action[]

export interface ActionsAliasInterface {
  all: {
    full: pb.Action[]
    type: {
      primary: Alias
      subitem: {
        mspHiddenTx: Alias
        hiddenTx: Alias
        history: Alias
        suggest: Alias
      }
    }
    level: {
      view: Alias
      create: Alias
      update: Alias
      suggest: Alias
      manageSuggestions: Alias
    }
  }
}

export const ActionsAlias: ActionsAliasInterface = {
  all: {
    full: [
      pb.Action.VIEW,
      pb.Action.CREATE,
      pb.Action.UPDATE,
      pb.Action.DELETE,
      pb.Action.SUGGEST_VIEW,
      pb.Action.SUGGEST_CREATE,
      pb.Action.SUGGEST_DELETE,
      pb.Action.SUGGEST_APPROVE,
      pb.Action.VIEW_HISTORY,
      pb.Action.VIEW_HIDDEN_TXS,
      pb.Action.UNHIDE_TX,
      pb.Action.HIDE_TX,
      pb.Action.VIEW_HISTORY,
      pb.Action.VIEW_MSP_HIDDEN_TX,
      pb.Action.HIDE_MSP_TX,
      pb.Action.UNHIDE_MSP_TX,
    ],

    type: {
      primary: [
        pb.Action.VIEW,
        pb.Action.CREATE,
        pb.Action.UPDATE,
        pb.Action.DELETE,
      ],
      subitem: {
        mspHiddenTx: [
          pb.Action.VIEW_MSP_HIDDEN_TX,
          pb.Action.HIDE_MSP_TX,
          pb.Action.UNHIDE_MSP_TX,
        ],
        hiddenTx: [
          pb.Action.VIEW_HIDDEN_TXS,
          pb.Action.UNHIDE_TX,
          pb.Action.HIDE_TX,
        ],
        history: [
          pb.Action.VIEW_HISTORY,
          pb.Action.VIEW_MSP_HIDDEN_TX,
          pb.Action.VIEW_HIDDEN_TXS,
          pb.Action.HIDE_TX,
          pb.Action.UNHIDE_TX,
          pb.Action.HIDE_MSP_TX,
          pb.Action.UNHIDE_MSP_TX,
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
        pb.Action.VIEW_MSP_HIDDEN_TX,
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
      suggest: [
        pb.Action.SUGGEST_VIEW,
        pb.Action.SUGGEST_CREATE,
      ],
      manageSuggestions: [
        pb.Action.SUGGEST_APPROVE,
        pb.Action.SUGGEST_DELETE,
      ],
    },
  },
} as const

export const actionsList: pb.Action[] = [
  pb.Action.VIEW,
  pb.Action.CREATE,
  pb.Action.UPDATE,
  pb.Action.DELETE,
  pb.Action.SUGGEST_VIEW,
  pb.Action.SUGGEST_CREATE,
  pb.Action.SUGGEST_DELETE,
  pb.Action.SUGGEST_APPROVE,
  pb.Action.VIEW_HISTORY,
  pb.Action.VIEW_HIDDEN_TXS,
  pb.Action.UNHIDE_TX,
  pb.Action.HIDE_TX,
  pb.Action.VIEW_HISTORY,
  pb.Action.VIEW_MSP_HIDDEN_TX,
  pb.Action.HIDE_MSP_TX,
  pb.Action.UNHIDE_MSP_TX,
] as const
