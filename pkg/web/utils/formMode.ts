export type FormMode =
  | "disabled"
  | "view"
  | "edit"
  | "suggest"
  | "update"
  | "create";

export const toModeColor = (mode: FormMode) => {
  switch (mode) {
    case "view":
      return "bg-green-200";
    case "update":
      return "bg-red-200";
    case "suggest":
      return "bg-orange-200";
    default:
      return "primary";
  }
};
