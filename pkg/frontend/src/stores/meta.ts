import { acceptHMRUpdate, defineStore } from 'pinia';

import { useAuthStore } from './auth';
import { useRouteParams } from '@vueuse/router';

import { schema } from 'src/lib/ccbio';
import { PlainMessage } from '@bufbuild/protobuf';

import { ccApi } from 'src/boot/axios';

export const useMetaStore = defineStore('meta', () => {
  const authStore = useAuthStore();

  // Raw Values
  const collections: Ref<PlainMessage<schema.Collection_List>> = ref(
    new schema.Collection_List()
  );

  const selectedCollectionId = useRouteParams('collectionId', '');

  const selectedCollection = computed(() => {
    return (
      collections.value.items.find(
        (item) => item.id?.collectionId === selectedCollectionId.value
      ) || new schema.Collection()
    );
  });
  const users: Ref<PlainMessage<schema.User_List>> = ref(
    new schema.User_List()
  );
  // Computed Values

  const collectionNames = computed(() => {
    return collections.value.items.map((c) => c.id?.collectionId || '');
  });

  const currentUser = computed(() => {
    console.log(authStore.username);
    return authStore.username;
  });

  // The roles of the current user, if no user default to public for all
  const currentRole = computed(() => {
    if (
      !authStore.isLoggedIn ||
      !authStore.blockchainUser ||
      !selectedCollection.value
    ) {
      return schema.Role.PUBLIC_UNSPECIFIED;
    }

    const role =
      authStore.blockchainUser.memberships[selectedCollectionId.value];

    return role || schema.Role.PUBLIC_UNSPECIFIED;
  });

  const currentEditPermissions = computed(() => {
    if (
      authStore.isLoggedIn &&
      selectedCollection.value.accessControl &&
      currentRole
    ) {
      const ac = selectedCollection.value.accessControl;
      const role = currentRole.value;
      return new schema.Permissions({
        roles: ac.roles?.edit.includes(role) || false,
        users: ac.users?.edit.includes(role) || false,
        specimen: false,
        primary: ac.primary?.edit.includes(role) || false,
        secondary: ac.secondary?.edit.includes(role) || false,
        taxon: ac.taxon?.edit.includes(role) || false,
        georeference: ac.georeference?.edit.includes(role) || false,
        images: ac.images?.edit.includes(role) || false,
        loans: ac.loans?.edit.includes(role) || false,
        grants: ac.grants?.edit.includes(role) || false,
      });
    }
  });
  const currentSuggestEditPermissions = computed(() => {
    if (
      authStore.isLoggedIn &&
      selectedCollection.value.accessControl &&
      currentRole
    ) {
      const ac = selectedCollection.value.accessControl;
      const role = currentRole.value;
      return new schema.Permissions({
        roles: false,
        users: false,
        specimen: false,
        primary: ac.primary?.suggestEdit.includes(role) || false,
        secondary: ac.secondary?.suggestEdit.includes(role) || false,
        taxon: ac.taxon?.suggestEdit.includes(role) || false,
        georeference: ac.georeference?.suggestEdit.includes(role) || false,
        images: ac.images?.suggestEdit.includes(role) || false,
        loans: ac.loans?.suggestEdit.includes(role) || false,
        grants: ac.grants?.suggestEdit.includes(role) || false,
      });
    }
  });

  const currentSuggestApprovePermissions = computed(() => {
    if (
      authStore.isLoggedIn &&
      selectedCollection.value.accessControl &&
      currentRole
    ) {
      const ac = selectedCollection.value.accessControl;
      const role = currentRole.value;
      return new schema.Permissions({
        roles: false,
        users: false,
        specimen: false,
        primary: ac.primary?.suggestApprove.includes(role) || false,
        secondary: ac.secondary?.suggestApprove.includes(role) || false,
        taxon: ac.taxon?.suggestApprove.includes(role) || false,
        georeference: ac.georeference?.suggestApprove.includes(role) || false,
        images: ac.images?.suggestApprove.includes(role) || false,
        loans: ac.loans?.suggestApprove.includes(role) || false,
        grants: ac.grants?.suggestApprove.includes(role) || false,
      });
    }
  });

  const currentSuggestRejectPermissions = computed(() => {
    if (
      authStore.isLoggedIn &&
      selectedCollection.value.accessControl &&
      currentRole
    ) {
      const ac = selectedCollection.value.accessControl;
      const role = currentRole.value;
      return new schema.Permissions({
        roles: false,
        users: false,
        specimen: false,
        primary: ac.primary?.suggestReject.includes(role) || false,
        secondary: ac.secondary?.suggestReject.includes(role) || false,
        taxon: ac.taxon?.suggestReject.includes(role) || false,
        georeference: ac.georeference?.suggestReject.includes(role) || false,
        images: ac.images?.suggestReject.includes(role) || false,
        loans: ac.loans?.suggestReject.includes(role) || false,
        grants: ac.grants?.suggestReject.includes(role) || false,
      });
    }
  });

  // Methods

  function getCollection(name: string) {
    if (collections.value.items.length === 0) {
      refresh();
      if (collections.value.items.length === 0) {
        return;
      }
    }
    return collections.value.items.find((c) => c.id?.collectionId === name);
  }

  function refresh() {
    if (authStore.isLoggedIn) {
      authStore.setHeaders();
    }

    ccApi.auth.GetUserList().then((res) => {
      users.value = res;
    });
    ccApi.auth.GetCollectionList().then((res) => {
      collections.value = res;
    });
  }

  // function setCurrentCollection(id: string) {
  //   selectedCollectionId.value = id;
  // }

  return {
    collections,
    collectionNames,
    currentUser,
    currentRole,
    currentEditPermissions,
    currentSuggestApprovePermissions,
    currentSuggestEditPermissions,
    currentSuggestRejectPermissions,
    users,
    refresh,
    selectedCollection,
    selectedCollectionId,
    getCollection,
  };
});
if (import.meta.hot) {
  import.meta.hot.accept(acceptHMRUpdate(useMetaStore, import.meta.hot));
}
