import * as state from '../schema/state_pb';

function MakeEmptySpecimenActions() {
  return new state.Collection_AccessControl_SpecimenActions({
    view: [],
    create: [],
    delete: [],
    hideTx: [],
  });
}

function MakeEmptySectionActions() {
  return new state.Collection_AccessControl_SectionActions({
    view: [],
    edit: [],
    suggestApprove: [],
    suggestEdit: [],
    suggestReject: [],
  });
}

function MakeEmptyCollectionAccessControlActions() {
  return new state.Collection_AccessControl_AccessControlActions({
    view: [],
    edit: [],
  });
}

function MakeEmptyCollectionAccessControl() {
  return new state.Collection_AccessControl({
    roles: MakeEmptyCollectionAccessControlActions(),
    users: MakeEmptyCollectionAccessControlActions(),
    specimen: MakeEmptySpecimenActions(),
    primary: MakeEmptySectionActions(),
    secondary: MakeEmptySectionActions(),
    taxon: MakeEmptySectionActions(),
    georeference: MakeEmptySectionActions(),
    images: MakeEmptySectionActions(),
    loans: MakeEmptySectionActions(),
    grants: MakeEmptySectionActions(),
  });
}

export function MakeEmptyCollection() {
  return new state.Collection({
    accessControl: MakeEmptyCollectionAccessControl(),
    id: new state.Collection_Id(),
  });
}

function MakeBaseSpecimenActions() {
  return new state.Collection_AccessControl_SpecimenActions({
    view: [0, 1, 2, 3, 4],
    create: [2, 3, 4],
    delete: [4],
    hideTx: [4],
  });
}

function MakeBaseSectionActions() {
  return new state.Collection_AccessControl_SectionActions({
    view: [0, 1, 2, 3, 4],
    edit: [2, 3, 4],
    suggestApprove: [3, 4],
    suggestEdit: [1, 2, 3, 4],
    suggestReject: [3, 4],
  });
}

function MakeBaseCollectionAccessControlActions() {
  return new state.Collection_AccessControl_AccessControlActions({
    view: [0, 1, 2, 3, 4],
    edit: [4],
  });
}

export function MakeBaseAccessControls() {
  return new state.Collection_AccessControl({
    roles: MakeBaseCollectionAccessControlActions(),
    users: MakeBaseCollectionAccessControlActions(),
    specimen: MakeBaseSpecimenActions(),
    primary: MakeBaseSectionActions(),
    secondary: MakeBaseSectionActions(),
    taxon: MakeBaseSectionActions(),
    georeference: MakeBaseSectionActions(),
    images: MakeBaseSectionActions(),
    loans: MakeBaseSectionActions(),
    grants: MakeBaseSectionActions(),
  });
}
