import { genImport, genExport, genSafeVariableName } from 'knitwork'
import fs from 'node:fs'
import { filename, normalizeAliases, resolveAlias } from 'pathe/utils'
import { resolve } from 'pathe'

export function BundleSchemas() {
  const schemaFolder = resolve(__dirname, '../schema/gen')
  const exportFile = resolve(__dirname, '../schema/index.ts')

  // get all files in schema folder
  const files = fs.readdirSync(schemaFolder)

  const importDefs: string[] = files.map((file) => {
    const name = filename(file).replaceAll('.', '_')

    const fileName = './gen/' + file

    return genImport(
      fileName,
      { name: '*', as: name },
      { attributes: { type: 'json' } },
    )
  })

  const names = files.map((file) => filename(file).replaceAll('.', '_'))

  const exportDefs = `export { ${names.join(', ')} }`

  const file = importDefs.join('\n') + '\n' + exportDefs

  fs.writeFileSync(exportFile, file)
}
