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
      return 'bg-blue-200 dark:bg-blue-800'
    case 'update':
      return 'bg-green-200 dark:bg-green-800'
    case 'suggest':
      return 'bg-orange-200 dark:bg-orange-800'
    default:
      return 'bg-blue-200 dark:bg-blue-800'
  }
}
