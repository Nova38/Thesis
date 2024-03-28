export type FormMode =
  | 'create'
  | 'disabled'
  | 'edit'
  | 'suggest'
  | 'update'
  | 'view'

export function toModeColor(mode: FormMode) {
  switch (mode) {
    case 'view':
      return 'bg-blue-200'
    case 'update':
      return 'bg-green-200'
    case 'suggest':
      return 'bg-orange-200'
    default:
      return 'bg-blue-200'
  }
}
