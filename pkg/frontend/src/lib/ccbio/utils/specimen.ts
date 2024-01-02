// import { schema } from '../ccSchema/';

import { schema } from '../';
import { Timestamp } from '@bufbuild/protobuf';
import type { AnyMessage, Message, PlainMessage } from '@bufbuild/protobuf';

type StringTypeKeys<Obj> = {
  [Property in keyof Obj]: Obj[Property] extends Timestamp ? string : Property;
}[keyof Obj];

export function PopulateNestedRequired<T>(obj: T) {
  const msg = obj as AnyMessage;

  for (const fieldInfo of msg.getType().fields.byNumber()) {
    if (fieldInfo.kind === 'message') {
      const nested =
        msg[fieldInfo.name] !== undefined
          ? msg[fieldInfo.name]
          : new fieldInfo.T();
      console.log(fieldInfo.name, fieldInfo.T.typeName, nested);

      msg[fieldInfo.localName] = PopulateNestedRequired(nested);
      console.log(msg[fieldInfo.name]);
    }
  }

  return obj;
}

// https://aishwarya2593.medium.com/create-wonders-with-advanced-typescript-types-525acf302770
type TimestampNestedObject<T> = T extends PlainMessage<Timestamp>
  ? string
  : {
      [Property in keyof T]-?: TimestampNestedObject<T[Property]>;
    };

export type FormSpecimen = TimestampNestedObject<PlainMessage<schema.Specimen>>;
//   ^?

export function ReplaceTimestamps(message: AnyMessage) {
  for (const fieldInfo of message.getType().fields.byNumber()) {
    if (fieldInfo.kind === 'message') {

      const nested =
        message[fieldInfo.name] !== undefined
          ? message[fieldInfo.name]
          : new fieldInfo.T();

      // return the path to the timestamp
      if (fieldInfo.T.typeName === 'google.protobuf.Timestamp') {
        // console.log(nested.toJson());


          console.log(fieldInfo);
          // message[fieldInfo.jsonName] = 'him';
          console.log(message)
          // message[fieldInfo.name] = '';

          message[fieldInfo.name] = nested.toJson();


      } else {
        message[fieldInfo.name] = nested;

        ReplaceTimestamps(nested);
      }
    }
  }

  console.log
  return message;
}

export function ReplaceDateStrings(message: AnyMessage, raw: unknown) {
  for (const fieldInfo of message.getType().fields.byNumber()) {
    const raw_field = (raw as Record<string, unknown>)[fieldInfo.name];

    if (fieldInfo.kind === 'message') {
      // return the path to the timestamp
      if (fieldInfo.T.typeName === 'google.protobuf.Timestamp') {
        message[fieldInfo.name] = Timestamp.fromJsonString(raw_field as string);
      } else {
        ReplaceDateStrings(message[fieldInfo.name], raw_field);
      }
    } else {
      message[fieldInfo.name] = raw_field;
    }
  }

  return message;
}

export function SchemaToFormSpecimen(specimen: schema.Specimen): FormSpecimen {
  const s = ReplaceTimestamps(specimen);

  return s.toJson({ emitDefaultValues: true }) as FormSpecimen;
}

export function FormSpecimenToSchema(specimen: FormSpecimen): schema.Specimen {
  const s = ReplaceDateStrings(new schema.Specimen(), specimen);

  return s as schema.Specimen;
}

export function MaybeSchemaToFormSpecimen(
  specimen: schema.Specimen | undefined
): FormSpecimen {
  if (specimen) {
    return SchemaToFormSpecimen(specimen);
  } else {
    return SchemaToFormSpecimen(new schema.Specimen());
  }
}

export function MaybePlainSchemaToForm(
  specimen: PlainMessage<schema.Specimen> | undefined
): FormSpecimen {
  if (specimen) {
    return SchemaToFormSpecimen(new schema.Specimen(specimen));
  } else {
    return SchemaToFormSpecimen(new schema.Specimen());
  }
}

// import { get } from 'lodash';

// export const MakeEmptySpecimen = () => {
//   return schema.Specimen.fromPartial({
//     id: '',
//     collection_id: '',
//     primary: schema.Specimen_Primary.create({}),
//     secondary: schema.Specimen_Secondary.create(),
//     taxon: schema.Specimen_Taxon.create(),
//     georeference: schema.Specimen_Georeference.create(),
//   });
// };

// // Make and initialize all sub-objects of a lection
// export const MakeEmptyCollection = () => {
//   return schema.Collection.fromPartial({
//     collection_id: '',
//     access_control: schema.Collection_AccessControl.fromPartial({
//       roles: schema.Collection_AccessControl_RoleActions.create(),
//       role_permissions: schema.Collection_AccessControl_RoleActions.create(),
//       users: schema.Collection_AccessControl_UserActions.create(),
//       specimen: schema.Collection_AccessControl_SpecimenActions.create(),
//       primary: schema.Collection_AccessControl_SectionActions.create(),
//       secondary: schema.Collection_AccessControl_SectionActions.create(),
//       taxon: schema.Collection_AccessControl_SectionActions.create(),
//       georeference: schema.Collection_AccessControl_SectionActions.create(),
//       images: schema.Collection_AccessControl_SectionActions.create(),
//       loans: schema.Collection_AccessControl_SectionActions.create(),
//       grants: schema.Collection_AccessControl_SectionActions.create(),
//       hidden: schema.Collection_AccessControl_SectionActions.create(),
//       suggested: schema.Collection_AccessControl_SuggestedActions.create(),
//     }),
//   });
// };

// export const MakeDefaultCollection = () => {
//   return schema.Collection.fromPartial({
//     collection_id: '',
//     access_control: schema.Collection_AccessControl.fromPartial({
//       roles: schema.Collection_AccessControl_RoleActions.fromPartial({
//         edit: [schema.Role.ROLE_MANAGER],
//         view: [schema.Role.ROLE_MANAGER],
//       }),
//       role_permissions: schema.Collection_AccessControl_RoleActions.fromPartial(
//         { edit: [schema.Role.ROLE_MANAGER], view: [schema.Role.ROLE_MANAGER] }
//       ),
//       users: schema.Collection_AccessControl_UserActions.fromPartial({
//         edit: [schema.Role.ROLE_MANAGER],
//         view: [schema.Role.ROLE_MANAGER],
//       }),
//       specimen: schema.Collection_AccessControl_SpecimenActions.fromPartial({
//         delete: [schema.Role.ROLE_MANAGER],
//         suggest: [schema.Role.ROLE_MANAGER],
//         create: [schema.Role.ROLE_MANAGER],
//         view: [schema.Role.ROLE_MANAGER],
//       }),
//       primary: schema.Collection_AccessControl_SectionActions.fromPartial({
//         edit: [schema.Role.ROLE_MANAGER],
//         view: [schema.Role.ROLE_MANAGER],
//         suggest: [schema.Role.ROLE_MANAGER],
//       }),
//       secondary: schema.Collection_AccessControl_SectionActions.fromPartial({
//         edit: [schema.Role.ROLE_MANAGER],
//         view: [schema.Role.ROLE_MANAGER],
//         suggest: [schema.Role.ROLE_MANAGER],
//       }),
//       taxon: schema.Collection_AccessControl_SectionActions.fromPartial({
//         edit: [schema.Role.ROLE_MANAGER],
//         view: [schema.Role.ROLE_MANAGER],
//         suggest: [schema.Role.ROLE_MANAGER],
//       }),
//       georeference: schema.Collection_AccessControl_SectionActions.fromPartial({
//         edit: [schema.Role.ROLE_MANAGER],
//         view: [schema.Role.ROLE_MANAGER],
//         suggest: [schema.Role.ROLE_MANAGER],
//       }),
//       images: schema.Collection_AccessControl_SectionActions.fromPartial({
//         edit: [schema.Role.ROLE_MANAGER],
//         view: [schema.Role.ROLE_MANAGER],
//         suggest: [schema.Role.ROLE_MANAGER],
//       }),
//       loans: schema.Collection_AccessControl_SectionActions.fromPartial({
//         edit: [schema.Role.ROLE_MANAGER],
//         view: [schema.Role.ROLE_MANAGER],
//         suggest: [schema.Role.ROLE_MANAGER],
//       }),
//       grants: schema.Collection_AccessControl_SectionActions.fromPartial({
//         edit: [schema.Role.ROLE_MANAGER],
//         view: [schema.Role.ROLE_MANAGER],
//         suggest: [schema.Role.ROLE_MANAGER],
//       }),
//       hidden: schema.Collection_AccessControl_SectionActions.fromPartial({
//         edit: [schema.Role.ROLE_MANAGER],
//         view: [schema.Role.ROLE_MANAGER],
//         suggest: [schema.Role.ROLE_MANAGER],
//       }),
//       suggested: schema.Collection_AccessControl_SuggestedActions.fromPartial({
//         approve: [schema.Role.ROLE_MANAGER],
//         reject: [schema.Role.ROLE_MANAGER],
//         view: [schema.Role.ROLE_MANAGER],
//       }),
//     }),
//   });
// };

// function formatDate(date: Date) {
//   return `${date.getMonth() + 1}/${date.getDate()}/${date.getFullYear()}`;
// }

// export function unFormatDate(date: string) {
//   // Split the date into parts
//   const parts = date.split('/');
//   if (parts.length !== 3) {
//     console.warn('Invalid date format');
//   }

//   // make a new date
//   const year = parseInt(parts[2], 10);
//   const month = parseInt(parts[0], 10) - 1;
//   const day = parseInt(parts[1], 10);

//   return new Date(year, month, day);
// }
// /**
//  * @function SpecimenToFormSpecimen
//  * @description Convert a specimen object from the Biochain API a format used in the display form
//  *
//  * @param specimen A specimen object from the Biochain API
//  */
// export function SpecimenToFormSpecimen(specimen: schema.Specimen) {
//   if (!specimen) {
//     return schema.Specimen.create({});
//   }

//   if (!specimen.primary) {
//     specimen.primary = schema.Specimen_Primary.create({});
//   } else {
//     // convert the dates
//     if (specimen.primary.catalog_date) {
//       specimen.primary.catalog_date = formatDate(
//         new Date(specimen.primary.catalog_date)
//       );
//     }
//     if (specimen.primary.determined_date) {
//       specimen.primary.determined_date = formatDate(
//         new Date(specimen.primary.determined_date)
//       );
//     }
//     if (specimen.primary.field_date) {
//       specimen.primary.field_date = formatDate(
//         new Date(specimen.primary.field_date)
//       );
//     }
//   }
//   if (!specimen.secondary) {
//     specimen.secondary = schema.Specimen_Secondary.create({});
//   }
//   if (!specimen.taxon) {
//     specimen.taxon = schema.Specimen_Taxon.create({});
//   }
//   if (!specimen.georeference) {
//     specimen.georeference = schema.Specimen_Georeference.create({});
//   }

//   console.log(specimen);
//   return specimen;
// }

// /**
//  * @function FormSpecimenToSpecimen
//  * @description Convert a the display form specimen object to a specimen object from the Biochain API
//  *
//  * This mostly just returns the dates back to the correct format for timestamps
//  *
//  * @param specimen A specimen object from the Biochain API
//  */
// export function FormSpecimenToSpecimen(specimen: schema.Specimen) {
//   if (!specimen.primary) {
//     specimen.primary = schema.Specimen_Primary.create({});
//   } else {
//     // convert the dates
//     if (specimen.primary.catalog_date) {
//       specimen.primary.catalog_date = unFormatDate(
//         specimen.primary.catalog_date
//       ).toISOString();
//     }
//     if (specimen.primary.determined_date) {
//       specimen.primary.determined_date = unFormatDate(
//         specimen.primary.determined_date
//       ).toISOString();
//     }
//     if (specimen.primary.field_date) {
//       specimen.primary.field_date = unFormatDate(
//         specimen.primary.field_date
//       ).toISOString();
//     }
//   }
//   if (!specimen.secondary) {
//     specimen.secondary = schema.Specimen_Secondary.create({});
//   }
//   if (!specimen.taxon) {
//     specimen.taxon = schema.Specimen_Taxon.create({});
//   }

//   return specimen;
// }
